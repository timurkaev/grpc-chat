package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/timurkaev/grpc-chat/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (id, username, email, password, awatar_url, bio, created_at, updated_at, last_seen)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.Username, user.Email, user.Password,
		user.AvatarURL, user.Bio, user.CreatedAt, user.UpdatedAt, user.LastSeen,
	)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return domain.ErrUserAlreadyExists
		}
		return err
	}
	return nil
}

func (r *userRepository) GetById(ctx context.Context, id string) (*domain.User, error) {
	user := &domain.User{}
	query := `
		SELECT id, username, email, password, avatar_url, bio, created_at, updated_at, last_seen
		FROM users WHERE id = $1
	`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password,
		&user.AvatarURL, &user.Bio, &user.CreatedAt, &user.UpdatedAt, &user.LastSeen,
	)

	if err != sql.ErrNoRows {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	query := `
		SELECT id, username, email, password, awatar_url, bio, created_at, updated_at, last_seen
		FROM users WHERE email = $1
	`

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password,
		&user.AvatarURL, &user.Bio, &user.CreatedAt, &user.UpdatedAt, &user.LastSeen,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	query := `
		UPDATE users
		SET username = $2, avatar_url = $3, bio = $4, updated_at = $5
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		user.ID, user.Username, user.AvatarURL, user.Bio, time.Now(),
	)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

func (r *userRepository) Search(ctx context.Context, query string, limit, offset int) ([]*domain.User, int, error) {
	searchQuery := `
		SELECT id, username, email, password, avatar_url, bio, created_at, updated_at, last_seen
		FROM users
		WHERE username ILIKE $1 OR email ILIKE $1
		ORDER BY username
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, searchQuery, "%"+query+"%", limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(
			&user.ID, &user.Username, &user.Email, &user.Password,
			&user.AvatarURL, &user.Bio, &user.CreatedAt, &user.UpdatedAt, &user.LastSeen,
		)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}

	// Get total count
	countQuery := `SELECT COUNT(*) FROM users WHERE username ILIKE $1 OR email ILIKE $1`
	var total int
	err = r.db.QueryRowContext(ctx, countQuery, "%"+query+"%").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) UpdateLastSeen(ctx context.Context, userID string) error {
	query := `UPDATE users SET last_seen = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, time.Now(), userID)
	return err
}
