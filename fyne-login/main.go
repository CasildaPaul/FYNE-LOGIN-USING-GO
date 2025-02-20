package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func Authenticate(username, password string) bool {
	return username == "casi" && password == "Lavender"
}

func ShowLoginWindow(myApp fyne.App) {
	myWindow := myApp.NewWindow("Login Page")
	myWindow.Resize(fyne.NewSize(300, 200))

	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Enter username...")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter password...")

	loginButton := widget.NewButton("Login", func() {
		username := usernameEntry.Text
		password := passwordEntry.Text

		if Authenticate(username, password) {
			dialog.ShowInformation("Login Successful", "Welcome, "+username, myWindow)
		} else {
			dialog.ShowError(fmt.Errorf("Invalid username or password"), myWindow)
		}
	})

	content := container.NewVBox(
		widget.NewLabel("Username:"),
		usernameEntry,
		widget.NewLabel("Password:"),
		passwordEntry,
		loginButton,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func main() {
	myApp := app.New()

	ShowLoginWindow(myApp)
}
