// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testOrthopedicSurgeriesActivityReviewSurgeries(t *testing.T) {
	t.Parallel()

	query := OrthopedicSurgeriesActivityReviewSurgeries()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
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

	count, err := OrthopedicSurgeriesActivityReviewSurgeries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OrthopedicSurgeriesActivityReviewSurgeries().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrthopedicSurgeriesActivityReviewSurgeries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrthopedicSurgeriesActivityReviewSurgerySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrthopedicSurgeriesActivityReviewSurgeries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OrthopedicSurgeriesActivityReviewSurgeryExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if OrthopedicSurgeriesActivityReviewSurgery exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrthopedicSurgeriesActivityReviewSurgeryExists to return true, but got false.")
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	orthopedicSurgeriesActivityReviewSurgeryFound, err := FindOrthopedicSurgeriesActivityReviewSurgery(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if orthopedicSurgeriesActivityReviewSurgeryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OrthopedicSurgeriesActivityReviewSurgeries().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OrthopedicSurgeriesActivityReviewSurgeries().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	orthopedicSurgeriesActivityReviewSurgeryOne := &OrthopedicSurgeriesActivityReviewSurgery{}
	orthopedicSurgeriesActivityReviewSurgeryTwo := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, orthopedicSurgeriesActivityReviewSurgeryOne, orthopedicSurgeriesActivityReviewSurgeryDBTypes, false, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}
	if err = randomize.Struct(seed, orthopedicSurgeriesActivityReviewSurgeryTwo, orthopedicSurgeriesActivityReviewSurgeryDBTypes, false, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = orthopedicSurgeriesActivityReviewSurgeryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = orthopedicSurgeriesActivityReviewSurgeryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrthopedicSurgeriesActivityReviewSurgeries().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	orthopedicSurgeriesActivityReviewSurgeryOne := &OrthopedicSurgeriesActivityReviewSurgery{}
	orthopedicSurgeriesActivityReviewSurgeryTwo := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, orthopedicSurgeriesActivityReviewSurgeryOne, orthopedicSurgeriesActivityReviewSurgeryDBTypes, false, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}
	if err = randomize.Struct(seed, orthopedicSurgeriesActivityReviewSurgeryTwo, orthopedicSurgeriesActivityReviewSurgeryDBTypes, false, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = orthopedicSurgeriesActivityReviewSurgeryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = orthopedicSurgeriesActivityReviewSurgeryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrthopedicSurgeriesActivityReviewSurgeries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func orthopedicSurgeriesActivityReviewSurgeryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrthopedicSurgeriesActivityReviewSurgery) error {
	*o = OrthopedicSurgeriesActivityReviewSurgery{}
	return nil
}

func orthopedicSurgeriesActivityReviewSurgeryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrthopedicSurgeriesActivityReviewSurgery) error {
	*o = OrthopedicSurgeriesActivityReviewSurgery{}
	return nil
}

func orthopedicSurgeriesActivityReviewSurgeryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OrthopedicSurgeriesActivityReviewSurgery) error {
	*o = OrthopedicSurgeriesActivityReviewSurgery{}
	return nil
}

func orthopedicSurgeriesActivityReviewSurgeryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrthopedicSurgeriesActivityReviewSurgery) error {
	*o = OrthopedicSurgeriesActivityReviewSurgery{}
	return nil
}

func orthopedicSurgeriesActivityReviewSurgeryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrthopedicSurgeriesActivityReviewSurgery) error {
	*o = OrthopedicSurgeriesActivityReviewSurgery{}
	return nil
}

func orthopedicSurgeriesActivityReviewSurgeryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrthopedicSurgeriesActivityReviewSurgery) error {
	*o = OrthopedicSurgeriesActivityReviewSurgery{}
	return nil
}

func orthopedicSurgeriesActivityReviewSurgeryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrthopedicSurgeriesActivityReviewSurgery) error {
	*o = OrthopedicSurgeriesActivityReviewSurgery{}
	return nil
}

func orthopedicSurgeriesActivityReviewSurgeryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrthopedicSurgeriesActivityReviewSurgery) error {
	*o = OrthopedicSurgeriesActivityReviewSurgery{}
	return nil
}

func orthopedicSurgeriesActivityReviewSurgeryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrthopedicSurgeriesActivityReviewSurgery) error {
	*o = OrthopedicSurgeriesActivityReviewSurgery{}
	return nil
}

func testOrthopedicSurgeriesActivityReviewSurgeriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OrthopedicSurgeriesActivityReviewSurgery{}
	o := &OrthopedicSurgeriesActivityReviewSurgery{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery object: %s", err)
	}

	AddOrthopedicSurgeriesActivityReviewSurgeryHook(boil.BeforeInsertHook, orthopedicSurgeriesActivityReviewSurgeryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	orthopedicSurgeriesActivityReviewSurgeryBeforeInsertHooks = []OrthopedicSurgeriesActivityReviewSurgeryHook{}

	AddOrthopedicSurgeriesActivityReviewSurgeryHook(boil.AfterInsertHook, orthopedicSurgeriesActivityReviewSurgeryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	orthopedicSurgeriesActivityReviewSurgeryAfterInsertHooks = []OrthopedicSurgeriesActivityReviewSurgeryHook{}

	AddOrthopedicSurgeriesActivityReviewSurgeryHook(boil.AfterSelectHook, orthopedicSurgeriesActivityReviewSurgeryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	orthopedicSurgeriesActivityReviewSurgeryAfterSelectHooks = []OrthopedicSurgeriesActivityReviewSurgeryHook{}

	AddOrthopedicSurgeriesActivityReviewSurgeryHook(boil.BeforeUpdateHook, orthopedicSurgeriesActivityReviewSurgeryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	orthopedicSurgeriesActivityReviewSurgeryBeforeUpdateHooks = []OrthopedicSurgeriesActivityReviewSurgeryHook{}

	AddOrthopedicSurgeriesActivityReviewSurgeryHook(boil.AfterUpdateHook, orthopedicSurgeriesActivityReviewSurgeryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	orthopedicSurgeriesActivityReviewSurgeryAfterUpdateHooks = []OrthopedicSurgeriesActivityReviewSurgeryHook{}

	AddOrthopedicSurgeriesActivityReviewSurgeryHook(boil.BeforeDeleteHook, orthopedicSurgeriesActivityReviewSurgeryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	orthopedicSurgeriesActivityReviewSurgeryBeforeDeleteHooks = []OrthopedicSurgeriesActivityReviewSurgeryHook{}

	AddOrthopedicSurgeriesActivityReviewSurgeryHook(boil.AfterDeleteHook, orthopedicSurgeriesActivityReviewSurgeryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	orthopedicSurgeriesActivityReviewSurgeryAfterDeleteHooks = []OrthopedicSurgeriesActivityReviewSurgeryHook{}

	AddOrthopedicSurgeriesActivityReviewSurgeryHook(boil.BeforeUpsertHook, orthopedicSurgeriesActivityReviewSurgeryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	orthopedicSurgeriesActivityReviewSurgeryBeforeUpsertHooks = []OrthopedicSurgeriesActivityReviewSurgeryHook{}

	AddOrthopedicSurgeriesActivityReviewSurgeryHook(boil.AfterUpsertHook, orthopedicSurgeriesActivityReviewSurgeryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	orthopedicSurgeriesActivityReviewSurgeryAfterUpsertHooks = []OrthopedicSurgeriesActivityReviewSurgeryHook{}
}

func testOrthopedicSurgeriesActivityReviewSurgeriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrthopedicSurgeriesActivityReviewSurgeries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(orthopedicSurgeriesActivityReviewSurgeryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OrthopedicSurgeriesActivityReviewSurgeries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeryToOneOrthopedicSurgeriesActivityReviewUsingOrthopedicSurgeriesActivityReview(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OrthopedicSurgeriesActivityReviewSurgery
	var foreign OrthopedicSurgeriesActivityReview

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, orthopedicSurgeriesActivityReviewSurgeryDBTypes, false, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, orthopedicSurgeriesActivityReviewDBTypes, false, orthopedicSurgeriesActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReview struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.OrthopedicSurgeriesActivityReviewID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.OrthopedicSurgeriesActivityReview().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := OrthopedicSurgeriesActivityReviewSurgerySlice{&local}
	if err = local.L.LoadOrthopedicSurgeriesActivityReview(ctx, tx, false, (*[]*OrthopedicSurgeriesActivityReviewSurgery)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.OrthopedicSurgeriesActivityReview == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.OrthopedicSurgeriesActivityReview = nil
	if err = local.L.LoadOrthopedicSurgeriesActivityReview(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.OrthopedicSurgeriesActivityReview == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeryToOneSurgeryUsingSurgery(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OrthopedicSurgeriesActivityReviewSurgery
	var foreign Surgery

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, orthopedicSurgeriesActivityReviewSurgeryDBTypes, false, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, surgeryDBTypes, false, surgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Surgery struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.SurgeryID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Surgery().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := OrthopedicSurgeriesActivityReviewSurgerySlice{&local}
	if err = local.L.LoadSurgery(ctx, tx, false, (*[]*OrthopedicSurgeriesActivityReviewSurgery)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Surgery == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Surgery = nil
	if err = local.L.LoadSurgery(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Surgery == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeryToOneSetOpOrthopedicSurgeriesActivityReviewUsingOrthopedicSurgeriesActivityReview(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OrthopedicSurgeriesActivityReviewSurgery
	var b, c OrthopedicSurgeriesActivityReview

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, orthopedicSurgeriesActivityReviewSurgeryDBTypes, false, strmangle.SetComplement(orthopedicSurgeriesActivityReviewSurgeryPrimaryKeyColumns, orthopedicSurgeriesActivityReviewSurgeryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, orthopedicSurgeriesActivityReviewDBTypes, false, strmangle.SetComplement(orthopedicSurgeriesActivityReviewPrimaryKeyColumns, orthopedicSurgeriesActivityReviewColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, orthopedicSurgeriesActivityReviewDBTypes, false, strmangle.SetComplement(orthopedicSurgeriesActivityReviewPrimaryKeyColumns, orthopedicSurgeriesActivityReviewColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*OrthopedicSurgeriesActivityReview{&b, &c} {
		err = a.SetOrthopedicSurgeriesActivityReview(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.OrthopedicSurgeriesActivityReview != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OrthopedicSurgeriesActivityReviewSurgeries[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OrthopedicSurgeriesActivityReviewID != x.ID {
			t.Error("foreign key was wrong value", a.OrthopedicSurgeriesActivityReviewID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OrthopedicSurgeriesActivityReviewID))
		reflect.Indirect(reflect.ValueOf(&a.OrthopedicSurgeriesActivityReviewID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrthopedicSurgeriesActivityReviewID != x.ID {
			t.Error("foreign key was wrong value", a.OrthopedicSurgeriesActivityReviewID, x.ID)
		}
	}
}
func testOrthopedicSurgeriesActivityReviewSurgeryToOneSetOpSurgeryUsingSurgery(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OrthopedicSurgeriesActivityReviewSurgery
	var b, c Surgery

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, orthopedicSurgeriesActivityReviewSurgeryDBTypes, false, strmangle.SetComplement(orthopedicSurgeriesActivityReviewSurgeryPrimaryKeyColumns, orthopedicSurgeriesActivityReviewSurgeryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, surgeryDBTypes, false, strmangle.SetComplement(surgeryPrimaryKeyColumns, surgeryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, surgeryDBTypes, false, strmangle.SetComplement(surgeryPrimaryKeyColumns, surgeryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Surgery{&b, &c} {
		err = a.SetSurgery(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Surgery != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OrthopedicSurgeriesActivityReviewSurgeries[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SurgeryID != x.ID {
			t.Error("foreign key was wrong value", a.SurgeryID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SurgeryID))
		reflect.Indirect(reflect.ValueOf(&a.SurgeryID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SurgeryID != x.ID {
			t.Error("foreign key was wrong value", a.SurgeryID, x.ID)
		}
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
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

func testOrthopedicSurgeriesActivityReviewSurgeriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrthopedicSurgeriesActivityReviewSurgerySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrthopedicSurgeriesActivityReviewSurgeries().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	orthopedicSurgeriesActivityReviewSurgeryDBTypes = map[string]string{`ID`: `uuid`, `SurgeryID`: `uuid`, `OrthopedicSurgeriesActivityReviewID`: `uuid`}
	_                                               = bytes.MinRead
)

func testOrthopedicSurgeriesActivityReviewSurgeriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(orthopedicSurgeriesActivityReviewSurgeryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(orthopedicSurgeriesActivityReviewSurgeryAllColumns) == len(orthopedicSurgeriesActivityReviewSurgeryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrthopedicSurgeriesActivityReviewSurgeries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true, orthopedicSurgeriesActivityReviewSurgeryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(orthopedicSurgeriesActivityReviewSurgeryAllColumns) == len(orthopedicSurgeriesActivityReviewSurgeryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true, orthopedicSurgeriesActivityReviewSurgeryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrthopedicSurgeriesActivityReviewSurgeries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true, orthopedicSurgeriesActivityReviewSurgeryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(orthopedicSurgeriesActivityReviewSurgeryAllColumns, orthopedicSurgeriesActivityReviewSurgeryPrimaryKeyColumns) {
		fields = orthopedicSurgeriesActivityReviewSurgeryAllColumns
	} else {
		fields = strmangle.SetComplement(
			orthopedicSurgeriesActivityReviewSurgeryAllColumns,
			orthopedicSurgeriesActivityReviewSurgeryPrimaryKeyColumns,
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

	slice := OrthopedicSurgeriesActivityReviewSurgerySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOrthopedicSurgeriesActivityReviewSurgeriesUpsert(t *testing.T) {
	t.Parallel()

	if len(orthopedicSurgeriesActivityReviewSurgeryAllColumns) == len(orthopedicSurgeriesActivityReviewSurgeryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OrthopedicSurgeriesActivityReviewSurgery{}
	if err = randomize.Struct(seed, &o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrthopedicSurgeriesActivityReviewSurgery: %s", err)
	}

	count, err := OrthopedicSurgeriesActivityReviewSurgeries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, orthopedicSurgeriesActivityReviewSurgeryDBTypes, false, orthopedicSurgeriesActivityReviewSurgeryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrthopedicSurgeriesActivityReviewSurgery struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrthopedicSurgeriesActivityReviewSurgery: %s", err)
	}

	count, err = OrthopedicSurgeriesActivityReviewSurgeries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
