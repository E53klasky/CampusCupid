CREATE TABLE Swipes (
    swipe_id VARCHAR(255) PRIMARY KEY,
    swiper_id VARCHAR(255) NOT NULL,
    swipee_id VARCHAR(255) NOT NULL,
    swipe_type ENUM('YES', 'NO') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (swiper_id) REFERENCES Users(user_id),
    FOREIGN KEY (swipee_id) REFERENCES Users(user_id),
    UNIQUE (swiper_id, swipee_id)
);