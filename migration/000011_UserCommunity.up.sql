CREATE TABLE user_communities (
    user_id        INT NOT NULL,
    community_id   INT NOT NULL,
    role           ENUM('user', 'admin', 'owner', 'banned') DEFAULT 'user',
    joined_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, community_id),
    FOREIGN KEY (user_id) REFERENCES user_data(user_id),
    FOREIGN KEY (community_id) REFERENCES communities(community_id)
);