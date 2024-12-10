CREATE TABLE IF NOT EXISTS threads (
    thread_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    community_id INT NOT NULL,
    tags TEXT,
    votes INT NOT NULL DEFAULT 0,
    content VARCHAR(1500) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status ENUM('active', 'inactive', 'banned') DEFAULT 'active', 
    FOREIGN KEY (user_id) REFERENCES user_data(user_id),
    FOREIGN KEY (community_id) REFERENCES communities(community_id)
);