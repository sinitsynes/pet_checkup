package languages

import (
	"context"
	"log"
	db "pet_checkup/db/language"
	"pet_checkup/pg_connection"
)

func dbSession(ctx context.Context) *db.Queries {
	conn, err := pg_connection.NewPG(ctx)
	if err != nil {
		log.Fatalf("can't connect to db: %q", err)
	}
	return db.New(conn)
}

func AddLanguage(ctx context.Context, name string) (string, error) {
	q := dbSession(ctx)
	res, err := q.CreateLanguage(ctx, name)
	if err != nil {
		log.Fatalf("can't create language: %q", err)
	}
	return res.Name, nil
}

func GetLanguages(ctx context.Context) ([]string, error) {
	q := dbSession(ctx)
	res, err := q.GetLanguages(ctx)
	return res, err
}
