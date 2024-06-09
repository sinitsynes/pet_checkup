package pets

import (
	"context"
	"log"
	db "pet_checkup/db/pets"
	postgres "pet_checkup/pg_connection"
)

func session(ctx context.Context) *db.Queries {
	conn, err := postgres.NewPG(ctx)
	if err != nil {
		log.Fatalf("err: %q", err)
	}
	return db.New(conn)
}

func CreatePet(ctx context.Context, arg db.CreatePetParams) error {
	q := session(ctx)
	return q.CreatePet(ctx, arg)
}

func EditPet(ctx context.Context, arg db.UpdatePetParams) error {
	q := session(ctx)
	return q.UpdatePet(ctx, arg)
}

func GetPet(ctx context.Context, id int32) (db.Pet, error) {
	q := session(ctx)
	return q.GetPet(ctx, id)
}

func DeletePet(ctx context.Context, id int32) error {
	q := session(ctx)
	return q.DeletePet(ctx, id)
}

func CreateWeightMeasure(ctx context.Context, arg db.CreateWeightMeasureParams) error {
	q := session(ctx)
	return q.CreateWeightMeasure(ctx, arg)
}

func DeleteWeightMeasure(ctx context.Context, arg db.RemoveWeightMeasureParams) error {
	q := session(ctx)
	return q.RemoveWeightMeasure(ctx, arg)
}
