CREATE Table carts(
    cart_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(100),
    Foreign Key (user_id) REFERENCES users(id)
)engine = innodb;