-- name: CreateRanking :execresult
INSERT INTO Rankings (user_id, like_count, dislike_count, attractiveness_score)
VALUES (?, 0, 0, 0.0);

-- name: GetRankingByUserId :one
SELECT * FROM Rankings WHERE user_id = ?;

-- name: UpdateRanking :execresult
UPDATE Rankings SET like_count = ?, dislike_count = ?, attractiveness_score = ? WHERE user_id = ?;