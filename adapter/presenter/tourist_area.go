package presenter

import (
	"fmt"
	"net/http"
	"practice/scraping/entity"
	"practice/scraping/usecase/port"
)

type TouristArea struct {
	w http.ResponseWriter
}

func NewTouristArea(w http.ResponseWriter) port.TouristAreaOutputPort {
	return &TouristArea{
		w: w,
	}
}

func (t *TouristArea) Render(ta *entity.TouristArea) {
	t.w.WriteHeader(http.StatusOK)
	fmt.Fprintf(t.w, ta.AccessContent)
}

func (t *TouristArea) RenderError(err error) {
	t.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(t.w, err)
}
