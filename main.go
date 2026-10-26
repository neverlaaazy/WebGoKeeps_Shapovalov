package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Главная</title>
			<link rel="stylesheet" href="/style.css">
		</head>
		<body>
			<div class="container">
				<h2>Добро пожаловать в <br><b>ЗАКЛАДКИ</b>!</h2>
				<b>Возможности:</b>
				<ul>
					<li>Просмотр;</li>
					<li>Добавление;</li>
					<li>Удаление.</li>
				</ul>
				<a href="/keeps">ВХОД</a>
			</div>
		</body>
		</html>
	`)
}
func keepsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Закладки</title>
			<link rel="stylesheet" href="/style.css">
		</head>
		<body>
			<div class="container">
				<a href="/">ВЫХОД</a>
			</div>
		</body>
		</html>
	`)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/keeps", keepsHandler)

	http.Handle("/style.css", http.FileServer(http.Dir(".")))

	fmt.Println("Сервер работает на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
