// SPDX-FileCopyrightText: 2021 Kaelan Thijs Fouwels <kaelan.thijs@fouwels.com>
//
// SPDX-License-Identifier: MIT

package open62541

// #cgo LDFLAGS: -L/usr/local/include -lopen62541
// #include <stdlib.h>
// #include "bridge.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func Initialise(port uint16, parentNodes_length uint32) error {

	i := C.Initialise(C.uint(port), C.uint(parentNodes_length))
	if i != 0 {
		return fmt.Errorf("exit with: %v", i)
	}
	return nil
}
func ListenAndServe() error {

	i := C.ListenAndServe()
	if i != 0 {
		return fmt.Errorf("exit with: %v", i)
	}
	return nil
}

func CreateBool(nodeID string, nodeName string, parentNodeID string, defaultValue bool) error {

	cnodeID := C.CString(nodeID)
	cnodeName := C.CString(nodeName)
	cparentNodeID := C.CString(parentNodeID)

	defer C.free(unsafe.Pointer(cnodeID))
	defer C.free(unsafe.Pointer(cnodeName))
	defer C.free(unsafe.Pointer(cparentNodeID))

	i := C.CreateBool(cnodeID, cnodeName, cparentNodeID, C.bool(defaultValue))
	if i != 0 {
		return fmt.Errorf("exit with: %v", i)
	}
	return nil
}

func CreateInt32(nodeID string, nodeName string, parentNodeID string, defaultValue int32) error {

	cnodeID := C.CString(nodeID)
	cnodeName := C.CString(nodeName)
	cparentNodeID := C.CString(parentNodeID)

	defer C.free(unsafe.Pointer(cnodeID))
	defer C.free(unsafe.Pointer(cnodeName))
	defer C.free(unsafe.Pointer(cparentNodeID))

	i := C.CreateInt32(cnodeID, cnodeName, cparentNodeID, C.int(defaultValue))
	if i != 0 {
		return fmt.Errorf("exit with: %v", i)
	}
	return nil
}

func CreateUInt32(nodeID string, nodeName string, parentNodeID string, defaultValue uint32) error {

	cnodeID := C.CString(nodeID)
	cnodeName := C.CString(nodeName)
	cparentNodeID := C.CString(parentNodeID)

	defer C.free(unsafe.Pointer(cnodeID))
	defer C.free(unsafe.Pointer(cnodeName))
	defer C.free(unsafe.Pointer(cparentNodeID))

	i := C.CreateUInt32(cnodeID, cnodeName, cparentNodeID, C.uint(defaultValue))
	if i != 0 {
		return fmt.Errorf("exit with: %v", i)
	}
	return nil
}

func CreateInt16(nodeID string, nodeName string, parentNodeID string, defaultValue int16) error {

	cnodeID := C.CString(nodeID)
	cnodeName := C.CString(nodeName)
	cparentNodeID := C.CString(parentNodeID)

	defer C.free(unsafe.Pointer(cnodeID))
	defer C.free(unsafe.Pointer(cnodeName))
	defer C.free(unsafe.Pointer(cparentNodeID))

	i := C.CreateInt16(cnodeID, cnodeName, cparentNodeID, C.short(defaultValue))
	if i != 0 {
		return fmt.Errorf("exit with: %v", i)
	}
	return nil
}

func CreateUInt16(nodeID string, nodeName string, parentNodeID string, defaultValue uint16) error {

	cnodeID := C.CString(nodeID)
	cnodeName := C.CString(nodeName)
	cparentNodeID := C.CString(parentNodeID)

	defer C.free(unsafe.Pointer(cnodeID))
	defer C.free(unsafe.Pointer(cnodeName))
	defer C.free(unsafe.Pointer(cparentNodeID))

	i := C.CreateUInt16(cnodeID, cnodeName, cparentNodeID, C.ushort(defaultValue))
	if i != 0 {
		return fmt.Errorf("exit with: %v", i)
	}
	return nil
}

func CreateFloat32(nodeID string, nodeName string, parentNodeID string, defaultValue float32) error {

	cnodeID := C.CString(nodeID)
	cnodeName := C.CString(nodeName)
	cparentNodeID := C.CString(parentNodeID)

	defer C.free(unsafe.Pointer(cnodeID))
	defer C.free(unsafe.Pointer(cnodeName))
	defer C.free(unsafe.Pointer(cparentNodeID))

	i := C.CreateFloat32(cnodeID, cnodeName, cparentNodeID, C.float(defaultValue))
	if i != 0 {
		return fmt.Errorf("exit with: %v", i)
	}
	return nil
}

func CreateFloat64(nodeID string, nodeName string, parentNodeID string, defaultValue float64) error {

	cnodeID := C.CString(nodeID)
	cnodeName := C.CString(nodeName)
	cparentNodeID := C.CString(parentNodeID)

	defer C.free(unsafe.Pointer(cnodeID))
	defer C.free(unsafe.Pointer(cnodeName))
	defer C.free(unsafe.Pointer(cparentNodeID))

	i := C.CreateFloat64(cnodeID, cnodeName, cparentNodeID, C.double(defaultValue))
	if i != 0 {
		return fmt.Errorf("exit with: %v", i)
	}
	return nil
}

func CreateObjectNode(nodeID string, nodeName string) error {

	cnodeID := C.CString(nodeID)
	cnodeName := C.CString(nodeName)

	defer C.free(unsafe.Pointer(cnodeID))
	defer C.free(unsafe.Pointer(cnodeName))

	i := C.CreateObjectNode(cnodeID, cnodeName)
	if i != 0 {
		return fmt.Errorf("exit with: %v", i)
	}
	return nil
}
