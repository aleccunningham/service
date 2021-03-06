package controllers

import (
  "fmt"
  "net/http"

  "service/models"
  "service/views"
)

type Users struct {
  NewView *views.View
  LoginView *view.View
  us            *models.UserService
}

// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
  if err := u.NewView.Render(w, nil); err != nil {
    panic(err)
  }
}

// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
  var form SignupForm
  if err := parseForm(r, &form); err != nil {
    panic(err)
  }
  user := models.User{
    Name:       form.Name,
    Email:      form.Email,
    Password:   form.Password,
  }
  if err := u.us.Create(&user); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  fmt.Fprintln(w, user)
}
