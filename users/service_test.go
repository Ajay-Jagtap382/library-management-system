package users

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Ajay-Jagtap382/library-management-system/app"
	"github.com/Ajay-Jagtap382/library-management-system/db/mocks"
)

func TestSuccessfullCreateService(t *testing.T) {
	app.InitLogger()
	l := app.GetLogger()

	sm := mocks.NewStorer(t)
	cs := NewService(sm, l)

	// ctx := context.Background()
	// assert := assert.New(t)

	var tests = []struct {
		ctx     context.Context
		req     CreateRequest
		err     error
		exp_err error
	}{
		{
			ctx:     context.Background(),
			req:     CreateRequest{},
			err:     nil,
			exp_err: nil,
		},
		// {
		// 	ctx:     context.Background(),
		// 	req:     CreateRequest{},
		// 	err:     errors.New("my_err"),
		// 	exp_err: errors.New("Error creating user"),
		// },
	}
	for _, test := range tests {
		sm.On("CreateUser", test.ctx, mock.Anything).Return(test.err)
		err := cs.Create(test.ctx, test.req)

		assert.ErrorIs(t, err, test.exp_err)
		// sm.AssertExpectations(t)
	}
}

// func TestCreateServiceWhenEmptyName(t *testing.T) {
// 	app.InitLogger()
// 	sm := &db.StorerMock{}
// 	l := app.GetLogger()
// 	cs := NewService(sm, l)

// 	ctx := context.Background()
// 	assert := assert.New(t)

// 	var test = struct {
// 		contx    context.Context
// 		req      CreateRequest
// 		expected error
// 	}{
// 		ctx,
// 		CreateRequest{Name: ""},
// 		errEmptyName,
// 	}

// 	assert.Equal(cs.Create(test.contx, test.req), test.expected)
// 	sm.AssertExpectations(t)
// }

// func TestCreateServiceWhenInternalError(t *testing.T) {
// 	app.InitLogger()
// 	sm := &db.StorerMock{}
// 	l := app.GetLogger()
// 	cs := NewService(sm, l)

// 	ctx := context.Background()
// 	assert := assert.New(t)

// 	var test = struct {
// 		contx    context.Context
// 		req      CreateRequest
// 		expected error
// 	}{
// 		ctx,
// 		CreateRequest{Name: "Sports"},
// 		errors.New("Internal Error"),
// 	}

// 	sm.On("CreateCategory", test.contx, mock.Anything).Return(errors.New("Internal Error"))
// 	assert.Equal(cs.Create(test.contx, test.req), test.expected)
// 	sm.AssertExpectations(t)
// }

// func TestSuccessfullListService(t *testing.T) {
// 	app.InitLogger()
// 	sm := &db.StorerMock{}
// 	l := app.GetLogger()
// 	cs := NewService(sm, l)

// 	ctx := context.Background()
// 	assert := assert.New(t)

// 	var test = struct {
// 		contx    context.Context
// 		expected error
// 	}{ctx, nil}
// 	sm.On("ListCategories", test.contx).Return(mock.Anything, nil)
// 	_, err := cs.List(test.contx)
// 	assert.Equal(err, test.expected)
// 	sm.AssertExpectations(t)
// }

// func TestListServiceWhenCategoriesNotExists(t *testing.T) {
// 	app.InitLogger()
// 	sm := &db.StorerMock{}
// 	l := app.GetLogger()
// 	cs := NewService(sm, l)

// 	ctx := context.Background()
// 	assert := assert.New(t)

// 	var test = struct {
// 		contx    context.Context
// 		expected error
// 	}{ctx, errNoCategories}
// 	sm.On("ListCategories", test.contx).Return(mock.Anything, db.ErrCategoryNotExist)
// 	_, err := cs.List(test.contx)
// 	assert.Equal(err, test.expected)
// 	sm.AssertExpectations(t)
// }

// func TestListServiceWhenInternalError(t *testing.T) {
// 	app.InitLogger()
// 	sm := &db.StorerMock{}
// 	l := app.GetLogger()
// 	cs := NewService(sm, l)

// 	ctx := context.Background()
// 	assert := assert.New(t)

// 	var test = struct {
// 		contx    context.Context
// 		expected error
// 	}{ctx, errors.New("Internal Error")}
// 	sm.On("ListCategories", test.contx).Return(mock.Anything, errors.New("Internal Error"))
// 	_, err := cs.List(test.contx)
// 	assert.Equal(err, test.expected)
// 	sm.AssertExpectations(t)
// }

// func TestSuccessfullUpdateService(t *testing.T) {
// 	app.InitLogger()
// 	sm := &db.StorerMock{}
// 	l := app.GetLogger()
// 	cs := NewService(sm, l)

// 	ctx := context.Background()
// 	assert := assert.New(t)

// 	var test = struct {
// 		contx    context.Context
// 		req      UpdateRequest
// 		expected error
// 	}{ctx, UpdateRequest{ID: "1", Name: "Sports"}, nil}
// 	sm.On("UpdateCategory", test.contx, mock.Anything).Return(nil)
// 	assert.Equal(cs.Update(test.contx, test.req), test.expected)
// 	sm.AssertExpectations(t)
// }

// func TestUpdateServiceWhenEmptyID(t *testing.T) {
// 	app.InitLogger()
// 	sm := &db.StorerMock{}
// 	l := app.GetLogger()
// 	cs := NewService(sm, l)

// 	ctx := context.Background()
// 	assert := assert.New(t)

// 	var test = struct {
// 		contx    context.Context
// 		req      UpdateRequest
// 		expected error
// 	}{ctx, UpdateRequest{ID: "", Name: "Sports"}, errEmptyID}
// 	assert.Equal(cs.Update(test.contx, test.req), test.expected)
// 	sm.AssertExpectations(t)
// }

// func TestUpdateServiceWhenEmptyName(t *testing.T) {
// 	app.InitLogger()
// 	sm := &db.StorerMock{}
// 	l := app.GetLogger()
// 	cs := NewService(sm, l)

// 	ctx := context.Background()
// 	assert := assert.New(t)

// 	var test = struct {
// 		contx    context.Context
// 		req      UpdateRequest
// 		expected error
// 	}{ctx, UpdateRequest{ID: "1", Name: ""}, errEmptyName}
// 	assert.Equal(cs.Update(test.contx, test.req), test.expected)
// 	sm.AssertExpectations(t)
// }

// func TestUpdateServiceWhenInternalError(t *testing.T) {
// 	app.InitLogger()
// 	sm := &db.StorerMock{}
// 	l := app.GetLogger()
// 	cs := NewService(sm, l)

// 	ctx := context.Background()
// 	assert := assert.New(t)

// 	var test = struct {
// 		contx    context.Context
// 		req      UpdateRequest
// 		expected error
// 	}{ctx, UpdateRequest{ID: "1", Name: "Sports"}, errors.New("Internal Error")}
// 	sm.On("UpdateCategory", test.contx, mock.Anything).Return(errors.New("Internal Error"))
// 	assert.Equal(cs.Update(test.contx, test.req), test.expected)
// 	sm.AssertExpectations(t)
// }

// func TestSuccessfullFindByIDService(t *testing.T) {
// 	app.InitLogger()
// 	sm := &db.StorerMock{}
// 	l := app.GetLogger()
// 	cs := NewService(sm, l)

// 	ctx := context.Background()
// 	assert := assert.New(t)

// 	var test = struct {
// 		contx    context.Context
// 		id       string
// 		expected error
// 	}{ctx, "1", nil}
// 	sm.On("FindCategoryByID", test.contx, test.id).Return(mock.Anything, nil)
// 	_, err := cs.FindByID(test.contx, test.id)
// 	assert.Equal(err, test.expected)
// 	sm.AssertExpectations(t)
// }

// func TestFindByIDServiceWhenCategoryNotExist(t *testing.T) {
// 	app.InitLogger()
// 	sm := &db.StorerMock{}
// 	l := app.GetLogger()
// 	cs := NewService(sm, l)

// 	ctx := context.Background()
// 	assert := assert.New(t)

// 	var test = struct {
// 		contx    context.Context
// 		id       string
// 		expected error
// 	}{ctx, "1", errNoCategoryId}
// 	sm.On("FindCategoryByID", test.contx, mock.Anything).Return(mock.Anything, db.ErrCategoryNotExist)
// 	_, err := cs.FindByID(test.contx, test.id)
// 	assert.Equal(err, test.expected)
// 	sm.AssertExpectations(t)
// }

// func TestFindByIDServiceWhenInternalError(t *testing.T) {
// 	app.InitLogger()
// 	sm := &db.StorerMock{}
// 	l := app.GetLogger()
// 	cs := NewService(sm, l)

// 	ctx := context.Background()
// 	assert := assert.New(t)

// 	var test = struct {
// 		contx    context.Context
// 		id       string
// 		expected error
// 	}{ctx, "1", errors.New("Internal Error")}
// 	sm.On("FindCategoryByID", test.contx, mock.Anything).Return(mock.Anything, errors.New("Internal Error"))
// 	_, err := cs.FindByID(test.contx, test.id)
// 	assert.Equal(err, test.expected)
// 	sm.AssertExpectations(t)
// }

// func TestSuccessfullDeleteByIDService(t *testing.T) {
// 	app.InitLogger()
// 	sm := &db.StorerMock{}
// 	l := app.GetLogger()
// 	cs := NewService(sm, l)

// 	ctx := context.Background()
// 	assert := assert.New(t)

// 	var test = struct {
// 		contx    context.Context
// 		id       string
// 		expected error
// 	}{ctx, "1", nil}
// 	sm.On("DeleteCategoryByID", test.contx, test.id).Return(nil)
// 	assert.Equal(cs.DeleteByID(test.contx, test.id), test.expected)
// 	sm.AssertExpectations(t)
// }

// func TestDeleteByIDServiceWhenCategoryNotExist(t *testing.T) {
// 	app.InitLogger()
// 	sm := &db.StorerMock{}
// 	l := app.GetLogger()
// 	cs := NewService(sm, l)

// 	ctx := context.Background()
// 	assert := assert.New(t)

// 	var test = struct {
// 		contx    context.Context
// 		id       string
// 		expected error
// 	}{ctx, "1", errNoCategoryId}
// 	sm.On("DeleteCategoryByID", test.contx, test.id).Return(db.ErrCategoryNotExist)
// 	assert.Equal(cs.DeleteByID(test.contx, test.id), test.expected)
// 	sm.AssertExpectations(t)
// }

// func TestDeleteByIDServiceWhenInternalError(t *testing.T) {
// 	app.InitLogger()
// 	sm := &db.StorerMock{}
// 	l := app.GetLogger()
// 	cs := NewService(sm, l)

// 	ctx := context.Background()
// 	assert := assert.New(t)

// 	var test = struct {
// 		contx    context.Context
// 		id       string
// 		expected error
// 	}{ctx, "1", errors.New("Internal Error")}
// 	sm.On("DeleteCategoryByID", test.contx, test.id).Return(errors.New("Internal Error"))
// 	assert.Equal(cs.DeleteByID(test.contx, test.id), test.expected)
// 	sm.AssertExpectations(t)
// }
