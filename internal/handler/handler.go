package handler

import "github.com/Thunderbirrd/CourseProject/internal/service"

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
