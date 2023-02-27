CREATE TABLE schools(
    schools_id bigserial PRIMARY KEY,
    name text NOT NULL, 
    district VARCHAR(100) NOT NULL,
    phone VARCHAR(25) NOT NULL
);

