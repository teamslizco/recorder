CREATE TABLE violations (
    id SERIAL PRIMARY KEY,
    code VARCHAR(16) NOT NULL UNIQUE,
    description VARCHAR(1024) NOT NULL UNIQUE,
    critical VARCHAR(32) NOT NULL,
    vermin BOOLEAN NOT NULL
);

CREATE TABLE grades (
    id SERIAL PRIMARY KEY,
    letter VARCHAR(16) NOT NULL UNIQUE
);

CREATE TABLE boros (
    id SERIAL PRIMARY KEY,
    name VARCHAR(16) NOT NULL UNIQUE
);

CREATE TABLE restaurants (
    id SERIAL PRIMARY KEY,
    camis INTEGER NOT NULL UNIQUE,
    dba   VARCHAR(64) NOT NULL,
    boro_id INTEGER NOT NULL REFERENCES boros,
    building INTEGER NOT NULL,
    street VARCHAR(64) NOT NULL,
    zip_code INTEGER NOT NULL,
    phone INTEGER NOT NULL,
    cuisine VARCHAR(128) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE inspections (
    id    SERIAL PRIMARY KEY,
    restaurant_id INTEGER NOT NULL REFERENCES restaurants,
    inspection_date date NOT NULL,
    action VARCHAR(256) NOT NULL,
    violation_id INTEGER NOT NULL REFERENCES violations,
    score INTEGER NOT NULL,
    grade_id INTEGER NOT NULL REFERENCES grades,
    grade_date date NOT NULL,
    record_date date NOT NULL,
    inspection_type VARCHAR(256) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
