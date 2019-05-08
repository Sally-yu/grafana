package user

import (
	"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/models"
)

// Upserts the user
func Upsert(
	externalUser *models.ExternalUserInfo,
	SignupAllowed bool,
) (*models.User, error) {
	// add/update user in grafana
	query := &models.UpsertUserCommand{
		ExternalUser:  externalUser,
		SignupAllowed: SignupAllowed,
	}

	err := bus.Dispatch(query)
	if err != nil {
		return nil, err
	}

	return query.Result, nil
}