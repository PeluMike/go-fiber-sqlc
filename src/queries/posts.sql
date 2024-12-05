-- 


-- name: CreatePost :one
INSERT INTO posts (post, user_id) VALUES ($1, $2)
RETURNING id, post, user_id, created_at::date, updated_at::date;

-- name: GetPostsByUserID :many
SELECT posts.id, posts.post, posts.user_id, to_char(posts.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at, to_char(posts.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at
FROM posts
LEFT JOIN 
  comments ON posts.id = comments.post_id
WHERE posts.user_id = $1; 