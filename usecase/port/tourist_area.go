package port

import (
	"context"
	"practice/scraping/entity"
)

type TouristAreaInputPort interface {
	GetTouristAreas(ctx context.Context)
}

type TouristAreaOutputPort interface {
	Render(*entity.TouristArea)
	RenderError(error)
}

type TouristAreaRepository interface {
	GetTouristAreas(ctx context.Context) (*entity.TouristArea, error)
}
