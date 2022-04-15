package gateway

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"practice/scraping/entity"
	"practice/scraping/usecase/port"
)

type TouristAreaRepository struct {
	conn *sql.DB
}

func NewTouristAreaRepository(conn *sql.DB) port.TouristAreaRepository {
	return &TouristAreaRepository{
		conn: conn,
	}
}

func (t *TouristAreaRepository) GetTouristAreas(ctx context.Context, url string) (*entity.TouristArea, error) {
	conn := t.GetDBConn()
	row := conn.QueryRowContext(ctx, "SELECT * FROM tourist")
	ta := entity.TouristArea{}
	err := row.Scan(ta)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("tourist area not found")
		}
		log.Println(err)
		return nil, errors.New("internal server error. adapter/gateway/GetTouristAreas")
	}
	return &ta, nil
}

func (t *TouristAreaRepository) GetDBConn() *sql.DB {
	return t.conn
}
