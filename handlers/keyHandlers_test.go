package handlers

import (
	"bytes"
	"goYS/model"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestBodyToKey(t *testing.T) {

	ts := []struct {
		txt string
		r   *http.Request
		k   *model.Key
		err bool
		exp *model.Key
	}{
		{
			txt: "nil request",
			err: true,
		},
		{
			txt: "empty request body",
			r:   &http.Request{},
			err: true,
		},
		{
			txt: "empty key",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString("{}")),
			},
			err: true,
		},
		{
			txt: "malformed data",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString(`{"id":12}`)),
			},
			k:   &model.Key{},
			err: true,
		},
	}

	for _, tc := range ts {
		t.Log(tc.txt)
		err := bodyToKey(tc.r, tc.k)
		if tc.err {
			if err == nil {
				t.Error("Expected error")
			}
			continue
		}
		if err != nil {
			t.Error("Unexpected error: ", err)
			continue
		}
		if !reflect.DeepEqual(tc.k, tc.exp) {
			t.Error("Unmarshalled data is diffrent")
			t.Error(tc.k)
			t.Error(tc.exp)
		}
	}
}
