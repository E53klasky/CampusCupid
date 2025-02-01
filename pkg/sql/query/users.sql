-- name: CreateUser :execresult
INSERT INTO Users (user_id, email, password, name, gender, age, latitude, longitude, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW()); 

-- name: GetUserByEmail :one
SELECT *
FROM Users
WHERE email = ?;

-- name: GetUserById :one
SELECT *
FROM Users
WHERE user_id = ?;

-- name: DiscoverUsersV1 :many
SELECT user_id, name, gender, age
FROM Users
WHERE user_id NOT IN (
    SELECT swipee_id
    FROM Swipes
    WHERE swiper_id = ?
)
AND user_id != ?;

-- name: DiscoverUsersV2 :many
SELECT u.user_id, u.name, u.gender, u.age, u.latitude, u.longitude, r.attractiveness_score
FROM Users u
LEFT JOIN Rankings r ON u.user_id = r.user_id
WHERE u.user_id NOT IN (
    SELECT swipee_id
    FROM Swipes
    WHERE swiper_id = ?
) 
AND u.user_id != ?
AND u.age >= COALESCE(sqlc.narg(greater_than_age), u.age)
AND u.age <= COALESCE(sqlc.narg(lower_than_age), u.age)
AND u.gender = COALESCE(sqlc.narg(gender), u.gender)
ORDER BY r.attractiveness_score DESC;