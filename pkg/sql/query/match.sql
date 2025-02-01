-- name: CreateMatch :execresult
INSERT INTO Matches (match_id, user1_id, user2_id, created_at)
VALUES (?, ?, ?, NOW());

-- name: GetMatches :many
SELECT *
FROM Matches AS m
WHERE m.user1_id = ? OR m.user2_id = ?;