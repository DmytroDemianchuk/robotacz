package rest

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	restErrors "github.com/dmytrodemianchuk/robotacz/pkg/rest/errors"

	"github.com/dmytrodemianchuk/robotacz/internal/domain"
	"github.com/gin-gonic/gin"
)

type PeopleService interface {
	List(ctx context.Context) (domain.ListPeople, error)
	Get(ctx context.Context, id int) (domain.People, error)
	Create(ctx context.Context, people domain.People) (domain.People, error)
	Update(ctx context.Context, id int, people domain.People) (domain.People, error)
	Delete(ctx context.Context, id int) error
}

type People struct {
	peopleService PeopleService
}

func NewPeople(peopleService PeopleService) *People {
	return &People{peopleService: peopleService}
}

func (p People) List(ctx *gin.Context) {
	music, err := p.peopleService.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		return
	}

	ctx.JSON(http.StatusOK, music)
}

func (p People) Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fields := map[string]string{"id": "should be an integer"}
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("validation error", fields))
		return
	}

	music, err := p.peopleService.Get(ctx, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			ctx.JSON(http.StatusNotFound, restErrors.NewNotFoundErr("people not found"))
		default:
			ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		}

		return
	}

	ctx.JSON(http.StatusOK, music)
}

func (p People) Create(ctx *gin.Context) {
	var music domain.Music
	if err := ctx.BindJSON(&music); err != nil {
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("cannot parse body", nil))
		return
	}

	createdMusic, err := m.musicService.Create(ctx, music)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		return
	}

	ctx.JSON(http.StatusCreated, createdMusic)
}

func (p People) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fields := map[string]string{"id": "should be an integer"}
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("validation error", fields))
		return
	}

	var music domain.Music
	if err := ctx.BindJSON(&music); err != nil {
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("cannot parse body", nil))
		return
	}

	updatedMusic, err := m.musicService.Update(ctx, id, music)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		return
	}

	ctx.JSON(http.StatusOK, updatedMusic)
}

func (p People) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fields := map[string]string{"id": "should be an integer"}
		ctx.JSON(http.StatusBadRequest, restErrors.NewBadRequestErr("validation error", fields))
		return
	}

	if err := p.peopleService.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, restErrors.NewInternalServerErr())
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
