-- 


-- name: CreatePostComment :one
INSERT INTO comments (post_id, user_id, comment) VALUES ($1, $2, $3)
RETURNING id::text, comment, user_id, post_id, to_char(created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at;

-- name: GetPostComment :many
SELECT id, post_id, user_id, comment, 
       to_char(created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at
FROM comments
WHERE post_id = $1;
