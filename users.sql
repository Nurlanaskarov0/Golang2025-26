CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    balance NUMERIC(10,2) DEFAULT 0
);

INSERT INTO users (name, email, balance) VALUES
('Alice', 'alice@example.com', 100.50),
('Bob', 'bob@example.com', 50.00),
('Charlie', 'charlie@example.com', 0.00);
