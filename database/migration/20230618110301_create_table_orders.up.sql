CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    id_product integer,
    id_customer UUID,
    code_bank integer,
    FOREIGN KEY (id_product) REFERENCES products(id),
    FOREIGN KEY (id_customer) REFERENCES customers(id),
    FOREIGN KEY (code_bank) REFERENCES banks(code)
);