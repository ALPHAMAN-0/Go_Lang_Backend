-- Table 1: Users (stores name)
CREATE TABLE IF NOT EXISTS users (
    id        SERIAL PRIMARY KEY,
    name      VARCHAR(100) NOT NULL,
    email     VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Table 2: Phones (stores phone number, linked to user)
CREATE TABLE IF NOT EXISTS phones (
    id          SERIAL PRIMARY KEY,
    user_id     INT REFERENCES users(id) ON DELETE CASCADE,
    phone_number VARCHAR(20) NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW()
);

-- Seed some data
INSERT INTO users (name, email) VALUES ('Alice', 'alice@example.com');
INSERT INTO users (name, email) VALUES ('Bob', 'bob@example.com');

INSERT INTO phones (user_id, phone_number) VALUES (1, '+1-555-0101');
INSERT INTO phones (user_id, phone_number) VALUES (2, '+1-555-0202');