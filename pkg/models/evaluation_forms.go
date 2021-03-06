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
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// EvaluationForm is an object representing the database table.
type EvaluationForm struct {
	ID           int               `boil:"id" json:"id" toml:"id" yaml:"id"`
	DepartmentID string            `boil:"department_id" json:"department_id" toml:"department_id" yaml:"department_id"`
	Name         string            `boil:"name" json:"name" toml:"name" yaml:"name"`
	Annotations  types.JSON        `boil:"annotations" json:"annotations" toml:"annotations" yaml:"annotations"`
	Difficulty   types.StringArray `boil:"difficulty" json:"difficulty" toml:"difficulty" yaml:"difficulty"`
	Citations    types.StringArray `boil:"citations" json:"citations" toml:"citations" yaml:"citations"`

	R *evaluationFormR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L evaluationFormL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EvaluationFormColumns = struct {
	ID           string
	DepartmentID string
	Name         string
	Annotations  string
	Difficulty   string
	Citations    string
}{
	ID:           "id",
	DepartmentID: "department_id",
	Name:         "name",
	Annotations:  "annotations",
	Difficulty:   "difficulty",
	Citations:    "citations",
}

var EvaluationFormTableColumns = struct {
	ID           string
	DepartmentID string
	Name         string
	Annotations  string
	Difficulty   string
	Citations    string
}{
	ID:           "evaluation_forms.id",
	DepartmentID: "evaluation_forms.department_id",
	Name:         "evaluation_forms.name",
	Annotations:  "evaluation_forms.annotations",
	Difficulty:   "evaluation_forms.difficulty",
	Citations:    "evaluation_forms.citations",
}

// Generated where

type whereHelpertypes_StringArray struct{ field string }

func (w whereHelpertypes_StringArray) EQ(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_StringArray) NEQ(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_StringArray) LT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_StringArray) LTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_StringArray) GT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_StringArray) GTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var EvaluationFormWhere = struct {
	ID           whereHelperint
	DepartmentID whereHelperstring
	Name         whereHelperstring
	Annotations  whereHelpertypes_JSON
	Difficulty   whereHelpertypes_StringArray
	Citations    whereHelpertypes_StringArray
}{
	ID:           whereHelperint{field: "\"evaluation_forms\".\"id\""},
	DepartmentID: whereHelperstring{field: "\"evaluation_forms\".\"department_id\""},
	Name:         whereHelperstring{field: "\"evaluation_forms\".\"name\""},
	Annotations:  whereHelpertypes_JSON{field: "\"evaluation_forms\".\"annotations\""},
	Difficulty:   whereHelpertypes_StringArray{field: "\"evaluation_forms\".\"difficulty\""},
	Citations:    whereHelpertypes_StringArray{field: "\"evaluation_forms\".\"citations\""},
}

// EvaluationFormRels is where relationship names are stored.
var EvaluationFormRels = struct {
	Department string
}{
	Department: "Department",
}

// evaluationFormR is where relationships are stored.
type evaluationFormR struct {
	Department *OrganizationalUnit `boil:"Department" json:"Department" toml:"Department" yaml:"Department"`
}

// NewStruct creates a new relationship struct
func (*evaluationFormR) NewStruct() *evaluationFormR {
	return &evaluationFormR{}
}

// evaluationFormL is where Load methods for each relationship are stored.
type evaluationFormL struct{}

var (
	evaluationFormAllColumns            = []string{"id", "department_id", "name", "annotations", "difficulty", "citations"}
	evaluationFormColumnsWithoutDefault = []string{"department_id", "name"}
	evaluationFormColumnsWithDefault    = []string{"id", "annotations", "difficulty", "citations"}
	evaluationFormPrimaryKeyColumns     = []string{"id"}
)

type (
	// EvaluationFormSlice is an alias for a slice of pointers to EvaluationForm.
	// This should almost always be used instead of []EvaluationForm.
	EvaluationFormSlice []*EvaluationForm
	// EvaluationFormHook is the signature for custom EvaluationForm hook methods
	EvaluationFormHook func(context.Context, boil.ContextExecutor, *EvaluationForm) error

	evaluationFormQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	evaluationFormType                 = reflect.TypeOf(&EvaluationForm{})
	evaluationFormMapping              = queries.MakeStructMapping(evaluationFormType)
	evaluationFormPrimaryKeyMapping, _ = queries.BindMapping(evaluationFormType, evaluationFormMapping, evaluationFormPrimaryKeyColumns)
	evaluationFormInsertCacheMut       sync.RWMutex
	evaluationFormInsertCache          = make(map[string]insertCache)
	evaluationFormUpdateCacheMut       sync.RWMutex
	evaluationFormUpdateCache          = make(map[string]updateCache)
	evaluationFormUpsertCacheMut       sync.RWMutex
	evaluationFormUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var evaluationFormBeforeInsertHooks []EvaluationFormHook
var evaluationFormBeforeUpdateHooks []EvaluationFormHook
var evaluationFormBeforeDeleteHooks []EvaluationFormHook
var evaluationFormBeforeUpsertHooks []EvaluationFormHook

var evaluationFormAfterInsertHooks []EvaluationFormHook
var evaluationFormAfterSelectHooks []EvaluationFormHook
var evaluationFormAfterUpdateHooks []EvaluationFormHook
var evaluationFormAfterDeleteHooks []EvaluationFormHook
var evaluationFormAfterUpsertHooks []EvaluationFormHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *EvaluationForm) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range evaluationFormBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *EvaluationForm) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range evaluationFormBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *EvaluationForm) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range evaluationFormBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *EvaluationForm) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range evaluationFormBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *EvaluationForm) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range evaluationFormAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *EvaluationForm) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range evaluationFormAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *EvaluationForm) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range evaluationFormAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *EvaluationForm) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range evaluationFormAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *EvaluationForm) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range evaluationFormAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEvaluationFormHook registers your hook function for all future operations.
func AddEvaluationFormHook(hookPoint boil.HookPoint, evaluationFormHook EvaluationFormHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		evaluationFormBeforeInsertHooks = append(evaluationFormBeforeInsertHooks, evaluationFormHook)
	case boil.BeforeUpdateHook:
		evaluationFormBeforeUpdateHooks = append(evaluationFormBeforeUpdateHooks, evaluationFormHook)
	case boil.BeforeDeleteHook:
		evaluationFormBeforeDeleteHooks = append(evaluationFormBeforeDeleteHooks, evaluationFormHook)
	case boil.BeforeUpsertHook:
		evaluationFormBeforeUpsertHooks = append(evaluationFormBeforeUpsertHooks, evaluationFormHook)
	case boil.AfterInsertHook:
		evaluationFormAfterInsertHooks = append(evaluationFormAfterInsertHooks, evaluationFormHook)
	case boil.AfterSelectHook:
		evaluationFormAfterSelectHooks = append(evaluationFormAfterSelectHooks, evaluationFormHook)
	case boil.AfterUpdateHook:
		evaluationFormAfterUpdateHooks = append(evaluationFormAfterUpdateHooks, evaluationFormHook)
	case boil.AfterDeleteHook:
		evaluationFormAfterDeleteHooks = append(evaluationFormAfterDeleteHooks, evaluationFormHook)
	case boil.AfterUpsertHook:
		evaluationFormAfterUpsertHooks = append(evaluationFormAfterUpsertHooks, evaluationFormHook)
	}
}

// One returns a single evaluationForm record from the query.
func (q evaluationFormQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EvaluationForm, error) {
	o := &EvaluationForm{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for evaluation_forms")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all EvaluationForm records from the query.
func (q evaluationFormQuery) All(ctx context.Context, exec boil.ContextExecutor) (EvaluationFormSlice, error) {
	var o []*EvaluationForm

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EvaluationForm slice")
	}

	if len(evaluationFormAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all EvaluationForm records in the query.
func (q evaluationFormQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count evaluation_forms rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q evaluationFormQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if evaluation_forms exists")
	}

	return count > 0, nil
}

// Department pointed to by the foreign key.
func (o *EvaluationForm) Department(mods ...qm.QueryMod) organizationalUnitQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.DepartmentID),
	}

	queryMods = append(queryMods, mods...)

	query := OrganizationalUnits(queryMods...)
	queries.SetFrom(query.Query, "\"organizational_units\"")

	return query
}

// LoadDepartment allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (evaluationFormL) LoadDepartment(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEvaluationForm interface{}, mods queries.Applicator) error {
	var slice []*EvaluationForm
	var object *EvaluationForm

	if singular {
		object = maybeEvaluationForm.(*EvaluationForm)
	} else {
		slice = *maybeEvaluationForm.(*[]*EvaluationForm)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &evaluationFormR{}
		}
		args = append(args, object.DepartmentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &evaluationFormR{}
			}

			for _, a := range args {
				if a == obj.DepartmentID {
					continue Outer
				}
			}

			args = append(args, obj.DepartmentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`organizational_units`),
		qm.WhereIn(`organizational_units.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load OrganizationalUnit")
	}

	var resultSlice []*OrganizationalUnit
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice OrganizationalUnit")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for organizational_units")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for organizational_units")
	}

	if len(evaluationFormAfterSelectHooks) != 0 {
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
		object.R.Department = foreign
		if foreign.R == nil {
			foreign.R = &organizationalUnitR{}
		}
		foreign.R.DepartmentEvaluationForms = append(foreign.R.DepartmentEvaluationForms, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.DepartmentID == foreign.ID {
				local.R.Department = foreign
				if foreign.R == nil {
					foreign.R = &organizationalUnitR{}
				}
				foreign.R.DepartmentEvaluationForms = append(foreign.R.DepartmentEvaluationForms, local)
				break
			}
		}
	}

	return nil
}

// SetDepartment of the evaluationForm to the related item.
// Sets o.R.Department to related.
// Adds o to related.R.DepartmentEvaluationForms.
func (o *EvaluationForm) SetDepartment(ctx context.Context, exec boil.ContextExecutor, insert bool, related *OrganizationalUnit) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"evaluation_forms\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"department_id"}),
		strmangle.WhereClause("\"", "\"", 2, evaluationFormPrimaryKeyColumns),
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

	o.DepartmentID = related.ID
	if o.R == nil {
		o.R = &evaluationFormR{
			Department: related,
		}
	} else {
		o.R.Department = related
	}

	if related.R == nil {
		related.R = &organizationalUnitR{
			DepartmentEvaluationForms: EvaluationFormSlice{o},
		}
	} else {
		related.R.DepartmentEvaluationForms = append(related.R.DepartmentEvaluationForms, o)
	}

	return nil
}

// EvaluationForms retrieves all the records using an executor.
func EvaluationForms(mods ...qm.QueryMod) evaluationFormQuery {
	mods = append(mods, qm.From("\"evaluation_forms\""))
	return evaluationFormQuery{NewQuery(mods...)}
}

// FindEvaluationForm retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEvaluationForm(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*EvaluationForm, error) {
	evaluationFormObj := &EvaluationForm{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"evaluation_forms\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, evaluationFormObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from evaluation_forms")
	}

	if err = evaluationFormObj.doAfterSelectHooks(ctx, exec); err != nil {
		return evaluationFormObj, err
	}

	return evaluationFormObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EvaluationForm) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no evaluation_forms provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(evaluationFormColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	evaluationFormInsertCacheMut.RLock()
	cache, cached := evaluationFormInsertCache[key]
	evaluationFormInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			evaluationFormAllColumns,
			evaluationFormColumnsWithDefault,
			evaluationFormColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(evaluationFormType, evaluationFormMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(evaluationFormType, evaluationFormMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"evaluation_forms\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"evaluation_forms\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into evaluation_forms")
	}

	if !cached {
		evaluationFormInsertCacheMut.Lock()
		evaluationFormInsertCache[key] = cache
		evaluationFormInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the EvaluationForm.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EvaluationForm) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	evaluationFormUpdateCacheMut.RLock()
	cache, cached := evaluationFormUpdateCache[key]
	evaluationFormUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			evaluationFormAllColumns,
			evaluationFormPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update evaluation_forms, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"evaluation_forms\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, evaluationFormPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(evaluationFormType, evaluationFormMapping, append(wl, evaluationFormPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update evaluation_forms row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for evaluation_forms")
	}

	if !cached {
		evaluationFormUpdateCacheMut.Lock()
		evaluationFormUpdateCache[key] = cache
		evaluationFormUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q evaluationFormQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for evaluation_forms")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for evaluation_forms")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EvaluationFormSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), evaluationFormPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"evaluation_forms\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, evaluationFormPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in evaluationForm slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all evaluationForm")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EvaluationForm) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no evaluation_forms provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(evaluationFormColumnsWithDefault, o)

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

	evaluationFormUpsertCacheMut.RLock()
	cache, cached := evaluationFormUpsertCache[key]
	evaluationFormUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			evaluationFormAllColumns,
			evaluationFormColumnsWithDefault,
			evaluationFormColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			evaluationFormAllColumns,
			evaluationFormPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert evaluation_forms, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(evaluationFormPrimaryKeyColumns))
			copy(conflict, evaluationFormPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"evaluation_forms\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(evaluationFormType, evaluationFormMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(evaluationFormType, evaluationFormMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert evaluation_forms")
	}

	if !cached {
		evaluationFormUpsertCacheMut.Lock()
		evaluationFormUpsertCache[key] = cache
		evaluationFormUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single EvaluationForm record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EvaluationForm) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EvaluationForm provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), evaluationFormPrimaryKeyMapping)
	sql := "DELETE FROM \"evaluation_forms\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from evaluation_forms")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for evaluation_forms")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q evaluationFormQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no evaluationFormQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from evaluation_forms")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for evaluation_forms")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EvaluationFormSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(evaluationFormBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), evaluationFormPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"evaluation_forms\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, evaluationFormPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from evaluationForm slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for evaluation_forms")
	}

	if len(evaluationFormAfterDeleteHooks) != 0 {
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
func (o *EvaluationForm) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEvaluationForm(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EvaluationFormSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EvaluationFormSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), evaluationFormPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"evaluation_forms\".* FROM \"evaluation_forms\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, evaluationFormPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EvaluationFormSlice")
	}

	*o = slice

	return nil
}

// EvaluationFormExists checks if the EvaluationForm row exists.
func EvaluationFormExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"evaluation_forms\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if evaluation_forms exists")
	}

	return exists, nil
}
