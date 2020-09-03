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

func testAutoWaterChanges(t *testing.T) {
	t.Parallel()

	query := AutoWaterChanges()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAutoWaterChangesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoWaterChange{}
	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
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

	count, err := AutoWaterChanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAutoWaterChangesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoWaterChange{}
	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AutoWaterChanges().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AutoWaterChanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAutoWaterChangesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoWaterChange{}
	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AutoWaterChangeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AutoWaterChanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAutoWaterChangesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoWaterChange{}
	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AutoWaterChangeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if AutoWaterChange exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AutoWaterChangeExists to return true, but got false.")
	}
}

func testAutoWaterChangesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoWaterChange{}
	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	autoWaterChangeFound, err := FindAutoWaterChange(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if autoWaterChangeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAutoWaterChangesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoWaterChange{}
	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AutoWaterChanges().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAutoWaterChangesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoWaterChange{}
	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AutoWaterChanges().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAutoWaterChangesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	autoWaterChangeOne := &AutoWaterChange{}
	autoWaterChangeTwo := &AutoWaterChange{}
	if err = randomize.Struct(seed, autoWaterChangeOne, autoWaterChangeDBTypes, false, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}
	if err = randomize.Struct(seed, autoWaterChangeTwo, autoWaterChangeDBTypes, false, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = autoWaterChangeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = autoWaterChangeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AutoWaterChanges().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAutoWaterChangesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	autoWaterChangeOne := &AutoWaterChange{}
	autoWaterChangeTwo := &AutoWaterChange{}
	if err = randomize.Struct(seed, autoWaterChangeOne, autoWaterChangeDBTypes, false, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}
	if err = randomize.Struct(seed, autoWaterChangeTwo, autoWaterChangeDBTypes, false, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = autoWaterChangeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = autoWaterChangeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AutoWaterChanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func autoWaterChangeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AutoWaterChange) error {
	*o = AutoWaterChange{}
	return nil
}

func autoWaterChangeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AutoWaterChange) error {
	*o = AutoWaterChange{}
	return nil
}

func autoWaterChangeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AutoWaterChange) error {
	*o = AutoWaterChange{}
	return nil
}

func autoWaterChangeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AutoWaterChange) error {
	*o = AutoWaterChange{}
	return nil
}

func autoWaterChangeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AutoWaterChange) error {
	*o = AutoWaterChange{}
	return nil
}

func autoWaterChangeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AutoWaterChange) error {
	*o = AutoWaterChange{}
	return nil
}

func autoWaterChangeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AutoWaterChange) error {
	*o = AutoWaterChange{}
	return nil
}

func autoWaterChangeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AutoWaterChange) error {
	*o = AutoWaterChange{}
	return nil
}

func autoWaterChangeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AutoWaterChange) error {
	*o = AutoWaterChange{}
	return nil
}

func testAutoWaterChangesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AutoWaterChange{}
	o := &AutoWaterChange{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange object: %s", err)
	}

	AddAutoWaterChangeHook(boil.BeforeInsertHook, autoWaterChangeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	autoWaterChangeBeforeInsertHooks = []AutoWaterChangeHook{}

	AddAutoWaterChangeHook(boil.AfterInsertHook, autoWaterChangeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	autoWaterChangeAfterInsertHooks = []AutoWaterChangeHook{}

	AddAutoWaterChangeHook(boil.AfterSelectHook, autoWaterChangeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	autoWaterChangeAfterSelectHooks = []AutoWaterChangeHook{}

	AddAutoWaterChangeHook(boil.BeforeUpdateHook, autoWaterChangeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	autoWaterChangeBeforeUpdateHooks = []AutoWaterChangeHook{}

	AddAutoWaterChangeHook(boil.AfterUpdateHook, autoWaterChangeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	autoWaterChangeAfterUpdateHooks = []AutoWaterChangeHook{}

	AddAutoWaterChangeHook(boil.BeforeDeleteHook, autoWaterChangeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	autoWaterChangeBeforeDeleteHooks = []AutoWaterChangeHook{}

	AddAutoWaterChangeHook(boil.AfterDeleteHook, autoWaterChangeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	autoWaterChangeAfterDeleteHooks = []AutoWaterChangeHook{}

	AddAutoWaterChangeHook(boil.BeforeUpsertHook, autoWaterChangeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	autoWaterChangeBeforeUpsertHooks = []AutoWaterChangeHook{}

	AddAutoWaterChangeHook(boil.AfterUpsertHook, autoWaterChangeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	autoWaterChangeAfterUpsertHooks = []AutoWaterChangeHook{}
}

func testAutoWaterChangesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoWaterChange{}
	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AutoWaterChanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAutoWaterChangesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoWaterChange{}
	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(autoWaterChangeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AutoWaterChanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAutoWaterChangeToManyAwcEvents(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AutoWaterChange
	var b, c AwcEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, autoWaterChangeDBTypes, true, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, awcEventDBTypes, false, awcEventColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, awcEventDBTypes, false, awcEventColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.AutoWaterChangeID = a.ID
	c.AutoWaterChangeID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.AwcEvents().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.AutoWaterChangeID == b.AutoWaterChangeID {
			bFound = true
		}
		if v.AutoWaterChangeID == c.AutoWaterChangeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AutoWaterChangeSlice{&a}
	if err = a.L.LoadAwcEvents(ctx, tx, false, (*[]*AutoWaterChange)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AwcEvents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AwcEvents = nil
	if err = a.L.LoadAwcEvents(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AwcEvents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAutoWaterChangeToManyAddOpAwcEvents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AutoWaterChange
	var b, c, d, e AwcEvent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, autoWaterChangeDBTypes, false, strmangle.SetComplement(autoWaterChangePrimaryKeyColumns, autoWaterChangeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AwcEvent{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, awcEventDBTypes, false, strmangle.SetComplement(awcEventPrimaryKeyColumns, awcEventColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*AwcEvent{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAwcEvents(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.AutoWaterChangeID {
			t.Error("foreign key was wrong value", a.ID, first.AutoWaterChangeID)
		}
		if a.ID != second.AutoWaterChangeID {
			t.Error("foreign key was wrong value", a.ID, second.AutoWaterChangeID)
		}

		if first.R.AutoWaterChange != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.AutoWaterChange != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AwcEvents[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AwcEvents[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AwcEvents().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAutoWaterChangeToOnePumpUsingWastePump(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local AutoWaterChange
	var foreign Pump

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, autoWaterChangeDBTypes, false, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, pumpDBTypes, false, pumpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pump struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.WastePumpID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.WastePump().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AutoWaterChangeSlice{&local}
	if err = local.L.LoadWastePump(ctx, tx, false, (*[]*AutoWaterChange)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.WastePump == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.WastePump = nil
	if err = local.L.LoadWastePump(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.WastePump == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAutoWaterChangeToOnePumpUsingFreshPump(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local AutoWaterChange
	var foreign Pump

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, autoWaterChangeDBTypes, false, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, pumpDBTypes, false, pumpColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pump struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.FreshPumpID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.FreshPump().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AutoWaterChangeSlice{&local}
	if err = local.L.LoadFreshPump(ctx, tx, false, (*[]*AutoWaterChange)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FreshPump == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FreshPump = nil
	if err = local.L.LoadFreshPump(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FreshPump == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAutoWaterChangeToOneSetOpPumpUsingWastePump(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AutoWaterChange
	var b, c Pump

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, autoWaterChangeDBTypes, false, strmangle.SetComplement(autoWaterChangePrimaryKeyColumns, autoWaterChangeColumnsWithoutDefault)...); err != nil {
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
		err = a.SetWastePump(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.WastePump != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.WastePumpAutoWaterChanges[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.WastePumpID != x.ID {
			t.Error("foreign key was wrong value", a.WastePumpID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.WastePumpID))
		reflect.Indirect(reflect.ValueOf(&a.WastePumpID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.WastePumpID != x.ID {
			t.Error("foreign key was wrong value", a.WastePumpID, x.ID)
		}
	}
}
func testAutoWaterChangeToOneSetOpPumpUsingFreshPump(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AutoWaterChange
	var b, c Pump

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, autoWaterChangeDBTypes, false, strmangle.SetComplement(autoWaterChangePrimaryKeyColumns, autoWaterChangeColumnsWithoutDefault)...); err != nil {
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
		err = a.SetFreshPump(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FreshPump != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FreshPumpAutoWaterChanges[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FreshPumpID != x.ID {
			t.Error("foreign key was wrong value", a.FreshPumpID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FreshPumpID))
		reflect.Indirect(reflect.ValueOf(&a.FreshPumpID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FreshPumpID != x.ID {
			t.Error("foreign key was wrong value", a.FreshPumpID, x.ID)
		}
	}
}

func testAutoWaterChangesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoWaterChange{}
	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
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

func testAutoWaterChangesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoWaterChange{}
	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AutoWaterChangeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAutoWaterChangesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutoWaterChange{}
	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AutoWaterChanges().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	autoWaterChangeDBTypes = map[string]string{`ID`: `UUID`, `FreshPumpID`: `UUID`, `WastePumpID`: `UUID`, `ExchangeRate`: `REAL`}
	_                      = bytes.MinRead
)

func testAutoWaterChangesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(autoWaterChangePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(autoWaterChangeAllColumns) == len(autoWaterChangePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AutoWaterChange{}
	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AutoWaterChanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true, autoWaterChangePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAutoWaterChangesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(autoWaterChangeAllColumns) == len(autoWaterChangePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AutoWaterChange{}
	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true, autoWaterChangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AutoWaterChanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, autoWaterChangeDBTypes, true, autoWaterChangePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AutoWaterChange struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(autoWaterChangeAllColumns, autoWaterChangePrimaryKeyColumns) {
		fields = autoWaterChangeAllColumns
	} else {
		fields = strmangle.SetComplement(
			autoWaterChangeAllColumns,
			autoWaterChangePrimaryKeyColumns,
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

	slice := AutoWaterChangeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
