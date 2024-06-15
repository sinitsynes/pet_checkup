package panels

import (
	"context"
	"log"
	db "pet_checkup/db/panels"
	postgres "pet_checkup/pg_connection"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func StringToRange(input string) (*pgtype.Range[pgtype.Numeric], error) {
	numericLower := new(pgtype.Numeric)
	numericUpper := new(pgtype.Numeric)

	bounds := strings.Split(input, "-")
	if errLower := numericLower.Scan(bounds[0]); errLower != nil {
		return nil, errLower
	}
	if errUpper := numericUpper.Scan(bounds[1]); errUpper != nil {
		return nil, errUpper
	}
	interval := new(pgtype.Range[pgtype.Numeric])
	interval.Lower = *numericLower
	interval.Upper = *numericUpper

	return interval, nil
}

func createMeasurentUnitParams(langID int32, name string) db.CreateMeasurementUnitParams {
	measurement := new(db.CreateMeasurementUnitParams)
	langData := pgtype.Int4{Int32: langID, Valid: true}
	measurement.LanguageID = langData
	measurement.Name = name
	return *measurement
}

func AddPanelMeasurement(ctx context.Context, apiRequest CreateMeasurementAPIRequest) error {
	conn, err := postgres.NewPG(ctx)
	if err != nil {
		log.Fatalf("err: %q", err)
	}
	session, err := conn.BeginTx(ctx, pgx.TxOptions{})
	defer session.Rollback(ctx)

	if err != nil {
		log.Fatalf("err: %q", err)
	}
	qtx := db.New(conn).WithTx(session)

	session.Begin(ctx)

	langID, langErr := qtx.GetLanguageID(ctx, apiRequest.LangLabel)
	if langErr != nil {
		log.Fatalf("langID %q", langErr)
	}
	unitParams := createMeasurentUnitParams(langID, apiRequest.Name)
	if err := qtx.CreateMeasurementUnit(ctx, unitParams); err != nil {
		log.Fatalf("unitParams: %q", err)
		return err
	}
	session.Commit(ctx)
	return nil
}
