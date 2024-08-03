package handler

import (
	"context"
	"net/http/httptest"
	"recipevault/api/util"
	"testing"
)

func TestGetCtxUserID_Valid(t *testing.T) {
	r := httptest.NewRequest("GET", "/api", nil)
	ctx := context.WithValue(r.Context(), util.CTX_USER_ID, 1)
	r = r.WithContext(ctx)

	id := getCtxUserID(r)

	if id != 1 {
		t.Error("Expected 1 got", id)
	}
}

func TestGetCtxUserID_Fail(t *testing.T) {
	r := httptest.NewRequest("GET", "/api", nil)
	var fail util.Key = "FAIL"
	ctx := context.WithValue(r.Context(), fail, 1)
	r = r.WithContext(ctx)

	id := getCtxUserID(r)

	if id != -1 {
		t.Error("Expected -1 got", id)
	}
}
