package Users

import (
	"calendly/schema/entities"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(name string, email string, password string) error {
	user := &entities.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	return mgm.Coll(user).Create(user)
}

func FindUserById(id string) (entities.User, error) {
	user := new(entities.User)
	coll := mgm.Coll(user)

	err := coll.FindByID(id, user)

	return *user, err
}

func FindUserByAttr(query primitive.M) (entities.User, error) {
	user := new(entities.User)

	err := mgm.Coll(user).First(query, user)

	return *user, err
}

func FindUsersByAttr(query primitive.M) ([]entities.User, error) {
	users := new([]entities.User)

	err := mgm.Coll(&entities.User{}).SimpleFind(users, query)

	return *users, err
}
