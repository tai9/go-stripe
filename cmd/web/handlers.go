package main

import "net/http"

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Hit the handler")
}
