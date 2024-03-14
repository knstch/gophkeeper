package common

import "errors"

const (
	ErrIntegrityViolation = `ERROR: duplicate key value violates unique constraint "credentials_pk" (SQLSTATE 23505)`
	ErrUserNotFound       = `record not found`
	ErrFieldIsEmpty       = "Значение не может быть пустым"
)

var (
	ErrBadEmail = errors.New("введите email в формате example@example.com")
	ErrBadPass  = errors.New("в пароле должна быть как минимум одна заглавная буква, цифра и длина от 8 символов")
)
