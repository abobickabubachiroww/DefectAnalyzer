package handlers

import "fmt"

type AppHandler struct{}

func (h *AppHandler) Greet(name string) string {
	return fmt.Sprintf("Привет, %s! (Отправлено из Go)", name)
}
