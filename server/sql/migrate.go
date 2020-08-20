package sql

import (
	"database/sql"
	"github.com/rubenv/sql-migrate"
)

// Migrate ensures the DB is in the expected shape
func Migrate(db *sql.DB, driver string) (int, error) {
	migrations := &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id: "000 initial",
				Up: []string{
					`CREATE TABLE firmatas(
						id UUID NOT NULL,
						serial_port TEXT NOT NULL,

						PRIMARY KEY (id)
					)`,
					`CREATE TABLE pumps(
						id UUID NOT NULL,
						firmata_id UUID NOT NULL,
						step_pin INT NOT NULL,
						dir_pin INT,
						en_pin INT,

						PRIMARY KEY (id)
						FOREIGN KEY(firmata_id) REFERENCES firmatas(id)
					)`,
					`CREATE TABLE calibrations(
						id UUID NOT NULL,
						pump_id UUID NOT NULL,
						target_volume REAL NOT NULL,
						measured_volume REAL NOT NULL,

						PRIMARY KEY (id)
						FOREIGN KEY(pump_id) REFERENCES pumps(id)
					)`,
				},
			},
		},
	}
	return migrate.Exec(db, driver, migrations, migrate.Up)
}
