package repository

import (
	"context"

	"github.com/dmytrodemianchuk/robotacz/internal/domain"
	"github.com/dmytrodemianchuk/robotacz/internal/repository/models"

	"github.com/jmoiron/sqlx"
)

type People struct {
	db *sqlx.DB
}

func NewPeople(db *sqlx.DB) *People {
	return &People{db: db}
}

func (p People) List(ctx context.Context) (domain.ListPeople, error) {
	var list []models.People
	if err := p.db.SelectContext(ctx, &list, "SELECT * FROM people"); err != nil {
		return nil, err
	}

	dlist := make(domain.ListPeople, 0, len(list))
	for _, people := range list {
		dlist = append(dlist, people.ToDomain())
	}

	return dlist, nil
}

func (p People) Get(ctx context.Context, id int) (domain.People, error) {
	var people models.People
	if err := p.db.GetContext(ctx, &people, "SELECT * FROM  music WHERE id=$1", id); err != nil {
		return domain.People{}, err
	}

	return people.ToDomain(), nil
}

func (p People) Create(ctx context.Context, people domain.People) (domain.People, error) {
	pPeople := models.People{
		Name:        people.Name,
		PhoneNumber: people.PhoneNumber,
		BirthYear:   people.BirthYear,
		Nationality: people.Nationality,
	}

	if err := p.db.QueryRowxContext(ctx, "INSERT INTO people (name, phone_number, birth_year, nationality) VALUES ($1, $2, $3, $4) RETURNING *", pPeople.Name, pPeople.PhoneNumber, pPeople.BirthYear, pPeople.Nationality).StructScan(&pPeople); err != nil {
		return domain.People{}, err
	}

	return pPeople.ToDomain(), nil
}

func (p People) Update(ctx context.Context, id int, people domain.People) (domain.People, error) {
	pPeople := models.People{
		Name:        people.Name,
		PhoneNumber: people.PhoneNumber,
		BirthYear:   people.BirthYear,
		Nationality: people.Nationality,
	}

	if err := p.db.QueryRowxContext(ctx, "UPDATE people SET name=$1, phone_number=$2, birth_year=$3, nationality=$4 WHERE id=$5 RETURNING *", pPeople.Name, pPeople.PhoneNumber, pPeople.BirthYear, pPeople.Nationality, id).StructScan(&pPeople); err != nil {
		return domain.Music{}, err
	}

	return pPeople.ToDomain(), nil
}

func (p People) Delete(ctx context.Context, id int) error {
	if _, err := p.db.ExecContext(ctx, "DELETE FROM people WHERE id=$1", id); err != nil {
		return err
	}
	return nil
}
