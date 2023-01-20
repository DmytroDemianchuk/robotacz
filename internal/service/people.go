package service

import (
	"context"

	"github.com/dmytrodemianchuk/robotacz/internal/domain"
)

type PeopleRepository interface {
	List(ctx context.Context) (domain.ListPeople, error)
	Get(ctx context.Context, id int) (domain.People, error)
	Create(ctx context.Context, people domain.People) (domain.People, error)
	Update(ctx context.Context, id int, people domain.People) (domain.People, error)
	Delete(ctx context.Context, id int) error
}

type People struct {
	peopleRepository PeopleRepository
}

func NewPeople(peopleRepository PeopleRepository) *People {
	return &People{peopleRepository: peopleRepository}
}

func (p People) List(ctx context.Context) (domain.ListPeople, error) {
	return p.peopleRepository.List(ctx)
}

func (p People) Get(ctx context.Context, id int) (domain.People, error) {
	return p.peopleRepository.Get(ctx, id)
}

func (p People) Create(ctx context.Context, people domain.People) (domain.People, error) {
	return p.peopleRepository.Create(ctx, people)
}

func (p People) Update(ctx context.Context, id int, people domain.People) (domain.People, error) {
	return p.peopleRepository.Update(ctx, id, people)
}

func (p People) Delete(ctx context.Context, id int) error {
	return p.peopleRepository.Delete(ctx, id)
}
