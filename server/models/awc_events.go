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

// AwcEvent is an object representing the database table.
type AwcEvent struct {
	ID                string `boil:"id" json:"id" toml:"id" yaml:"id"`
	AutoWaterChangeID string `boil:"auto_water_change_id" json:"auto_water_change_id" toml:"auto_water_change_id" yaml:"auto_water_change_id"`
	Timestamp         int64  `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	Kind              string `boil:"kind" json:"kind" toml:"kind" yaml:"kind"`
	Data              string `boil:"data" json:"data" toml:"data" yaml:"data"`

	R *awcEventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L awcEventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AwcEventColumns = struct {
	ID                string
	AutoWaterChangeID string
	Timestamp         string
	Kind              string
	Data              string
}{
	ID:                "id",
	AutoWaterChangeID: "auto_water_change_id",
	Timestamp:         "timestamp",
	Kind:              "kind",
	Data:              "data",
}

// Generated where

var AwcEventWhere = struct {
	ID                whereHelperstring
	AutoWaterChangeID whereHelperstring
	Timestamp         whereHelperint64
	Kind              whereHelperstring
	Data              whereHelperstring
}{
	ID:                whereHelperstring{field: "\"awc_events\".\"id\""},
	AutoWaterChangeID: whereHelperstring{field: "\"awc_events\".\"auto_water_change_id\""},
	Timestamp:         whereHelperint64{field: "\"awc_events\".\"timestamp\""},
	Kind:              whereHelperstring{field: "\"awc_events\".\"kind\""},
	Data:              whereHelperstring{field: "\"awc_events\".\"data\""},
}

// AwcEventRels is where relationship names are stored.
var AwcEventRels = struct {
	AutoWaterChange string
}{
	AutoWaterChange: "AutoWaterChange",
}

// awcEventR is where relationships are stored.
type awcEventR struct {
	AutoWaterChange *AutoWaterChange `boil:"AutoWaterChange" json:"AutoWaterChange" toml:"AutoWaterChange" yaml:"AutoWaterChange"`
}

// NewStruct creates a new relationship struct
func (*awcEventR) NewStruct() *awcEventR {
	return &awcEventR{}
}

// awcEventL is where Load methods for each relationship are stored.
type awcEventL struct{}

var (
	awcEventAllColumns            = []string{"id", "auto_water_change_id", "timestamp", "kind", "data"}
	awcEventColumnsWithoutDefault = []string{"id", "auto_water_change_id", "timestamp", "kind", "data"}
	awcEventColumnsWithDefault    = []string{}
	awcEventPrimaryKeyColumns     = []string{"id"}
)

type (
	// AwcEventSlice is an alias for a slice of pointers to AwcEvent.
	// This should generally be used opposed to []AwcEvent.
	AwcEventSlice []*AwcEvent
	// AwcEventHook is the signature for custom AwcEvent hook methods
	AwcEventHook func(context.Context, boil.ContextExecutor, *AwcEvent) error

	awcEventQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	awcEventType                 = reflect.TypeOf(&AwcEvent{})
	awcEventMapping              = queries.MakeStructMapping(awcEventType)
	awcEventPrimaryKeyMapping, _ = queries.BindMapping(awcEventType, awcEventMapping, awcEventPrimaryKeyColumns)
	awcEventInsertCacheMut       sync.RWMutex
	awcEventInsertCache          = make(map[string]insertCache)
	awcEventUpdateCacheMut       sync.RWMutex
	awcEventUpdateCache          = make(map[string]updateCache)
	awcEventUpsertCacheMut       sync.RWMutex
	awcEventUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var awcEventBeforeInsertHooks []AwcEventHook
var awcEventBeforeUpdateHooks []AwcEventHook
var awcEventBeforeDeleteHooks []AwcEventHook
var awcEventBeforeUpsertHooks []AwcEventHook

var awcEventAfterInsertHooks []AwcEventHook
var awcEventAfterSelectHooks []AwcEventHook
var awcEventAfterUpdateHooks []AwcEventHook
var awcEventAfterDeleteHooks []AwcEventHook
var awcEventAfterUpsertHooks []AwcEventHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AwcEvent) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range awcEventBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AwcEvent) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range awcEventBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AwcEvent) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range awcEventBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AwcEvent) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range awcEventBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AwcEvent) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range awcEventAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AwcEvent) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range awcEventAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AwcEvent) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range awcEventAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AwcEvent) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range awcEventAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AwcEvent) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range awcEventAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAwcEventHook registers your hook function for all future operations.
func AddAwcEventHook(hookPoint boil.HookPoint, awcEventHook AwcEventHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		awcEventBeforeInsertHooks = append(awcEventBeforeInsertHooks, awcEventHook)
	case boil.BeforeUpdateHook:
		awcEventBeforeUpdateHooks = append(awcEventBeforeUpdateHooks, awcEventHook)
	case boil.BeforeDeleteHook:
		awcEventBeforeDeleteHooks = append(awcEventBeforeDeleteHooks, awcEventHook)
	case boil.BeforeUpsertHook:
		awcEventBeforeUpsertHooks = append(awcEventBeforeUpsertHooks, awcEventHook)
	case boil.AfterInsertHook:
		awcEventAfterInsertHooks = append(awcEventAfterInsertHooks, awcEventHook)
	case boil.AfterSelectHook:
		awcEventAfterSelectHooks = append(awcEventAfterSelectHooks, awcEventHook)
	case boil.AfterUpdateHook:
		awcEventAfterUpdateHooks = append(awcEventAfterUpdateHooks, awcEventHook)
	case boil.AfterDeleteHook:
		awcEventAfterDeleteHooks = append(awcEventAfterDeleteHooks, awcEventHook)
	case boil.AfterUpsertHook:
		awcEventAfterUpsertHooks = append(awcEventAfterUpsertHooks, awcEventHook)
	}
}

// One returns a single awcEvent record from the query.
func (q awcEventQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AwcEvent, error) {
	o := &AwcEvent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for awc_events")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AwcEvent records from the query.
func (q awcEventQuery) All(ctx context.Context, exec boil.ContextExecutor) (AwcEventSlice, error) {
	var o []*AwcEvent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AwcEvent slice")
	}

	if len(awcEventAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AwcEvent records in the query.
func (q awcEventQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count awc_events rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q awcEventQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if awc_events exists")
	}

	return count > 0, nil
}

// AutoWaterChange pointed to by the foreign key.
func (o *AwcEvent) AutoWaterChange(mods ...qm.QueryMod) autoWaterChangeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.AutoWaterChangeID),
	}

	queryMods = append(queryMods, mods...)

	query := AutoWaterChanges(queryMods...)
	queries.SetFrom(query.Query, "\"auto_water_changes\"")

	return query
}

// LoadAutoWaterChange allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (awcEventL) LoadAutoWaterChange(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAwcEvent interface{}, mods queries.Applicator) error {
	var slice []*AwcEvent
	var object *AwcEvent

	if singular {
		object = maybeAwcEvent.(*AwcEvent)
	} else {
		slice = *maybeAwcEvent.(*[]*AwcEvent)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &awcEventR{}
		}
		args = append(args, object.AutoWaterChangeID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &awcEventR{}
			}

			for _, a := range args {
				if a == obj.AutoWaterChangeID {
					continue Outer
				}
			}

			args = append(args, obj.AutoWaterChangeID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`auto_water_changes`),
		qm.WhereIn(`auto_water_changes.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AutoWaterChange")
	}

	var resultSlice []*AutoWaterChange
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AutoWaterChange")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for auto_water_changes")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for auto_water_changes")
	}

	if len(awcEventAfterSelectHooks) != 0 {
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
		object.R.AutoWaterChange = foreign
		if foreign.R == nil {
			foreign.R = &autoWaterChangeR{}
		}
		foreign.R.AwcEvents = append(foreign.R.AwcEvents, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AutoWaterChangeID == foreign.ID {
				local.R.AutoWaterChange = foreign
				if foreign.R == nil {
					foreign.R = &autoWaterChangeR{}
				}
				foreign.R.AwcEvents = append(foreign.R.AwcEvents, local)
				break
			}
		}
	}

	return nil
}

// SetAutoWaterChange of the awcEvent to the related item.
// Sets o.R.AutoWaterChange to related.
// Adds o to related.R.AwcEvents.
func (o *AwcEvent) SetAutoWaterChange(ctx context.Context, exec boil.ContextExecutor, insert bool, related *AutoWaterChange) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"awc_events\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"auto_water_change_id"}),
		strmangle.WhereClause("\"", "\"", 0, awcEventPrimaryKeyColumns),
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

	o.AutoWaterChangeID = related.ID
	if o.R == nil {
		o.R = &awcEventR{
			AutoWaterChange: related,
		}
	} else {
		o.R.AutoWaterChange = related
	}

	if related.R == nil {
		related.R = &autoWaterChangeR{
			AwcEvents: AwcEventSlice{o},
		}
	} else {
		related.R.AwcEvents = append(related.R.AwcEvents, o)
	}

	return nil
}

// AwcEvents retrieves all the records using an executor.
func AwcEvents(mods ...qm.QueryMod) awcEventQuery {
	mods = append(mods, qm.From("\"awc_events\""))
	return awcEventQuery{NewQuery(mods...)}
}

// FindAwcEvent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAwcEvent(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*AwcEvent, error) {
	awcEventObj := &AwcEvent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"awc_events\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, awcEventObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from awc_events")
	}

	return awcEventObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AwcEvent) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no awc_events provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(awcEventColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	awcEventInsertCacheMut.RLock()
	cache, cached := awcEventInsertCache[key]
	awcEventInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			awcEventAllColumns,
			awcEventColumnsWithDefault,
			awcEventColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(awcEventType, awcEventMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(awcEventType, awcEventMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"awc_events\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"awc_events\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"awc_events\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, awcEventPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into awc_events")
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
		return errors.Wrap(err, "models: unable to populate default values for awc_events")
	}

CacheNoHooks:
	if !cached {
		awcEventInsertCacheMut.Lock()
		awcEventInsertCache[key] = cache
		awcEventInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AwcEvent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AwcEvent) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	awcEventUpdateCacheMut.RLock()
	cache, cached := awcEventUpdateCache[key]
	awcEventUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			awcEventAllColumns,
			awcEventPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update awc_events, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"awc_events\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, awcEventPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(awcEventType, awcEventMapping, append(wl, awcEventPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update awc_events row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for awc_events")
	}

	if !cached {
		awcEventUpdateCacheMut.Lock()
		awcEventUpdateCache[key] = cache
		awcEventUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q awcEventQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for awc_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for awc_events")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AwcEventSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), awcEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"awc_events\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, awcEventPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in awcEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all awcEvent")
	}
	return rowsAff, nil
}

// Delete deletes a single AwcEvent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AwcEvent) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AwcEvent provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), awcEventPrimaryKeyMapping)
	sql := "DELETE FROM \"awc_events\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from awc_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for awc_events")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q awcEventQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no awcEventQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from awc_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for awc_events")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AwcEventSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(awcEventBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), awcEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"awc_events\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, awcEventPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from awcEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for awc_events")
	}

	if len(awcEventAfterDeleteHooks) != 0 {
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
func (o *AwcEvent) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAwcEvent(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AwcEventSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AwcEventSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), awcEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"awc_events\".* FROM \"awc_events\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, awcEventPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AwcEventSlice")
	}

	*o = slice

	return nil
}

// AwcEventExists checks if the AwcEvent row exists.
func AwcEventExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"awc_events\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if awc_events exists")
	}

	return exists, nil
}
