package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	Name    string  `json:"name"`
	Age     string  `json:"age"`
	Friends []*User `json:"friends"`
}

func main() {
	fmt.Println(help())
	fmt.Println(yourMetods())
	var req string
	for {
		fmt.Println("Ожидаю ваш запрос")
		_, er := fmt.Scan(&req)
		if er != nil {
			fmt.Println("Попробуй ещё.")
		}
		num := strings.LastIndex(req, "/")
		if num == 8 {
			// getFriends(req)
			continue
		}
		switch req {
		case "/create":
			create()
		case "/make_friends":
			make_friends()
		case "/user":
			deleteUs()
		case "/user_id":
			// user_id()
		case "/help":
			yourMetods()
		default:
			fmt.Println("Такой запрос сервер не обрабатывает")
		}
	}

}

func create() {
	var name string
	var age int
	for{
		fmt.Println("Введите имя пользователя")
		if _, er := fmt.Scan(&name); er != nil{
			fmt.Println("Что-то пошло не так, попробуй ещё.")
			continue
		}
		if _, er := strconv.Atoi(name); er == nil{
			fmt.Println("Сомневаюсь, что твоё имя состоит только из цифр.")
			continue
		}
		break 
	}
	for{
		fmt.Println("Введите возраст вашего пользователя")
		if _, er := fmt.Scan(&age); er != nil{
			fmt.Println("Укажи возраст числом.")
			continue
		}
		break 
	}

	us := &User{name, strconv.Itoa(age), []*User{}}
	// users := sevice{Store: "Петяб привет"}

	// Кодируем структуру User в JSON
	data, err := json.Marshal(us)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Создаем новый запрос
	req, err := http.NewRequest("POST", "http://localhost:8080/create", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	// Устанавливаем заголовок с типом данных в теле запроса
	req.Header.Set("Content-Type", "application/json")
	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	// Выводим ответ от сервера
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func make_friends() {
	fmt.Println("Введите ID пользователей, который просится в друзья")
	var sourse int
	if _, er := fmt.Scan(&sourse); er != nil {
		fmt.Println("Нужно вводить число.")
		return
	}
	fmt.Println("Введите ID пользователей, который принимает в друзья")
	var target int
	if _, er := fmt.Scan(&target); er != nil {
		fmt.Println("Нужно вводить число.")
		return
	}
	if sourse == target{
		fmt.Println("Не может пользователь добавить себе в друзья себя же")
		return
	}
	// Кодируем структуру в JSON
	// data := []byte(`{"source_id":"1","target_id":"2"}`)
	strac := map[string]int{"source_id":sourse,"target_id":target}
	data, err := json.Marshal(strac)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Создаем новый запрос
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/make_friends", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	// Устанавливаем заголовок с типом данных в теле запроса
	req.Header.Set("Content-Type", "application/json")
	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	// Выводим ответ от сервера
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

// DELETE
func deleteUs(){
	fmt.Println("Введите ID пользоватея, которого нужно удалить")
	var id int
	if _, er := fmt.Scan(&id); er != nil {
		fmt.Println("Нужно вводить число.")
		return
	}
	strac := map[string]int{"target_id":id}
	data, err := json.Marshal(strac)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Создаем новый запрос
	req, err := http.NewRequest("DELETE", "http://localhost:8080/user", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	// Устанавливаем заголовок с типом данных в теле запроса
	req.Header.Set("Content-Type", "application/json")
	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	// Выводим ответ от сервера
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func help() string {
	return "Эта программа позволяет добавлять пользователей в базу, подружить пользователей между собой, удалять его из базы,\n" +
		"посмотреть на друзей конкретного пользователя, обновить возраст пользователю.\n"
}
func yourMetods() string {
	return "Запросы, которыми вы можете воспользоваться:\n" +
		"/create - добавить пользователя\n" +
		"/make_friends - подружить двоих пользователей\n" +
		"/user - удаляет пользователя\n" +
		"/friends/user_id - получить друзей указанного пользователя\n" +
		"/user_id - обновляет возраст пользователя\n" +
		"/help - просмотреть доступные запросы\n"+
		"/get_all - посмотреть всех пользователей\n"
}
