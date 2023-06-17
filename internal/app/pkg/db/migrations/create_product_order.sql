CREATE TABLE product_order (
                         id SERIAL PRIMARY KEY,
                         product_id INT NOT NULL,
                         order_id INT NOT NULL,
                         quantity INT NOT NULL,
                         created_at TIMESTAMP DEFAULT NOW()
);
