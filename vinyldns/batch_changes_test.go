/*
Copyright 2018 Comcast Cable Communications Management, LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vinyldns

import (
	"testing"

	"github.com/gobs/pretty"
)

func TestBatchRecordChanges(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/zones/batchrecordchanges",
			code:     200,
			body:     batchRecordChangesJSON,
		},
	})

	defer server.Close()

	changes, err := client.BatchRecordChanges()
	if err != nil {
		t.Log(pretty.PrettyFormat(changes))
		t.Error(err)
	}

	c := changes[0]
	if c.UserName != "vinyl201" {
		t.Error("Expected BatchRecordChanges[0].UserName to be 'vinyl201'")
	}
	if c.TotalChanges != 5 {
		t.Error("Expected BatchRecordChanges[0].TotalChanges to be '5'")
	}
}

func TestBatchRecordChange(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/zones/batchrecordchanges/123",
			code:     200,
			body:     batchRecordChangeJSON,
		},
	})

	defer server.Close()

	change, err := client.BatchRecordChange("123")
	if err != nil {
		t.Log(pretty.PrettyFormat(change))
		t.Error(err)
	}

	c := change.Changes[0]
	if c.RecordName != "parent.com." {
		t.Error("Expected BatchRecordChange.Changes[0].RecordName to be 'parent.com.'")
	}
	if c.ZoneName != "parent.com." {
		t.Error("Expected BatchRecordChange.Changes[0].ZoneName to be 'parent.com.'")
	}
}

func TestBatchRecordChangeCreate(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/zones/batchrecordchanges",
			code:     200,
			body:     batchRecordChangeCreateJSON,
		},
	})

	defer server.Close()

	change := &BatchRecordChange{}
	changeResult, err := client.BatchRecordChangeCreate(change)
	if err != nil {
		t.Log(pretty.PrettyFormat(changeResult))
		t.Error(err)
	}

	c := changeResult.Changes[0]
	if c.ChangeType != "Add" {
		t.Error("Expected BatchRecordChangeCreate.Changes[0].ChangeType to be 'Add'")
	}
	if changeResult.Comments != "this is optional" {
		t.Error("Expected BatchRecordChangeCreate.Comments to be 'this is optional'")
	}
}
