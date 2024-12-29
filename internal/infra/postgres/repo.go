package postgres

import (
	"context"
	"time"

	"app/internal/domain"
	sqlcgen "app/internal/infra/postgres/queries"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostgresRepo struct {
	queries *sqlcgen.Queries
	ctx     context.Context
}

func InitPostgresRepo(connStr string, ctx context.Context) (*PostgresRepo, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	queries := sqlcgen.New(conn)

	return &PostgresRepo{
		queries: queries,
		ctx:     ctx,
	}, nil
}

func (r *PostgresRepo) GetUserByLogin(login string) (*domain.User, error) {
	user, err := r.queries.GetUserByLogin(r.ctx, login)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        user.ID,
		Login:     user.Login,
		PwdHash:   user.PwdHash,
		CreatedAt: user.CreatedAt.Time.Unix(),
	}, nil
}

func (r *PostgresRepo) CreateUser(user *domain.User) error {
	_, err := r.queries.CreateUser(r.ctx, sqlcgen.CreateUserParams{
		Login:     user.Login,
		PwdHash:   user.PwdHash,
		CreatedAt: pgtype.Timestamp{Time: time.Unix(user.CreatedAt, 0)},
	})

	return err
}
