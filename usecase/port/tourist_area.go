package port

import (
	"context"
	"practice/scraping/entity"
)

type TouristAreaInputPort interface {
	GetTouristAreas(ctx context.Context, url string)
}

type TouristAreaOutputPort interface {
	Render(*entity.TouristArea)
	RenderError(error)
}

type TouristAreaRepository interface {
	GetTouristAreas(ctx context.Context, url string) (*entity.TouristArea, error)
}
