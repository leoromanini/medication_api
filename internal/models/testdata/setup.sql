USE test_healthcare;

CREATE TABLE medications (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    dosage VARCHAR(100) NOT NULL,
    form VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT 1,
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_update TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
CREATE INDEX idx_medications_is_active ON medications(is_active);

INSERT INTO medications (name, dosage, form) VALUES (
    'Amoxicillin',
    '250 mg',
    'Tablet'
);

INSERT INTO medications (name, dosage, form) VALUES (
    'Ozempic',
    '0.25 mg',
    'Pen'
);

INSERT INTO medications (name, dosage, form) VALUES (
    'Citalopram',
    '30 mg',
    'Capsule'
);
