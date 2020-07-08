// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"testing"

	"github.com/json-iterator/go"
)

func TestPrefixQuery(t *testing.T) {
	q := NewPrefixQuery("user", "ki")
	src, err := q.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := jsoniter.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"prefix":{"user":"ki"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestPrefixQueryWithOptions(t *testing.T) {
	q := NewPrefixQuery("user", "ki")
	q = q.QueryName("my_query_name")
	src, err := q.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := jsoniter.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"prefix":{"user":{"_name":"my_query_name","value":"ki"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
