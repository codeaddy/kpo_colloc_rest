CREATE TABLE product (
                      id SERIAL PRIMARY KEY,
                      name VARCHAR(100) NOT NULL,
                      description TEXT,
                      price INT NOT NULL,
                      quantity INT NOT NULL,
                      created_at TIMESTAMP DEFAULT NOW()
);
