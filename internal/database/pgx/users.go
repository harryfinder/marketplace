package pgx

import (
	"context"
	"marketplace/internal/models"
	"time"
)

func (d *db) CreateUser(ctx context.Context, user models.User) (users models.User, err error) {

	sqlQuery := `Insert into users(full_name,email,password,number_phone,active_phone,role,stocks,active_email,created_at,updated_at)values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning email;
		`
	err = d.postgres.QueryRow(ctx, sqlQuery, user.FullName, user.Email, user.Password, user.NumberPhone, user.ActivePhone, user.Role, user.Status, user.ActiveEmail, time.Now(), time.Now()).Scan(&users.Email)

	return
}

func (d *db) GetRoleUser(ctx context.Context, users *models.User) (*models.ResponseUserCheckingKey, error) {
	sqlQuery := `SELECT
		id,
		role
		FROM 
		users
		WHERE
		email =$1;
		`

	rows, err := d.postgres.Query(ctx, sqlQuery, users.Email)
	if err != nil {
		return nil, err
	}

	var user = &models.ResponseUserCheckingKey{}
	for rows.Next() {
		if err := rows.Scan(
			&user.UserId,
			&user.Role,
		); err != nil {
			return nil, err
		}

	}

	return user, nil
}
func (d *db) GetAllCheckEmails(ctx context.Context, check models.CheckEmail) (*models.CheckEmail, error) {
	sqlQuery := `SELECT
	id,
	key,
	user_id,
	email,
	status,
	FROM users
	WHERE
	user_id = $1
	`
	rows, err := d.postgres.Query(ctx, sqlQuery, check.UserId)
	if err != nil {
		return nil, err
	}
	var user = &models.CheckEmail{}
	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Key,
			&user.UserId,
			&user.Status,
		); err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (d *db) CreateCheckEmail(ctx context.Context, ch *models.CheckEmail) (*models.CheckEmail, error) {
	sqlQuery := `INSERT INTO check_emails
		(key,
		user_id,
		email,
		status)
		VALUES
		($1,$2,$3,$4)
		`
	rows, err := d.postgres.Query(ctx,
		sqlQuery,
		ch.Key,
		ch.UserId,
		ch.Email,
		ch.Status,
	)
	if err != nil {
		return nil, err
	}

	var users = &models.CheckEmail{}
	for rows.Next() {
		if err = rows.Scan(
			&ch.Key,
			&ch.UserId,
			&ch.Email,
			&ch.Status,
		); err != nil {
			return nil, err
		}
	}

	return users, nil
}
func (d *db) UpdateCheckEmail(ctx context.Context, ch *models.CheckEmail) error {
	sqlQuery := `Update check_emails
		SET key = $1,
		email 	= $2,
		WHERE user_id = $3;
		`
	_, err := d.postgres.Exec(ctx,
		sqlQuery,
		ch.Key,
		ch.Email,
		ch.UserId,
	)
	if err != nil {
		return err
	}

	return nil
}
func (d *db) GetUsersByID(ctx context.Context, id int64) (*models.User, error) {
	sqlQuery := `SELECT
		full_name,
		role,
		email,
		number_phone,
		stocks,
		FROM users
		WHERE
		id = $1
		AND active_email =true
		`

	rows, err := d.postgres.Query(ctx, sqlQuery, id)
	if err != nil {
		return nil, err
	}
	var user = &models.User{}
	for rows.Next() {
		if err = rows.Scan(
			&user.FullName,
			&user.Role,
			&user.Email,
			&user.NumberPhone,
		); err != nil {
			return nil, err
		}
		var count int64

		sqlQuery = `SELECT
			COUNT(*)
		FROM users
		WHERE
		id = $1
		`
		err = d.postgres.QueryRow(ctx, sqlQuery, id).Scan(&count)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (d *db) GetUser(ctx context.Context, login string, password string) (user models.ResponseUser, err error) {
	sqlQuery := `SELECT
		id,
		role
		FROM users
		WHERE
		email = $1 AND password = $2 AND active_email = true;`

	err = d.postgres.QueryRow(ctx, sqlQuery, login, password).Scan(&user.Id, &user.Role)

	return
}
