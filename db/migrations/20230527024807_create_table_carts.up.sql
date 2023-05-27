CREATE Table carts(
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    Foreign Key (products_id) REFERENCES products(id),
    quantity int NOT NULL
)engine = innodb;