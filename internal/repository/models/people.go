package models

import "github.com/dmytrodemianchuk/robotacz/internal/domain"

type People struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	PhoneNumber string `db:"phone_number"`
	BirthYear   int    `db:"birth_year"`
	Nationality string `db:"nationality"`
}

func (p People) ToDomain() domain.People {
	return domain.People{
		ID:          p.ID,
		Name:        p.Name,
		PhoneNumber: p.PhoneNumber,
		BirthYear:   p.BirthYear,
		Nationality: p.Nationality,
	}
}
