-- name: GetRoom :one
SELECT "id", "theme" FROM room WHERE id = $1;

-- name: GetRooms :many
SELECT "id", "theme" FROM room;

-- name: InsertRoom :one
INSERT INTO room ("theme") VALUES ($1) RETURNING "id";

-- name: GetMessage :one
SELECT "id", "room_id",  "message", "reaction_count", "answered"
FROM message 
WHERE id = $1;

-- name: GetRoomMessages :many
SELECT "id", "room_id", "message", "reaction_count", 'answered'
FROM message
WHERE room_id = $1;

-- name: InsertMessage :one
INSERT INTO message 
    ("room_id", "message") VALUES
    ($1, $2)
RETURNING "id";

-- name: ReactToMessage :one
UPDATE message
SET 
    reaction_count = reaction_count + 1
WHERE
    id = $1
RETURNING "reaction_count";

-- name: RemoveReactionFromMessage :one
UPDATE message
SET 
    reaction_count = reaction_count - 1
WHERE
    id = $1
RETURNING "reaction_count";

-- name: MarkMessageAnswered :exec
UPDATE message
SET
    answered = true
WHERE
    id = $1;

