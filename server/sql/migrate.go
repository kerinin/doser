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
					`CREATE TABLE water_level_sensors(
						id UUID NOT NULL,
						firmata_id UUID NOT NULL,
						pin INT NOT NULL,
						kind TEXT NOT NULL,

						PRIMARY KEY (id)
						FOREIGN KEY(firmata_id) REFERENCES firmatas(id)
					)`,
					`CREATE TABLE auto_top_offs(
						id UUID NOT NULL,
						pump_id UUID NOT NULL,
						fill_rate REAL NOT NULL,
						fill_frequency TEXT NOT NULL,
						max_fill_volume REAL NOT NULL,

						PRIMARY KEY (id)
						FOREIGN KEY(pump_id) REFERENCES pumps(id)
					)`,
					`CREATE TABLE auto_top_offs_water_level_sensors(
						auto_top_off_id UUID NOT NULL,
						water_level_sensor_id UUID NOT NULL,
						PRIMARY KEY (auto_top_off_id, water_level_sensor_id)
						FOREIGN KEY(auto_top_off_id) REFERENCES auto_top_offs(id)
						FOREIGN KEY(water_level_sensor_id) REFERENCES water_level_sensors(id)
					)`,
					`CREATE TABLE auto_water_changes(
						id UUID NOT NULL,
						fresh_pump_id UUID NOT NULL,
						waste_pump_id UUID NOT NULL,
						exchange_rate REAL NOT NULL,

						PRIMARY KEY (id)
						FOREIGN KEY(fresh_pump_id) REFERENCES pumps(id)
						FOREIGN KEY(waste_pump_id) REFERENCES pumps(id)
					)`,
					`CREATE TABLE dosers(
						id UUID NOT NULL,
						PRIMARY KEY (id)
					)`,
					`CREATE TABLE doser_components(
						id UUID NOT NULL,
						doser_id UUID NOT NULL,
						pump_id UUID NOT NULL,
						dose_rate REAL NOT NULL,

						PRIMARY KEY (id)
						FOREIGN KEY(doser_id) REFERENCES dosers(id)
						FOREIGN KEY(pump_id) REFERENCES pumps(id)
					)`,
				},
			},
		},
	}
	return migrate.Exec(db, driver, migrations, migrate.Up)
}
