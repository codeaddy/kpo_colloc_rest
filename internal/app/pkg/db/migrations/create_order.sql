CREATE TABLE "order" (
                         id SERIAL PRIMARY KEY,
                         user_id INT NOT NULL,
                         status TEXT NOT NULL,
                         created_at TIMESTAMP DEFAULT NOW()
);
