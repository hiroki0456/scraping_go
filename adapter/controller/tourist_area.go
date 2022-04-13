package controller

import (
	"database/sql"
	"net/http"
	"practice/scraping/usecase/port"
)

type TouristArea struct {
	OutputFactory func(w http.ResponseWriter) port.TouristAreaOutputPort
	InputFactory  func(o port.TouristAreaOutputPort, t port.TouristAreaRepository) port.TouristAreaInputPort
	RepoFactory   func(c *sql.DB) port.TouristAreaRepository
	Conn          *sql.DB
}

func (t *TouristArea) GetTouristAreas(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	outputPort := t.OutputFactory(w)
	repository := t.RepoFactory(t.Conn)
	inputPort := t.InputFactory(outputPort, repository)
	inputPort.GetTouristAreas(ctx)

}
