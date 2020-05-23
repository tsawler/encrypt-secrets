package main

import (
	"context"
	"time"
)

func (app *application) updateSecret(pref string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var value string

	query := `select preference from preferences p where p.name = $1`
	row := app.db.QueryRowContext(ctx, query, pref)

	err := row.Scan(
		&value,
	)

	if err != nil {
		return err
	}
	encrypted, err := Encrypt(value, []byte(key))
	stmt := "update preferences p set preference = $1 where p.name = $2"

	_, err = app.db.ExecContext(ctx, stmt, encrypted, pref)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) updateSecretMysql(pref string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var value string

	query := `select preference from preferences p where p.name = ?`
	row := app.db.QueryRowContext(ctx, query, pref)

	err := row.Scan(
		&value,
	)

	if err != nil {
		return err
	}
	encrypted, err := Encrypt(value, []byte(key))
	stmt := "update preferences p set preference = ? where p.name = ?"

	_, err = app.db.ExecContext(ctx, stmt, encrypted, pref)
	if err != nil {
		return err
	}

	return nil
}

