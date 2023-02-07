CREATE TABLE IF NOT EXISTS activities (
      id BIGINT PRIMARY KEY AUTO_INCREMENT,
      title VARCHAR(255) NOT NULL,
      email VARCHAR(255) NOT NULL,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)