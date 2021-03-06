// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testWaterLevelSensors(t *testing.T) {
	t.Parallel()

	query := WaterLevelSensors()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testWaterLevelSensorsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WaterLevelSensor{}
	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WaterLevelSensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWaterLevelSensorsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WaterLevelSensor{}
	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := WaterLevelSensors().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WaterLevelSensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWaterLevelSensorsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WaterLevelSensor{}
	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WaterLevelSensorSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WaterLevelSensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWaterLevelSensorsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WaterLevelSensor{}
	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := WaterLevelSensorExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if WaterLevelSensor exists: %s", err)
	}
	if !e {
		t.Errorf("Expected WaterLevelSensorExists to return true, but got false.")
	}
}

func testWaterLevelSensorsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WaterLevelSensor{}
	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	waterLevelSensorFound, err := FindWaterLevelSensor(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if waterLevelSensorFound == nil {
		t.Error("want a record, got nil")
	}
}

func testWaterLevelSensorsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WaterLevelSensor{}
	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = WaterLevelSensors().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testWaterLevelSensorsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WaterLevelSensor{}
	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := WaterLevelSensors().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testWaterLevelSensorsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	waterLevelSensorOne := &WaterLevelSensor{}
	waterLevelSensorTwo := &WaterLevelSensor{}
	if err = randomize.Struct(seed, waterLevelSensorOne, waterLevelSensorDBTypes, false, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}
	if err = randomize.Struct(seed, waterLevelSensorTwo, waterLevelSensorDBTypes, false, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = waterLevelSensorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = waterLevelSensorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := WaterLevelSensors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testWaterLevelSensorsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	waterLevelSensorOne := &WaterLevelSensor{}
	waterLevelSensorTwo := &WaterLevelSensor{}
	if err = randomize.Struct(seed, waterLevelSensorOne, waterLevelSensorDBTypes, false, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}
	if err = randomize.Struct(seed, waterLevelSensorTwo, waterLevelSensorDBTypes, false, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = waterLevelSensorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = waterLevelSensorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WaterLevelSensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func waterLevelSensorBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *WaterLevelSensor) error {
	*o = WaterLevelSensor{}
	return nil
}

func waterLevelSensorAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *WaterLevelSensor) error {
	*o = WaterLevelSensor{}
	return nil
}

func waterLevelSensorAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *WaterLevelSensor) error {
	*o = WaterLevelSensor{}
	return nil
}

func waterLevelSensorBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *WaterLevelSensor) error {
	*o = WaterLevelSensor{}
	return nil
}

func waterLevelSensorAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *WaterLevelSensor) error {
	*o = WaterLevelSensor{}
	return nil
}

func waterLevelSensorBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *WaterLevelSensor) error {
	*o = WaterLevelSensor{}
	return nil
}

func waterLevelSensorAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *WaterLevelSensor) error {
	*o = WaterLevelSensor{}
	return nil
}

func waterLevelSensorBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *WaterLevelSensor) error {
	*o = WaterLevelSensor{}
	return nil
}

func waterLevelSensorAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *WaterLevelSensor) error {
	*o = WaterLevelSensor{}
	return nil
}

func testWaterLevelSensorsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &WaterLevelSensor{}
	o := &WaterLevelSensor{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, false); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor object: %s", err)
	}

	AddWaterLevelSensorHook(boil.BeforeInsertHook, waterLevelSensorBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	waterLevelSensorBeforeInsertHooks = []WaterLevelSensorHook{}

	AddWaterLevelSensorHook(boil.AfterInsertHook, waterLevelSensorAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	waterLevelSensorAfterInsertHooks = []WaterLevelSensorHook{}

	AddWaterLevelSensorHook(boil.AfterSelectHook, waterLevelSensorAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	waterLevelSensorAfterSelectHooks = []WaterLevelSensorHook{}

	AddWaterLevelSensorHook(boil.BeforeUpdateHook, waterLevelSensorBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	waterLevelSensorBeforeUpdateHooks = []WaterLevelSensorHook{}

	AddWaterLevelSensorHook(boil.AfterUpdateHook, waterLevelSensorAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	waterLevelSensorAfterUpdateHooks = []WaterLevelSensorHook{}

	AddWaterLevelSensorHook(boil.BeforeDeleteHook, waterLevelSensorBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	waterLevelSensorBeforeDeleteHooks = []WaterLevelSensorHook{}

	AddWaterLevelSensorHook(boil.AfterDeleteHook, waterLevelSensorAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	waterLevelSensorAfterDeleteHooks = []WaterLevelSensorHook{}

	AddWaterLevelSensorHook(boil.BeforeUpsertHook, waterLevelSensorBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	waterLevelSensorBeforeUpsertHooks = []WaterLevelSensorHook{}

	AddWaterLevelSensorHook(boil.AfterUpsertHook, waterLevelSensorAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	waterLevelSensorAfterUpsertHooks = []WaterLevelSensorHook{}
}

func testWaterLevelSensorsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WaterLevelSensor{}
	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WaterLevelSensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWaterLevelSensorsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WaterLevelSensor{}
	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(waterLevelSensorColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := WaterLevelSensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWaterLevelSensorToManyAutoTopOffs(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WaterLevelSensor
	var b, c AutoTopOff

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, waterLevelSensorDBTypes, true, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, autoTopOffDBTypes, false, autoTopOffColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, autoTopOffDBTypes, false, autoTopOffColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"auto_top_offs_water_level_sensors\" (\"water_level_sensor_id\", \"auto_top_off_id\") values (?, ?)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"auto_top_offs_water_level_sensors\" (\"water_level_sensor_id\", \"auto_top_off_id\") values (?, ?)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.AutoTopOffs().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := WaterLevelSensorSlice{&a}
	if err = a.L.LoadAutoTopOffs(ctx, tx, false, (*[]*WaterLevelSensor)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AutoTopOffs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AutoTopOffs = nil
	if err = a.L.LoadAutoTopOffs(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AutoTopOffs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testWaterLevelSensorToManyAddOpAutoTopOffs(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WaterLevelSensor
	var b, c, d, e AutoTopOff

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, waterLevelSensorDBTypes, false, strmangle.SetComplement(waterLevelSensorPrimaryKeyColumns, waterLevelSensorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AutoTopOff{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, autoTopOffDBTypes, false, strmangle.SetComplement(autoTopOffPrimaryKeyColumns, autoTopOffColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*AutoTopOff{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAutoTopOffs(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.WaterLevelSensors[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.WaterLevelSensors[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.AutoTopOffs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AutoTopOffs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AutoTopOffs().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testWaterLevelSensorToManySetOpAutoTopOffs(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WaterLevelSensor
	var b, c, d, e AutoTopOff

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, waterLevelSensorDBTypes, false, strmangle.SetComplement(waterLevelSensorPrimaryKeyColumns, waterLevelSensorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AutoTopOff{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, autoTopOffDBTypes, false, strmangle.SetComplement(autoTopOffPrimaryKeyColumns, autoTopOffColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetAutoTopOffs(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.AutoTopOffs().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetAutoTopOffs(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.AutoTopOffs().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.WaterLevelSensors) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.WaterLevelSensors) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.WaterLevelSensors[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.WaterLevelSensors[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.AutoTopOffs[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.AutoTopOffs[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testWaterLevelSensorToManyRemoveOpAutoTopOffs(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WaterLevelSensor
	var b, c, d, e AutoTopOff

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, waterLevelSensorDBTypes, false, strmangle.SetComplement(waterLevelSensorPrimaryKeyColumns, waterLevelSensorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AutoTopOff{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, autoTopOffDBTypes, false, strmangle.SetComplement(autoTopOffPrimaryKeyColumns, autoTopOffColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddAutoTopOffs(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.AutoTopOffs().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveAutoTopOffs(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.AutoTopOffs().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.WaterLevelSensors) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.WaterLevelSensors) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.WaterLevelSensors[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.WaterLevelSensors[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.AutoTopOffs) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.AutoTopOffs[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.AutoTopOffs[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testWaterLevelSensorToOneFirmataUsingFirmatum(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local WaterLevelSensor
	var foreign Firmata

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, waterLevelSensorDBTypes, true, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, firmataDBTypes, false, firmataColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Firmata struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.FirmataID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Firmatum().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := WaterLevelSensorSlice{&local}
	if err = local.L.LoadFirmatum(ctx, tx, false, (*[]*WaterLevelSensor)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Firmatum == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Firmatum = nil
	if err = local.L.LoadFirmatum(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Firmatum == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testWaterLevelSensorToOneSetOpFirmataUsingFirmatum(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WaterLevelSensor
	var b, c Firmata

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, waterLevelSensorDBTypes, false, strmangle.SetComplement(waterLevelSensorPrimaryKeyColumns, waterLevelSensorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, firmataDBTypes, false, strmangle.SetComplement(firmataPrimaryKeyColumns, firmataColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, firmataDBTypes, false, strmangle.SetComplement(firmataPrimaryKeyColumns, firmataColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Firmata{&b, &c} {
		err = a.SetFirmatum(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Firmatum != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FirmatumWaterLevelSensors[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.FirmataID, x.ID) {
			t.Error("foreign key was wrong value", a.FirmataID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FirmataID))
		reflect.Indirect(reflect.ValueOf(&a.FirmataID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.FirmataID, x.ID) {
			t.Error("foreign key was wrong value", a.FirmataID, x.ID)
		}
	}
}

func testWaterLevelSensorToOneRemoveOpFirmataUsingFirmatum(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WaterLevelSensor
	var b Firmata

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, waterLevelSensorDBTypes, false, strmangle.SetComplement(waterLevelSensorPrimaryKeyColumns, waterLevelSensorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, firmataDBTypes, false, strmangle.SetComplement(firmataPrimaryKeyColumns, firmataColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetFirmatum(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveFirmatum(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Firmatum().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Firmatum != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.FirmataID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.FirmatumWaterLevelSensors) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testWaterLevelSensorsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WaterLevelSensor{}
	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWaterLevelSensorsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WaterLevelSensor{}
	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WaterLevelSensorSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWaterLevelSensorsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WaterLevelSensor{}
	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := WaterLevelSensors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	waterLevelSensorDBTypes = map[string]string{`ID`: `UUID`, `FirmataID`: `UUID`, `Pin`: `INT`, `Kind`: `TEXT`, `DetectionThreshold`: `INT`, `Invert`: `BOOLEAN`, `Name`: `TEXT`}
	_                       = bytes.MinRead
)

func testWaterLevelSensorsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(waterLevelSensorPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(waterLevelSensorAllColumns) == len(waterLevelSensorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &WaterLevelSensor{}
	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WaterLevelSensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true, waterLevelSensorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testWaterLevelSensorsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(waterLevelSensorAllColumns) == len(waterLevelSensorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &WaterLevelSensor{}
	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WaterLevelSensors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, waterLevelSensorDBTypes, true, waterLevelSensorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WaterLevelSensor struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(waterLevelSensorAllColumns, waterLevelSensorPrimaryKeyColumns) {
		fields = waterLevelSensorAllColumns
	} else {
		fields = strmangle.SetComplement(
			waterLevelSensorAllColumns,
			waterLevelSensorPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := WaterLevelSensorSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
