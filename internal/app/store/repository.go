package store

import (
	"time"

	"github.com/spolyakovs/racing-backend-ifmo/internal/app/model"
)

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByUsername(string) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	Update(*model.User) error
	Delete(int) error
}

type TeamRepository interface {
	Create(*model.Team) error
	Find(int) (*model.Team, error)
	FindByName(string) (*model.Team, error)
	Update(*model.Team) error
	Delete(int) error
}

type DriverRepository interface {
	Create(*model.Driver) error
	Find(int) (*model.Driver, error)
	Update(*model.Driver) error
	Delete(int) error
}

type RaceRepository interface {
	Create(*model.Race) error
	Find(int) (*model.Race, error)
	FindByDate(time.Time) (*model.Race, error)
	Update(*model.Race) error
	Delete(int) error
}

type TeamDriverContractRepository interface {
	Create(*model.TeamDriverContract) error
	Find(int) (*model.TeamDriverContract, error)
	Update(*model.TeamDriverContract) error
	Delete(int) error
}

type RaceResultRepository interface {
	Create(*model.RaceResult) error
	Find(int) (*model.RaceResult, error)
	Update(*model.RaceResult) error
	Delete(int) error
}
