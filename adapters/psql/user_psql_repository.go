package psql

import (
	"context"
	"fmt"
	"it-test/app/query"

	"github.com/google/uuid"

	"github.com/go-pg/pg/v10"
)

type User struct {
	//nolint:structcheck,gocritic,unused
	tableName struct{} `sql:"users"`
	ID        string   `sql:"id,notnull"`
	UserName  string   `sql:"user_name,notnull"`
	LastName  string   `sql:"last_name,notnull"`
	FirstName string   `sql:"first_name,notnull"`
	Password  string   `sql:"password,notnull"`
	Email     string   `sql:"email,notnull"`
	Mobile    string   `sql:"mobile,notnull"`
	ASZF      bool     `sql:"aszf,notnull"`
}

type UserPSQLRepository struct {
	db *pg.DB
}

func NewUserPSQLRepository(db *pg.DB) *UserPSQLRepository {
	return &UserPSQLRepository{db: db}
}

func (r *UserPSQLRepository) GetUserCount(
	ctx context.Context) (int, error) {
	user := new(User)
	count, err := r.db.WithContext(ctx).
		Model(user).
		Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *UserPSQLRepository) GetUserList(
	ctx context.Context, params query.GetUserList) (res []query.GetUserListItem, err error) {
	if params.Limit <= 0 {
		res = make([]query.GetUserListItem, 0)
		return
	}
	users := make([]User, 0)
	q := r.db.WithContext(ctx).
		Model(&users).
		Limit(params.Limit).
		Offset(params.Limit * params.PageIndex)
	if params.OrderBy != "" {
		if params.Order == "" {
			params.Order = "asc"
		}
		q = q.OrderExpr(fmt.Sprintf("%s %s", params.OrderBy, params.Order))
	}
	if params.EmailFilter != nil {
		q = q.Where("email = ?", params.EmailFilter)
	}
	if err = q.Select(); err != nil {
		return
	}
	res = make([]query.GetUserListItem, 0)
	for _, user := range users {
		uid, _ := uuid.Parse(user.ID)
		res = append(res, query.GetUserListItem{
			Aszf:      user.ASZF,
			Email:     user.Email,
			FirstName: user.FirstName,
			Id:        uid,
			LastName:  user.LastName,
			UserName:  user.UserName,
			Mobile:    user.Mobile,
		})
	}

	return
}

func (r *UserPSQLRepository) CreateUser(
	ctx context.Context, params query.CreateDbUser) (res query.GetUser, err error) {
	user := User{
		UserName:  params.UserName,
		LastName:  params.LastName,
		FirstName: params.FirstName,
		Password:  params.Password,
		Email:     params.Email,
		Mobile:    params.Mobile,
		ASZF:      params.Aszf,
	}
	_, err = r.db.WithContext(ctx).
		Model(&user).
		Returning("*").Insert()
	if err != nil {
		return
	}
	uid, _ := uuid.Parse(user.ID)
	res = query.GetUser{
		Aszf:      user.ASZF,
		Email:     user.Email,
		FirstName: user.FirstName,
		Id:        uid,
		LastName:  user.LastName,
		Mobile:    user.Mobile,
		UserName:  user.UserName,
	}
	return
}

func (r *UserPSQLRepository) UpdateUser(
	ctx context.Context, params query.UpdateDbUser) (res query.GetUser, err error) {
	user := User{
		ID:        params.Id.String(),
		UserName:  params.UserName,
		LastName:  params.LastName,
		FirstName: params.FirstName,
		Password:  params.Password,
		Mobile:    params.Mobile,
	}
	_, err = r.db.WithContext(ctx).
		Model(&user).
		Column("user_name").
		Column("last_name").
		Column("first_name").
		Column("password").
		Column("mobile").
		WherePK().
		Update()
	if err != nil {
		return
	}
	err = r.db.WithContext(ctx).
		Model(&user).
		WherePK().
		Select()
	if err != nil {
		return
	}
	uid, _ := uuid.Parse(user.ID)
	res = query.GetUser{
		Aszf:      user.ASZF,
		Email:     user.Email,
		FirstName: user.FirstName,
		Id:        uid,
		LastName:  user.LastName,
		Mobile:    user.Mobile,
		UserName:  user.UserName,
	}
	return
}
