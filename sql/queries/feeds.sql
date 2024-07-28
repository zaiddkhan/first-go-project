
-- name: CreateFeed :one
insert into  feeds (
    id,created_at ,updated_at , name ,url, user_id
 ) VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetNextFeedsToFetch :many
SELECT * from feeds
ORDER BY last_fetched_at asc nulls first
limit $1;

-- name: MakeFeedAsFetched :one
UPDATE feeds
set last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;
