package controller

import (
	"fmt"
	"net/http"
	"practice/scraping/usecase/port"
)

type TouristArea struct {
	OutputFactory func(w http.ResponseWriter) port.TouristAreaOutputPort
	InputFactory  func(o port.TouristAreaOutputPort, t port.TouristAreaRepository) port.TouristAreaInputPort
	RepoFactory   func(url string) port.TouristAreaRepository
	URL           string
}

func (t *TouristArea) GetTouristAreas(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	outputPort := t.OutputFactory(w)
	repository := t.RepoFactory(t.URL)
	inputPort := t.InputFactory(outputPort, repository)
	inputPort.GetTouristAreas(ctx, fmt.Sprintf("%s%s", r.Host, r.URL.Path))

}
