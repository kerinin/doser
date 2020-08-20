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

// Firmata is an object representing the database table.
type Firmata struct {
	ID         string `boil:"id" json:"id" toml:"id" yaml:"id"`
	SerialPort string `boil:"serial_port" json:"serial_port" toml:"serial_port" yaml:"serial_port"`

	R *firmataR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L firmataL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FirmataColumns = struct {
	ID         string
	SerialPort string
}{
	ID:         "id",
	SerialPort: "serial_port",
}

// Generated where

var FirmataWhere = struct {
	ID         whereHelperstring
	SerialPort whereHelperstring
}{
	ID:         whereHelperstring{field: "\"firmatas\".\"id\""},
	SerialPort: whereHelperstring{field: "\"firmatas\".\"serial_port\""},
}

// FirmataRels is where relationship names are stored.
var FirmataRels = struct {
	FirmatumPumps string
}{
	FirmatumPumps: "FirmatumPumps",
}

// firmataR is where relationships are stored.
type firmataR struct {
	FirmatumPumps PumpSlice `boil:"FirmatumPumps" json:"FirmatumPumps" toml:"FirmatumPumps" yaml:"FirmatumPumps"`
}

// NewStruct creates a new relationship struct
func (*firmataR) NewStruct() *firmataR {
	return &firmataR{}
}

// firmataL is where Load methods for each relationship are stored.
type firmataL struct{}

var (
	firmataAllColumns            = []string{"id", "serial_port"}
	firmataColumnsWithoutDefault = []string{"id", "serial_port"}
	firmataColumnsWithDefault    = []string{}
	firmataPrimaryKeyColumns     = []string{"id"}
)

type (
	// FirmataSlice is an alias for a slice of pointers to Firmata.
	// This should generally be used opposed to []Firmata.
	FirmataSlice []*Firmata
	// FirmataHook is the signature for custom Firmata hook methods
	FirmataHook func(context.Context, boil.ContextExecutor, *Firmata) error

	firmataQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	firmataType                 = reflect.TypeOf(&Firmata{})
	firmataMapping              = queries.MakeStructMapping(firmataType)
	firmataPrimaryKeyMapping, _ = queries.BindMapping(firmataType, firmataMapping, firmataPrimaryKeyColumns)
	firmataInsertCacheMut       sync.RWMutex
	firmataInsertCache          = make(map[string]insertCache)
	firmataUpdateCacheMut       sync.RWMutex
	firmataUpdateCache          = make(map[string]updateCache)
	firmataUpsertCacheMut       sync.RWMutex
	firmataUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var firmataBeforeInsertHooks []FirmataHook
var firmataBeforeUpdateHooks []FirmataHook
var firmataBeforeDeleteHooks []FirmataHook
var firmataBeforeUpsertHooks []FirmataHook

var firmataAfterInsertHooks []FirmataHook
var firmataAfterSelectHooks []FirmataHook
var firmataAfterUpdateHooks []FirmataHook
var firmataAfterDeleteHooks []FirmataHook
var firmataAfterUpsertHooks []FirmataHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Firmata) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmataBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Firmata) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmataBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Firmata) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmataBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Firmata) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmataBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Firmata) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmataAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Firmata) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmataAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Firmata) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmataAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Firmata) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmataAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Firmata) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range firmataAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFirmataHook registers your hook function for all future operations.
func AddFirmataHook(hookPoint boil.HookPoint, firmataHook FirmataHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		firmataBeforeInsertHooks = append(firmataBeforeInsertHooks, firmataHook)
	case boil.BeforeUpdateHook:
		firmataBeforeUpdateHooks = append(firmataBeforeUpdateHooks, firmataHook)
	case boil.BeforeDeleteHook:
		firmataBeforeDeleteHooks = append(firmataBeforeDeleteHooks, firmataHook)
	case boil.BeforeUpsertHook:
		firmataBeforeUpsertHooks = append(firmataBeforeUpsertHooks, firmataHook)
	case boil.AfterInsertHook:
		firmataAfterInsertHooks = append(firmataAfterInsertHooks, firmataHook)
	case boil.AfterSelectHook:
		firmataAfterSelectHooks = append(firmataAfterSelectHooks, firmataHook)
	case boil.AfterUpdateHook:
		firmataAfterUpdateHooks = append(firmataAfterUpdateHooks, firmataHook)
	case boil.AfterDeleteHook:
		firmataAfterDeleteHooks = append(firmataAfterDeleteHooks, firmataHook)
	case boil.AfterUpsertHook:
		firmataAfterUpsertHooks = append(firmataAfterUpsertHooks, firmataHook)
	}
}

// One returns a single firmata record from the query.
func (q firmataQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Firmata, error) {
	o := &Firmata{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for firmatas")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Firmata records from the query.
func (q firmataQuery) All(ctx context.Context, exec boil.ContextExecutor) (FirmataSlice, error) {
	var o []*Firmata

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Firmata slice")
	}

	if len(firmataAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Firmata records in the query.
func (q firmataQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count firmatas rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q firmataQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if firmatas exists")
	}

	return count > 0, nil
}

// FirmatumPumps retrieves all the pump's Pumps with an executor via firmata_id column.
func (o *Firmata) FirmatumPumps(mods ...qm.QueryMod) pumpQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"pumps\".\"firmata_id\"=?", o.ID),
	)

	query := Pumps(queryMods...)
	queries.SetFrom(query.Query, "\"pumps\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"pumps\".*"})
	}

	return query
}

// LoadFirmatumPumps allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (firmataL) LoadFirmatumPumps(ctx context.Context, e boil.ContextExecutor, singular bool, maybeFirmata interface{}, mods queries.Applicator) error {
	var slice []*Firmata
	var object *Firmata

	if singular {
		object = maybeFirmata.(*Firmata)
	} else {
		slice = *maybeFirmata.(*[]*Firmata)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &firmataR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &firmataR{}
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
		qm.From(`pumps`),
		qm.WhereIn(`pumps.firmata_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load pumps")
	}

	var resultSlice []*Pump
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice pumps")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on pumps")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for pumps")
	}

	if len(pumpAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.FirmatumPumps = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &pumpR{}
			}
			foreign.R.Firmatum = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.FirmataID {
				local.R.FirmatumPumps = append(local.R.FirmatumPumps, foreign)
				if foreign.R == nil {
					foreign.R = &pumpR{}
				}
				foreign.R.Firmatum = local
				break
			}
		}
	}

	return nil
}

// AddFirmatumPumps adds the given related objects to the existing relationships
// of the firmatas, optionally inserting them as new records.
// Appends related to o.R.FirmatumPumps.
// Sets related.R.Firmatum appropriately.
func (o *Firmata) AddFirmatumPumps(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Pump) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.FirmataID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"pumps\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"firmata_id"}),
				strmangle.WhereClause("\"", "\"", 0, pumpPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.FirmataID = o.ID
		}
	}

	if o.R == nil {
		o.R = &firmataR{
			FirmatumPumps: related,
		}
	} else {
		o.R.FirmatumPumps = append(o.R.FirmatumPumps, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &pumpR{
				Firmatum: o,
			}
		} else {
			rel.R.Firmatum = o
		}
	}
	return nil
}

// Firmatas retrieves all the records using an executor.
func Firmatas(mods ...qm.QueryMod) firmataQuery {
	mods = append(mods, qm.From("\"firmatas\""))
	return firmataQuery{NewQuery(mods...)}
}

// FindFirmata retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFirmata(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Firmata, error) {
	firmataObj := &Firmata{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"firmatas\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, firmataObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from firmatas")
	}

	return firmataObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Firmata) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no firmatas provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(firmataColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	firmataInsertCacheMut.RLock()
	cache, cached := firmataInsertCache[key]
	firmataInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			firmataAllColumns,
			firmataColumnsWithDefault,
			firmataColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(firmataType, firmataMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(firmataType, firmataMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"firmatas\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"firmatas\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"firmatas\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, firmataPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into firmatas")
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
		return errors.Wrap(err, "models: unable to populate default values for firmatas")
	}

CacheNoHooks:
	if !cached {
		firmataInsertCacheMut.Lock()
		firmataInsertCache[key] = cache
		firmataInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Firmata.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Firmata) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	firmataUpdateCacheMut.RLock()
	cache, cached := firmataUpdateCache[key]
	firmataUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			firmataAllColumns,
			firmataPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update firmatas, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"firmatas\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, firmataPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(firmataType, firmataMapping, append(wl, firmataPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update firmatas row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for firmatas")
	}

	if !cached {
		firmataUpdateCacheMut.Lock()
		firmataUpdateCache[key] = cache
		firmataUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q firmataQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for firmatas")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for firmatas")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FirmataSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), firmataPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"firmatas\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, firmataPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in firmata slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all firmata")
	}
	return rowsAff, nil
}

// Delete deletes a single Firmata record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Firmata) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Firmata provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), firmataPrimaryKeyMapping)
	sql := "DELETE FROM \"firmatas\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from firmatas")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for firmatas")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q firmataQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no firmataQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from firmatas")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for firmatas")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FirmataSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(firmataBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), firmataPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"firmatas\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, firmataPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from firmata slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for firmatas")
	}

	if len(firmataAfterDeleteHooks) != 0 {
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
func (o *Firmata) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFirmata(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FirmataSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FirmataSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), firmataPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"firmatas\".* FROM \"firmatas\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, firmataPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FirmataSlice")
	}

	*o = slice

	return nil
}

// FirmataExists checks if the Firmata row exists.
func FirmataExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"firmatas\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if firmatas exists")
	}

	return exists, nil
}
