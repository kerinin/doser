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

func testAutoTopOffs(t *testing.T) {
	t.Parallel()

	query := AutoTopOffs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAutoTopOffsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoTopOff{}
	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
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

	count, err := AutoTopOffs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAutoTopOffsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoTopOff{}
	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AutoTopOffs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AutoTopOffs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAutoTopOffsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoTopOff{}
	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AutoTopOffSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AutoTopOffs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAutoTopOffsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoTopOff{}
	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AutoTopOffExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if AutoTopOff exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AutoTopOffExists to return true, but got false.")
	}
}

func testAutoTopOffsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoTopOff{}
	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	autoTopOffFound, err := FindAutoTopOff(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if autoTopOffFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAutoTopOffsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoTopOff{}
	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AutoTopOffs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAutoTopOffsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoTopOff{}
	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AutoTopOffs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAutoTopOffsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	autoTopOffOne := &AutoTopOff{}
	autoTopOffTwo := &AutoTopOff{}
	if err = randomize.Struct(seed, autoTopOffOne, autoTopOffDBTypes, false, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}
	if err = randomize.Struct(seed, autoTopOffTwo, autoTopOffDBTypes, false, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = autoTopOffOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = autoTopOffTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AutoTopOffs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAutoTopOffsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	autoTopOffOne := &AutoTopOff{}
	autoTopOffTwo := &AutoTopOff{}
	if err = randomize.Struct(seed, autoTopOffOne, autoTopOffDBTypes, false, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}
	if err = randomize.Struct(seed, autoTopOffTwo, autoTopOffDBTypes, false, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = autoTopOffOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = autoTopOffTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AutoTopOffs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func autoTopOffBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AutoTopOff) error {
	*o = AutoTopOff{}
	return nil
}

func autoTopOffAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AutoTopOff) error {
	*o = AutoTopOff{}
	return nil
}

func autoTopOffAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AutoTopOff) error {
	*o = AutoTopOff{}
	return nil
}

func autoTopOffBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AutoTopOff) error {
	*o = AutoTopOff{}
	return nil
}

func autoTopOffAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AutoTopOff) error {
	*o = AutoTopOff{}
	return nil
}

func autoTopOffBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AutoTopOff) error {
	*o = AutoTopOff{}
	return nil
}

func autoTopOffAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AutoTopOff) error {
	*o = AutoTopOff{}
	return nil
}

func autoTopOffBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AutoTopOff) error {
	*o = AutoTopOff{}
	return nil
}

func autoTopOffAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AutoTopOff) error {
	*o = AutoTopOff{}
	return nil
}

func testAutoTopOffsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AutoTopOff{}
	o := &AutoTopOff{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, autoTopOffDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AutoTopOff object: %s", err)
	}

	AddAutoTopOffHook(boil.BeforeInsertHook, autoTopOffBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	autoTopOffBeforeInsertHooks = []AutoTopOffHook{}

	AddAutoTopOffHook(boil.AfterInsertHook, autoTopOffAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	autoTopOffAfterInsertHooks = []AutoTopOffHook{}

	AddAutoTopOffHook(boil.AfterSelectHook, autoTopOffAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	autoTopOffAfterSelectHooks = []AutoTopOffHook{}

	AddAutoTopOffHook(boil.BeforeUpdateHook, autoTopOffBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	autoTopOffBeforeUpdateHooks = []AutoTopOffHook{}

	AddAutoTopOffHook(boil.AfterUpdateHook, autoTopOffAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	autoTopOffAfterUpdateHooks = []AutoTopOffHook{}

	AddAutoTopOffHook(boil.BeforeDeleteHook, autoTopOffBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	autoTopOffBeforeDeleteHooks = []AutoTopOffHook{}

	AddAutoTopOffHook(boil.AfterDeleteHook, autoTopOffAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	autoTopOffAfterDeleteHooks = []AutoTopOffHook{}

	AddAutoTopOffHook(boil.BeforeUpsertHook, autoTopOffBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	autoTopOffBeforeUpsertHooks = []AutoTopOffHook{}

	AddAutoTopOffHook(boil.AfterUpsertHook, autoTopOffAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	autoTopOffAfterUpsertHooks = []AutoTopOffHook{}
}

func testAutoTopOffsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoTopOff{}
	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AutoTopOffs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAutoTopOffsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoTopOff{}
	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(autoTopOffColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AutoTopOffs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAutoTopOffToManyAtoEvents(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AutoTopOff
	var b, c AtoEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, autoTopOffDBTypes, true, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, atoEventDBTypes, false, atoEventColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, atoEventDBTypes, false, atoEventColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.AutoTopOffID = a.ID
	c.AutoTopOffID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.AtoEvents().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.AutoTopOffID == b.AutoTopOffID {
			bFound = true
		}
		if v.AutoTopOffID == c.AutoTopOffID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AutoTopOffSlice{&a}
	if err = a.L.LoadAtoEvents(ctx, tx, false, (*[]*AutoTopOff)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AtoEvents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AtoEvents = nil
	if err = a.L.LoadAtoEvents(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AtoEvents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAutoTopOffToManyWaterLevelSensors(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AutoTopOff
	var b, c WaterLevelSensor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, autoTopOffDBTypes, true, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, waterLevelSensorDBTypes, false, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, waterLevelSensorDBTypes, false, waterLevelSensorColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"auto_top_offs_water_level_sensors\" (\"auto_top_off_id\", \"water_level_sensor_id\") values (?, ?)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"auto_top_offs_water_level_sensors\" (\"auto_top_off_id\", \"water_level_sensor_id\") values (?, ?)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.WaterLevelSensors().All(ctx, tx)
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

	slice := AutoTopOffSlice{&a}
	if err = a.L.LoadWaterLevelSensors(ctx, tx, false, (*[]*AutoTopOff)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.WaterLevelSensors); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.WaterLevelSensors = nil
	if err = a.L.LoadWaterLevelSensors(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.WaterLevelSensors); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAutoTopOffToManyAddOpAtoEvents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AutoTopOff
	var b, c, d, e AtoEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, autoTopOffDBTypes, false, strmangle.SetComplement(autoTopOffPrimaryKeyColumns, autoTopOffColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AtoEvent{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, atoEventDBTypes, false, strmangle.SetComplement(atoEventPrimaryKeyColumns, atoEventColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*AtoEvent{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAtoEvents(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.AutoTopOffID {
			t.Error("foreign key was wrong value", a.ID, first.AutoTopOffID)
		}
		if a.ID != second.AutoTopOffID {
			t.Error("foreign key was wrong value", a.ID, second.AutoTopOffID)
		}

		if first.R.AutoTopOff != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.AutoTopOff != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AtoEvents[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AtoEvents[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AtoEvents().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAutoTopOffToManyAddOpWaterLevelSensors(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AutoTopOff
	var b, c, d, e WaterLevelSensor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, autoTopOffDBTypes, false, strmangle.SetComplement(autoTopOffPrimaryKeyColumns, autoTopOffColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*WaterLevelSensor{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, waterLevelSensorDBTypes, false, strmangle.SetComplement(waterLevelSensorPrimaryKeyColumns, waterLevelSensorColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*WaterLevelSensor{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddWaterLevelSensors(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.AutoTopOffs[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.AutoTopOffs[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.WaterLevelSensors[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.WaterLevelSensors[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.WaterLevelSensors().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testAutoTopOffToManySetOpWaterLevelSensors(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AutoTopOff
	var b, c, d, e WaterLevelSensor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, autoTopOffDBTypes, false, strmangle.SetComplement(autoTopOffPrimaryKeyColumns, autoTopOffColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*WaterLevelSensor{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, waterLevelSensorDBTypes, false, strmangle.SetComplement(waterLevelSensorPrimaryKeyColumns, waterLevelSensorColumnsWithoutDefault)...); err != nil {
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

	err = a.SetWaterLevelSensors(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.WaterLevelSensors().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetWaterLevelSensors(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.WaterLevelSensors().Count(ctx, tx)
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
	// if len(b.R.AutoTopOffs) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.AutoTopOffs) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.AutoTopOffs[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.AutoTopOffs[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.WaterLevelSensors[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.WaterLevelSensors[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testAutoTopOffToManyRemoveOpWaterLevelSensors(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AutoTopOff
	var b, c, d, e WaterLevelSensor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, autoTopOffDBTypes, false, strmangle.SetComplement(autoTopOffPrimaryKeyColumns, autoTopOffColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*WaterLevelSensor{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, waterLevelSensorDBTypes, false, strmangle.SetComplement(waterLevelSensorPrimaryKeyColumns, waterLevelSensorColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddWaterLevelSensors(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.WaterLevelSensors().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveWaterLevelSensors(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.WaterLevelSensors().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.AutoTopOffs) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.AutoTopOffs) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.AutoTopOffs[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.AutoTopOffs[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.WaterLevelSensors) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.WaterLevelSensors[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.WaterLevelSensors[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testAutoTopOffToOnePumpUsingPump(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local AutoTopOff
	var foreign Pump

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, autoTopOffDBTypes, false, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, pumpDBTypes, false, pumpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pump struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PumpID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Pump().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AutoTopOffSlice{&local}
	if err = local.L.LoadPump(ctx, tx, false, (*[]*AutoTopOff)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Pump == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Pump = nil
	if err = local.L.LoadPump(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Pump == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAutoTopOffToOneSetOpPumpUsingPump(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AutoTopOff
	var b, c Pump

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, autoTopOffDBTypes, false, strmangle.SetComplement(autoTopOffPrimaryKeyColumns, autoTopOffColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pumpDBTypes, false, strmangle.SetComplement(pumpPrimaryKeyColumns, pumpColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pumpDBTypes, false, strmangle.SetComplement(pumpPrimaryKeyColumns, pumpColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Pump{&b, &c} {
		err = a.SetPump(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Pump != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AutoTopOffs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PumpID != x.ID {
			t.Error("foreign key was wrong value", a.PumpID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PumpID))
		reflect.Indirect(reflect.ValueOf(&a.PumpID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PumpID != x.ID {
			t.Error("foreign key was wrong value", a.PumpID, x.ID)
		}
	}
}

func testAutoTopOffsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoTopOff{}
	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
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

func testAutoTopOffsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoTopOff{}
	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AutoTopOffSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAutoTopOffsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoTopOff{}
	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AutoTopOffs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	autoTopOffDBTypes = map[string]string{`ID`: `UUID`, `PumpID`: `UUID`, `FillRate`: `REAL`, `FillInterval`: `INT`, `MaxFillVolume`: `REAL`}
	_                 = bytes.MinRead
)

func testAutoTopOffsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(autoTopOffPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(autoTopOffAllColumns) == len(autoTopOffPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AutoTopOff{}
	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AutoTopOffs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true, autoTopOffPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAutoTopOffsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(autoTopOffAllColumns) == len(autoTopOffPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AutoTopOff{}
	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true, autoTopOffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AutoTopOffs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, autoTopOffDBTypes, true, autoTopOffPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AutoTopOff struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(autoTopOffAllColumns, autoTopOffPrimaryKeyColumns) {
		fields = autoTopOffAllColumns
	} else {
		fields = strmangle.SetComplement(
			autoTopOffAllColumns,
			autoTopOffPrimaryKeyColumns,
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

	slice := AutoTopOffSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
