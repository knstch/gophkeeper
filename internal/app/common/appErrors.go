package common

import "errors"

const (
	ErrLength = "значение не может быть больше"
)

var (
	ErrIntegrityViolation = errors.New(`ERROR: duplicate key value violates unique constraint "credentials_pk" (SQLSTATE 23505)`)
	ErrUserNotFound       = errors.New(`record not found`)
	ErrBadEmail           = errors.New("введите email в формате example@example.com")
	ErrBadPass            = errors.New("в пароле должна быть как минимум одна заглавная буква, цифра и длина от 8 символов")
	ErrNotLoggedIn        = errors.New("зарегестрируйтесь в сервисе или войдите в свой аккаунт")
	ErrInvalidToken       = errors.New("токен невалиден")
	ErrNoDataWereFound    = errors.New("данные не найдены")
	ErrFieldIsEmpty       = errors.New("значение не может быть пустым")
	ErrTextDouble         = errors.New("этот заголовок уже использован")
	ErrBadCardDate        = errors.New("неверно указана дата на карте")
	ErrFileSize           = errors.New("размер файла не может быть больше 5 мб")
	ErrNotInt             = errors.New("переданное значение не число")
	ErrWrongCvv           = errors.New("неверный cvv")
)
