-- name: CreateSwipe :execresult
INSERT INTO Swipes (swipe_id, swiper_id, swipee_id, swipe_type, created_at) 
VALUES (?, ?, ?, ?, NOW());

-- name: CheckMatch :one
SELECT EXISTS (
    SELECT 1 FROM Swipes 
    WHERE swiper_id = ? AND swipee_id = ? AND swipe_type = 'YES'
) AS has_swiped_yes;