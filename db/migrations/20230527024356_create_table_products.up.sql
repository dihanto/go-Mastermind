CREATE Table products(
    id int not NULL AUTO_INCREMENT,
    name VARCHAR(100) not  NULL,
    price int not NULL,
    quantity int not NULL,
    category_id int,
    PRIMARY KEY(id),
    Foreign Key (category_id) REFERENCES categories(id)
)Engine = InnoDb;