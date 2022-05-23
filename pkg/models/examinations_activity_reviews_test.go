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

func testExaminationsActivityReviews(t *testing.T) {
	t.Parallel()

	query := ExaminationsActivityReviews()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testExaminationsActivityReviewsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
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

	count, err := ExaminationsActivityReviews().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testExaminationsActivityReviewsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ExaminationsActivityReviews().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ExaminationsActivityReviews().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testExaminationsActivityReviewsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ExaminationsActivityReviewSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ExaminationsActivityReviews().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testExaminationsActivityReviewsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ExaminationsActivityReviewExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ExaminationsActivityReview exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ExaminationsActivityReviewExists to return true, but got false.")
	}
}

func testExaminationsActivityReviewsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	examinationsActivityReviewFound, err := FindExaminationsActivityReview(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if examinationsActivityReviewFound == nil {
		t.Error("want a record, got nil")
	}
}

func testExaminationsActivityReviewsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ExaminationsActivityReviews().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testExaminationsActivityReviewsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ExaminationsActivityReviews().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testExaminationsActivityReviewsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	examinationsActivityReviewOne := &ExaminationsActivityReview{}
	examinationsActivityReviewTwo := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, examinationsActivityReviewOne, examinationsActivityReviewDBTypes, false, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}
	if err = randomize.Struct(seed, examinationsActivityReviewTwo, examinationsActivityReviewDBTypes, false, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = examinationsActivityReviewOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = examinationsActivityReviewTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ExaminationsActivityReviews().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testExaminationsActivityReviewsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	examinationsActivityReviewOne := &ExaminationsActivityReview{}
	examinationsActivityReviewTwo := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, examinationsActivityReviewOne, examinationsActivityReviewDBTypes, false, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}
	if err = randomize.Struct(seed, examinationsActivityReviewTwo, examinationsActivityReviewDBTypes, false, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = examinationsActivityReviewOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = examinationsActivityReviewTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ExaminationsActivityReviews().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func examinationsActivityReviewBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ExaminationsActivityReview) error {
	*o = ExaminationsActivityReview{}
	return nil
}

func examinationsActivityReviewAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ExaminationsActivityReview) error {
	*o = ExaminationsActivityReview{}
	return nil
}

func examinationsActivityReviewAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ExaminationsActivityReview) error {
	*o = ExaminationsActivityReview{}
	return nil
}

func examinationsActivityReviewBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ExaminationsActivityReview) error {
	*o = ExaminationsActivityReview{}
	return nil
}

func examinationsActivityReviewAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ExaminationsActivityReview) error {
	*o = ExaminationsActivityReview{}
	return nil
}

func examinationsActivityReviewBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ExaminationsActivityReview) error {
	*o = ExaminationsActivityReview{}
	return nil
}

func examinationsActivityReviewAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ExaminationsActivityReview) error {
	*o = ExaminationsActivityReview{}
	return nil
}

func examinationsActivityReviewBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ExaminationsActivityReview) error {
	*o = ExaminationsActivityReview{}
	return nil
}

func examinationsActivityReviewAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ExaminationsActivityReview) error {
	*o = ExaminationsActivityReview{}
	return nil
}

func testExaminationsActivityReviewsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ExaminationsActivityReview{}
	o := &ExaminationsActivityReview{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview object: %s", err)
	}

	AddExaminationsActivityReviewHook(boil.BeforeInsertHook, examinationsActivityReviewBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	examinationsActivityReviewBeforeInsertHooks = []ExaminationsActivityReviewHook{}

	AddExaminationsActivityReviewHook(boil.AfterInsertHook, examinationsActivityReviewAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	examinationsActivityReviewAfterInsertHooks = []ExaminationsActivityReviewHook{}

	AddExaminationsActivityReviewHook(boil.AfterSelectHook, examinationsActivityReviewAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	examinationsActivityReviewAfterSelectHooks = []ExaminationsActivityReviewHook{}

	AddExaminationsActivityReviewHook(boil.BeforeUpdateHook, examinationsActivityReviewBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	examinationsActivityReviewBeforeUpdateHooks = []ExaminationsActivityReviewHook{}

	AddExaminationsActivityReviewHook(boil.AfterUpdateHook, examinationsActivityReviewAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	examinationsActivityReviewAfterUpdateHooks = []ExaminationsActivityReviewHook{}

	AddExaminationsActivityReviewHook(boil.BeforeDeleteHook, examinationsActivityReviewBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	examinationsActivityReviewBeforeDeleteHooks = []ExaminationsActivityReviewHook{}

	AddExaminationsActivityReviewHook(boil.AfterDeleteHook, examinationsActivityReviewAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	examinationsActivityReviewAfterDeleteHooks = []ExaminationsActivityReviewHook{}

	AddExaminationsActivityReviewHook(boil.BeforeUpsertHook, examinationsActivityReviewBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	examinationsActivityReviewBeforeUpsertHooks = []ExaminationsActivityReviewHook{}

	AddExaminationsActivityReviewHook(boil.AfterUpsertHook, examinationsActivityReviewAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	examinationsActivityReviewAfterUpsertHooks = []ExaminationsActivityReviewHook{}
}

func testExaminationsActivityReviewsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ExaminationsActivityReviews().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testExaminationsActivityReviewsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(examinationsActivityReviewColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ExaminationsActivityReviews().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testExaminationsActivityReviewToManyExaminationsActivitiesReviews(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ExaminationsActivityReview
	var b, c ExaminationsActivitiesReview

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, examinationsActivityReviewDBTypes, true, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, examinationsActivitiesReviewDBTypes, false, examinationsActivitiesReviewColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, examinationsActivitiesReviewDBTypes, false, examinationsActivitiesReviewColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.ExaminationsActivityReviewsID, a.ID)
	queries.Assign(&c.ExaminationsActivityReviewsID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ExaminationsActivitiesReviews().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.ExaminationsActivityReviewsID, b.ExaminationsActivityReviewsID) {
			bFound = true
		}
		if queries.Equal(v.ExaminationsActivityReviewsID, c.ExaminationsActivityReviewsID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ExaminationsActivityReviewSlice{&a}
	if err = a.L.LoadExaminationsActivitiesReviews(ctx, tx, false, (*[]*ExaminationsActivityReview)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ExaminationsActivitiesReviews); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ExaminationsActivitiesReviews = nil
	if err = a.L.LoadExaminationsActivitiesReviews(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ExaminationsActivitiesReviews); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testExaminationsActivityReviewToManyAddOpExaminationsActivitiesReviews(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ExaminationsActivityReview
	var b, c, d, e ExaminationsActivitiesReview

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, examinationsActivityReviewDBTypes, false, strmangle.SetComplement(examinationsActivityReviewPrimaryKeyColumns, examinationsActivityReviewColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ExaminationsActivitiesReview{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, examinationsActivitiesReviewDBTypes, false, strmangle.SetComplement(examinationsActivitiesReviewPrimaryKeyColumns, examinationsActivitiesReviewColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*ExaminationsActivitiesReview{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddExaminationsActivitiesReviews(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.ExaminationsActivityReviewsID) {
			t.Error("foreign key was wrong value", a.ID, first.ExaminationsActivityReviewsID)
		}
		if !queries.Equal(a.ID, second.ExaminationsActivityReviewsID) {
			t.Error("foreign key was wrong value", a.ID, second.ExaminationsActivityReviewsID)
		}

		if first.R.ExaminationsActivityReview != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ExaminationsActivityReview != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ExaminationsActivitiesReviews[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ExaminationsActivitiesReviews[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ExaminationsActivitiesReviews().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testExaminationsActivityReviewToManySetOpExaminationsActivitiesReviews(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ExaminationsActivityReview
	var b, c, d, e ExaminationsActivitiesReview

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, examinationsActivityReviewDBTypes, false, strmangle.SetComplement(examinationsActivityReviewPrimaryKeyColumns, examinationsActivityReviewColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ExaminationsActivitiesReview{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, examinationsActivitiesReviewDBTypes, false, strmangle.SetComplement(examinationsActivitiesReviewPrimaryKeyColumns, examinationsActivitiesReviewColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetExaminationsActivitiesReviews(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ExaminationsActivitiesReviews().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetExaminationsActivitiesReviews(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ExaminationsActivitiesReviews().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ExaminationsActivityReviewsID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ExaminationsActivityReviewsID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.ExaminationsActivityReviewsID) {
		t.Error("foreign key was wrong value", a.ID, d.ExaminationsActivityReviewsID)
	}
	if !queries.Equal(a.ID, e.ExaminationsActivityReviewsID) {
		t.Error("foreign key was wrong value", a.ID, e.ExaminationsActivityReviewsID)
	}

	if b.R.ExaminationsActivityReview != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.ExaminationsActivityReview != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.ExaminationsActivityReview != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.ExaminationsActivityReview != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.ExaminationsActivitiesReviews[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.ExaminationsActivitiesReviews[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testExaminationsActivityReviewToManyRemoveOpExaminationsActivitiesReviews(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ExaminationsActivityReview
	var b, c, d, e ExaminationsActivitiesReview

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, examinationsActivityReviewDBTypes, false, strmangle.SetComplement(examinationsActivityReviewPrimaryKeyColumns, examinationsActivityReviewColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ExaminationsActivitiesReview{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, examinationsActivitiesReviewDBTypes, false, strmangle.SetComplement(examinationsActivitiesReviewPrimaryKeyColumns, examinationsActivitiesReviewColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddExaminationsActivitiesReviews(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ExaminationsActivitiesReviews().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveExaminationsActivitiesReviews(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ExaminationsActivitiesReviews().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ExaminationsActivityReviewsID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ExaminationsActivityReviewsID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.ExaminationsActivityReview != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.ExaminationsActivityReview != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.ExaminationsActivityReview != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.ExaminationsActivityReview != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.ExaminationsActivitiesReviews) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.ExaminationsActivitiesReviews[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.ExaminationsActivitiesReviews[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testExaminationsActivityReviewToOneUserUsingSupervisorUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ExaminationsActivityReview
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, examinationsActivityReviewDBTypes, false, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.SupervisorUserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.SupervisorUser().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ExaminationsActivityReviewSlice{&local}
	if err = local.L.LoadSupervisorUser(ctx, tx, false, (*[]*ExaminationsActivityReview)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.SupervisorUser == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SupervisorUser = nil
	if err = local.L.LoadSupervisorUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.SupervisorUser == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testExaminationsActivityReviewToOneSetOpUserUsingSupervisorUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ExaminationsActivityReview
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, examinationsActivityReviewDBTypes, false, strmangle.SetComplement(examinationsActivityReviewPrimaryKeyColumns, examinationsActivityReviewColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetSupervisorUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SupervisorUser != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SupervisorUserExaminationsActivityReviews[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SupervisorUserID != x.ID {
			t.Error("foreign key was wrong value", a.SupervisorUserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SupervisorUserID))
		reflect.Indirect(reflect.ValueOf(&a.SupervisorUserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SupervisorUserID != x.ID {
			t.Error("foreign key was wrong value", a.SupervisorUserID, x.ID)
		}
	}
}

func testExaminationsActivityReviewsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
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

func testExaminationsActivityReviewsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ExaminationsActivityReviewSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testExaminationsActivityReviewsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ExaminationsActivityReviews().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	examinationsActivityReviewDBTypes = map[string]string{`ID`: `uuid`, `SupervisorUserID`: `uuid`, `DisplayName`: `character varying`, `CreatedAt`: `timestamp with time zone`, `Annotations`: `jsonb`, `Comment`: `character varying`}
	_                                 = bytes.MinRead
)

func testExaminationsActivityReviewsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(examinationsActivityReviewPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(examinationsActivityReviewAllColumns) == len(examinationsActivityReviewPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ExaminationsActivityReviews().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true, examinationsActivityReviewPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testExaminationsActivityReviewsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(examinationsActivityReviewAllColumns) == len(examinationsActivityReviewPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ExaminationsActivityReview{}
	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true, examinationsActivityReviewColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ExaminationsActivityReviews().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, examinationsActivityReviewDBTypes, true, examinationsActivityReviewPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(examinationsActivityReviewAllColumns, examinationsActivityReviewPrimaryKeyColumns) {
		fields = examinationsActivityReviewAllColumns
	} else {
		fields = strmangle.SetComplement(
			examinationsActivityReviewAllColumns,
			examinationsActivityReviewPrimaryKeyColumns,
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

	slice := ExaminationsActivityReviewSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testExaminationsActivityReviewsUpsert(t *testing.T) {
	t.Parallel()

	if len(examinationsActivityReviewAllColumns) == len(examinationsActivityReviewPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ExaminationsActivityReview{}
	if err = randomize.Struct(seed, &o, examinationsActivityReviewDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ExaminationsActivityReview: %s", err)
	}

	count, err := ExaminationsActivityReviews().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, examinationsActivityReviewDBTypes, false, examinationsActivityReviewPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ExaminationsActivityReview struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ExaminationsActivityReview: %s", err)
	}

	count, err = ExaminationsActivityReviews().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}