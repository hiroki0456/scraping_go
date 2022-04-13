package interactor

import (
	"context"
	"practice/scraping/usecase/port"
)

type TouristArea struct {
	OutputPort      port.TouristAreaOutputPort
	TouristAreaRepo port.TouristAreaRepository
}

func NewTouristAreaInputPort(outputPort port.TouristAreaOutputPort, touristAreaRepository port.TouristAreaRepository) port.TouristAreaInputPort {
	return &TouristArea{
		OutputPort:      outputPort,
		TouristAreaRepo: touristAreaRepository,
	}
}

func (t *TouristArea) GetTouristAreas(ctx context.Context) {
	ta, err := t.TouristAreaRepo.GetTouristAreas(ctx)
	if err != nil {
		t.OutputPort.RenderError(err)
		return
	}
	t.OutputPort.Render(ta)
}
