CREATE TABLE IF NOT EXISTS todos (
      id BIGINT PRIMARY KEY AUTO_INCREMENT,
      activity_group_id BIGINT,
      title VARCHAR(255) NOT NULL,
      priority VARCHAR(255) NOT NULL,
      is_active BOOLEAN DEFAULT false NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)