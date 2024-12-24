USE healthcare;

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
CREATE INDEX idx_medications_created ON medications(created);
-- TODO: Refine index;

-- Add some dummy records so our application is already populated when start
INSERT INTO medications (name, dosage, form) VALUES (
    'Paracetamol',
    '500 mg',
    'Tablet'
);

INSERT INTO medications (name, dosage, form) VALUES (
    'Lexapro',
    '10 mg',
    'Tablet'
);

INSERT INTO medications (name, dosage, form) VALUES (
    'Melatonin',
    '5 mg',
    'Capsule'
);

-- Granting just the required permissions for the user that our REST API will use
CREATE USER 'web'@'%';
GRANT ALL PRIVILEGES ON healthcare.* TO 'web'@'%';
ALTER USER 'web'@'%' IDENTIFIED BY 'password';
FLUSH PRIVILEGES;
