-- Application resources init

USE healthcare;

CREATE TABLE medications (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    dosage VARCHAR(20) NOT NULL,
    form VARCHAR(20) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT 1,
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_update TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
CREATE INDEX idx_medications_is_active ON medications(is_active);
-- TODO: In a real case scenario, index would demand more analisys to be set;

-- Add some dummy records so our application is already populated when start
-- Holpefully this will give a little help when testing
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
GRANT SELECT, INSERT, UPDATE ON healthcare.medications TO 'web'@'%';
ALTER USER 'web'@'%' IDENTIFIED BY 'password';

-- Integration tests resources init
CREATE DATABASE test_healthcare CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'test_web'@'%';
GRANT ALL PRIVILEGES ON test_healthcare.* TO 'test_web'@'%';
ALTER USER 'test_web'@'%' IDENTIFIED BY 'test_password';

FLUSH PRIVILEGES;
