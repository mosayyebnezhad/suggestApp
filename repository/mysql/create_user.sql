CREATE TABLE  users (
    id int primary key auto_increment,
    name VARCHAR(255) not null ,
    phone_number VARCHAR(255) not null unique,
    password VARCHAR(255) not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
