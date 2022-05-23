// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
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

// DopsEvaluationsSurgery is an object representing the database table.
type DopsEvaluationsSurgery struct {
	ID               string `boil:"id" json:"id" toml:"id" yaml:"id"`
	SurgeryID        string `boil:"surgery_id" json:"surgery_id" toml:"surgery_id" yaml:"surgery_id"`
	DopsEvaluationID string `boil:"dops_evaluation_id" json:"dops_evaluation_id" toml:"dops_evaluation_id" yaml:"dops_evaluation_id"`

	R *dopsEvaluationsSurgeryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dopsEvaluationsSurgeryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DopsEvaluationsSurgeryColumns = struct {
	ID               string
	SurgeryID        string
	DopsEvaluationID string
}{
	ID:               "id",
	SurgeryID:        "surgery_id",
	DopsEvaluationID: "dops_evaluation_id",
}

var DopsEvaluationsSurgeryTableColumns = struct {
	ID               string
	SurgeryID        string
	DopsEvaluationID string
}{
	ID:               "dops_evaluations_surgeries.id",
	SurgeryID:        "dops_evaluations_surgeries.surgery_id",
	DopsEvaluationID: "dops_evaluations_surgeries.dops_evaluation_id",
}

// Generated where

var DopsEvaluationsSurgeryWhere = struct {
	ID               whereHelperstring
	SurgeryID        whereHelperstring
	DopsEvaluationID whereHelperstring
}{
	ID:               whereHelperstring{field: "\"dops_evaluations_surgeries\".\"id\""},
	SurgeryID:        whereHelperstring{field: "\"dops_evaluations_surgeries\".\"surgery_id\""},
	DopsEvaluationID: whereHelperstring{field: "\"dops_evaluations_surgeries\".\"dops_evaluation_id\""},
}

// DopsEvaluationsSurgeryRels is where relationship names are stored.
var DopsEvaluationsSurgeryRels = struct {
	DopsEvaluation string
	Surgery        string
}{
	DopsEvaluation: "DopsEvaluation",
	Surgery:        "Surgery",
}

// dopsEvaluationsSurgeryR is where relationships are stored.
type dopsEvaluationsSurgeryR struct {
	DopsEvaluation *DopsEvaluation `boil:"DopsEvaluation" json:"DopsEvaluation" toml:"DopsEvaluation" yaml:"DopsEvaluation"`
	Surgery        *Surgery        `boil:"Surgery" json:"Surgery" toml:"Surgery" yaml:"Surgery"`
}

// NewStruct creates a new relationship struct
func (*dopsEvaluationsSurgeryR) NewStruct() *dopsEvaluationsSurgeryR {
	return &dopsEvaluationsSurgeryR{}
}

// dopsEvaluationsSurgeryL is where Load methods for each relationship are stored.
type dopsEvaluationsSurgeryL struct{}

var (
	dopsEvaluationsSurgeryAllColumns            = []string{"id", "surgery_id", "dops_evaluation_id"}
	dopsEvaluationsSurgeryColumnsWithoutDefault = []string{"surgery_id", "dops_evaluation_id"}
	dopsEvaluationsSurgeryColumnsWithDefault    = []string{"id"}
	dopsEvaluationsSurgeryPrimaryKeyColumns     = []string{"id"}
)

type (
	// DopsEvaluationsSurgerySlice is an alias for a slice of pointers to DopsEvaluationsSurgery.
	// This should almost always be used instead of []DopsEvaluationsSurgery.
	DopsEvaluationsSurgerySlice []*DopsEvaluationsSurgery
	// DopsEvaluationsSurgeryHook is the signature for custom DopsEvaluationsSurgery hook methods
	DopsEvaluationsSurgeryHook func(context.Context, boil.ContextExecutor, *DopsEvaluationsSurgery) error

	dopsEvaluationsSurgeryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dopsEvaluationsSurgeryType                 = reflect.TypeOf(&DopsEvaluationsSurgery{})
	dopsEvaluationsSurgeryMapping              = queries.MakeStructMapping(dopsEvaluationsSurgeryType)
	dopsEvaluationsSurgeryPrimaryKeyMapping, _ = queries.BindMapping(dopsEvaluationsSurgeryType, dopsEvaluationsSurgeryMapping, dopsEvaluationsSurgeryPrimaryKeyColumns)
	dopsEvaluationsSurgeryInsertCacheMut       sync.RWMutex
	dopsEvaluationsSurgeryInsertCache          = make(map[string]insertCache)
	dopsEvaluationsSurgeryUpdateCacheMut       sync.RWMutex
	dopsEvaluationsSurgeryUpdateCache          = make(map[string]updateCache)
	dopsEvaluationsSurgeryUpsertCacheMut       sync.RWMutex
	dopsEvaluationsSurgeryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dopsEvaluationsSurgeryBeforeInsertHooks []DopsEvaluationsSurgeryHook
var dopsEvaluationsSurgeryBeforeUpdateHooks []DopsEvaluationsSurgeryHook
var dopsEvaluationsSurgeryBeforeDeleteHooks []DopsEvaluationsSurgeryHook
var dopsEvaluationsSurgeryBeforeUpsertHooks []DopsEvaluationsSurgeryHook

var dopsEvaluationsSurgeryAfterInsertHooks []DopsEvaluationsSurgeryHook
var dopsEvaluationsSurgeryAfterSelectHooks []DopsEvaluationsSurgeryHook
var dopsEvaluationsSurgeryAfterUpdateHooks []DopsEvaluationsSurgeryHook
var dopsEvaluationsSurgeryAfterDeleteHooks []DopsEvaluationsSurgeryHook
var dopsEvaluationsSurgeryAfterUpsertHooks []DopsEvaluationsSurgeryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DopsEvaluationsSurgery) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dopsEvaluationsSurgeryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DopsEvaluationsSurgery) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dopsEvaluationsSurgeryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DopsEvaluationsSurgery) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dopsEvaluationsSurgeryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DopsEvaluationsSurgery) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dopsEvaluationsSurgeryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DopsEvaluationsSurgery) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dopsEvaluationsSurgeryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DopsEvaluationsSurgery) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dopsEvaluationsSurgeryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DopsEvaluationsSurgery) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dopsEvaluationsSurgeryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DopsEvaluationsSurgery) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dopsEvaluationsSurgeryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DopsEvaluationsSurgery) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dopsEvaluationsSurgeryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDopsEvaluationsSurgeryHook registers your hook function for all future operations.
func AddDopsEvaluationsSurgeryHook(hookPoint boil.HookPoint, dopsEvaluationsSurgeryHook DopsEvaluationsSurgeryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		dopsEvaluationsSurgeryBeforeInsertHooks = append(dopsEvaluationsSurgeryBeforeInsertHooks, dopsEvaluationsSurgeryHook)
	case boil.BeforeUpdateHook:
		dopsEvaluationsSurgeryBeforeUpdateHooks = append(dopsEvaluationsSurgeryBeforeUpdateHooks, dopsEvaluationsSurgeryHook)
	case boil.BeforeDeleteHook:
		dopsEvaluationsSurgeryBeforeDeleteHooks = append(dopsEvaluationsSurgeryBeforeDeleteHooks, dopsEvaluationsSurgeryHook)
	case boil.BeforeUpsertHook:
		dopsEvaluationsSurgeryBeforeUpsertHooks = append(dopsEvaluationsSurgeryBeforeUpsertHooks, dopsEvaluationsSurgeryHook)
	case boil.AfterInsertHook:
		dopsEvaluationsSurgeryAfterInsertHooks = append(dopsEvaluationsSurgeryAfterInsertHooks, dopsEvaluationsSurgeryHook)
	case boil.AfterSelectHook:
		dopsEvaluationsSurgeryAfterSelectHooks = append(dopsEvaluationsSurgeryAfterSelectHooks, dopsEvaluationsSurgeryHook)
	case boil.AfterUpdateHook:
		dopsEvaluationsSurgeryAfterUpdateHooks = append(dopsEvaluationsSurgeryAfterUpdateHooks, dopsEvaluationsSurgeryHook)
	case boil.AfterDeleteHook:
		dopsEvaluationsSurgeryAfterDeleteHooks = append(dopsEvaluationsSurgeryAfterDeleteHooks, dopsEvaluationsSurgeryHook)
	case boil.AfterUpsertHook:
		dopsEvaluationsSurgeryAfterUpsertHooks = append(dopsEvaluationsSurgeryAfterUpsertHooks, dopsEvaluationsSurgeryHook)
	}
}

// One returns a single dopsEvaluationsSurgery record from the query.
func (q dopsEvaluationsSurgeryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DopsEvaluationsSurgery, error) {
	o := &DopsEvaluationsSurgery{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for dops_evaluations_surgeries")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DopsEvaluationsSurgery records from the query.
func (q dopsEvaluationsSurgeryQuery) All(ctx context.Context, exec boil.ContextExecutor) (DopsEvaluationsSurgerySlice, error) {
	var o []*DopsEvaluationsSurgery

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DopsEvaluationsSurgery slice")
	}

	if len(dopsEvaluationsSurgeryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DopsEvaluationsSurgery records in the query.
func (q dopsEvaluationsSurgeryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count dops_evaluations_surgeries rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dopsEvaluationsSurgeryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if dops_evaluations_surgeries exists")
	}

	return count > 0, nil
}

// DopsEvaluation pointed to by the foreign key.
func (o *DopsEvaluationsSurgery) DopsEvaluation(mods ...qm.QueryMod) dopsEvaluationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.DopsEvaluationID),
	}

	queryMods = append(queryMods, mods...)

	query := DopsEvaluations(queryMods...)
	queries.SetFrom(query.Query, "\"dops_evaluations\"")

	return query
}

// Surgery pointed to by the foreign key.
func (o *DopsEvaluationsSurgery) Surgery(mods ...qm.QueryMod) surgeryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SurgeryID),
	}

	queryMods = append(queryMods, mods...)

	query := Surgeries(queryMods...)
	queries.SetFrom(query.Query, "\"surgeries\"")

	return query
}

// LoadDopsEvaluation allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dopsEvaluationsSurgeryL) LoadDopsEvaluation(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDopsEvaluationsSurgery interface{}, mods queries.Applicator) error {
	var slice []*DopsEvaluationsSurgery
	var object *DopsEvaluationsSurgery

	if singular {
		object = maybeDopsEvaluationsSurgery.(*DopsEvaluationsSurgery)
	} else {
		slice = *maybeDopsEvaluationsSurgery.(*[]*DopsEvaluationsSurgery)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dopsEvaluationsSurgeryR{}
		}
		args = append(args, object.DopsEvaluationID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dopsEvaluationsSurgeryR{}
			}

			for _, a := range args {
				if a == obj.DopsEvaluationID {
					continue Outer
				}
			}

			args = append(args, obj.DopsEvaluationID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`dops_evaluations`),
		qm.WhereIn(`dops_evaluations.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DopsEvaluation")
	}

	var resultSlice []*DopsEvaluation
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DopsEvaluation")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for dops_evaluations")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dops_evaluations")
	}

	if len(dopsEvaluationsSurgeryAfterSelectHooks) != 0 {
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
		object.R.DopsEvaluation = foreign
		if foreign.R == nil {
			foreign.R = &dopsEvaluationR{}
		}
		foreign.R.DopsEvaluationsSurgeries = append(foreign.R.DopsEvaluationsSurgeries, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.DopsEvaluationID == foreign.ID {
				local.R.DopsEvaluation = foreign
				if foreign.R == nil {
					foreign.R = &dopsEvaluationR{}
				}
				foreign.R.DopsEvaluationsSurgeries = append(foreign.R.DopsEvaluationsSurgeries, local)
				break
			}
		}
	}

	return nil
}

// LoadSurgery allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dopsEvaluationsSurgeryL) LoadSurgery(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDopsEvaluationsSurgery interface{}, mods queries.Applicator) error {
	var slice []*DopsEvaluationsSurgery
	var object *DopsEvaluationsSurgery

	if singular {
		object = maybeDopsEvaluationsSurgery.(*DopsEvaluationsSurgery)
	} else {
		slice = *maybeDopsEvaluationsSurgery.(*[]*DopsEvaluationsSurgery)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dopsEvaluationsSurgeryR{}
		}
		args = append(args, object.SurgeryID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dopsEvaluationsSurgeryR{}
			}

			for _, a := range args {
				if a == obj.SurgeryID {
					continue Outer
				}
			}

			args = append(args, obj.SurgeryID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`surgeries`),
		qm.WhereIn(`surgeries.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Surgery")
	}

	var resultSlice []*Surgery
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Surgery")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for surgeries")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for surgeries")
	}

	if len(dopsEvaluationsSurgeryAfterSelectHooks) != 0 {
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
		object.R.Surgery = foreign
		if foreign.R == nil {
			foreign.R = &surgeryR{}
		}
		foreign.R.DopsEvaluationsSurgeries = append(foreign.R.DopsEvaluationsSurgeries, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SurgeryID == foreign.ID {
				local.R.Surgery = foreign
				if foreign.R == nil {
					foreign.R = &surgeryR{}
				}
				foreign.R.DopsEvaluationsSurgeries = append(foreign.R.DopsEvaluationsSurgeries, local)
				break
			}
		}
	}

	return nil
}

// SetDopsEvaluation of the dopsEvaluationsSurgery to the related item.
// Sets o.R.DopsEvaluation to related.
// Adds o to related.R.DopsEvaluationsSurgeries.
func (o *DopsEvaluationsSurgery) SetDopsEvaluation(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DopsEvaluation) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"dops_evaluations_surgeries\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"dops_evaluation_id"}),
		strmangle.WhereClause("\"", "\"", 2, dopsEvaluationsSurgeryPrimaryKeyColumns),
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

	o.DopsEvaluationID = related.ID
	if o.R == nil {
		o.R = &dopsEvaluationsSurgeryR{
			DopsEvaluation: related,
		}
	} else {
		o.R.DopsEvaluation = related
	}

	if related.R == nil {
		related.R = &dopsEvaluationR{
			DopsEvaluationsSurgeries: DopsEvaluationsSurgerySlice{o},
		}
	} else {
		related.R.DopsEvaluationsSurgeries = append(related.R.DopsEvaluationsSurgeries, o)
	}

	return nil
}

// SetSurgery of the dopsEvaluationsSurgery to the related item.
// Sets o.R.Surgery to related.
// Adds o to related.R.DopsEvaluationsSurgeries.
func (o *DopsEvaluationsSurgery) SetSurgery(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Surgery) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"dops_evaluations_surgeries\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"surgery_id"}),
		strmangle.WhereClause("\"", "\"", 2, dopsEvaluationsSurgeryPrimaryKeyColumns),
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

	o.SurgeryID = related.ID
	if o.R == nil {
		o.R = &dopsEvaluationsSurgeryR{
			Surgery: related,
		}
	} else {
		o.R.Surgery = related
	}

	if related.R == nil {
		related.R = &surgeryR{
			DopsEvaluationsSurgeries: DopsEvaluationsSurgerySlice{o},
		}
	} else {
		related.R.DopsEvaluationsSurgeries = append(related.R.DopsEvaluationsSurgeries, o)
	}

	return nil
}

// DopsEvaluationsSurgeries retrieves all the records using an executor.
func DopsEvaluationsSurgeries(mods ...qm.QueryMod) dopsEvaluationsSurgeryQuery {
	mods = append(mods, qm.From("\"dops_evaluations_surgeries\""))
	return dopsEvaluationsSurgeryQuery{NewQuery(mods...)}
}

// FindDopsEvaluationsSurgery retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDopsEvaluationsSurgery(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*DopsEvaluationsSurgery, error) {
	dopsEvaluationsSurgeryObj := &DopsEvaluationsSurgery{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"dops_evaluations_surgeries\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, dopsEvaluationsSurgeryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from dops_evaluations_surgeries")
	}

	if err = dopsEvaluationsSurgeryObj.doAfterSelectHooks(ctx, exec); err != nil {
		return dopsEvaluationsSurgeryObj, err
	}

	return dopsEvaluationsSurgeryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DopsEvaluationsSurgery) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no dops_evaluations_surgeries provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dopsEvaluationsSurgeryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dopsEvaluationsSurgeryInsertCacheMut.RLock()
	cache, cached := dopsEvaluationsSurgeryInsertCache[key]
	dopsEvaluationsSurgeryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dopsEvaluationsSurgeryAllColumns,
			dopsEvaluationsSurgeryColumnsWithDefault,
			dopsEvaluationsSurgeryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dopsEvaluationsSurgeryType, dopsEvaluationsSurgeryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dopsEvaluationsSurgeryType, dopsEvaluationsSurgeryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"dops_evaluations_surgeries\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"dops_evaluations_surgeries\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into dops_evaluations_surgeries")
	}

	if !cached {
		dopsEvaluationsSurgeryInsertCacheMut.Lock()
		dopsEvaluationsSurgeryInsertCache[key] = cache
		dopsEvaluationsSurgeryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DopsEvaluationsSurgery.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DopsEvaluationsSurgery) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	dopsEvaluationsSurgeryUpdateCacheMut.RLock()
	cache, cached := dopsEvaluationsSurgeryUpdateCache[key]
	dopsEvaluationsSurgeryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dopsEvaluationsSurgeryAllColumns,
			dopsEvaluationsSurgeryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update dops_evaluations_surgeries, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"dops_evaluations_surgeries\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dopsEvaluationsSurgeryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dopsEvaluationsSurgeryType, dopsEvaluationsSurgeryMapping, append(wl, dopsEvaluationsSurgeryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update dops_evaluations_surgeries row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for dops_evaluations_surgeries")
	}

	if !cached {
		dopsEvaluationsSurgeryUpdateCacheMut.Lock()
		dopsEvaluationsSurgeryUpdateCache[key] = cache
		dopsEvaluationsSurgeryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dopsEvaluationsSurgeryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for dops_evaluations_surgeries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for dops_evaluations_surgeries")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DopsEvaluationsSurgerySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dopsEvaluationsSurgeryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"dops_evaluations_surgeries\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, dopsEvaluationsSurgeryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in dopsEvaluationsSurgery slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all dopsEvaluationsSurgery")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DopsEvaluationsSurgery) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no dops_evaluations_surgeries provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dopsEvaluationsSurgeryColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	dopsEvaluationsSurgeryUpsertCacheMut.RLock()
	cache, cached := dopsEvaluationsSurgeryUpsertCache[key]
	dopsEvaluationsSurgeryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dopsEvaluationsSurgeryAllColumns,
			dopsEvaluationsSurgeryColumnsWithDefault,
			dopsEvaluationsSurgeryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			dopsEvaluationsSurgeryAllColumns,
			dopsEvaluationsSurgeryPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert dops_evaluations_surgeries, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dopsEvaluationsSurgeryPrimaryKeyColumns))
			copy(conflict, dopsEvaluationsSurgeryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"dops_evaluations_surgeries\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(dopsEvaluationsSurgeryType, dopsEvaluationsSurgeryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dopsEvaluationsSurgeryType, dopsEvaluationsSurgeryMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert dops_evaluations_surgeries")
	}

	if !cached {
		dopsEvaluationsSurgeryUpsertCacheMut.Lock()
		dopsEvaluationsSurgeryUpsertCache[key] = cache
		dopsEvaluationsSurgeryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DopsEvaluationsSurgery record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DopsEvaluationsSurgery) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DopsEvaluationsSurgery provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dopsEvaluationsSurgeryPrimaryKeyMapping)
	sql := "DELETE FROM \"dops_evaluations_surgeries\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from dops_evaluations_surgeries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for dops_evaluations_surgeries")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dopsEvaluationsSurgeryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no dopsEvaluationsSurgeryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from dops_evaluations_surgeries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for dops_evaluations_surgeries")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DopsEvaluationsSurgerySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(dopsEvaluationsSurgeryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dopsEvaluationsSurgeryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"dops_evaluations_surgeries\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dopsEvaluationsSurgeryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from dopsEvaluationsSurgery slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for dops_evaluations_surgeries")
	}

	if len(dopsEvaluationsSurgeryAfterDeleteHooks) != 0 {
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
func (o *DopsEvaluationsSurgery) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDopsEvaluationsSurgery(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DopsEvaluationsSurgerySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DopsEvaluationsSurgerySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dopsEvaluationsSurgeryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"dops_evaluations_surgeries\".* FROM \"dops_evaluations_surgeries\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dopsEvaluationsSurgeryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DopsEvaluationsSurgerySlice")
	}

	*o = slice

	return nil
}

// DopsEvaluationsSurgeryExists checks if the DopsEvaluationsSurgery row exists.
func DopsEvaluationsSurgeryExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"dops_evaluations_surgeries\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if dops_evaluations_surgeries exists")
	}

	return exists, nil
}