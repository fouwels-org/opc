// SPDX-FileCopyrightText: 2021 Kaelan Thijs Fouwels <kaelan.thijs@fouwels.com>
//
// SPDX-License-Identifier: MIT

package open62541

import (
	"testing"
)

func TestCreateRunServer(t *testing.T) {
	err := Initialise(2000, 0)
	if err != nil {
		t.Fatalf("failed to create server: %v", err)
	}

	err = CreateObjectNode("TAGS", "Tag List")
	if err != nil {
		t.Fatalf("failed to create node: %v", err)
	}

	err = CreateBool("BOOLA", "Boolean A", "TAGS", true)
	if err != nil {
		t.Fatalf("failed to create bool: %v", err)
	}

	err = CreateInt32("INT32A", "Int32 A", "TAGS", -100)
	if err != nil {
		t.Fatalf("failed to create bool: %v", err)
	}

	err = CreateUInt32("UINT32A", "UInt32 A", "TAGS", 100)
	if err != nil {
		t.Fatalf("failed to create uint32: %v", err)
	}

	err = CreateInt16("INT16A", "Int16 A", "TAGS", 50)
	if err != nil {
		t.Fatalf("failed to create int16: %v", err)
	}

	err = CreateUInt16("UINT16A", "UInt16 A", "TAGS", 50)
	if err != nil {
		t.Fatalf("failed to create uint16: %v", err)
	}

	err = CreateFloat32("FLOAT32A", "Float32 A", "TAGS", 12.5)
	if err != nil {
		t.Fatalf("failed to create float32: %v", err)
	}

	err = CreateFloat64("FLOAT64A", "Float64 A", "TAGS", 25.7)
	if err != nil {
		t.Fatalf("failed to create uint16: %v", err)
	}

	err = ListenAndServe()
	if err != nil {
		t.Fatalf("failed to start server: %v", err)
	}
}
