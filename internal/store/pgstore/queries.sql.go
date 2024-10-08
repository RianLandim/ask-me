// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package pgstore

import (
	"context"

	"github.com/google/uuid"
)

const getMessage = `-- name: GetMessage :one
SELECT "id", "room_id",  "message", "reaction_count", "answered"
FROM message 
WHERE id = $1
`

func (q *Queries) GetMessage(ctx context.Context, id uuid.UUID) (Message, error) {
	row := q.db.QueryRow(ctx, getMessage, id)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.RoomID,
		&i.Message,
		&i.ReactionCount,
		&i.Answered,
	)
	return i, err
}

const getRoom = `-- name: GetRoom :one
SELECT "id", "theme" FROM room WHERE id = $1
`

func (q *Queries) GetRoom(ctx context.Context, id uuid.UUID) (Room, error) {
	row := q.db.QueryRow(ctx, getRoom, id)
	var i Room
	err := row.Scan(&i.ID, &i.Theme)
	return i, err
}

const getRoomMessages = `-- name: GetRoomMessages :many
SELECT "id", "room_id", "message", "reaction_count", 'answered'
FROM message
WHERE room_id = $1
`

type GetRoomMessagesRow struct {
	ID            uuid.UUID
	RoomID        uuid.UUID
	Message       string
	ReactionCount int64
	Column5       string
}

func (q *Queries) GetRoomMessages(ctx context.Context, roomID uuid.UUID) ([]GetRoomMessagesRow, error) {
	rows, err := q.db.Query(ctx, getRoomMessages, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRoomMessagesRow
	for rows.Next() {
		var i GetRoomMessagesRow
		if err := rows.Scan(
			&i.ID,
			&i.RoomID,
			&i.Message,
			&i.ReactionCount,
			&i.Column5,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRooms = `-- name: GetRooms :many
SELECT "id", "theme" FROM room
`

func (q *Queries) GetRooms(ctx context.Context) ([]Room, error) {
	rows, err := q.db.Query(ctx, getRooms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Room
	for rows.Next() {
		var i Room
		if err := rows.Scan(&i.ID, &i.Theme); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertMessage = `-- name: InsertMessage :one
INSERT INTO message 
    ("room_id", "message") VALUES
    ($1, $2)
RETURNING "id"
`

type InsertMessageParams struct {
	RoomID  uuid.UUID
	Message string
}

func (q *Queries) InsertMessage(ctx context.Context, arg InsertMessageParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertMessage, arg.RoomID, arg.Message)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const insertRoom = `-- name: InsertRoom :one
INSERT INTO room ("theme") VALUES ($1) RETURNING "id"
`

func (q *Queries) InsertRoom(ctx context.Context, theme string) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertRoom, theme)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const markMessageAnswered = `-- name: MarkMessageAnswered :exec
UPDATE message
SET
    answered = true
WHERE
    id = $1
`

func (q *Queries) MarkMessageAnswered(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, markMessageAnswered, id)
	return err
}

const reactToMessage = `-- name: ReactToMessage :one
UPDATE message
SET 
    reaction_count = reaction_count + 1
WHERE
    id = $1
RETURNING "reaction_count"
`

func (q *Queries) ReactToMessage(ctx context.Context, id uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, reactToMessage, id)
	var reaction_count int64
	err := row.Scan(&reaction_count)
	return reaction_count, err
}

const removeReactionFromMessage = `-- name: RemoveReactionFromMessage :one
UPDATE message
SET 
    reaction_count = reaction_count - 1
WHERE
    id = $1
RETURNING "reaction_count"
`

func (q *Queries) RemoveReactionFromMessage(ctx context.Context, id uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, removeReactionFromMessage, id)
	var reaction_count int64
	err := row.Scan(&reaction_count)
	return reaction_count, err
}
