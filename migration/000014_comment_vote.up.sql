CREATE TABLE comment_vote (
    user_id   INT NOT NULL,
    comment_id   INT NOT NULL,
    vote      INT,
    voted_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, comment_id),
    FOREIGN KEY (user_id) REFERENCES user_data(user_id),
    FOREIGN KEY (comment_id) REFERENCES comments(comment_id)
);