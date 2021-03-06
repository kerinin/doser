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

func testAwcEvents(t *testing.T) {
	t.Parallel()

	query := AwcEvents()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAwcEventsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AwcEvent{}
	if err = randomize.Struct(seed, o, awcEventDBTypes, true, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
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

	count, err := AwcEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAwcEventsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AwcEvent{}
	if err = randomize.Struct(seed, o, awcEventDBTypes, true, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AwcEvents().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AwcEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAwcEventsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AwcEvent{}
	if err = randomize.Struct(seed, o, awcEventDBTypes, true, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AwcEventSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AwcEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAwcEventsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AwcEvent{}
	if err = randomize.Struct(seed, o, awcEventDBTypes, true, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AwcEventExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if AwcEvent exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AwcEventExists to return true, but got false.")
	}
}

func testAwcEventsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AwcEvent{}
	if err = randomize.Struct(seed, o, awcEventDBTypes, true, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	awcEventFound, err := FindAwcEvent(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if awcEventFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAwcEventsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AwcEvent{}
	if err = randomize.Struct(seed, o, awcEventDBTypes, true, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AwcEvents().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAwcEventsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AwcEvent{}
	if err = randomize.Struct(seed, o, awcEventDBTypes, true, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AwcEvents().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAwcEventsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	awcEventOne := &AwcEvent{}
	awcEventTwo := &AwcEvent{}
	if err = randomize.Struct(seed, awcEventOne, awcEventDBTypes, false, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}
	if err = randomize.Struct(seed, awcEventTwo, awcEventDBTypes, false, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = awcEventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = awcEventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AwcEvents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAwcEventsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	awcEventOne := &AwcEvent{}
	awcEventTwo := &AwcEvent{}
	if err = randomize.Struct(seed, awcEventOne, awcEventDBTypes, false, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}
	if err = randomize.Struct(seed, awcEventTwo, awcEventDBTypes, false, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = awcEventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = awcEventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AwcEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func awcEventBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AwcEvent) error {
	*o = AwcEvent{}
	return nil
}

func awcEventAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AwcEvent) error {
	*o = AwcEvent{}
	return nil
}

func awcEventAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AwcEvent) error {
	*o = AwcEvent{}
	return nil
}

func awcEventBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AwcEvent) error {
	*o = AwcEvent{}
	return nil
}

func awcEventAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AwcEvent) error {
	*o = AwcEvent{}
	return nil
}

func awcEventBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AwcEvent) error {
	*o = AwcEvent{}
	return nil
}

func awcEventAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AwcEvent) error {
	*o = AwcEvent{}
	return nil
}

func awcEventBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AwcEvent) error {
	*o = AwcEvent{}
	return nil
}

func awcEventAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AwcEvent) error {
	*o = AwcEvent{}
	return nil
}

func testAwcEventsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AwcEvent{}
	o := &AwcEvent{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, awcEventDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AwcEvent object: %s", err)
	}

	AddAwcEventHook(boil.BeforeInsertHook, awcEventBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	awcEventBeforeInsertHooks = []AwcEventHook{}

	AddAwcEventHook(boil.AfterInsertHook, awcEventAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	awcEventAfterInsertHooks = []AwcEventHook{}

	AddAwcEventHook(boil.AfterSelectHook, awcEventAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	awcEventAfterSelectHooks = []AwcEventHook{}

	AddAwcEventHook(boil.BeforeUpdateHook, awcEventBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	awcEventBeforeUpdateHooks = []AwcEventHook{}

	AddAwcEventHook(boil.AfterUpdateHook, awcEventAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	awcEventAfterUpdateHooks = []AwcEventHook{}

	AddAwcEventHook(boil.BeforeDeleteHook, awcEventBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	awcEventBeforeDeleteHooks = []AwcEventHook{}

	AddAwcEventHook(boil.AfterDeleteHook, awcEventAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	awcEventAfterDeleteHooks = []AwcEventHook{}

	AddAwcEventHook(boil.BeforeUpsertHook, awcEventBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	awcEventBeforeUpsertHooks = []AwcEventHook{}

	AddAwcEventHook(boil.AfterUpsertHook, awcEventAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	awcEventAfterUpsertHooks = []AwcEventHook{}
}

func testAwcEventsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AwcEvent{}
	if err = randomize.Struct(seed, o, awcEventDBTypes, true, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AwcEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAwcEventsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AwcEvent{}
	if err = randomize.Struct(seed, o, awcEventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(awcEventColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AwcEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAwcEventToOneAutoWaterChangeUsingAutoWaterChange(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local AwcEvent
	var foreign AutoWaterChange

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, awcEventDBTypes, false, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, autoWaterChangeDBTypes, false, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.AutoWaterChangeID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.AutoWaterChange().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AwcEventSlice{&local}
	if err = local.L.LoadAutoWaterChange(ctx, tx, false, (*[]*AwcEvent)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.AutoWaterChange == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.AutoWaterChange = nil
	if err = local.L.LoadAutoWaterChange(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.AutoWaterChange == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAwcEventToOneSetOpAutoWaterChangeUsingAutoWaterChange(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AwcEvent
	var b, c AutoWaterChange

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, awcEventDBTypes, false, strmangle.SetComplement(awcEventPrimaryKeyColumns, awcEventColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, autoWaterChangeDBTypes, false, strmangle.SetComplement(autoWaterChangePrimaryKeyColumns, autoWaterChangeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, autoWaterChangeDBTypes, false, strmangle.SetComplement(autoWaterChangePrimaryKeyColumns, autoWaterChangeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*AutoWaterChange{&b, &c} {
		err = a.SetAutoWaterChange(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.AutoWaterChange != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AwcEvents[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AutoWaterChangeID != x.ID {
			t.Error("foreign key was wrong value", a.AutoWaterChangeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AutoWaterChangeID))
		reflect.Indirect(reflect.ValueOf(&a.AutoWaterChangeID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AutoWaterChangeID != x.ID {
			t.Error("foreign key was wrong value", a.AutoWaterChangeID, x.ID)
		}
	}
}

func testAwcEventsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AwcEvent{}
	if err = randomize.Struct(seed, o, awcEventDBTypes, true, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
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

func testAwcEventsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AwcEvent{}
	if err = randomize.Struct(seed, o, awcEventDBTypes, true, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AwcEventSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAwcEventsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AwcEvent{}
	if err = randomize.Struct(seed, o, awcEventDBTypes, true, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AwcEvents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	awcEventDBTypes = map[string]string{`ID`: `UUID`, `AutoWaterChangeID`: `UUID`, `Timestamp`: `INT`, `Kind`: `STRING`, `Data`: `STRING`}
	_               = bytes.MinRead
)

func testAwcEventsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(awcEventPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(awcEventAllColumns) == len(awcEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AwcEvent{}
	if err = randomize.Struct(seed, o, awcEventDBTypes, true, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AwcEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, awcEventDBTypes, true, awcEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAwcEventsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(awcEventAllColumns) == len(awcEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AwcEvent{}
	if err = randomize.Struct(seed, o, awcEventDBTypes, true, awcEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AwcEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, awcEventDBTypes, true, awcEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AwcEvent struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(awcEventAllColumns, awcEventPrimaryKeyColumns) {
		fields = awcEventAllColumns
	} else {
		fields = strmangle.SetComplement(
			awcEventAllColumns,
			awcEventPrimaryKeyColumns,
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

	slice := AwcEventSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
