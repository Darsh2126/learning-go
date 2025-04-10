package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Darsh2126/learning-go/internal/types"
	"github.com/Darsh2126/learning-go/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// request validation
		if err := validator.New().Struct(student); err != nil {
			validateErrors := err.(validator.ValidationErrors)
			response.WriteJSON(w, http.StatusBadRequest, response.ValidationError(validateErrors))
			return
		}

		// w.Write([]byte("Welcome to learning go"))
		response.WriteJSON(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
