package main

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/knstch/gophkeeper/cmd/config"
	"github.com/knstch/gophkeeper/internal/app/approuter"
	"github.com/knstch/gophkeeper/internal/app/common"
	"github.com/knstch/gophkeeper/internal/app/handler"
	"github.com/knstch/gophkeeper/internal/app/storage/psql"
	"github.com/stretchr/testify/assert"
)

func emailGenerator() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 10)
	for i := 0; i < 10; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result) + "@gmail.com"
}

type testUser struct {
	email    string
	password string
	cookie   []*http.Cookie
}

type want struct {
	statusCode  int
	contentType string
	body        string
	extraBody   string
}

type request struct {
	contentType string
	body        string
	extraBody   string
	cookie      []*http.Cookie
}

type tests struct {
	name   string
	want   want
	reqest request
}

var (
	testUserCorrectOne = testUser{
		email:    emailGenerator(),
		password: "Xer@0101",
	}
	testUserCorrectTwo = testUser{
		email:    emailGenerator(),
		password: "1234@!aBoba",
	}
	testUserCorrectThree = testUser{
		email:    emailGenerator(),
		password: "piskaLol@228",
	}
)

func TestSignUp(t *testing.T) {
	tests := []tests{
		{
			name: "#1 sign up normal acc",
			want: want{
				statusCode:  200,
				contentType: "application/json",
				body:        `{"message":"вы успешно зарегестрировались!"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "` + testUserCorrectOne.email + `","password": "` + testUserCorrectOne.password + `"}`,
			},
		},
		{
			name: "#2 sign up normal acc",
			want: want{
				statusCode:  200,
				contentType: "application/json",
				body:        `{"message":"вы успешно зарегестрировались!"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "` + testUserCorrectTwo.email + `","password": "` + testUserCorrectTwo.password + `"}`,
			},
		},
		{
			name: "#3 sign up normal acc",
			want: want{
				statusCode:  200,
				contentType: "application/json",
				body:        `{"message":"вы успешно зарегестрировались!"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "` + testUserCorrectThree.email + `","password": "` + testUserCorrectThree.password + `"}`,
			},
		},
		{
			name: "#4 sign up no email acc",
			want: want{
				statusCode:  400,
				contentType: "application/json",
				body:        `{"error":"email: значение не может быть пустым."}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "","password": "asdadsW23@"}`,
			},
		},
		{
			name: "#5 sign up no password acc",
			want: want{
				statusCode:  400,
				contentType: "application/json",
				body:        `{"error":"password: значение не может быть пустым."}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "` + emailGenerator() + `","password": ""}`,
			},
		},
		{
			name: "#6 sign up bad email one",
			want: want{
				statusCode:  400,
				contentType: "application/json",
				body:        `{"error":"email: введите email в формате example@example.com."}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "emailgmail.com","password": "Xer@0101"}`,
			},
		},
		{
			name: "#7 sign up bad email two",
			want: want{
				statusCode:  400,
				contentType: "application/json",
				body:        `{"error":"email: введите email в формате example@example.com."}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "email@gmail","password": "Xer@0101"}`,
			},
		},
		{
			name: "#8 sign up bad email three",
			want: want{
				statusCode:  400,
				contentType: "application/json",
				body:        `{"error":"email: введите email в формате example@example.com."}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "@gmail.com","password": "Xer@0101"}`,
			},
		},
		{
			name: "#9 sign up bad password one",
			want: want{
				statusCode:  400,
				contentType: "application/json",
				body:        `{"error":"password: в пароле должна быть как минимум одна заглавная буква, цифра и длина от 8 символов."}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "` + emailGenerator() + `","password": "asd"}`,
			},
		},
		{
			name: "#10 sign up bad password two",
			want: want{
				statusCode:  400,
				contentType: "application/json",
				body:        `{"error":"password: в пароле должна быть как минимум одна заглавная буква, цифра и длина от 8 символов."}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "` + emailGenerator() + `","password": "12345678"}`,
			},
		},
		{
			name: "#11 sign up bad password three",
			want: want{
				statusCode:  400,
				contentType: "application/json",
				body:        `{"error":"password: в пароле должна быть как минимум одна заглавная буква, цифра и длина от 8 символов."}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "` + emailGenerator() + `","password": "qwertyui"}`,
			},
		},
		{
			name: "#12 sign up bad password four",
			want: want{
				statusCode:  400,
				contentType: "application/json",
				body:        `{"error":"password: в пароле должна быть как минимум одна заглавная буква, цифра и длина от 8 символов."}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "` + emailGenerator() + `","password": "@!#$%!@#$%"}`,
			},
		},
		{
			name: "#13 sign up normal acc",
			want: want{
				statusCode:  400,
				contentType: "application/json",
				body:        `{"error":"email: значение не может быть пустым; password: значение не может быть пустым."}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "","password": ""}`,
			},
		},
		{
			name: "#14 sign up repeat good acc",
			want: want{
				statusCode:  409,
				contentType: "application/json",
				body:        `{"error":"эта почта уже занята"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "` + testUserCorrectOne.email + `","password": "` + testUserCorrectOne.password + `"}`,
			},
		},
	}

	config.ParseConfig()
	psqlStorage, err := psql.NewPsqlStorage(config.ReadyConfig.DSN)
	if err != nil {
		log.Error(err)
	}

	handlers := handler.NewHandler(psqlStorage)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	approuter.InitRouter(app, handlers, psqlStorage)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer([]byte(tt.reqest.body)))

			rr, err := app.Test(req, 100)
			if err != nil {
				log.Error(err)
			}

			body, err := io.ReadAll(rr.Body)
			if err != nil {
				log.Error(err)
			}
			switch tt.name {
			case "#1 sign up normal acc":
				testUserCorrectOne.cookie = rr.Cookies()
				assert.Equal(t, 1, len(testUserCorrectOne.cookie))
			case "#2 sign up normal acc":
				testUserCorrectTwo.cookie = rr.Cookies()
				assert.Equal(t, 1, len(testUserCorrectOne.cookie))
			case "#3 sign up normal acc":
				testUserCorrectThree.cookie = rr.Cookies()
				assert.Equal(t, 1, len(testUserCorrectOne.cookie))
			}
			assert.Equal(t, tt.want.statusCode, rr.StatusCode)
			assert.Equal(t, tt.want.contentType, rr.Header.Get("Content-Type"))
			assert.JSONEq(t, tt.want.body, string(body))
		})
	}
}

func TestAuth(t *testing.T) {
	tests := []tests{
		{
			name: "#1 sign in normal acc",
			want: want{
				statusCode:  200,
				contentType: "application/json",
				body:        `{"message":"вы успешно залогинились!"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "` + testUserCorrectOne.email + `","password": "` + testUserCorrectOne.password + `"}`,
			},
		},
		{
			name: "#2 sign up normal acc",
			want: want{
				statusCode:  200,
				contentType: "application/json",
				body:        `{"message":"вы успешно залогинились!"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "` + testUserCorrectTwo.email + `","password": "` + testUserCorrectTwo.password + `"}`,
			},
		},
		{
			name: "#3 sign in bad pass",
			want: want{
				statusCode:  404,
				contentType: "application/json",
				body:        `{"message":"неверная почта или пароль"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "` + testUserCorrectOne.email + `","password": "fasdfas"}`,
			},
		},
		{
			name: "#4 sign up bad email",
			want: want{
				statusCode:  404,
				contentType: "application/json",
				body:        `{"message":"неверная почта или пароль"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "dsadad","password": "` + testUserCorrectTwo.password + `"}`,
			},
		},
		{
			name: "#5 sign up bad email and pass",
			want: want{
				statusCode:  404,
				contentType: "application/json",
				body:        `{"message":"неверная почта или пароль"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "asdasdad","password": "asdasda"}`,
			},
		},
		{
			name: "#6 sign up empty email",
			want: want{
				statusCode:  404,
				contentType: "application/json",
				body:        `{"message":"неверная почта или пароль"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "","password": "` + testUserCorrectTwo.password + `"}`,
			},
		},
		{
			name: "#7 sign up empty pass",
			want: want{
				statusCode:  404,
				contentType: "application/json",
				body:        `{"message":"неверная почта или пароль"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"email": "` + testUserCorrectOne.email + `","password": ""}`,
			},
		},
	}
	psqlStorage, err := psql.NewPsqlStorage(config.ReadyConfig.DSN)
	if err != nil {
		log.Error(err)
	}

	handlers := handler.NewHandler(psqlStorage)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	approuter.InitRouter(app, handlers, psqlStorage)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/auth", bytes.NewBuffer([]byte(tt.reqest.body)))

			rr, err := app.Test(req, 100)
			if err != nil {
				log.Error(err)
			}

			body, err := io.ReadAll(rr.Body)
			if err != nil {
				log.Error(err)
			}
			switch tt.name {
			case "#1 sign in normal acc":
				testUserCorrectOne.cookie = rr.Cookies()
				assert.Equal(t, 1, len(testUserCorrectOne.cookie))
			case "#2 sign in normal acc":
				testUserCorrectTwo.cookie = rr.Cookies()
				assert.Equal(t, 1, len(testUserCorrectOne.cookie))
			}
			assert.Equal(t, tt.want.statusCode, rr.StatusCode)
			assert.Equal(t, tt.want.contentType, rr.Header.Get("Content-Type"))
			assert.JSONEq(t, tt.want.body, string(body))
		})
	}
}

var secretsFromUsers map[string]string = map[string]string{
	"userOneSecretServiceOne": "userOneTest1",
	"userTwoSecretServiceOne": "userTwoTest1",
	"userTwoSecretServiceTwo": "userTwoTest2",
}

func TestStorePrivate(t *testing.T) {
	tests := []tests{
		{
			name: "#1 post good secret user one",
			want: want{
				statusCode:  201,
				contentType: "application/json",
				body:        `{"message": "данны успешно сохранены!"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"service": "` + secretsFromUsers["userOneSecretServiceOne"] + `","login": "test","password": "test","metadata": "MET"}`,
				cookie:      testUserCorrectOne.cookie,
			},
		},
		{
			name: "#2 post good secret user two",
			want: want{
				statusCode:  201,
				contentType: "application/json",
				body:        `{"message": "данны успешно сохранены!"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"service": "` + secretsFromUsers["userTwoSecretServiceOne"] + `","login": "test","password": "test","metadata": "MET"}`,
				cookie:      testUserCorrectTwo.cookie,
			},
		},
		{
			name: "#3 post secret without login",
			want: want{
				statusCode:  403,
				contentType: "application/json",
				body:        `{"message": "необходимо зарегестрироваться или авторизоваться"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"service": "huypizda1","login": "pizda","password": "aboba","metadata": "govno zhopa"}`,
			},
		},
		{
			name: "#4 post bad secret user one service empty",
			want: want{
				statusCode:  400,
				contentType: "application/json",
				body:        `{"error": "поле не может быть пустым"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"service": "","login": "pizda","password": "aboba","metadata": "govno zhopa"}`,
				cookie:      testUserCorrectTwo.cookie,
			},
		},
		{
			name: "#5 post bad secret user one login empty",
			want: want{
				statusCode:  400,
				contentType: "application/json",
				body:        `{"error": "поле не может быть пустым"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"service": "huypizda1","login": "","password": "aboba","metadata": "govno zhopa"}`,
				cookie:      testUserCorrectTwo.cookie,
			},
		},
		{
			name: "#6 post bad secret user one password empty",
			want: want{
				statusCode:  400,
				contentType: "application/json",
				body:        `{"error": "поле не может быть пустым"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"service": "huypizda1","login": "pizda","password": "","metadata": "govno zhopa"}`,
				cookie:      testUserCorrectTwo.cookie,
			},
		},
		{
			name: "#7 post good secret user two metadata empty",
			want: want{
				statusCode:  201,
				contentType: "application/json",
				body:        `{"message": "данны успешно сохранены!"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"service": "` + secretsFromUsers["userTwoSecretServiceTwo"] + `","login": "pizda","password": "aboba","metadata": ""}`,
				cookie:      testUserCorrectTwo.cookie,
			},
		},
	}
	psqlStorage, err := psql.NewPsqlStorage(config.ReadyConfig.DSN)
	if err != nil {
		log.Error(err)
	}

	handlers := handler.NewHandler(psqlStorage)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	approuter.InitRouter(app, handlers, psqlStorage)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/secret", bytes.NewBuffer([]byte(tt.reqest.body)))
			if tt.name != "#3 post secret without login" {
				req.AddCookie(tt.reqest.cookie[0])
			}

			rr, err := app.Test(req, 100)
			if err != nil {
				log.Error(err)
			}

			body, err := io.ReadAll(rr.Body)
			if err != nil {
				log.Error(err)
			}

			assert.Equal(t, tt.want.statusCode, rr.StatusCode)
			assert.Equal(t, tt.want.contentType, rr.Header.Get("Content-Type"))
			assert.JSONEq(t, tt.want.body, string(body))
		})
	}
}

func TestGetAllPrivate(t *testing.T) {
	tests := []tests{
		{
			name: "#1 get all secrets user one",
			want: want{
				statusCode:  200,
				contentType: "application/json",
				body:        "1",
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				cookie:      testUserCorrectOne.cookie,
			},
		},
		{
			name: "#2 get all secrets user two",
			want: want{
				statusCode:  200,
				contentType: "application/json",
				body:        "2",
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				cookie:      testUserCorrectTwo.cookie,
			},
		},
		{
			name: "#3 get all secrets without login",
			want: want{
				statusCode:  403,
				contentType: "application/json",
				body:        `{"message": "необходимо зарегестрироваться или авторизоваться"}`,
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				body:        `{"message": "необходимо зарегестрироваться или авторизоваться"}`,
			},
		},
		{
			name: "#4 get all secrets user three",
			want: want{
				statusCode:  204,
				contentType: "application/json",
				body:        "0",
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				cookie:      testUserCorrectThree.cookie,
			},
		},
	}
	psqlStorage, err := psql.NewPsqlStorage(config.ReadyConfig.DSN)
	if err != nil {
		log.Error(err)
	}

	handlers := handler.NewHandler(psqlStorage)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	approuter.InitRouter(app, handlers, psqlStorage)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/secret", bytes.NewBuffer([]byte(tt.reqest.body)))
			if tt.name != "#3 get all secrets without login" {
				req.AddCookie(tt.reqest.cookie[0])
			}

			rr, err := app.Test(req, 100)
			if err != nil {
				log.Error(err)
			}

			body, err := io.ReadAll(rr.Body)
			if err != nil {
				log.Error(err)
			}

			var respDataToObj = common.AllSecrets{}
			if tt.name != "#3 get all secrets without login" {
				if err := json.Unmarshal(body, &respDataToObj); err != nil {
					log.Error(err)
				}
			}

			lenOfSecrets := strconv.Itoa(len(respDataToObj.Secrets))

			assert.Equal(t, tt.want.statusCode, rr.StatusCode)
			assert.Equal(t, tt.want.contentType, rr.Header.Get("Content-Type"))
			if tt.name != "#3 get all secrets without login" {
				assert.Equal(t, tt.want.body, lenOfSecrets)
				return
			}
			assert.JSONEq(t, tt.want.body, tt.reqest.body)
		})
	}
}

func TestGetPrivateSecret(t *testing.T) {
	tests := []tests{
		{
			name: "#1 get good secret user one",
			want: want{
				statusCode:  200,
				contentType: "application/json",
				body:        secretsFromUsers["userOneSecretServiceOne"],
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				cookie:      testUserCorrectOne.cookie,
				body:        secretsFromUsers["userOneSecretServiceOne"],
			},
		},
		{
			name: "#2 get wrong secret user one",
			want: want{
				statusCode:  204,
				contentType: "application/json",
				body:        "0",
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				cookie:      testUserCorrectOne.cookie,
				body:        secretsFromUsers["userTwoSecretServiceOne"],
			},
		},
		{
			name: "#3 get good secret user two",
			want: want{
				statusCode:  200,
				contentType: "application/json",
				body:        secretsFromUsers["userTwoSecretServiceOne"],
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				cookie:      testUserCorrectTwo.cookie,
				body:        secretsFromUsers["userTwoSecretServiceOne"],
			},
		},
	}

	psqlStorage, err := psql.NewPsqlStorage(config.ReadyConfig.DSN)
	if err != nil {
		log.Error(err)
	}

	handlers := handler.NewHandler(psqlStorage)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	approuter.InitRouter(app, handlers, psqlStorage)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", `/secret/`+tt.reqest.body+``, bytes.NewBuffer([]byte(tt.reqest.body)))
			req.AddCookie(tt.reqest.cookie[0])

			rr, err := app.Test(req, 100)
			if err != nil {
				log.Error(err)
			}

			body, err := io.ReadAll(rr.Body)
			if err != nil {
				log.Error(err)
			}

			var respDataToObj common.AllSecrets
			if err := json.Unmarshal(body, &respDataToObj); err != nil {
				log.Error(err)
			}

			assert.Equal(t, tt.want.statusCode, rr.StatusCode)
			assert.Equal(t, tt.want.contentType, rr.Header.Get("Content-Type"))
			if len(respDataToObj.Secrets) != 0 {
				assert.Equal(t, tt.want.body, respDataToObj.Secrets[0].Service)
			}
		})
	}
}

var secretsFromUsersToPut map[string]string = map[string]string{
	"userOneSecretServiceOne": "userOneTestPut",
	"userTwoSecretServiceOne": "userTwoTestPut",
	"userTwoSecretServiceTwo": "userTwoTestPut",
}

func TestPutPrivateSecret(t *testing.T) {
	tests := []tests{
		{
			name: "#1 put good secret user one",
			want: want{
				statusCode:  200,
				contentType: "application/json",
				body:        secretsFromUsers["userOneSecretServiceOne"],
				extraBody:   secretsFromUsersToPut["userOneSecretServiceOne"],
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				cookie:      testUserCorrectOne.cookie,
				body:        secretsFromUsers["userOneSecretServiceOne"],
				extraBody:   secretsFromUsersToPut["userOneSecretServiceOne"],
			},
		},
		{
			name: "#2 put good secret user two",
			want: want{
				statusCode:  200,
				contentType: "application/json",
				body:        secretsFromUsers["userTwoSecretServiceOne"],
				extraBody:   secretsFromUsersToPut["userTwoSecretServiceOne"],
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				cookie:      testUserCorrectTwo.cookie,
				body:        secretsFromUsers["userTwoSecretServiceTwo"],
				extraBody:   secretsFromUsersToPut["userTwoSecretServiceTwo"],
			},
		},
	}

	psqlStorage, err := psql.NewPsqlStorage(config.ReadyConfig.DSN)
	if err != nil {
		log.Error(err)
	}

	handlers := handler.NewHandler(psqlStorage)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	approuter.InitRouter(app, handlers, psqlStorage)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", `/secret/`+tt.reqest.body+``, nil)
			req.AddCookie(tt.reqest.cookie[0])

			rr, err := app.Test(req, 100)
			if err != nil {
				log.Error(err)
			}

			body, err := io.ReadAll(rr.Body)
			if err != nil {
				log.Error(err)
			}

			var respDataToObj common.AllSecrets
			if err := json.Unmarshal(body, &respDataToObj); err != nil {
				log.Error(err)
			}

			respDataToObj.Secrets[0].Service = tt.reqest.extraBody
			preparedObjToChange, err := json.Marshal(respDataToObj.Secrets[0])
			if err != nil {
				log.Error(err)
			}

			req = httptest.NewRequest("PUT", `/secret/`, bytes.NewBuffer([]byte(preparedObjToChange)))
			req.AddCookie(tt.reqest.cookie[0])

			_, err = app.Test(req, 100)
			if err != nil {
				log.Error(err)
			}

			req = httptest.NewRequest("GET", `/secret/`+tt.reqest.extraBody+``, nil)
			req.AddCookie(tt.reqest.cookie[0])

			_, err = app.Test(req, 100)
			if err != nil {
				log.Error(err)
			}

			assert.Equal(t, tt.want.statusCode, rr.StatusCode)
			assert.Equal(t, tt.want.contentType, rr.Header.Get("Content-Type"))
			if len(respDataToObj.Secrets) != 0 {
				assert.Equal(t, tt.want.extraBody, respDataToObj.Secrets[0].Service)
			}
		})
	}
}

func TestDeletePrivateSecret(t *testing.T) {
	tests := []tests{
		{
			name: "#1 delete good secret user one",
			want: want{
				statusCode:  200,
				contentType: "application/json",
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				cookie:      testUserCorrectOne.cookie,
				body:        secretsFromUsersToPut["userOneSecretServiceOne"],
			},
		},
		{
			name: "#2 delete good secret user two",
			want: want{
				statusCode:  200,
				contentType: "application/json",
			},
			reqest: request{
				contentType: "application/json; charset=utf-8",
				cookie:      testUserCorrectTwo.cookie,
				body:        secretsFromUsersToPut["userTwoSecretServiceTwo"],
			},
		},
	}

	psqlStorage, err := psql.NewPsqlStorage(config.ReadyConfig.DSN)
	if err != nil {
		log.Error(err)
	}

	handlers := handler.NewHandler(psqlStorage)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	approuter.InitRouter(app, handlers, psqlStorage)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", `/secret/`+tt.reqest.body+``, nil)
			req.AddCookie(tt.reqest.cookie[0])

			rr, err := app.Test(req, 100)
			if err != nil {
				log.Error(err)
			}

			body, err := io.ReadAll(rr.Body)
			if err != nil {
				log.Error(err)
			}

			var respDataToObj common.AllSecrets
			if err := json.Unmarshal(body, &respDataToObj); err != nil {
				log.Error(err)
			}

			req = httptest.NewRequest("DELETE", `/secret/`+respDataToObj.Secrets[0].Uuid+``, nil)
			req.AddCookie(tt.reqest.cookie[0])

			_, err = app.Test(req, 100)
			if err != nil {
				log.Error(err)
			}

			req = httptest.NewRequest("GET", `/secret/`+tt.reqest.body+``, nil)
			req.AddCookie(tt.reqest.cookie[0])

			rr, err = app.Test(req, 100)
			if err != nil {
				log.Error(err)
			}

			body, err = io.ReadAll(rr.Body)
			if err != nil {
				log.Error(err)
			}
			var newRespDataToObj common.AllSecrets
			if err := json.Unmarshal(body, &newRespDataToObj); err != nil {
				log.Error(err)
			}

			assert.Equal(t, 0, len(newRespDataToObj.Secrets))
			assert.Equal(t, 204, rr.StatusCode)
		})
	}
}
