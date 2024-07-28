
-- name: CreateUser :one
INSERT into users(id ,created_at,updated_at,name,api_key)
values ($1,$2,$3,$4,
        encode(sha256(random()::text::bytea),'hex')
       )
returning *;

-- name: GetUserByApiKey :one
SELECT  * from users where api_key = $1;