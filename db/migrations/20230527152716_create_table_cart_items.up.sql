CREATE Table cart_items (
    cart_item_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    cart_id int not NULL,
    Foreign Key (cart_id) REFERENCES carts(cart_id),
    product_id int NOT NULL,
    Foreign Key (product_id) REFERENCES products(product_id),
    quantity int NOT NULL
)Engine=InnoDB;