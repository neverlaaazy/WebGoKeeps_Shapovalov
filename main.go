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
				<h2>Добро пожаловать в <br><b>4KEEPS</b>!</h2>
				<b>Возможности:</b>
				<ul>
					<li>Просмотр;</li>
					<li>Добавление;</li>
					<li>Удаление.</li>
				</ul>
				<a class="conBtn" href="/keeps">ВХОД</a>
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
			<div class="container conKeeps">
				<a class="conBtn" href="/showkeeps">Посмотреть</a>
				<a class="conBtn" href="/addkeep">Добавить</a>
				<a class="conBtnExit" href="/">ВЫХОД</a>
			</div>
		</body>
		</html>
	`)
}
func addKeepHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Добавить закладку</title>
			<link rel="stylesheet" href="/style.css">
		</head>
		<body>
			<div class="container">
				<form method="POST">
					<label for="keep">Текст заметки</label>
					<textarea id="keep" class="formInputText" name="keep"></textarea>
					<div class="conDivFormBtns">
						<input class="formBtn" type="submit" value="Добавить">
						<a class="conBtnExit" href="/keeps">Отмена</a>
					</div>
				</form>
			</div>
		</body>
		</html>
	`)
}
func showKeepsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Показать закладки</title>
			<link rel="stylesheet" href="/style.css">
		</head>
		<body>
			<div class="container">
				<a class="conBtnExit" href="/keeps">Назад</a>
			</div>
		</body>
		</html>
	`)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/keeps", keepsHandler)
	http.HandleFunc("/addkeep", addKeepHandler)
	http.HandleFunc("/showkeeps", showKeepsHandler)

	http.Handle("/style.css", http.FileServer(http.Dir(".")))

	fmt.Println("Сервер работает на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
