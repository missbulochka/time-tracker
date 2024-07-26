package storage

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)

const (
	CreateUser = `-- name: CreateUser :one
	INSERT INTO users
	(
		passport_number,
		surname,
		name,
		patronymic,
		adress
	) VALUES (
		$1, $2, $3, $4, $5
	)`

	DeleteUser = `-- name: DeleteUser :one
	DELTE FROM users WHERE user_id = $1
	`
	// TODO: GetUsersInfo

	// TODO: UpdateUser

	// TODO: GetUser

	CreateTimer = `-- name: CreateTimer :one
	INSERT INTO user_tasks
	(
		user_id,
		task_id
	) VALUES (
		$1, $2
	)`

	StopTimer = `-- name: StopTimer :one
	UPDATE user_tasks
	SET end_time = LOCALTIMESTAMP
	WHERE users_tasks_pk = $1`
)
