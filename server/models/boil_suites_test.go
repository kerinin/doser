// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("AtoEvents", testAtoEvents)
	t.Run("AutoTopOffs", testAutoTopOffs)
	t.Run("AutoWaterChanges", testAutoWaterChanges)
	t.Run("AwcEvents", testAwcEvents)
	t.Run("Calibrations", testCalibrations)
	t.Run("DoserComponents", testDoserComponents)
	t.Run("Dosers", testDosers)
	t.Run("Doses", testDoses)
	t.Run("Firmatas", testFirmatas)
	t.Run("GorpMigrations", testGorpMigrations)
	t.Run("Pumps", testPumps)
	t.Run("WaterLevelSensors", testWaterLevelSensors)
}

func TestDelete(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsDelete)
	t.Run("AutoTopOffs", testAutoTopOffsDelete)
	t.Run("AutoWaterChanges", testAutoWaterChangesDelete)
	t.Run("AwcEvents", testAwcEventsDelete)
	t.Run("Calibrations", testCalibrationsDelete)
	t.Run("DoserComponents", testDoserComponentsDelete)
	t.Run("Dosers", testDosersDelete)
	t.Run("Doses", testDosesDelete)
	t.Run("Firmatas", testFirmatasDelete)
	t.Run("GorpMigrations", testGorpMigrationsDelete)
	t.Run("Pumps", testPumpsDelete)
	t.Run("WaterLevelSensors", testWaterLevelSensorsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsQueryDeleteAll)
	t.Run("AutoTopOffs", testAutoTopOffsQueryDeleteAll)
	t.Run("AutoWaterChanges", testAutoWaterChangesQueryDeleteAll)
	t.Run("AwcEvents", testAwcEventsQueryDeleteAll)
	t.Run("Calibrations", testCalibrationsQueryDeleteAll)
	t.Run("DoserComponents", testDoserComponentsQueryDeleteAll)
	t.Run("Dosers", testDosersQueryDeleteAll)
	t.Run("Doses", testDosesQueryDeleteAll)
	t.Run("Firmatas", testFirmatasQueryDeleteAll)
	t.Run("GorpMigrations", testGorpMigrationsQueryDeleteAll)
	t.Run("Pumps", testPumpsQueryDeleteAll)
	t.Run("WaterLevelSensors", testWaterLevelSensorsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsSliceDeleteAll)
	t.Run("AutoTopOffs", testAutoTopOffsSliceDeleteAll)
	t.Run("AutoWaterChanges", testAutoWaterChangesSliceDeleteAll)
	t.Run("AwcEvents", testAwcEventsSliceDeleteAll)
	t.Run("Calibrations", testCalibrationsSliceDeleteAll)
	t.Run("DoserComponents", testDoserComponentsSliceDeleteAll)
	t.Run("Dosers", testDosersSliceDeleteAll)
	t.Run("Doses", testDosesSliceDeleteAll)
	t.Run("Firmatas", testFirmatasSliceDeleteAll)
	t.Run("GorpMigrations", testGorpMigrationsSliceDeleteAll)
	t.Run("Pumps", testPumpsSliceDeleteAll)
	t.Run("WaterLevelSensors", testWaterLevelSensorsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsExists)
	t.Run("AutoTopOffs", testAutoTopOffsExists)
	t.Run("AutoWaterChanges", testAutoWaterChangesExists)
	t.Run("AwcEvents", testAwcEventsExists)
	t.Run("Calibrations", testCalibrationsExists)
	t.Run("DoserComponents", testDoserComponentsExists)
	t.Run("Dosers", testDosersExists)
	t.Run("Doses", testDosesExists)
	t.Run("Firmatas", testFirmatasExists)
	t.Run("GorpMigrations", testGorpMigrationsExists)
	t.Run("Pumps", testPumpsExists)
	t.Run("WaterLevelSensors", testWaterLevelSensorsExists)
}

func TestFind(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsFind)
	t.Run("AutoTopOffs", testAutoTopOffsFind)
	t.Run("AutoWaterChanges", testAutoWaterChangesFind)
	t.Run("AwcEvents", testAwcEventsFind)
	t.Run("Calibrations", testCalibrationsFind)
	t.Run("DoserComponents", testDoserComponentsFind)
	t.Run("Dosers", testDosersFind)
	t.Run("Doses", testDosesFind)
	t.Run("Firmatas", testFirmatasFind)
	t.Run("GorpMigrations", testGorpMigrationsFind)
	t.Run("Pumps", testPumpsFind)
	t.Run("WaterLevelSensors", testWaterLevelSensorsFind)
}

func TestBind(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsBind)
	t.Run("AutoTopOffs", testAutoTopOffsBind)
	t.Run("AutoWaterChanges", testAutoWaterChangesBind)
	t.Run("AwcEvents", testAwcEventsBind)
	t.Run("Calibrations", testCalibrationsBind)
	t.Run("DoserComponents", testDoserComponentsBind)
	t.Run("Dosers", testDosersBind)
	t.Run("Doses", testDosesBind)
	t.Run("Firmatas", testFirmatasBind)
	t.Run("GorpMigrations", testGorpMigrationsBind)
	t.Run("Pumps", testPumpsBind)
	t.Run("WaterLevelSensors", testWaterLevelSensorsBind)
}

func TestOne(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsOne)
	t.Run("AutoTopOffs", testAutoTopOffsOne)
	t.Run("AutoWaterChanges", testAutoWaterChangesOne)
	t.Run("AwcEvents", testAwcEventsOne)
	t.Run("Calibrations", testCalibrationsOne)
	t.Run("DoserComponents", testDoserComponentsOne)
	t.Run("Dosers", testDosersOne)
	t.Run("Doses", testDosesOne)
	t.Run("Firmatas", testFirmatasOne)
	t.Run("GorpMigrations", testGorpMigrationsOne)
	t.Run("Pumps", testPumpsOne)
	t.Run("WaterLevelSensors", testWaterLevelSensorsOne)
}

func TestAll(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsAll)
	t.Run("AutoTopOffs", testAutoTopOffsAll)
	t.Run("AutoWaterChanges", testAutoWaterChangesAll)
	t.Run("AwcEvents", testAwcEventsAll)
	t.Run("Calibrations", testCalibrationsAll)
	t.Run("DoserComponents", testDoserComponentsAll)
	t.Run("Dosers", testDosersAll)
	t.Run("Doses", testDosesAll)
	t.Run("Firmatas", testFirmatasAll)
	t.Run("GorpMigrations", testGorpMigrationsAll)
	t.Run("Pumps", testPumpsAll)
	t.Run("WaterLevelSensors", testWaterLevelSensorsAll)
}

func TestCount(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsCount)
	t.Run("AutoTopOffs", testAutoTopOffsCount)
	t.Run("AutoWaterChanges", testAutoWaterChangesCount)
	t.Run("AwcEvents", testAwcEventsCount)
	t.Run("Calibrations", testCalibrationsCount)
	t.Run("DoserComponents", testDoserComponentsCount)
	t.Run("Dosers", testDosersCount)
	t.Run("Doses", testDosesCount)
	t.Run("Firmatas", testFirmatasCount)
	t.Run("GorpMigrations", testGorpMigrationsCount)
	t.Run("Pumps", testPumpsCount)
	t.Run("WaterLevelSensors", testWaterLevelSensorsCount)
}

func TestHooks(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsHooks)
	t.Run("AutoTopOffs", testAutoTopOffsHooks)
	t.Run("AutoWaterChanges", testAutoWaterChangesHooks)
	t.Run("AwcEvents", testAwcEventsHooks)
	t.Run("Calibrations", testCalibrationsHooks)
	t.Run("DoserComponents", testDoserComponentsHooks)
	t.Run("Dosers", testDosersHooks)
	t.Run("Doses", testDosesHooks)
	t.Run("Firmatas", testFirmatasHooks)
	t.Run("GorpMigrations", testGorpMigrationsHooks)
	t.Run("Pumps", testPumpsHooks)
	t.Run("WaterLevelSensors", testWaterLevelSensorsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsInsert)
	t.Run("AtoEvents", testAtoEventsInsertWhitelist)
	t.Run("AutoTopOffs", testAutoTopOffsInsert)
	t.Run("AutoTopOffs", testAutoTopOffsInsertWhitelist)
	t.Run("AutoWaterChanges", testAutoWaterChangesInsert)
	t.Run("AutoWaterChanges", testAutoWaterChangesInsertWhitelist)
	t.Run("AwcEvents", testAwcEventsInsert)
	t.Run("AwcEvents", testAwcEventsInsertWhitelist)
	t.Run("Calibrations", testCalibrationsInsert)
	t.Run("Calibrations", testCalibrationsInsertWhitelist)
	t.Run("DoserComponents", testDoserComponentsInsert)
	t.Run("DoserComponents", testDoserComponentsInsertWhitelist)
	t.Run("Dosers", testDosersInsert)
	t.Run("Dosers", testDosersInsertWhitelist)
	t.Run("Doses", testDosesInsert)
	t.Run("Doses", testDosesInsertWhitelist)
	t.Run("Firmatas", testFirmatasInsert)
	t.Run("Firmatas", testFirmatasInsertWhitelist)
	t.Run("GorpMigrations", testGorpMigrationsInsert)
	t.Run("GorpMigrations", testGorpMigrationsInsertWhitelist)
	t.Run("Pumps", testPumpsInsert)
	t.Run("Pumps", testPumpsInsertWhitelist)
	t.Run("WaterLevelSensors", testWaterLevelSensorsInsert)
	t.Run("WaterLevelSensors", testWaterLevelSensorsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("AtoEventToAutoTopOffUsingAutoTopOff", testAtoEventToOneAutoTopOffUsingAutoTopOff)
	t.Run("AutoTopOffToPumpUsingPump", testAutoTopOffToOnePumpUsingPump)
	t.Run("AutoWaterChangeToPumpUsingWastePump", testAutoWaterChangeToOnePumpUsingWastePump)
	t.Run("AutoWaterChangeToPumpUsingFreshPump", testAutoWaterChangeToOnePumpUsingFreshPump)
	t.Run("AwcEventToAutoWaterChangeUsingAutoWaterChange", testAwcEventToOneAutoWaterChangeUsingAutoWaterChange)
	t.Run("CalibrationToPumpUsingPump", testCalibrationToOnePumpUsingPump)
	t.Run("DoserComponentToPumpUsingPump", testDoserComponentToOnePumpUsingPump)
	t.Run("DoserComponentToDoserUsingDoser", testDoserComponentToOneDoserUsingDoser)
	t.Run("DoseToPumpUsingPump", testDoseToOnePumpUsingPump)
	t.Run("PumpToFirmataUsingFirmatum", testPumpToOneFirmataUsingFirmatum)
	t.Run("WaterLevelSensorToFirmataUsingFirmatum", testWaterLevelSensorToOneFirmataUsingFirmatum)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AutoTopOffToAtoEvents", testAutoTopOffToManyAtoEvents)
	t.Run("AutoTopOffToWaterLevelSensors", testAutoTopOffToManyWaterLevelSensors)
	t.Run("AutoWaterChangeToAwcEvents", testAutoWaterChangeToManyAwcEvents)
	t.Run("DoserToDoserComponents", testDoserToManyDoserComponents)
	t.Run("FirmataToFirmatumPumps", testFirmataToManyFirmatumPumps)
	t.Run("FirmataToFirmatumWaterLevelSensors", testFirmataToManyFirmatumWaterLevelSensors)
	t.Run("PumpToAutoTopOffs", testPumpToManyAutoTopOffs)
	t.Run("PumpToWastePumpAutoWaterChanges", testPumpToManyWastePumpAutoWaterChanges)
	t.Run("PumpToFreshPumpAutoWaterChanges", testPumpToManyFreshPumpAutoWaterChanges)
	t.Run("PumpToCalibrations", testPumpToManyCalibrations)
	t.Run("PumpToDoserComponents", testPumpToManyDoserComponents)
	t.Run("PumpToDoses", testPumpToManyDoses)
	t.Run("WaterLevelSensorToAutoTopOffs", testWaterLevelSensorToManyAutoTopOffs)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("AtoEventToAutoTopOffUsingAtoEvents", testAtoEventToOneSetOpAutoTopOffUsingAutoTopOff)
	t.Run("AutoTopOffToPumpUsingAutoTopOffs", testAutoTopOffToOneSetOpPumpUsingPump)
	t.Run("AutoWaterChangeToPumpUsingWastePumpAutoWaterChanges", testAutoWaterChangeToOneSetOpPumpUsingWastePump)
	t.Run("AutoWaterChangeToPumpUsingFreshPumpAutoWaterChanges", testAutoWaterChangeToOneSetOpPumpUsingFreshPump)
	t.Run("AwcEventToAutoWaterChangeUsingAwcEvents", testAwcEventToOneSetOpAutoWaterChangeUsingAutoWaterChange)
	t.Run("CalibrationToPumpUsingCalibrations", testCalibrationToOneSetOpPumpUsingPump)
	t.Run("DoserComponentToPumpUsingDoserComponents", testDoserComponentToOneSetOpPumpUsingPump)
	t.Run("DoserComponentToDoserUsingDoserComponents", testDoserComponentToOneSetOpDoserUsingDoser)
	t.Run("DoseToPumpUsingDoses", testDoseToOneSetOpPumpUsingPump)
	t.Run("PumpToFirmataUsingFirmatumPumps", testPumpToOneSetOpFirmataUsingFirmatum)
	t.Run("WaterLevelSensorToFirmataUsingFirmatumWaterLevelSensors", testWaterLevelSensorToOneSetOpFirmataUsingFirmatum)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("WaterLevelSensorToFirmataUsingFirmatumWaterLevelSensors", testWaterLevelSensorToOneRemoveOpFirmataUsingFirmatum)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AutoTopOffToAtoEvents", testAutoTopOffToManyAddOpAtoEvents)
	t.Run("AutoTopOffToWaterLevelSensors", testAutoTopOffToManyAddOpWaterLevelSensors)
	t.Run("AutoWaterChangeToAwcEvents", testAutoWaterChangeToManyAddOpAwcEvents)
	t.Run("DoserToDoserComponents", testDoserToManyAddOpDoserComponents)
	t.Run("FirmataToFirmatumPumps", testFirmataToManyAddOpFirmatumPumps)
	t.Run("FirmataToFirmatumWaterLevelSensors", testFirmataToManyAddOpFirmatumWaterLevelSensors)
	t.Run("PumpToAutoTopOffs", testPumpToManyAddOpAutoTopOffs)
	t.Run("PumpToWastePumpAutoWaterChanges", testPumpToManyAddOpWastePumpAutoWaterChanges)
	t.Run("PumpToFreshPumpAutoWaterChanges", testPumpToManyAddOpFreshPumpAutoWaterChanges)
	t.Run("PumpToCalibrations", testPumpToManyAddOpCalibrations)
	t.Run("PumpToDoserComponents", testPumpToManyAddOpDoserComponents)
	t.Run("PumpToDoses", testPumpToManyAddOpDoses)
	t.Run("WaterLevelSensorToAutoTopOffs", testWaterLevelSensorToManyAddOpAutoTopOffs)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("AutoTopOffToWaterLevelSensors", testAutoTopOffToManySetOpWaterLevelSensors)
	t.Run("FirmataToFirmatumWaterLevelSensors", testFirmataToManySetOpFirmatumWaterLevelSensors)
	t.Run("WaterLevelSensorToAutoTopOffs", testWaterLevelSensorToManySetOpAutoTopOffs)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("AutoTopOffToWaterLevelSensors", testAutoTopOffToManyRemoveOpWaterLevelSensors)
	t.Run("FirmataToFirmatumWaterLevelSensors", testFirmataToManyRemoveOpFirmatumWaterLevelSensors)
	t.Run("WaterLevelSensorToAutoTopOffs", testWaterLevelSensorToManyRemoveOpAutoTopOffs)
}

func TestReload(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsReload)
	t.Run("AutoTopOffs", testAutoTopOffsReload)
	t.Run("AutoWaterChanges", testAutoWaterChangesReload)
	t.Run("AwcEvents", testAwcEventsReload)
	t.Run("Calibrations", testCalibrationsReload)
	t.Run("DoserComponents", testDoserComponentsReload)
	t.Run("Dosers", testDosersReload)
	t.Run("Doses", testDosesReload)
	t.Run("Firmatas", testFirmatasReload)
	t.Run("GorpMigrations", testGorpMigrationsReload)
	t.Run("Pumps", testPumpsReload)
	t.Run("WaterLevelSensors", testWaterLevelSensorsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsReloadAll)
	t.Run("AutoTopOffs", testAutoTopOffsReloadAll)
	t.Run("AutoWaterChanges", testAutoWaterChangesReloadAll)
	t.Run("AwcEvents", testAwcEventsReloadAll)
	t.Run("Calibrations", testCalibrationsReloadAll)
	t.Run("DoserComponents", testDoserComponentsReloadAll)
	t.Run("Dosers", testDosersReloadAll)
	t.Run("Doses", testDosesReloadAll)
	t.Run("Firmatas", testFirmatasReloadAll)
	t.Run("GorpMigrations", testGorpMigrationsReloadAll)
	t.Run("Pumps", testPumpsReloadAll)
	t.Run("WaterLevelSensors", testWaterLevelSensorsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsSelect)
	t.Run("AutoTopOffs", testAutoTopOffsSelect)
	t.Run("AutoWaterChanges", testAutoWaterChangesSelect)
	t.Run("AwcEvents", testAwcEventsSelect)
	t.Run("Calibrations", testCalibrationsSelect)
	t.Run("DoserComponents", testDoserComponentsSelect)
	t.Run("Dosers", testDosersSelect)
	t.Run("Doses", testDosesSelect)
	t.Run("Firmatas", testFirmatasSelect)
	t.Run("GorpMigrations", testGorpMigrationsSelect)
	t.Run("Pumps", testPumpsSelect)
	t.Run("WaterLevelSensors", testWaterLevelSensorsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsUpdate)
	t.Run("AutoTopOffs", testAutoTopOffsUpdate)
	t.Run("AutoWaterChanges", testAutoWaterChangesUpdate)
	t.Run("AwcEvents", testAwcEventsUpdate)
	t.Run("Calibrations", testCalibrationsUpdate)
	t.Run("DoserComponents", testDoserComponentsUpdate)
	t.Run("Dosers", testDosersUpdate)
	t.Run("Doses", testDosesUpdate)
	t.Run("Firmatas", testFirmatasUpdate)
	t.Run("GorpMigrations", testGorpMigrationsUpdate)
	t.Run("Pumps", testPumpsUpdate)
	t.Run("WaterLevelSensors", testWaterLevelSensorsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("AtoEvents", testAtoEventsSliceUpdateAll)
	t.Run("AutoTopOffs", testAutoTopOffsSliceUpdateAll)
	t.Run("AutoWaterChanges", testAutoWaterChangesSliceUpdateAll)
	t.Run("AwcEvents", testAwcEventsSliceUpdateAll)
	t.Run("Calibrations", testCalibrationsSliceUpdateAll)
	t.Run("DoserComponents", testDoserComponentsSliceUpdateAll)
	t.Run("Dosers", testDosersSliceUpdateAll)
	t.Run("Doses", testDosesSliceUpdateAll)
	t.Run("Firmatas", testFirmatasSliceUpdateAll)
	t.Run("GorpMigrations", testGorpMigrationsSliceUpdateAll)
	t.Run("Pumps", testPumpsSliceUpdateAll)
	t.Run("WaterLevelSensors", testWaterLevelSensorsSliceUpdateAll)
}
