CREATE TABLE IF NOT EXISTS cat_matches(
    id BIGINT PRIMARY KEY,
    issued_id BIGINT REFERENCES users(id),
    match_cat_id BIGINT REFERENCES cats(id),
    user_cat_id BIGINT REFERENCES cats(id),
    message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (issued_id) REFERENCES users(id),
    FOREIGN KEY (match_cat_id) REFERENCES cats(id),
    FOREIGN KEY (user_cat_id) REFERENCES cats(id)
);