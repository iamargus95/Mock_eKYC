BEGIN;

CREATE TABLE IF NOT EXISTS clients(
    id BIGSERIAL PRIMARY KEY,
    client_name VARCHAR(50) UNIQUE NOT NULL,
    client_email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS plans(
    id BIGSERIAL PRIMARY KEY,
    client_id INT NOT NULL,
    plan VARCHAR(10),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_clients FOREIGN KEY(client_id) REFERENCES clients(id)
);

COMMIT;