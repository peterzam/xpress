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

CREATE TABLE pickups(
    id INT AUTO_INCREMENT PRIMARY KEY,
    size VARCHAR(50) NOT NULL,
    type VARCHAR(20) NOT NULL,
    note VARCHAR(50),
    src_id INT,
    FOREIGN KEY (src_id) REFERENCES users(id)
);

CREATE TABLE complaints(
    id INT AUTO_INCREMENT PRIMARY KEY,
    note VARCHAR(50),
    src_id INT,
    FOREIGN KEY (src_id) REFERENCES users(id)
);

INSERT INTO users (phone, name, address, password, role) VALUES
('00000000','Admin','Yangon Main Street','dd4b21e9ef71e1291183a46b913ae6f2','admin'),
('11111111','Aung Aung','Mandalay Second Street','1bbd886460827015e5d605ed44252251','user'),
('22222222','Bo Bo','Yangon 3rd Street','bae5e3208a3c700e3db642b6631e95b9','user'),
('33333333','Chan Chan','Mandalay Block 1','d27d320c27c3033b7883347d8beca317','inactive'),
('44444444','Daung Daung','Yangon Block 2','b857eed5c9405c1f2b98048aae506792','user');

INSERT INTO pickups(size, type, note, src_id) VALUES
('2XL', 'normal', 'Function-based impactful support', 3),
('XS', 'normal', 'Configurable multimedia migration', 5);

INSERT INTO complaints(note, src_id) VALUES
('Slow support', 4),
('Configurable multimedia migration', 3);

INSERT INTO packages(code, dest_name, dest_addr, dest_phone, size, type, note, status, src_id, created_at) VALUES
('NUTLA', 'Purcell', '30 Farwell Terrace', '7071080999', '2XL', 'normal', 'Function-based impactful support', 'Received at Dawei Office', 3, 1655405185),
('DRTSH', 'Johann', '94389 East Way', '8619810583', 'XS', 'normal', 'Configurable multimedia migration', 'sent', 5, 1642581403),
('PDSAY', 'Yoshiko', '23309 Pepper Wood Street', '5936090467', '3XL', 'normal', 'Cross-group systematic Graphical User Interface', 'sent', 1, 1662706734),
('CWETJ', 'Michele', '011 Forest Dale Crossing', '7963184975', 'XL', 'business', 'Automated fresh-thinking moderator', 'sent', 4, 1642077003),
('JNFTT', 'Abigael', '325 8th Trail', '6838041438', 'XL', 'normal', 'Universal web-enabled Graphical User Interface', 'Received at Monywa Office', 1, 1637102322),
('VTPLI', 'Micheal', '925 Charing Cross Center', '7588735466', 'M', 'business', 'Versatile 5th generation open architecture', 'sent', 1, 1663454348),
('GFDGF', 'Ade', '02680 Crest Line Street', '4525002567', 'M', 'business', 'Configurable 4th generation encryption', 'Received at Dawei Office', 3, 1659821376),
('ZAEMG', 'Gayle', '21 Bunker Hill Place', '8075908344', '3XL', 'express', 'Enhanced uniform complexity', 'Received at Dawei Office', 3, 1642353903),
('RAPVG', 'Gleda', '0338 Hoffman Drive', '5301593785', '2XL', 'express', 'Optimized 6th generation time-frame', 'sent', 1, 1637100583),
('DPYHU', 'Quentin', '2 Miller Trail', '7482117095', 'S', 'business', 'Reduced bi-directional project', 'sent', 2, 1643781920),
('FCTPU', 'Inessa', '6681 Nelson Lane', '1779332352', '2XL', 'business', 'Function-based intermediate hub', 'sent', 2, 1660581401),
('UXRHE', 'Tommy', '9 Michigan Road', '3353518657', 'L', 'business', 'Open-source content-based conglomeration', 'sent', 2, 1643275700),
('NIIVU', 'Paola', '3 Cottonwood Way', '1856342370', 'XL', 'normal', 'Optimized logistical focus group', 'Received at Bago Office', 4, 1656796375),
('SHQDE', 'Roze', '79 Steensland Plaza', '7367948855', 'S', 'normal', 'Innovative optimal superstructure', 'sent', 1, 1667126584),
('JDXWK', 'Winny', '25370 Kinsman Drive', '2887089928', 'L', 'express', 'Progressive object-oriented local area network', 'sent', 4, 1656495543),
('PXWUG', 'Teddie', '756 Farwell Center', '3854821594', '2XL', 'business', 'Customer-focused stable encoding', 'Received at Mandalay Office', 2, 1651562511),
('WEBER', 'Thorpe', '4 Brown Hill', '4737394773', 'M', 'normal', 'Robust systematic forecast', 'sent', 1, 1635477190),
('QUYYX', 'Kellby', '24 Granby Park', '8773455005', 'XS', 'business', 'Grass-roots contextually-based hub', 'Received at Dawei Office', 1, 1651461632),
('CLXDS', 'Alene', '3039 Nelson Circle', '9951012561', 'M', 'normal', 'Compatible tertiary structure', 'sent', 3, 1651753051),
('RYSGT', 'Harbert', '1743 Sherman Drive', '8383790030', 'S', 'normal', 'Upgradable mission-critical hardware', 'sent', 3, 1635678872),
('THFRM', 'Westbrooke', '0945 Bashford Pass', '4225607894', 'L', 'business', 'Ergonomic clear-thinking groupware', 'Received at Monywa Office', 1, 1641817470),
('QZESJ', 'Allegra', '4 Westport Park', '6173112104', 'S', 'normal', 'Profound needs-based intranet', 'Received at Dawei Office', 3, 1642966253),
('SPDTG', 'Sasha', '39 Stoughton Point', '7295036237', 'XL', 'normal', 'Multi-layered coherent migration', 'sent', 2, 1653021154),
('JQYTF', 'Clerc', '31418 Packers Center', '8002600045', '2XL', 'normal', 'Organized encompassing function', 'Received at Bago Office', 1, 1656088190),
('OYWSI', 'Leonardo', '07 Shelley Park', '1069059474', '2XL', 'business', 'Ameliorated dedicated Graphic Interface', 'sent', 1, 1647979285),
('QMLBL', 'Jessie', '15 Barnett Park', '7301214355', 'L', 'express', 'Function-based multi-state infrastructure', 'Received at Bago Office', 5, 1640841733),
('EIWUB', 'Boone', '177 Canary Drive', '5693150597', '2XL', 'business', 'Balanced 5th generation firmware', 'Received at Dawei Office', 5, 1660539904),
('SWOYV', 'Mair', '8690 Chive Road', '2717393211', 'XL', 'normal', 'Re-contextualized discrete contingency', 'Received at Monywa Office', 1, 1662049187),
('ZWYBK', 'Helyn', '69 Shasta Hill', '5717854211', 'L', 'normal', 'Expanded value-added matrices', 'sent', 2, 1660032783),
('YPKIY', 'Madelon', '4479 Delladonna Place', '9746539209', '2XL', 'express', 'Business-focused zero defect hierarchy', 'sent', 5, 1635750540),
('UZRXE', 'Thorvald', '4 Cardinal Terrace', '7019998408', '2XL', 'normal', 'Triple-buffered encompassing process improvement', 'sent', 1, 1638594683),
('BWMPX', 'Muriel', '155 Blackbird Trail', '6881456136', 'XS', 'express', 'Persistent composite local area network', 'Received at Yangon Office', 5, 1641706697),
('RGTJT', 'Obed', '6 Kennedy Pass', '1129248777', 'L', 'business', 'Grass-roots tertiary orchestration', 'sent', 5, 1660054956),
('HBGIB', 'Dollie', '8 Nobel Center', '6218327282', 'L', 'normal', 'Cross-group executive architecture', 'Received at Bago Office', 4, 1649713868),
('BAWRF', 'Vera', '86649 Marquette Trail', '9392940408', 'M', 'express', 'Total background database', 'Received at Yangon Office', 1, 1651910395),
('RTKOR', 'Elsbeth', '87 Kenwood Pass', '1691392770', '3XL', 'express', 'Automated tertiary moratorium', 'Received at Monywa Office', 5, 1652147022),
('KQKYG', 'Marshall', '76 Manley Point', '2805857593', '3XL', 'business', 'Multi-tiered value-added project', 'Received at Dawei Office', 5, 1645834295),
('WNGJN', 'Jemmy', '71 Duke Place', '9364402880', '3XL', 'business', 'Networked modular functionalities', 'Received at Monywa Office', 2, 1644125818),
('AYWIB', 'Damara', '46039 Thompson Center', '9642012010', '2XL', 'normal', 'Organic asynchronous projection', 'sent', 4, 1636243373),
('PZVYE', 'Monique', '5686 Iowa Hill', '9433364994', '2XL', 'express', 'Customer-focused impactful Graphic Interface', 'sent', 4, 1640774024),
('LBEXV', 'Kali', '78525 Hagan Alley', '7512731956', 'M', 'business', 'Business-focused system-worthy database', 'sent', 3, 1642354713),
('KTAUT', 'Sile', '9037 Reinke Park', '4909988703', 'XS', 'express', 'Profit-focused logistical productivity', 'sent', 5, 1646457858),
('HXIXZ', 'Yoshi', '32597 Clemons Point', '2138474364', '3XL', 'business', 'Multi-tiered executive ability', 'Received at Dawei Office', 1, 1641303587),
('TENCX', 'Sandy', '236 Division Point', '6523144030', 'L', 'express', 'Adaptive attitude-oriented interface', 'sent', 4, 1647911419),
('WUFRW', 'Celle', '95 Calypso Way', '3881925938', 'XL', 'express', 'Intuitive bifurcated orchestration', 'Received at Dawei Office', 4, 1649019839),
('DBSPG', 'Aubry', '45 Mesta Alley', '6297430651', 'XS', 'express', 'Down-sized cohesive toolset', 'Received at Dawei Office', 5, 1663198283),
('AVVDR', 'Basia', '562 Weeping Birch Street', '7018093805', 'S', 'express', 'Profound holistic encryption', 'sent', 2, 1645171899),
('XJJIM', 'Rakel', '46 Clove Plaza', '7492754450', 'XL', 'normal', 'Balanced grid-enabled moderator', 'sent', 1, 1662567656),
('JLJEP', 'Virgil', '66 Hanover Avenue', '9052505670', '2XL', 'normal', 'Ergonomic cohesive contingency', 'sent', 1, 1644024993),
('GKWHL', 'Rozina', '23287 Huxley Center', '3462854543', 'XS', 'business', 'Robust discrete installation', 'Received at Dawei Office', 2, 1633602243),
('V2N3RC','Aung','Pyin Oo Lwin ','01010101','20cm x 50cm x 2cm','express','Heavy Books',"sent",3,1665384385),
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