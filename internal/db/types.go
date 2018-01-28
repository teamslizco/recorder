package db

import (
	"time"
)

type Inspection struct {
	ID             uint32    `json:"id,omitempty"`
	RestaurantID   uint32    `json:"restaurant_id,omitempty"`
	InspectionType string    `json:"inspection_type,omitempty"`
	Action         string    `json:"action,omitmepty"`
	ViolationID    uint32    `json:violation_id,omitempty"`
	Score          string    `json:"score,omitempty"`
	GradeID        uint32    `json:"grade,omitmepty"`
	GradeDate      time.Time `json:"grade_date,omitempty"`
	InspectionDate time.Time `json:"inspection_date,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

type Restaurant struct {
	ID                 uint32    `json:"id,omitempty"`
	Camis              string    `json:"camis,omitempty"`
	DBA                string    `json:"dba,omitempty"`
	BoroID             uint32    `json:"boro_id,omitempty"`
	Building           string    `json:"building,omitempty"`
	Street             string    `json:"street,omitempty"`
	ZipCode            int       `json:"zip_code,omitempty"`
	Phone              int       `json:"phone,omitempty"`
	CuisineDescription string    `json:"cuisine_description,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
}

type Boro struct {
	ID   uint32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Grade struct {
	ID   uint32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Violation struct {
	ID          uint32    `json:"id,omitempty"`
	Code        string    `json:"code,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
