package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type DB interface {
	CreateInspection(*CreateInspectionInput) (*Inspection, error)
	CreateRestaurant(*CreateRestaurantInput) (*Restaurant, error)
	CreateViolation(*CreateViolationInput) (*Violation, error)
}

type db struct {
	db     *gorm.DB
	logger logrus.FieldLogger
}

type CreateInspectionInput struct {
	CuisineDescription   *string `json:"cuisine_description,omitempty"`
	DBA                  *string `json:"dba,omitempty"`
	Boro                 *string `json:"boro,omitempty"`
	InspectionDate       *string `json:"inspection_date,omitempty"`
	Building             *string `json:"building,omitempty"`
	ZipCode              *string `json:"zipcode,omitempty"`
	Score                *string `json:"score,omitempty"`
	Phone                *string `json:"phone,omitempty"`
	Street               *string `json:"street,omitempty"`
	Grade                *string `json:"grade,omitmepty"`
	CriticalFlag         *string `json:"critical_flag,omitmepty"`
	Camis                *string `json:"camis,omitmepty"`
	Action               *string `json:"action,omitmepty"`
	ViolationCode        *string `json:"violation_code,omitempty"`
	ViolationDescription *string `json:"violation_description,omitempty"`
	GradeDate            *string `json:"grade_date,omitempty"`
	InspectionType       *string `json:"inspection_type,omitempty"`
}

func (db *db) CreateInspection(i *CreateInspectionInput) (*Inspection, error) {

	// Look-up restaurant, if none matching exists, create one
	var restaurant *Restaurant
	err := db.db.Where("camis = ?", i.Camis).Find(restaurant).Error
	if err != nil {
		db.logger.WithField("camis", i.Camis).Error(err)
		return nil, errors.Wrap(err, "error retrieving restaurant by camis")
	}
	if restaurant.ID == 0 {
		restaurant, err = db.CreateRestaurant(&CreateRestaurantInput{
			Camis:              i.Camis,
			DBA:                i.DBA,
			Boro:               i.Boro,
			Building:           i.Building,
			Street:             i.Street,
			ZipCode:            i.ZipCode,
			Phone:              i.Phone,
			CuisineDescription: i.CuisineDescription,
		})
		if err != nil {
			db.logger.WithField("camis", i.Camis).Error(err)
			return nil, errors.Wrap(err, fmt.Sprintf("error creating restaurant: %s", i.Camis))
		}
	}

	// Look-up violation, if none matching exists, create one
	var violation *Violation
	err = db.db.Where("code = ?", i.ViolationCode).Find(violation).Error
	if err != nil {
		db.logger.WithField("violation_code", i.ViolationCode).Error(err)
		return nil, errors.Wrap(err, fmt.Sprintf("could not query for violation by code: %s", i.ViolationCode))
	}
	if violation.ID == 0 {
		violation, err = db.CreateViolation(&CreateViolationInput{
			Code:        i.ViolationCode,
			Description: i.ViolationDescription,
		})
		if err != nil {
			db.logger.WithField("violation_code", i.ViolationCode).Error(err)
			return nil, errors.Wrap(err, fmt.Sprintf("error creating violation: %s", i.ViolationCode))
		}
	}

	var grade Grade
	err = db.db.Where("name = ?", i.Grade).Find(&grade).Error
	if err != nil {
		db.logger.WithField("grade", i.Grade).Error(err)
		return nil, errors.Wrap(err, "error retrieving grade by name")
	}

	inspectionDate, err := strToDate(i.InspectionDate)
	if err != nil {
		return nil, errors.Wrap(err, "cannot convert inspection date")
	}
	gradeDate, err := strToDate(i.GradeDate)
	if err != nil {
		return nil, errors.Wrap(err, "cannot convert grade date")
	}

	inspec := Inspection{
		RestaurantID:   restaurant.ID,
		InspectionType: stringP(i.InspectionType),
		Action:         stringP(i.Action),
		ViolationID:    violation.ID,
		Score:          stringP(i.Score),
		GradeID:        grade.ID,
		GradeDate:      gradeDate,
		InspectionDate: inspectionDate,
	}

	if err = db.db.Create(&inspec).Error; err != nil {
		db.logger.WithFields(logrus.Fields{
			"restaurant_id":   inspec.RestaurantID,
			"inspection_date": inspec.InspectionDate,
			"action":          inspec.Action,
			"violation_id":    inspec.ViolationID,
			"score":           inspec.Score,
			"grade_id":        inspec.GradeID,
			"grade_date":      inspec.GradeDate,
			"inspection_type": inspec.InspectionType,
		}).Error(err)
		return nil, errors.Wrap(err, "error creating inspection")
	}

	return &inspec, nil
}

type CreateRestaurantInput struct {
	Camis              *string `json:"camis,omitempty"`
	DBA                *string `json:"dba,omitempty"`
	Boro               *string `json:"boro,omitempty"`
	Building           *string `json:"building,omitempty"`
	Street             *string `json:"street,omitempty"`
	ZipCode            *string `json:"zip_code,omitempty"`
	Phone              *string `json:"phone,omitempty"`
	CuisineDescription *string `json:"cuisine_description,omitempty"`
}

func (db *db) CreateRestaurant(i *CreateRestaurantInput) (*Restaurant, error) {
	var boro Boro
	err := db.db.Where("name = ?", i.Boro).Find(&boro).Error
	if err != nil {
		db.logger.WithField("name", i.Boro).Error(err)
		return nil, errors.Wrap(err, "error retrieving boro by name")
	}
	if boro.ID == 0 {
		err := errors.New("no boro found with provided name")
		db.logger.WithField("boro", i.Boro).Error(err)
		return nil, err
	}

	restaurant := Restaurant{
		Camis:              stringP(i.Camis),
		DBA:                stringP(i.DBA),
		BoroID:             boro.ID,
		Building:           stringP(i.Building),
		Street:             stringP(i.Street),
		ZipCode:            atoi(i.ZipCode),
		Phone:              atoi(i.Phone),
		CuisineDescription: stringP(i.CuisineDescription),
	}
	if err := db.db.Create(&restaurant).Error; err != nil {
		db.logger.WithFields(logrus.Fields{
			"camis":               restaurant.Camis,
			"dba":                 restaurant.DBA,
			"boro_id":             restaurant.BoroID,
			"building":            restaurant.Building,
			"street":              restaurant.Street,
			"zip_code":            restaurant.ZipCode,
			"phone":               restaurant.Phone,
			"cuisine_description": restaurant.CuisineDescription,
		}).Error(err)
		return nil, errors.Wrap(err, "encountered error creating restaurant")
	}

	return &restaurant, nil
}

type CreateViolationInput struct {
	ID          uint32  `json:"id,omitempty"`
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (db *db) CreateViolation(i *CreateViolationInput) (*Violation, error) {
	violation := Violation{
		Code:        stringP(i.Code),
		Description: stringP(i.Description),
	}
	if err := db.db.Create(&violation).Error; err != nil {
		return nil, errors.Wrap(err, "encountered error creating violation")
	}

	return &violation, nil
}
