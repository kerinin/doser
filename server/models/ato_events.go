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

// AtoEvent is an object representing the database table.
type AtoEvent struct {
	ID           string `boil:"id" json:"id" toml:"id" yaml:"id"`
	AutoTopOffID string `boil:"auto_top_off_id" json:"auto_top_off_id" toml:"auto_top_off_id" yaml:"auto_top_off_id"`
	Timestamp    int64  `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	Kind         string `boil:"kind" json:"kind" toml:"kind" yaml:"kind"`
	Data         string `boil:"data" json:"data" toml:"data" yaml:"data"`

	R *atoEventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L atoEventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AtoEventColumns = struct {
	ID           string
	AutoTopOffID string
	Timestamp    string
	Kind         string
	Data         string
}{
	ID:           "id",
	AutoTopOffID: "auto_top_off_id",
	Timestamp:    "timestamp",
	Kind:         "kind",
	Data:         "data",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var AtoEventWhere = struct {
	ID           whereHelperstring
	AutoTopOffID whereHelperstring
	Timestamp    whereHelperint64
	Kind         whereHelperstring
	Data         whereHelperstring
}{
	ID:           whereHelperstring{field: "\"ato_events\".\"id\""},
	AutoTopOffID: whereHelperstring{field: "\"ato_events\".\"auto_top_off_id\""},
	Timestamp:    whereHelperint64{field: "\"ato_events\".\"timestamp\""},
	Kind:         whereHelperstring{field: "\"ato_events\".\"kind\""},
	Data:         whereHelperstring{field: "\"ato_events\".\"data\""},
}

// AtoEventRels is where relationship names are stored.
var AtoEventRels = struct {
	AutoTopOff string
}{
	AutoTopOff: "AutoTopOff",
}

// atoEventR is where relationships are stored.
type atoEventR struct {
	AutoTopOff *AutoTopOff `boil:"AutoTopOff" json:"AutoTopOff" toml:"AutoTopOff" yaml:"AutoTopOff"`
}

// NewStruct creates a new relationship struct
func (*atoEventR) NewStruct() *atoEventR {
	return &atoEventR{}
}

// atoEventL is where Load methods for each relationship are stored.
type atoEventL struct{}

var (
	atoEventAllColumns            = []string{"id", "auto_top_off_id", "timestamp", "kind", "data"}
	atoEventColumnsWithoutDefault = []string{"id", "auto_top_off_id", "timestamp", "kind", "data"}
	atoEventColumnsWithDefault    = []string{}
	atoEventPrimaryKeyColumns     = []string{"id"}
)

type (
	// AtoEventSlice is an alias for a slice of pointers to AtoEvent.
	// This should generally be used opposed to []AtoEvent.
	AtoEventSlice []*AtoEvent
	// AtoEventHook is the signature for custom AtoEvent hook methods
	AtoEventHook func(context.Context, boil.ContextExecutor, *AtoEvent) error

	atoEventQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	atoEventType                 = reflect.TypeOf(&AtoEvent{})
	atoEventMapping              = queries.MakeStructMapping(atoEventType)
	atoEventPrimaryKeyMapping, _ = queries.BindMapping(atoEventType, atoEventMapping, atoEventPrimaryKeyColumns)
	atoEventInsertCacheMut       sync.RWMutex
	atoEventInsertCache          = make(map[string]insertCache)
	atoEventUpdateCacheMut       sync.RWMutex
	atoEventUpdateCache          = make(map[string]updateCache)
	atoEventUpsertCacheMut       sync.RWMutex
	atoEventUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var atoEventBeforeInsertHooks []AtoEventHook
var atoEventBeforeUpdateHooks []AtoEventHook
var atoEventBeforeDeleteHooks []AtoEventHook
var atoEventBeforeUpsertHooks []AtoEventHook

var atoEventAfterInsertHooks []AtoEventHook
var atoEventAfterSelectHooks []AtoEventHook
var atoEventAfterUpdateHooks []AtoEventHook
var atoEventAfterDeleteHooks []AtoEventHook
var atoEventAfterUpsertHooks []AtoEventHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AtoEvent) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range atoEventBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AtoEvent) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range atoEventBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AtoEvent) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range atoEventBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AtoEvent) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range atoEventBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AtoEvent) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range atoEventAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AtoEvent) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range atoEventAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AtoEvent) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range atoEventAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AtoEvent) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range atoEventAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AtoEvent) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range atoEventAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAtoEventHook registers your hook function for all future operations.
func AddAtoEventHook(hookPoint boil.HookPoint, atoEventHook AtoEventHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		atoEventBeforeInsertHooks = append(atoEventBeforeInsertHooks, atoEventHook)
	case boil.BeforeUpdateHook:
		atoEventBeforeUpdateHooks = append(atoEventBeforeUpdateHooks, atoEventHook)
	case boil.BeforeDeleteHook:
		atoEventBeforeDeleteHooks = append(atoEventBeforeDeleteHooks, atoEventHook)
	case boil.BeforeUpsertHook:
		atoEventBeforeUpsertHooks = append(atoEventBeforeUpsertHooks, atoEventHook)
	case boil.AfterInsertHook:
		atoEventAfterInsertHooks = append(atoEventAfterInsertHooks, atoEventHook)
	case boil.AfterSelectHook:
		atoEventAfterSelectHooks = append(atoEventAfterSelectHooks, atoEventHook)
	case boil.AfterUpdateHook:
		atoEventAfterUpdateHooks = append(atoEventAfterUpdateHooks, atoEventHook)
	case boil.AfterDeleteHook:
		atoEventAfterDeleteHooks = append(atoEventAfterDeleteHooks, atoEventHook)
	case boil.AfterUpsertHook:
		atoEventAfterUpsertHooks = append(atoEventAfterUpsertHooks, atoEventHook)
	}
}

// One returns a single atoEvent record from the query.
func (q atoEventQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AtoEvent, error) {
	o := &AtoEvent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ato_events")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AtoEvent records from the query.
func (q atoEventQuery) All(ctx context.Context, exec boil.ContextExecutor) (AtoEventSlice, error) {
	var o []*AtoEvent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AtoEvent slice")
	}

	if len(atoEventAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AtoEvent records in the query.
func (q atoEventQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ato_events rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q atoEventQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ato_events exists")
	}

	return count > 0, nil
}

// AutoTopOff pointed to by the foreign key.
func (o *AtoEvent) AutoTopOff(mods ...qm.QueryMod) autoTopOffQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.AutoTopOffID),
	}

	queryMods = append(queryMods, mods...)

	query := AutoTopOffs(queryMods...)
	queries.SetFrom(query.Query, "\"auto_top_offs\"")

	return query
}

// LoadAutoTopOff allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (atoEventL) LoadAutoTopOff(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAtoEvent interface{}, mods queries.Applicator) error {
	var slice []*AtoEvent
	var object *AtoEvent

	if singular {
		object = maybeAtoEvent.(*AtoEvent)
	} else {
		slice = *maybeAtoEvent.(*[]*AtoEvent)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &atoEventR{}
		}
		args = append(args, object.AutoTopOffID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &atoEventR{}
			}

			for _, a := range args {
				if a == obj.AutoTopOffID {
					continue Outer
				}
			}

			args = append(args, obj.AutoTopOffID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`auto_top_offs`),
		qm.WhereIn(`auto_top_offs.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AutoTopOff")
	}

	var resultSlice []*AutoTopOff
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AutoTopOff")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for auto_top_offs")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for auto_top_offs")
	}

	if len(atoEventAfterSelectHooks) != 0 {
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
		object.R.AutoTopOff = foreign
		if foreign.R == nil {
			foreign.R = &autoTopOffR{}
		}
		foreign.R.AtoEvents = append(foreign.R.AtoEvents, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AutoTopOffID == foreign.ID {
				local.R.AutoTopOff = foreign
				if foreign.R == nil {
					foreign.R = &autoTopOffR{}
				}
				foreign.R.AtoEvents = append(foreign.R.AtoEvents, local)
				break
			}
		}
	}

	return nil
}

// SetAutoTopOff of the atoEvent to the related item.
// Sets o.R.AutoTopOff to related.
// Adds o to related.R.AtoEvents.
func (o *AtoEvent) SetAutoTopOff(ctx context.Context, exec boil.ContextExecutor, insert bool, related *AutoTopOff) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"ato_events\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"auto_top_off_id"}),
		strmangle.WhereClause("\"", "\"", 0, atoEventPrimaryKeyColumns),
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

	o.AutoTopOffID = related.ID
	if o.R == nil {
		o.R = &atoEventR{
			AutoTopOff: related,
		}
	} else {
		o.R.AutoTopOff = related
	}

	if related.R == nil {
		related.R = &autoTopOffR{
			AtoEvents: AtoEventSlice{o},
		}
	} else {
		related.R.AtoEvents = append(related.R.AtoEvents, o)
	}

	return nil
}

// AtoEvents retrieves all the records using an executor.
func AtoEvents(mods ...qm.QueryMod) atoEventQuery {
	mods = append(mods, qm.From("\"ato_events\""))
	return atoEventQuery{NewQuery(mods...)}
}

// FindAtoEvent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAtoEvent(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*AtoEvent, error) {
	atoEventObj := &AtoEvent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"ato_events\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, atoEventObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ato_events")
	}

	return atoEventObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AtoEvent) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ato_events provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(atoEventColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	atoEventInsertCacheMut.RLock()
	cache, cached := atoEventInsertCache[key]
	atoEventInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			atoEventAllColumns,
			atoEventColumnsWithDefault,
			atoEventColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(atoEventType, atoEventMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(atoEventType, atoEventMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"ato_events\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"ato_events\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"ato_events\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, atoEventPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into ato_events")
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
		return errors.Wrap(err, "models: unable to populate default values for ato_events")
	}

CacheNoHooks:
	if !cached {
		atoEventInsertCacheMut.Lock()
		atoEventInsertCache[key] = cache
		atoEventInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AtoEvent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AtoEvent) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	atoEventUpdateCacheMut.RLock()
	cache, cached := atoEventUpdateCache[key]
	atoEventUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			atoEventAllColumns,
			atoEventPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update ato_events, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"ato_events\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, atoEventPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(atoEventType, atoEventMapping, append(wl, atoEventPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update ato_events row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for ato_events")
	}

	if !cached {
		atoEventUpdateCacheMut.Lock()
		atoEventUpdateCache[key] = cache
		atoEventUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q atoEventQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for ato_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for ato_events")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AtoEventSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), atoEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"ato_events\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, atoEventPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in atoEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all atoEvent")
	}
	return rowsAff, nil
}

// Delete deletes a single AtoEvent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AtoEvent) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AtoEvent provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), atoEventPrimaryKeyMapping)
	sql := "DELETE FROM \"ato_events\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from ato_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for ato_events")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q atoEventQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no atoEventQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ato_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ato_events")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AtoEventSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(atoEventBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), atoEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"ato_events\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, atoEventPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from atoEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ato_events")
	}

	if len(atoEventAfterDeleteHooks) != 0 {
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
func (o *AtoEvent) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAtoEvent(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AtoEventSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AtoEventSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), atoEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"ato_events\".* FROM \"ato_events\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, atoEventPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AtoEventSlice")
	}

	*o = slice

	return nil
}

// AtoEventExists checks if the AtoEvent row exists.
func AtoEventExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"ato_events\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ato_events exists")
	}

	return exists, nil
}
