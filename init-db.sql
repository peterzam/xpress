USE 'xpress_db';

CREATE TABLE users(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(30) NOT NULL UNIQUE,
    password VARCHAR(32) NOT NULL,
    role VARCHAR(10) NOT NULL,
    active BOOLEAN
);

INSERT INTO users (name, password, role, active) VALUES
('admin','21232f297a57a5a743894a0e4a801fc3','admin',true),
('user','ee11cbb19052e40b07aac0ca060c23ee','user',false);