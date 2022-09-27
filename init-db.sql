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

CREATE TABLE packages(
    id INT AUTO_INCREMENT PRIMARY KEY,
    code VARCHAR(10) NOT NULL UNIQUE,
    dest_name VARCHAR(30) NOT NULL,
    dest_addr VARCHAR(50) NOT NULL,
    dest_phone VARCHAR(20) NOT NULL,
    size VARCHAR(50) NOT NULL,
    type VARCHAR(20) NOT NULL,
    note VARCHAR(50),
    active BOOLEAN,
    src_id INT,
    FOREIGN KEY (src_id) REFERENCES users(id)
);


INSERT INTO users (phone, name, address, password, role, active) VALUES
('00000000','admin','Yangon','21232f297a57a5a743894a0e4a801fc3','admin',true),
('11111111','user','Mandalay','ee11cbb19052e40b07aac0ca060c23ee','user',false);

INSERT INTO packages(code, dest_name, dest_addr, dest_phone, size, type, note, active, src_id) VALUES
('V2N3RC','Aung','Pyin Oo Lwin','01010101','20cm x 50cm x 2cm','express','books',true,1),
('UYTZWM','Kyaw','Bago','02020202','50cm x 50cm x 30cm','normal','',true,2);