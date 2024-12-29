package main

import (
	"time"

	"github.com/daniela2001-png/rss_aggregator_project/internal/database"
	"github.com/google/uuid"
)

type ResponseUser struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func ConvertDataBaseUserToResponseUser(dbUser database.User) ResponseUser {
	return ResponseUser{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}
