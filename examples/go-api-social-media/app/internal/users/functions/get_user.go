package users_functions

import (
	"api-social-media/app/core/db"
	responses "api-social-media/app/core/helpers/response"
	"api-social-media/app/models"
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	conn := db.SetupDB()
	defer conn.Close(context.Background())

	// Extrair o ID do usuário da rota
	id := mux.Vars(r)["id"]
	userID, err := uuid.Parse(id)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, errors.New("ID de usuário inválido"))
		return
	}

	// Consultar o banco de dados para obter o usuário pelo ID
	var user models.User
	query := "SELECT * FROM users WHERE id = $1"
	err = conn.QueryRow(context.Background(), query, userID).Scan(
		&user.ID, &user.Name, &user.Nick, &user.Email, &user.Pass, &user.CreatedAt,
	)
	if err != nil {
		responses.Err(w, http.StatusNotFound, errors.New("usuário não encontrado"))
		return
	}

	responses.JSON(w, http.StatusOK, user)
}
