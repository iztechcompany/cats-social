BEGIN;

CREATE TYPE enum_race AS ENUM (
    'Persian',
    'Maine Coon',
    'Siamese',
    'Ragdoll',
    'Bengal',
    'Sphynx',
    'British Shorthair',
    'Abyssinian',
    'Scottish Fold',
    'Birman'
);
CREATE TYPE enum_sex AS ENUM (
    'male',
    'female'
);
CREATE TABLE IF NOT EXISTS cats(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    race enum_race NOT NULL,
    sex enum_sex NOT NULL,
    age_in_month INTEGER CHECK (age_in_month BETWEEN 1 AND 120082) NOT NULL,
    description VARCHAR(200),
    image_urls TEXT[] NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
COMMIT;