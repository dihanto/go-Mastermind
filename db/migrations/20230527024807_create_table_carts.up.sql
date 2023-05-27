CREATE Table carts(
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    products_id INT,
    Foreign Key (products_id) REFERENCES products(id),
    quantity int
)engine = innodb;