CREATE TABLE logs (
    id SERIAL PRIMARY KEY,
    method VARCHAR(10),
    path VARCHAR(255),
    status_code INT,
    duration INTERVAL
);