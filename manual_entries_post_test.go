package bexio_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestManualEntriesPost(t *testing.T) {
	req := client.NewManualEntriesPostRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
