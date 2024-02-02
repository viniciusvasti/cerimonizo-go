package dto

import (
	"time"
	"viniciusvasti/cerimonize/application"
)

type WeddingDTO struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Date   string  `json:"date"`
	Budget float64 `json:"budget"`
	Status string  `json:"status"`
}

func (w *WeddingDTO) Bind(wedding application.WeddingInterface) {
	w.ID = wedding.GetId()
	w.Name = wedding.GetName()
	w.Date = wedding.GetDate().String()
	w.Budget = wedding.GetBudget()
	w.Status = wedding.GetStatus()
}

func BindAll(weddings []application.WeddingInterface) []WeddingDTO {
	var weddingsDTO []WeddingDTO
	for _, wedding := range weddings {
		var weddingDTO WeddingDTO
		weddingDTO.Bind(wedding)
		weddingsDTO = append(weddingsDTO, weddingDTO)
	}
	return weddingsDTO
}

func (w WeddingDTO) ConvertToEntity() (*application.Wedding, error) {
	date, err := time.Parse("2006-01-02 15:04:05-07:00", w.Date)
	if err != nil {
		return nil, err
	}
	return &application.Wedding{
		ID:     w.ID,
		Name:   w.Name,
		Date:   date,
		Budget: w.Budget,
		Status: w.Status,
	}, nil
}
