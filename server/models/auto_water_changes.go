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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// AutoWaterChange is an object representing the database table.
type AutoWaterChange struct {
	ID           string  `boil:"id" json:"id" toml:"id" yaml:"id"`
	FreshPumpID  string  `boil:"fresh_pump_id" json:"fresh_pump_id" toml:"fresh_pump_id" yaml:"fresh_pump_id"`
	WastePumpID  string  `boil:"waste_pump_id" json:"waste_pump_id" toml:"waste_pump_id" yaml:"waste_pump_id"`
	ExchangeRate float64 `boil:"exchange_rate" json:"exchange_rate" toml:"exchange_rate" yaml:"exchange_rate"`

	R *autoWaterChangeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L autoWaterChangeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AutoWaterChangeColumns = struct {
	ID           string
	FreshPumpID  string
	WastePumpID  string
	ExchangeRate string
}{
	ID:           "id",
	FreshPumpID:  "fresh_pump_id",
	WastePumpID:  "waste_pump_id",
	ExchangeRate: "exchange_rate",
}

// Generated where

var AutoWaterChangeWhere = struct {
	ID           whereHelperstring
	FreshPumpID  whereHelperstring
	WastePumpID  whereHelperstring
	ExchangeRate whereHelperfloat64
}{
	ID:           whereHelperstring{field: "\"auto_water_changes\".\"id\""},
	FreshPumpID:  whereHelperstring{field: "\"auto_water_changes\".\"fresh_pump_id\""},
	WastePumpID:  whereHelperstring{field: "\"auto_water_changes\".\"waste_pump_id\""},
	ExchangeRate: whereHelperfloat64{field: "\"auto_water_changes\".\"exchange_rate\""},
}

// AutoWaterChangeRels is where relationship names are stored.
var AutoWaterChangeRels = struct {
	WastePump string
	FreshPump string
}{
	WastePump: "WastePump",
	FreshPump: "FreshPump",
}

// autoWaterChangeR is where relationships are stored.
type autoWaterChangeR struct {
	WastePump *Pump `boil:"WastePump" json:"WastePump" toml:"WastePump" yaml:"WastePump"`
	FreshPump *Pump `boil:"FreshPump" json:"FreshPump" toml:"FreshPump" yaml:"FreshPump"`
}

// NewStruct creates a new relationship struct
func (*autoWaterChangeR) NewStruct() *autoWaterChangeR {
	return &autoWaterChangeR{}
}

// autoWaterChangeL is where Load methods for each relationship are stored.
type autoWaterChangeL struct{}

var (
	autoWaterChangeAllColumns            = []string{"id", "fresh_pump_id", "waste_pump_id", "exchange_rate"}
	autoWaterChangeColumnsWithoutDefault = []string{"id", "fresh_pump_id", "waste_pump_id", "exchange_rate"}
	autoWaterChangeColumnsWithDefault    = []string{}
	autoWaterChangePrimaryKeyColumns     = []string{"id"}
)

type (
	// AutoWaterChangeSlice is an alias for a slice of pointers to AutoWaterChange.
	// This should generally be used opposed to []AutoWaterChange.
	AutoWaterChangeSlice []*AutoWaterChange
	// AutoWaterChangeHook is the signature for custom AutoWaterChange hook methods
	AutoWaterChangeHook func(context.Context, boil.ContextExecutor, *AutoWaterChange) error

	autoWaterChangeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	autoWaterChangeType                 = reflect.TypeOf(&AutoWaterChange{})
	autoWaterChangeMapping              = queries.MakeStructMapping(autoWaterChangeType)
	autoWaterChangePrimaryKeyMapping, _ = queries.BindMapping(autoWaterChangeType, autoWaterChangeMapping, autoWaterChangePrimaryKeyColumns)
	autoWaterChangeInsertCacheMut       sync.RWMutex
	autoWaterChangeInsertCache          = make(map[string]insertCache)
	autoWaterChangeUpdateCacheMut       sync.RWMutex
	autoWaterChangeUpdateCache          = make(map[string]updateCache)
	autoWaterChangeUpsertCacheMut       sync.RWMutex
	autoWaterChangeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var autoWaterChangeBeforeInsertHooks []AutoWaterChangeHook
var autoWaterChangeBeforeUpdateHooks []AutoWaterChangeHook
var autoWaterChangeBeforeDeleteHooks []AutoWaterChangeHook
var autoWaterChangeBeforeUpsertHooks []AutoWaterChangeHook

var autoWaterChangeAfterInsertHooks []AutoWaterChangeHook
var autoWaterChangeAfterSelectHooks []AutoWaterChangeHook
var autoWaterChangeAfterUpdateHooks []AutoWaterChangeHook
var autoWaterChangeAfterDeleteHooks []AutoWaterChangeHook
var autoWaterChangeAfterUpsertHooks []AutoWaterChangeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AutoWaterChange) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range autoWaterChangeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AutoWaterChange) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range autoWaterChangeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AutoWaterChange) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range autoWaterChangeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AutoWaterChange) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range autoWaterChangeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AutoWaterChange) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range autoWaterChangeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AutoWaterChange) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range autoWaterChangeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AutoWaterChange) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range autoWaterChangeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AutoWaterChange) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range autoWaterChangeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AutoWaterChange) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range autoWaterChangeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAutoWaterChangeHook registers your hook function for all future operations.
func AddAutoWaterChangeHook(hookPoint boil.HookPoint, autoWaterChangeHook AutoWaterChangeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		autoWaterChangeBeforeInsertHooks = append(autoWaterChangeBeforeInsertHooks, autoWaterChangeHook)
	case boil.BeforeUpdateHook:
		autoWaterChangeBeforeUpdateHooks = append(autoWaterChangeBeforeUpdateHooks, autoWaterChangeHook)
	case boil.BeforeDeleteHook:
		autoWaterChangeBeforeDeleteHooks = append(autoWaterChangeBeforeDeleteHooks, autoWaterChangeHook)
	case boil.BeforeUpsertHook:
		autoWaterChangeBeforeUpsertHooks = append(autoWaterChangeBeforeUpsertHooks, autoWaterChangeHook)
	case boil.AfterInsertHook:
		autoWaterChangeAfterInsertHooks = append(autoWaterChangeAfterInsertHooks, autoWaterChangeHook)
	case boil.AfterSelectHook:
		autoWaterChangeAfterSelectHooks = append(autoWaterChangeAfterSelectHooks, autoWaterChangeHook)
	case boil.AfterUpdateHook:
		autoWaterChangeAfterUpdateHooks = append(autoWaterChangeAfterUpdateHooks, autoWaterChangeHook)
	case boil.AfterDeleteHook:
		autoWaterChangeAfterDeleteHooks = append(autoWaterChangeAfterDeleteHooks, autoWaterChangeHook)
	case boil.AfterUpsertHook:
		autoWaterChangeAfterUpsertHooks = append(autoWaterChangeAfterUpsertHooks, autoWaterChangeHook)
	}
}

// One returns a single autoWaterChange record from the query.
func (q autoWaterChangeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AutoWaterChange, error) {
	o := &AutoWaterChange{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for auto_water_changes")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AutoWaterChange records from the query.
func (q autoWaterChangeQuery) All(ctx context.Context, exec boil.ContextExecutor) (AutoWaterChangeSlice, error) {
	var o []*AutoWaterChange

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AutoWaterChange slice")
	}

	if len(autoWaterChangeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AutoWaterChange records in the query.
func (q autoWaterChangeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count auto_water_changes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q autoWaterChangeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if auto_water_changes exists")
	}

	return count > 0, nil
}

// WastePump pointed to by the foreign key.
func (o *AutoWaterChange) WastePump(mods ...qm.QueryMod) pumpQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.WastePumpID),
	}

	queryMods = append(queryMods, mods...)

	query := Pumps(queryMods...)
	queries.SetFrom(query.Query, "\"pumps\"")

	return query
}

// FreshPump pointed to by the foreign key.
func (o *AutoWaterChange) FreshPump(mods ...qm.QueryMod) pumpQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.FreshPumpID),
	}

	queryMods = append(queryMods, mods...)

	query := Pumps(queryMods...)
	queries.SetFrom(query.Query, "\"pumps\"")

	return query
}

// LoadWastePump allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (autoWaterChangeL) LoadWastePump(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAutoWaterChange interface{}, mods queries.Applicator) error {
	var slice []*AutoWaterChange
	var object *AutoWaterChange

	if singular {
		object = maybeAutoWaterChange.(*AutoWaterChange)
	} else {
		slice = *maybeAutoWaterChange.(*[]*AutoWaterChange)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &autoWaterChangeR{}
		}
		args = append(args, object.WastePumpID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &autoWaterChangeR{}
			}

			for _, a := range args {
				if a == obj.WastePumpID {
					continue Outer
				}
			}

			args = append(args, obj.WastePumpID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`pumps`),
		qm.WhereIn(`pumps.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Pump")
	}

	var resultSlice []*Pump
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Pump")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for pumps")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for pumps")
	}

	if len(autoWaterChangeAfterSelectHooks) != 0 {
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
		object.R.WastePump = foreign
		if foreign.R == nil {
			foreign.R = &pumpR{}
		}
		foreign.R.WastePumpAutoWaterChanges = append(foreign.R.WastePumpAutoWaterChanges, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.WastePumpID == foreign.ID {
				local.R.WastePump = foreign
				if foreign.R == nil {
					foreign.R = &pumpR{}
				}
				foreign.R.WastePumpAutoWaterChanges = append(foreign.R.WastePumpAutoWaterChanges, local)
				break
			}
		}
	}

	return nil
}

// LoadFreshPump allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (autoWaterChangeL) LoadFreshPump(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAutoWaterChange interface{}, mods queries.Applicator) error {
	var slice []*AutoWaterChange
	var object *AutoWaterChange

	if singular {
		object = maybeAutoWaterChange.(*AutoWaterChange)
	} else {
		slice = *maybeAutoWaterChange.(*[]*AutoWaterChange)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &autoWaterChangeR{}
		}
		args = append(args, object.FreshPumpID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &autoWaterChangeR{}
			}

			for _, a := range args {
				if a == obj.FreshPumpID {
					continue Outer
				}
			}

			args = append(args, obj.FreshPumpID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`pumps`),
		qm.WhereIn(`pumps.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Pump")
	}

	var resultSlice []*Pump
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Pump")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for pumps")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for pumps")
	}

	if len(autoWaterChangeAfterSelectHooks) != 0 {
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
		object.R.FreshPump = foreign
		if foreign.R == nil {
			foreign.R = &pumpR{}
		}
		foreign.R.FreshPumpAutoWaterChanges = append(foreign.R.FreshPumpAutoWaterChanges, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.FreshPumpID == foreign.ID {
				local.R.FreshPump = foreign
				if foreign.R == nil {
					foreign.R = &pumpR{}
				}
				foreign.R.FreshPumpAutoWaterChanges = append(foreign.R.FreshPumpAutoWaterChanges, local)
				break
			}
		}
	}

	return nil
}

// SetWastePump of the autoWaterChange to the related item.
// Sets o.R.WastePump to related.
// Adds o to related.R.WastePumpAutoWaterChanges.
func (o *AutoWaterChange) SetWastePump(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Pump) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"auto_water_changes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"waste_pump_id"}),
		strmangle.WhereClause("\"", "\"", 0, autoWaterChangePrimaryKeyColumns),
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

	o.WastePumpID = related.ID
	if o.R == nil {
		o.R = &autoWaterChangeR{
			WastePump: related,
		}
	} else {
		o.R.WastePump = related
	}

	if related.R == nil {
		related.R = &pumpR{
			WastePumpAutoWaterChanges: AutoWaterChangeSlice{o},
		}
	} else {
		related.R.WastePumpAutoWaterChanges = append(related.R.WastePumpAutoWaterChanges, o)
	}

	return nil
}

// SetFreshPump of the autoWaterChange to the related item.
// Sets o.R.FreshPump to related.
// Adds o to related.R.FreshPumpAutoWaterChanges.
func (o *AutoWaterChange) SetFreshPump(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Pump) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"auto_water_changes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"fresh_pump_id"}),
		strmangle.WhereClause("\"", "\"", 0, autoWaterChangePrimaryKeyColumns),
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

	o.FreshPumpID = related.ID
	if o.R == nil {
		o.R = &autoWaterChangeR{
			FreshPump: related,
		}
	} else {
		o.R.FreshPump = related
	}

	if related.R == nil {
		related.R = &pumpR{
			FreshPumpAutoWaterChanges: AutoWaterChangeSlice{o},
		}
	} else {
		related.R.FreshPumpAutoWaterChanges = append(related.R.FreshPumpAutoWaterChanges, o)
	}

	return nil
}

// AutoWaterChanges retrieves all the records using an executor.
func AutoWaterChanges(mods ...qm.QueryMod) autoWaterChangeQuery {
	mods = append(mods, qm.From("\"auto_water_changes\""))
	return autoWaterChangeQuery{NewQuery(mods...)}
}

// FindAutoWaterChange retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAutoWaterChange(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*AutoWaterChange, error) {
	autoWaterChangeObj := &AutoWaterChange{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"auto_water_changes\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, autoWaterChangeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from auto_water_changes")
	}

	return autoWaterChangeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AutoWaterChange) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no auto_water_changes provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(autoWaterChangeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	autoWaterChangeInsertCacheMut.RLock()
	cache, cached := autoWaterChangeInsertCache[key]
	autoWaterChangeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			autoWaterChangeAllColumns,
			autoWaterChangeColumnsWithDefault,
			autoWaterChangeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(autoWaterChangeType, autoWaterChangeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(autoWaterChangeType, autoWaterChangeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"auto_water_changes\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"auto_water_changes\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"auto_water_changes\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, autoWaterChangePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into auto_water_changes")
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
		return errors.Wrap(err, "models: unable to populate default values for auto_water_changes")
	}

CacheNoHooks:
	if !cached {
		autoWaterChangeInsertCacheMut.Lock()
		autoWaterChangeInsertCache[key] = cache
		autoWaterChangeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AutoWaterChange.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AutoWaterChange) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	autoWaterChangeUpdateCacheMut.RLock()
	cache, cached := autoWaterChangeUpdateCache[key]
	autoWaterChangeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			autoWaterChangeAllColumns,
			autoWaterChangePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update auto_water_changes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"auto_water_changes\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, autoWaterChangePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(autoWaterChangeType, autoWaterChangeMapping, append(wl, autoWaterChangePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update auto_water_changes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for auto_water_changes")
	}

	if !cached {
		autoWaterChangeUpdateCacheMut.Lock()
		autoWaterChangeUpdateCache[key] = cache
		autoWaterChangeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q autoWaterChangeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for auto_water_changes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for auto_water_changes")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AutoWaterChangeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), autoWaterChangePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"auto_water_changes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, autoWaterChangePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in autoWaterChange slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all autoWaterChange")
	}
	return rowsAff, nil
}

// Delete deletes a single AutoWaterChange record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AutoWaterChange) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AutoWaterChange provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), autoWaterChangePrimaryKeyMapping)
	sql := "DELETE FROM \"auto_water_changes\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from auto_water_changes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for auto_water_changes")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q autoWaterChangeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no autoWaterChangeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from auto_water_changes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for auto_water_changes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AutoWaterChangeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(autoWaterChangeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), autoWaterChangePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"auto_water_changes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, autoWaterChangePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from autoWaterChange slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for auto_water_changes")
	}

	if len(autoWaterChangeAfterDeleteHooks) != 0 {
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
func (o *AutoWaterChange) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAutoWaterChange(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AutoWaterChangeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AutoWaterChangeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), autoWaterChangePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"auto_water_changes\".* FROM \"auto_water_changes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, autoWaterChangePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AutoWaterChangeSlice")
	}

	*o = slice

	return nil
}

// AutoWaterChangeExists checks if the AutoWaterChange row exists.
func AutoWaterChangeExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"auto_water_changes\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if auto_water_changes exists")
	}

	return exists, nil
}
