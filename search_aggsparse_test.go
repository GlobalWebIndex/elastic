package elastic

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	f, err := os.Open("/Users/lulu/Documents/gwi/src/github.com/GlobalWebIndex/core-next/jsontest/x03")
	if err != nil {
		t.Error(err)
		return
	}
	var buf = new(bytes.Buffer)
	_, err = buf.ReadFrom(f)
	if err != nil {
		t.Error(err)
		return
	}

	var out SearchResult
	err = jsoniter.Unmarshal(buf.Bytes(), &out)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Aggregation metric-: %+v \n", out.Aggregations["metric-"])
	fmt.Printf("Aggregation metric-/: %+v \n", out.Aggregations["metric-"].Aggregations["by_segment"])
	fmt.Printf("TEST done \n")
}
