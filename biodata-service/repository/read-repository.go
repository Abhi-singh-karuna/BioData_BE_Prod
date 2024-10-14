package repository

import (
	"myapp/model"

	"github.com/pkg/errors"
)

func (r *Database) ValidateUserVerified(userId string) (bool, *model.User, error) {
	row, err := r.db.Query(QueryValidateUserVerified, userId)
	if err != nil {
		return false, nil, errors.Wrap(err, "repository.user.ValidateUserVerified.query")
	}
	var user model.User
	for row.Next() {
		err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNo, &user.IsVerified)
		if err != nil {
			return false, nil, errors.Wrap(err, "repository.user.ValidateUserVerified.rows.Scan")
		}
	}
	r.logger.Infof("repository.user.ValidateUserVerified %v", user)

	if !user.IsVerified {
		return false, &user, nil
	}
	return true, &user, nil
}
