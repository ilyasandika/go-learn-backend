CREATE TABLE user_profiles (
                               user_id INT NOT NULL,
                               full_name VARCHAR(255) NULL,
                               gender ENUM('male', 'female') NULL,
                               birthdate DATE NULL,
                               phone_number VARCHAR(255) NULL,
                               address TEXT NULL,
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
                               PRIMARY KEY (user_id),
                               CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
