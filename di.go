package main

import (
	"errors"
	"fmt"
	"net/http"
)

// Logger интерфейс для логирования.
type Logger interface {
	Log(message string)
}

// LoggerAdapter функциональный тип, реализующий Logger.
type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// LogOutput реализация логирования.
func LogOutput(message string) {
	fmt.Println(message)
}

// DataStore интерфейс для работы с хранилищем данных.
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// SimpleDataStore простая реализация хранилища данных.
type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// NewSimpleDataStore создает экземпляр SimpleDataStore.
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

// Logic интерфейс бизнес-логики.
type Logic interface {
	SayHello(userID string) (string, error)
}

// SimpleLogic реализация бизнес-логики.
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

// NewSimpleLogic создает экземпляр SimpleLogic.
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

// Controller структура для обработки HTTP-запросов.
type Controller struct {
	l     Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

// NewController создает экземпляр Controller.
func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

func main() {
	// Настраиваем зависимости.
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)

	// Настраиваем HTTP-роут.
	http.HandleFunc("/hello", c.SayHello)

	// Запускаем сервер.
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
