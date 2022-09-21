USE 'xpress_db';

CREATE TABLE User(
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    user_name VARCHAR(30) NOT NULL UNIQUE,
    user_password VARCHAR(32) NOT NULL,
    user_role VARCHAR(10) NOT NULL,
    user_active BOOLEAN
);

INSERT INTO User (user_name, user_password, user_role) VALUES
('admin','21232f297a57a5a743894a0e4a801fc3','admin',true),
('user','ee11cbb19052e40b07aac0ca060c23ee','user',false);