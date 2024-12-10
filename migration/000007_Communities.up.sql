CREATE TABLE communities (
    community_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    communities_media VARCHAR(255) DEFAULT '/communities/Default-communities-pic.jpg',
    communities_banner VARCHAR(255) DEFAULT '/communities/Default-communities-banner.jpg',
    total_joined INT NOT NULL DEFAULT 1,
    FOREIGN KEY (user_id) REFERENCES user_data(user_id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);