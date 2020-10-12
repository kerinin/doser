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
						baud INT NOT NULL,

						PRIMARY KEY (id)
					)`,
					`CREATE TABLE pumps(
						id UUID NOT NULL,
						firmata_id UUID NOT NULL,
						device_id INT NOT NULL,
						step_pin INT NOT NULL,
						dir_pin INT,
						en_pin INT,

						PRIMARY KEY (id)
						FOREIGN KEY(firmata_id) REFERENCES firmatas(id)
					)`,
					`CREATE TABLE calibrations(
						id UUID NOT NULL,
						pump_id UUID NOT NULL,
						steps INT NOT NULL,
						volume REAL NOT NULL,

						PRIMARY KEY (id)
						FOREIGN KEY(pump_id) REFERENCES pumps(id)
					)`,
					`CREATE TABLE water_level_sensors(
						id UUID NOT NULL,
						firmata_id UUID,
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
			&migrate.Migration{
				Id: "001 allow-analog",
				Up: []string{
					`ALTER TABLE water_level_sensors ADD COLUMN detection_threshold INT`,
				},
			},
			&migrate.Migration{
				Id: "002 interval-not-cron",
				Up: []string{
					`DROP TABLE auto_top_offs`,
					`DROP TABLE auto_top_offs_water_level_sensors`,
					`CREATE TABLE auto_top_offs(
						id UUID NOT NULL,
						pump_id UUID NOT NULL,
						fill_rate REAL NOT NULL,
						fill_interval INT NOT NULL,
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
				},
			},
			&migrate.Migration{
				Id: "003 pump-acceleration",
				Up: []string{
					`ALTER TABLE pumps ADD COLUMN acceleration REAL`,
				},
			},
			&migrate.Migration{
				Id: "004 events",
				Up: []string{
					`CREATE TABLE ato_events(
						id UUID NOT NULL,
						auto_top_off_id UUID NOT NULL,
						timestamp INT NOT NULL,
						kind STRING NOT NULL,
						data STRING NOT NULL,

						PRIMARY KEY (id)
						FOREIGN KEY(auto_top_off_id) REFERENCES auto_top_offs(id)
					)`,
					`CREATE TABLE awc_events(
						id UUID NOT NULL,
						auto_water_change_id UUID NOT NULL,
						timestamp INT NOT NULL,
						kind STRING NOT NULL,
						data STRING NOT NULL,

						PRIMARY KEY (id)
						FOREIGN KEY(auto_water_change_id) REFERENCES auto_water_changes(id)
					)`,
				},
			},
			&migrate.Migration{
				Id: "005 doses",
				Up: []string{
					`CREATE TABLE doses(
						id UUID NOT NULL,
						pump_id UUID NOT NULL,
						timestamp INT NOT NULL,
						volume REAL NOT NULL,
						message STRING,

						PRIMARY KEY (id)
						FOREIGN KEY(pump_id) REFERENCES pumps(id)
					)`,
				},
			},
			&migrate.Migration{
				Id: "006 enable",
				Up: []string{
					`ALTER TABLE auto_top_offs ADD COLUMN enabled BOOLEAN NOT NULL DEFAULT 1`,
					`ALTER TABLE auto_water_changes ADD COLUMN enabled BOOLEAN NOT NULL DEFAULT 1`,
					`ALTER TABLE dosers ADD COLUMN enabled BOOLEAN NOT NULL DEFAULT 1`,
				},
			},
			&migrate.Migration{
				Id: "007 sensorinvert",
				Up: []string{
					`ALTER TABLE water_level_sensors ADD COLUMN invert BOOLEAN NOT NULL DEFAULT 0`,
				},
			},
			&migrate.Migration{
				Id: "008 names",
				Up: []string{
					`ALTER TABLE firmatas ADD COLUMN name TEXT`,
					`ALTER TABLE pumps ADD COLUMN name TEXT`,
					`ALTER TABLE water_level_sensors ADD COLUMN name TEXT`,
					`ALTER TABLE auto_top_offs ADD COLUMN name TEXT`,
					`ALTER TABLE auto_water_changes ADD COLUMN name TEXT`,
					`ALTER TABLE dosers ADD COLUMN name TEXT`,
				},
			},
		},
	}
	return migrate.Exec(db, driver, migrations, migrate.Up)
}
