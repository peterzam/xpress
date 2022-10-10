USE 'xpress_db';

CREATE TABLE users(
    id INT AUTO_INCREMENT PRIMARY KEY,
    phone VARCHAR(20) NOT NULL UNIQUE,
    name VARCHAR(30) NOT NULL,
    address VARCHAR(50) NOT NULL,
    password VARCHAR(32) NOT NULL,
    role VARCHAR(10) NOT NULL
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
    status VARCHAR(50) NOT NULL,
    src_id INT,
    created_at INT NOT NULL,
    FOREIGN KEY (src_id) REFERENCES users(id)
);

CREATE TABLE offices(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    location_lat VARCHAR(10) NOT NULL,
    location_lng VARCHAR(10) NOT NULL,
    address VARCHAR(50) NOT NULL,
    phone VARCHAR(20) NOT NULL
);


INSERT INTO users (phone, name, address, password, role) VALUES
('00000000','Admin','Yangon Main Street','dd4b21e9ef71e1291183a46b913ae6f2','admin'),
('11111111','Aung Aung','Mandalay Second Street','1bbd886460827015e5d605ed44252251','user'),
('22222222','Bo Bo','Yangon 3rd Street','bae5e3208a3c700e3db642b6631e95b9','user'),
('33333333','Cha Cha','Mandalay Block 1','d27d320c27c3033b7883347d8beca317','inactive'),
('44444444','Daung Daung','Yangon Block 2','b857eed5c9405c1f2b98048aae506792','user');

INSERT INTO packages(code, dest_name, dest_addr, dest_phone, size, type, note, status, src_id, created_at) VALUES
('V2N3RC','Aung','Pyin Oo Lwin ','01010101','20cm x 50cm x 2cm','express','Heavy Books',"Send",3,1665384385),
('UYTZWM','Kyaw','Bago','02020202','50cm x 50cm x 30cm','normal','Fragile Items',"Received at Mandalay Office",2,1665384385);

INSERT INTO offices(name, location_lat, location_lng, address, phone) VALUES
('Yangon Main Office','16.8409','96.1735','Yangon Main Street','00100100'),
('Mandalay Branch','21.9588','96.1735','Mandalay Main Street','00200200'),
('Nay Pyi Taw Branch','19.7450','96.12972','Nay Pyi Taw Main Street','00300300'),
('Mawlamyine Branch','16.49051','97.62825','Mawlamyine Main Street','00400400'),
('Bago Branch','17.33521','96.48135','Bago Main Street','00500500'),
('Pathein Branch','16.77919','94.73212','Pathein Main Street','00600600'),
('Monywa Branch','	22.10856','95.13583','Monywa Main Street','00700700'),
('Sittwe Branch','20.144444','92.896942','Sittwe Main Street','00800800'),
('Meiktila Branch','20.87776','95.85844','Meiktila Main Street','00900900'),
('Taunggyi Branch','20.78919','97.03776','Taunggyi Main Street','01001000'),
('Magway Branch','20.14956','94.93246','Magway Main Street','01101100'),
('Myeik Branch','12.43954','98.60028','Myeik Main Street','01201200'),
('Dawei Branch','14.0823','98.19151','Dawei Main Street','01301300'),
('Myitkyina Branch','25.38327','97.39637','Myitkyina Main Street','01401400'),
('Lashio Branch','22.9359','97.7498','Lashio Main Street','01501500');