CREATE TABLE IF NOT EXISTS user_data (
    user_id INT AUTO_INCREMENT PRIMARY KEY,         
    username VARCHAR(50) NOT NULL UNIQUE,           
    password_hash VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    address TEXT NOT NULL,             
    email VARCHAR(100) NOT NULL UNIQUE,              
    first_name VARCHAR(50) NOT NULL,                     
    last_name VARCHAR(50) NOT NULL,
    description VARCHAR(255) NOT NULL,
    social_links VARCHAR(255) NOT NULL,   
    profile_url VARCHAR(255) DEFAULT '/profile/Default-profile-pic.jpg',                                                
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    status ENUM('active', 'inactive', 'banned') DEFAULT 'active', 
    CONSTRAINT email_format CHECK (email LIKE '%_@__%.__%')
);