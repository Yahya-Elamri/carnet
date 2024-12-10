CREATE TABLE thread_vote (
    user_id   INT NOT NULL,
    thread_id   INT NOT NULL,
    vote      INT,
    voted_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, thread_id),
    FOREIGN KEY (user_id) REFERENCES user_data(user_id),
    FOREIGN KEY (thread_id) REFERENCES threads(thread_id)
);