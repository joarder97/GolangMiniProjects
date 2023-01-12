package controllers

import "net/http"

func Logout(w http.ResponseWriter, r *http.Request) {
	// Clear the cookie.
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
