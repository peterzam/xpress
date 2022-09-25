USE 'xpress_db';

CREATE TABLE users(
    id INT AUTO_INCREMENT PRIMARY KEY,
    phone VARCHAR(20) NOT NULL UNIQUE,
    name VARCHAR(30) NOT NULL,
    address VARCHAR(50) NOT NULL,
    password VARCHAR(32) NOT NULL,
    role VARCHAR(10) NOT NULL,
    active BOOLEAN
);

INSERT INTO users (phone, name, address, password, role, active) VALUES
('00000000','admin','Yangon','21232f297a57a5a743894a0e4a801fc3','admin',true),
('11111111','user','Mandalay','ee11cbb19052e40b07aac0ca060c23ee','user',false);