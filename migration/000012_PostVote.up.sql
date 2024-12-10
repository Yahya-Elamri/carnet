CREATE TABLE post_vote (
    user_id   INT NOT NULL,
    post_id   INT NOT NULL,
    vote      INT,
    voted_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, post_id),
    FOREIGN KEY (user_id) REFERENCES user_data(user_id),
    FOREIGN KEY (post_id) REFERENCES posts(post_id)
);