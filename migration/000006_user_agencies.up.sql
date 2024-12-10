CREATE TABLE IF NOT EXISTS user_agencies (
    user_id INT NOT NULL,
    agency_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, agency_id),
    FOREIGN KEY (user_id) REFERENCES user_data (user_id) ON DELETE CASCADE,
    FOREIGN KEY (agency_id) REFERENCES agencies_data (id) ON DELETE CASCADE
);