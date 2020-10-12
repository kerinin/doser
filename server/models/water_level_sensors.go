// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// WaterLevelSensor is an object representing the database table.
type WaterLevelSensor struct {
	ID                 string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	FirmataID          null.String `boil:"firmata_id" json:"firmata_id,omitempty" toml:"firmata_id" yaml:"firmata_id,omitempty"`
	Pin                int64       `boil:"pin" json:"pin" toml:"pin" yaml:"pin"`
	Kind               string      `boil:"kind" json:"kind" toml:"kind" yaml:"kind"`
	DetectionThreshold null.Int64  `boil:"detection_threshold" json:"detection_threshold,omitempty" toml:"detection_threshold" yaml:"detection_threshold,omitempty"`
	Invert             bool        `boil:"invert" json:"invert" toml:"invert" yaml:"invert"`
	Name               null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`

	R *waterLevelSensorR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L waterLevelSensorL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var WaterLevelSensorColumns = struct {
	ID                 string
	FirmataID          string
	Pin                string
	Kind               string
	DetectionThreshold string
	Invert             string
	Name               string
}{
	ID:                 "id",
	FirmataID:          "firmata_id",
	Pin:                "pin",
	Kind:               "kind",
	DetectionThreshold: "detection_threshold",
	Invert:             "invert",
	Name:               "name",
}

// Generated where

var WaterLevelSensorWhere = struct {
	ID                 whereHelperstring
	FirmataID          whereHelpernull_String
	Pin                whereHelperint64
	Kind               whereHelperstring
	DetectionThreshold whereHelpernull_Int64
	Invert             whereHelperbool
	Name               whereHelpernull_String
}{
	ID:                 whereHelperstring{field: "\"water_level_sensors\".\"id\""},
	FirmataID:          whereHelpernull_String{field: "\"water_level_sensors\".\"firmata_id\""},
	Pin:                whereHelperint64{field: "\"water_level_sensors\".\"pin\""},
	Kind:               whereHelperstring{field: "\"water_level_sensors\".\"kind\""},
	DetectionThreshold: whereHelpernull_Int64{field: "\"water_level_sensors\".\"detection_threshold\""},
	Invert:             whereHelperbool{field: "\"water_level_sensors\".\"invert\""},
	Name:               whereHelpernull_String{field: "\"water_level_sensors\".\"name\""},
}

// WaterLevelSensorRels is where relationship names are stored.
var WaterLevelSensorRels = struct {
	Firmatum    string
	AutoTopOffs string
}{
	Firmatum:    "Firmatum",
	AutoTopOffs: "AutoTopOffs",
}

// waterLevelSensorR is where relationships are stored.
type waterLevelSensorR struct {
	Firmatum    *Firmata        `boil:"Firmatum" json:"Firmatum" toml:"Firmatum" yaml:"Firmatum"`
	AutoTopOffs AutoTopOffSlice `boil:"AutoTopOffs" json:"AutoTopOffs" toml:"AutoTopOffs" yaml:"AutoTopOffs"`
}

// NewStruct creates a new relationship struct
func (*waterLevelSensorR) NewStruct() *waterLevelSensorR {
	return &waterLevelSensorR{}
}

// waterLevelSensorL is where Load methods for each relationship are stored.
type waterLevelSensorL struct{}

var (
	waterLevelSensorAllColumns            = []string{"id", "firmata_id", "pin", "kind", "detection_threshold", "invert", "name"}
	waterLevelSensorColumnsWithoutDefault = []string{"id", "firmata_id", "pin", "kind", "detection_threshold", "name"}
	waterLevelSensorColumnsWithDefault    = []string{"invert"}
	waterLevelSensorPrimaryKeyColumns     = []string{"id"}
)

type (
	// WaterLevelSensorSlice is an alias for a slice of pointers to WaterLevelSensor.
	// This should generally be used opposed to []WaterLevelSensor.
	WaterLevelSensorSlice []*WaterLevelSensor
	// WaterLevelSensorHook is the signature for custom WaterLevelSensor hook methods
	WaterLevelSensorHook func(context.Context, boil.ContextExecutor, *WaterLevelSensor) error

	waterLevelSensorQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	waterLevelSensorType                 = reflect.TypeOf(&WaterLevelSensor{})
	waterLevelSensorMapping              = queries.MakeStructMapping(waterLevelSensorType)
	waterLevelSensorPrimaryKeyMapping, _ = queries.BindMapping(waterLevelSensorType, waterLevelSensorMapping, waterLevelSensorPrimaryKeyColumns)
	waterLevelSensorInsertCacheMut       sync.RWMutex
	waterLevelSensorInsertCache          = make(map[string]insertCache)
	waterLevelSensorUpdateCacheMut       sync.RWMutex
	waterLevelSensorUpdateCache          = make(map[string]updateCache)
	waterLevelSensorUpsertCacheMut       sync.RWMutex
	waterLevelSensorUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var waterLevelSensorBeforeInsertHooks []WaterLevelSensorHook
var waterLevelSensorBeforeUpdateHooks []WaterLevelSensorHook
var waterLevelSensorBeforeDeleteHooks []WaterLevelSensorHook
var waterLevelSensorBeforeUpsertHooks []WaterLevelSensorHook

var waterLevelSensorAfterInsertHooks []WaterLevelSensorHook
var waterLevelSensorAfterSelectHooks []WaterLevelSensorHook
var waterLevelSensorAfterUpdateHooks []WaterLevelSensorHook
var waterLevelSensorAfterDeleteHooks []WaterLevelSensorHook
var waterLevelSensorAfterUpsertHooks []WaterLevelSensorHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *WaterLevelSensor) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waterLevelSensorBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *WaterLevelSensor) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waterLevelSensorBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *WaterLevelSensor) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waterLevelSensorBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *WaterLevelSensor) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waterLevelSensorBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *WaterLevelSensor) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waterLevelSensorAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *WaterLevelSensor) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waterLevelSensorAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *WaterLevelSensor) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waterLevelSensorAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *WaterLevelSensor) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waterLevelSensorAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *WaterLevelSensor) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waterLevelSensorAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddWaterLevelSensorHook registers your hook function for all future operations.
func AddWaterLevelSensorHook(hookPoint boil.HookPoint, waterLevelSensorHook WaterLevelSensorHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		waterLevelSensorBeforeInsertHooks = append(waterLevelSensorBeforeInsertHooks, waterLevelSensorHook)
	case boil.BeforeUpdateHook:
		waterLevelSensorBeforeUpdateHooks = append(waterLevelSensorBeforeUpdateHooks, waterLevelSensorHook)
	case boil.BeforeDeleteHook:
		waterLevelSensorBeforeDeleteHooks = append(waterLevelSensorBeforeDeleteHooks, waterLevelSensorHook)
	case boil.BeforeUpsertHook:
		waterLevelSensorBeforeUpsertHooks = append(waterLevelSensorBeforeUpsertHooks, waterLevelSensorHook)
	case boil.AfterInsertHook:
		waterLevelSensorAfterInsertHooks = append(waterLevelSensorAfterInsertHooks, waterLevelSensorHook)
	case boil.AfterSelectHook:
		waterLevelSensorAfterSelectHooks = append(waterLevelSensorAfterSelectHooks, waterLevelSensorHook)
	case boil.AfterUpdateHook:
		waterLevelSensorAfterUpdateHooks = append(waterLevelSensorAfterUpdateHooks, waterLevelSensorHook)
	case boil.AfterDeleteHook:
		waterLevelSensorAfterDeleteHooks = append(waterLevelSensorAfterDeleteHooks, waterLevelSensorHook)
	case boil.AfterUpsertHook:
		waterLevelSensorAfterUpsertHooks = append(waterLevelSensorAfterUpsertHooks, waterLevelSensorHook)
	}
}

// One returns a single waterLevelSensor record from the query.
func (q waterLevelSensorQuery) One(ctx context.Context, exec boil.ContextExecutor) (*WaterLevelSensor, error) {
	o := &WaterLevelSensor{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for water_level_sensors")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all WaterLevelSensor records from the query.
func (q waterLevelSensorQuery) All(ctx context.Context, exec boil.ContextExecutor) (WaterLevelSensorSlice, error) {
	var o []*WaterLevelSensor

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to WaterLevelSensor slice")
	}

	if len(waterLevelSensorAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all WaterLevelSensor records in the query.
func (q waterLevelSensorQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count water_level_sensors rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q waterLevelSensorQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if water_level_sensors exists")
	}

	return count > 0, nil
}

// Firmatum pointed to by the foreign key.
func (o *WaterLevelSensor) Firmatum(mods ...qm.QueryMod) firmataQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.FirmataID),
	}

	queryMods = append(queryMods, mods...)

	query := Firmatas(queryMods...)
	queries.SetFrom(query.Query, "\"firmatas\"")

	return query
}

// AutoTopOffs retrieves all the auto_top_off's AutoTopOffs with an executor.
func (o *WaterLevelSensor) AutoTopOffs(mods ...qm.QueryMod) autoTopOffQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"auto_top_offs_water_level_sensors\" on \"auto_top_offs\".\"id\" = \"auto_top_offs_water_level_sensors\".\"auto_top_off_id\""),
		qm.Where("\"auto_top_offs_water_level_sensors\".\"water_level_sensor_id\"=?", o.ID),
	)

	query := AutoTopOffs(queryMods...)
	queries.SetFrom(query.Query, "\"auto_top_offs\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"auto_top_offs\".*"})
	}

	return query
}

// LoadFirmatum allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (waterLevelSensorL) LoadFirmatum(ctx context.Context, e boil.ContextExecutor, singular bool, maybeWaterLevelSensor interface{}, mods queries.Applicator) error {
	var slice []*WaterLevelSensor
	var object *WaterLevelSensor

	if singular {
		object = maybeWaterLevelSensor.(*WaterLevelSensor)
	} else {
		slice = *maybeWaterLevelSensor.(*[]*WaterLevelSensor)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &waterLevelSensorR{}
		}
		if !queries.IsNil(object.FirmataID) {
			args = append(args, object.FirmataID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &waterLevelSensorR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.FirmataID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.FirmataID) {
				args = append(args, obj.FirmataID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`firmatas`),
		qm.WhereIn(`firmatas.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Firmata")
	}

	var resultSlice []*Firmata
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Firmata")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for firmatas")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for firmatas")
	}

	if len(waterLevelSensorAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Firmatum = foreign
		if foreign.R == nil {
			foreign.R = &firmataR{}
		}
		foreign.R.FirmatumWaterLevelSensors = append(foreign.R.FirmatumWaterLevelSensors, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.FirmataID, foreign.ID) {
				local.R.Firmatum = foreign
				if foreign.R == nil {
					foreign.R = &firmataR{}
				}
				foreign.R.FirmatumWaterLevelSensors = append(foreign.R.FirmatumWaterLevelSensors, local)
				break
			}
		}
	}

	return nil
}

// LoadAutoTopOffs allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (waterLevelSensorL) LoadAutoTopOffs(ctx context.Context, e boil.ContextExecutor, singular bool, maybeWaterLevelSensor interface{}, mods queries.Applicator) error {
	var slice []*WaterLevelSensor
	var object *WaterLevelSensor

	if singular {
		object = maybeWaterLevelSensor.(*WaterLevelSensor)
	} else {
		slice = *maybeWaterLevelSensor.(*[]*WaterLevelSensor)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &waterLevelSensorR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &waterLevelSensorR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.Select("\"auto_top_offs\".*, \"a\".\"water_level_sensor_id\""),
		qm.From("\"auto_top_offs\""),
		qm.InnerJoin("\"auto_top_offs_water_level_sensors\" as \"a\" on \"auto_top_offs\".\"id\" = \"a\".\"auto_top_off_id\""),
		qm.WhereIn("\"a\".\"water_level_sensor_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load auto_top_offs")
	}

	var resultSlice []*AutoTopOff

	var localJoinCols []string
	for results.Next() {
		one := new(AutoTopOff)
		var localJoinCol string

		err = results.Scan(&one.ID, &one.PumpID, &one.FillRate, &one.FillInterval, &one.MaxFillVolume, &one.Enabled, &one.Name, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for auto_top_offs")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice auto_top_offs")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on auto_top_offs")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for auto_top_offs")
	}

	if len(autoTopOffAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.AutoTopOffs = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &autoTopOffR{}
			}
			foreign.R.WaterLevelSensors = append(foreign.R.WaterLevelSensors, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.AutoTopOffs = append(local.R.AutoTopOffs, foreign)
				if foreign.R == nil {
					foreign.R = &autoTopOffR{}
				}
				foreign.R.WaterLevelSensors = append(foreign.R.WaterLevelSensors, local)
				break
			}
		}
	}

	return nil
}

// SetFirmatum of the waterLevelSensor to the related item.
// Sets o.R.Firmatum to related.
// Adds o to related.R.FirmatumWaterLevelSensors.
func (o *WaterLevelSensor) SetFirmatum(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Firmata) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"water_level_sensors\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"firmata_id"}),
		strmangle.WhereClause("\"", "\"", 0, waterLevelSensorPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.FirmataID, related.ID)
	if o.R == nil {
		o.R = &waterLevelSensorR{
			Firmatum: related,
		}
	} else {
		o.R.Firmatum = related
	}

	if related.R == nil {
		related.R = &firmataR{
			FirmatumWaterLevelSensors: WaterLevelSensorSlice{o},
		}
	} else {
		related.R.FirmatumWaterLevelSensors = append(related.R.FirmatumWaterLevelSensors, o)
	}

	return nil
}

// RemoveFirmatum relationship.
// Sets o.R.Firmatum to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *WaterLevelSensor) RemoveFirmatum(ctx context.Context, exec boil.ContextExecutor, related *Firmata) error {
	var err error

	queries.SetScanner(&o.FirmataID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("firmata_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Firmatum = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.FirmatumWaterLevelSensors {
		if queries.Equal(o.FirmataID, ri.FirmataID) {
			continue
		}

		ln := len(related.R.FirmatumWaterLevelSensors)
		if ln > 1 && i < ln-1 {
			related.R.FirmatumWaterLevelSensors[i] = related.R.FirmatumWaterLevelSensors[ln-1]
		}
		related.R.FirmatumWaterLevelSensors = related.R.FirmatumWaterLevelSensors[:ln-1]
		break
	}
	return nil
}

// AddAutoTopOffs adds the given related objects to the existing relationships
// of the water_level_sensor, optionally inserting them as new records.
// Appends related to o.R.AutoTopOffs.
// Sets related.R.WaterLevelSensors appropriately.
func (o *WaterLevelSensor) AddAutoTopOffs(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*AutoTopOff) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"auto_top_offs_water_level_sensors\" (\"water_level_sensor_id\", \"auto_top_off_id\") values (?, ?)"
		values := []interface{}{o.ID, rel.ID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, query)
			fmt.Fprintln(writer, values)
		}
		_, err = exec.ExecContext(ctx, query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &waterLevelSensorR{
			AutoTopOffs: related,
		}
	} else {
		o.R.AutoTopOffs = append(o.R.AutoTopOffs, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &autoTopOffR{
				WaterLevelSensors: WaterLevelSensorSlice{o},
			}
		} else {
			rel.R.WaterLevelSensors = append(rel.R.WaterLevelSensors, o)
		}
	}
	return nil
}

// SetAutoTopOffs removes all previously related items of the
// water_level_sensor replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.WaterLevelSensors's AutoTopOffs accordingly.
// Replaces o.R.AutoTopOffs with related.
// Sets related.R.WaterLevelSensors's AutoTopOffs accordingly.
func (o *WaterLevelSensor) SetAutoTopOffs(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*AutoTopOff) error {
	query := "delete from \"auto_top_offs_water_level_sensors\" where \"water_level_sensor_id\" = ?"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeAutoTopOffsFromWaterLevelSensorsSlice(o, related)
	if o.R != nil {
		o.R.AutoTopOffs = nil
	}
	return o.AddAutoTopOffs(ctx, exec, insert, related...)
}

// RemoveAutoTopOffs relationships from objects passed in.
// Removes related items from R.AutoTopOffs (uses pointer comparison, removal does not keep order)
// Sets related.R.WaterLevelSensors.
func (o *WaterLevelSensor) RemoveAutoTopOffs(ctx context.Context, exec boil.ContextExecutor, related ...*AutoTopOff) error {
	var err error
	query := fmt.Sprintf(
		"delete from \"auto_top_offs_water_level_sensors\" where \"water_level_sensor_id\" = ? and \"auto_top_off_id\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeAutoTopOffsFromWaterLevelSensorsSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.AutoTopOffs {
			if rel != ri {
				continue
			}

			ln := len(o.R.AutoTopOffs)
			if ln > 1 && i < ln-1 {
				o.R.AutoTopOffs[i] = o.R.AutoTopOffs[ln-1]
			}
			o.R.AutoTopOffs = o.R.AutoTopOffs[:ln-1]
			break
		}
	}

	return nil
}

func removeAutoTopOffsFromWaterLevelSensorsSlice(o *WaterLevelSensor, related []*AutoTopOff) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.WaterLevelSensors {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.WaterLevelSensors)
			if ln > 1 && i < ln-1 {
				rel.R.WaterLevelSensors[i] = rel.R.WaterLevelSensors[ln-1]
			}
			rel.R.WaterLevelSensors = rel.R.WaterLevelSensors[:ln-1]
			break
		}
	}
}

// WaterLevelSensors retrieves all the records using an executor.
func WaterLevelSensors(mods ...qm.QueryMod) waterLevelSensorQuery {
	mods = append(mods, qm.From("\"water_level_sensors\""))
	return waterLevelSensorQuery{NewQuery(mods...)}
}

// FindWaterLevelSensor retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindWaterLevelSensor(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*WaterLevelSensor, error) {
	waterLevelSensorObj := &WaterLevelSensor{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"water_level_sensors\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, waterLevelSensorObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from water_level_sensors")
	}

	return waterLevelSensorObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *WaterLevelSensor) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no water_level_sensors provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(waterLevelSensorColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	waterLevelSensorInsertCacheMut.RLock()
	cache, cached := waterLevelSensorInsertCache[key]
	waterLevelSensorInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			waterLevelSensorAllColumns,
			waterLevelSensorColumnsWithDefault,
			waterLevelSensorColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(waterLevelSensorType, waterLevelSensorMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(waterLevelSensorType, waterLevelSensorMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"water_level_sensors\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"water_level_sensors\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"water_level_sensors\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, waterLevelSensorPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into water_level_sensors")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for water_level_sensors")
	}

CacheNoHooks:
	if !cached {
		waterLevelSensorInsertCacheMut.Lock()
		waterLevelSensorInsertCache[key] = cache
		waterLevelSensorInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the WaterLevelSensor.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *WaterLevelSensor) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	waterLevelSensorUpdateCacheMut.RLock()
	cache, cached := waterLevelSensorUpdateCache[key]
	waterLevelSensorUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			waterLevelSensorAllColumns,
			waterLevelSensorPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update water_level_sensors, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"water_level_sensors\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, waterLevelSensorPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(waterLevelSensorType, waterLevelSensorMapping, append(wl, waterLevelSensorPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update water_level_sensors row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for water_level_sensors")
	}

	if !cached {
		waterLevelSensorUpdateCacheMut.Lock()
		waterLevelSensorUpdateCache[key] = cache
		waterLevelSensorUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q waterLevelSensorQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for water_level_sensors")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for water_level_sensors")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o WaterLevelSensorSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), waterLevelSensorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"water_level_sensors\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, waterLevelSensorPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in waterLevelSensor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all waterLevelSensor")
	}
	return rowsAff, nil
}

// Delete deletes a single WaterLevelSensor record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *WaterLevelSensor) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no WaterLevelSensor provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), waterLevelSensorPrimaryKeyMapping)
	sql := "DELETE FROM \"water_level_sensors\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from water_level_sensors")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for water_level_sensors")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q waterLevelSensorQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no waterLevelSensorQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from water_level_sensors")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for water_level_sensors")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o WaterLevelSensorSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(waterLevelSensorBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), waterLevelSensorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"water_level_sensors\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, waterLevelSensorPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from waterLevelSensor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for water_level_sensors")
	}

	if len(waterLevelSensorAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *WaterLevelSensor) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindWaterLevelSensor(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *WaterLevelSensorSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := WaterLevelSensorSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), waterLevelSensorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"water_level_sensors\".* FROM \"water_level_sensors\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, waterLevelSensorPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in WaterLevelSensorSlice")
	}

	*o = slice

	return nil
}

// WaterLevelSensorExists checks if the WaterLevelSensor row exists.
func WaterLevelSensorExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"water_level_sensors\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if water_level_sensors exists")
	}

	return exists, nil
}
