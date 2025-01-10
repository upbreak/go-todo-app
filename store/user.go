package store

import (
	"context"
	"fmt"
	"github.com/upbreak/go-todo-app/entity"
)

func (r *Repository) GetUserValid(ctx context.Context, db Queryer, id string, pwMd5 string) (entity.User, error) {
	user := entity.User{}

	sql := `SELECT
		    t1.USER_ID
		FROM
			COMMON.V_BIZ_USER_INFO t1
		WHERE
		    t1.USER_ID = :1
			AND t1.USER_PWD = :2`

	if err := db.GetContext(ctx, &user, sql, id, pwMd5); err != nil {
		return user, fmt.Errorf("store/user.go/db.GetContext(): %w", err)
	}
	return user, nil
}
