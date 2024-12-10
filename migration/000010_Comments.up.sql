CREATE TABLE comments (
    comment_id INT AUTO_INCREMENT PRIMARY KEY,
    thread_id INT REFERENCES threads(thread_id) ON DELETE CASCADE,  -- comment on a thread (optional)
    post_id INT REFERENCES posts(post_id) ON DELETE CASCADE,      -- comment on a post (optional)
    parent_comment_id INT REFERENCES comments(comment_id) ON DELETE CASCADE,  -- self-referencing, allows for replies to comments
    user_id INT REFERENCES user_data(user_id),  -- creator of the comment
    content TEXT,
    votes INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);