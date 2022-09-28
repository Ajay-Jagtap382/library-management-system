package users_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ajay-Jagtap382/library-management-system/users"
	"github.com/Ajay-Jagtap382/library-management-system/users/mocks"
	"github.com/stretchr/testify/mock"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func makeHTTPCall(handler http.HandlerFunc, method, path, body string) (rr *httptest.ResponseRecorder) {
	request := []byte(body)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(request))
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return
}

// Create:
func TestSuccessfullCreate(t *testing.T) {
	cs := &mocks.Service{}
	cs.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("Sucess"))

	rr := makeHTTPCall(users.CreateUser(cs), http.MethodPost, "/users", `{
	"first_name": "abhijeet",
    "email": "abhijeet@gmail.com",
    "last_name": "dhumal",
    "gender": "Male",
    "mobile_num": "9956861238",
    "password": "abhi",
    "role": "admin"}`)

	fmt.Println("Account Success")

	checkResponseCode(t, http.StatusCreated, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenInvalidRequestBody(t *testing.T) {
	cs := &mocks.Service{}

	rr := makeHTTPCall(users.CreateUser(cs), http.MethodPost, "/users", `{
		"first_name": "abhijeet"
		"email": "abhi@gmail.com",
		"last_name": "dhumal",
		"gender": "Male",
		"mobile_num": "9956861238",
		"password": "abhi",
		"role": "admin"}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenEmptyName(t *testing.T) {
	cs := &mocks.Service{}
	// cs.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("User first name must be present"))

	rr := makeHTTPCall(users.CreateUser(cs), http.MethodPost, "/users", `{
		"first_name":"",
		"email": "abhi@gmail.com",
		"last_name":"",
		"gender": "Male",
		"mobile_num": "9956861238",
		"password": "abhi",
		"role": "admin"}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

// List :
func TestSuccessfullList(t *testing.T) {
	var user users.ListResponse
	cs := &mocks.Service{}
	cs.On("List", mock.Anything).Return(user, nil)

	rr := makeHTTPCall(users.GetUser(cs), http.MethodGet, "/users", "")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

// func TestListWhenNoCategories(t *testing.T) {
// 	cs := &CategoryServiceMock{}
// 	cs.On("List", mock.Anything).Return(mock.Anything, errNoCategories)

// 	rr := makeHTTPCall(List(cs), http.MethodGet, "/categories", "")

// 	checkResponseCode(t, http.StatusNotFound, rr.Code)
// 	cs.AssertExpectations(t)
// }

func TestListInternalError(t *testing.T) {
	var user users.ListResponse
	cs := &mocks.Service{}
	cs.On("List", mock.Anything).Return(user, errors.New("Internal Error"))

	rr := makeHTTPCall(users.GetUser(cs), http.MethodGet, "/users", "")

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

// func TestSuccessfullFindByID(t *testing.T) {
// 	cs := &CategoryServiceMock{}
// 	cs.On("FindByID", mock.Anything, mock.Anything).Return(mock.Anything, nil)

// 	rr := makeHTTPCall(FindByID(cs), http.MethodGet, "/categories/1", "")

// 	checkResponseCode(t, http.StatusOK, rr.Code)
// 	cs.AssertExpectations(t)
// }

// func TestFindByIDWhenIDNotExist(t *testing.T) {
// 	cs := &CategoryServiceMock{}
// 	cs.On("FindByID", mock.Anything, mock.Anything).Return(mock.Anything, errNoCategoryId)

// 	rr := makeHTTPCall(FindByID(cs), http.MethodGet, "/categories/1", "")

// 	checkResponseCode(t, http.StatusNotFound, rr.Code)
// 	cs.AssertExpectations(t)
// }

// func TestFindByIDWhenInternalError(t *testing.T) {
// 	cs := &CategoryServiceMock{}
// 	cs.On("FindByID", mock.Anything, mock.Anything).Return(mock.Anything, errors.New("Internal Error"))

// 	rr := makeHTTPCall(FindByID(cs), http.MethodGet, "/categories/1", "")

// 	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
// 	cs.AssertExpectations(t)
// }

// //DeleteByID
// func TestSuccessfullDeleteByID(t *testing.T) {
// 	cs := &CategoryServiceMock{}

// 	cs.On("DeleteByID", mock.Anything, mock.Anything).Return(nil)

// 	rr := makeHTTPCall(DeleteByID(cs), http.MethodDelete, "/categories/1", "")

// 	checkResponseCode(t, http.StatusOK, rr.Code)
// 	cs.AssertExpectations(t)
// }

// func TestDeleteByIDWhenIDNotExist(t *testing.T) {
// 	cs := &CategoryServiceMock{}
// 	cs.On("DeleteByID", mock.Anything, mock.Anything).Return(errNoCategoryId)

// 	rr := makeHTTPCall(DeleteByID(cs), http.MethodDelete, "/categories/1", "")

// 	checkResponseCode(t, http.StatusNotFound, rr.Code)
// 	cs.AssertExpectations(t)
// }

// func TestDeleteByIDWhenInternalError(t *testing.T) {
// 	cs := &CategoryServiceMock{}
// 	cs.On("DeleteByID", mock.Anything, mock.Anything).Return(errors.New("Internal Error"))

// 	rr := makeHTTPCall(DeleteByID(cs), http.MethodDelete, "/categories/1", "")

// 	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
// 	cs.AssertExpectations(t)
// }

// func TestSuccessfullUpdate(t *testing.T) {
// 	cs := &CategoryServiceMock{}
// 	cs.On("Update", mock.Anything, mock.Anything).Return(nil)

// 	rr := makeHTTPCall(Update(cs), http.MethodPut, "/categories", `{"id":"1", "name":"sports"}`)

// 	checkResponseCode(t, http.StatusOK, rr.Code)
// 	cs.AssertExpectations(t)
// }

// func TestUpdateWhenInvalidRequestBody(t *testing.T) {
// 	cs := &CategoryServiceMock{}

// 	rr := makeHTTPCall(Update(cs), http.MethodPut, "/categories", `{"id":"1", "name":"sports",}`)

// 	checkResponseCode(t, http.StatusBadRequest, rr.Code)
// 	cs.AssertExpectations(t)
// }

// func TestUpdateWhenEmptyID(t *testing.T) {
// 	cs := &CategoryServiceMock{}
// 	cs.On("Update", mock.Anything, mock.Anything).Return(errEmptyID)

// 	rr := makeHTTPCall(Update(cs), http.MethodPut, "/categories", `{"name":"Sports"}`)

// 	checkResponseCode(t, http.StatusBadRequest, rr.Code)
// 	cs.AssertExpectations(t)
// }

// func TestUpdateWhenEmptyName(t *testing.T) {
// 	cs := &CategoryServiceMock{}
// 	cs.On("Update", mock.Anything, mock.Anything).Return(errEmptyName)

// 	rr := makeHTTPCall(Update(cs), http.MethodPut, "/categories", `{"id":"1"}`)

// 	checkResponseCode(t, http.StatusBadRequest, rr.Code)
// 	cs.AssertExpectations(t)
// }

// func TestUpdateWhenInternalError(t *testing.T) {
// 	cs := &CategoryServiceMock{}
// 	cs.On("Update", mock.Anything, mock.Anything).Return(errors.New("Internal Error"))

// 	rr := makeHTTPCall(Update(cs), http.MethodPut, "/categories", `{"id":"1"}`)

// 	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
// 	cs.AssertExpectations(t)
// }
