package handlers

import (
	"fmt"
	"net/http"
	"todoapi/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.PrintMessage("HomeHandler called")
	fmt.Fprintf(w, "Welcome to the Home Page!")
}
