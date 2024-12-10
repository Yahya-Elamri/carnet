CREATE TABLE IF NOT EXISTS car_listings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    agency_id INT NOT NULL,
    make VARCHAR(50) NOT NULL,
    model VARCHAR(50) NOT NULL,
    year INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    image_url VARCHAR(255),
    description TEXT,
    fuel_type VARCHAR(20), -- e.g., petrol, diesel, electric
    transmission VARCHAR(20), -- e.g., manual, automatic
    mileage INT, -- in kilometers or miles
    seats INT,
    color VARCHAR(30),
    status VARCHAR(20) NOT NULL DEFAULT 'available', -- e.g., available, rented, maintenance
    category TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (agency_id) REFERENCES agencies_data (id) ON DELETE CASCADE
);