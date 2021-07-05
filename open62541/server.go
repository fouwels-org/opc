// SPDX-FileCopyrightText: 2021 Belcan Advanced Solutions
//
// SPDX-License-Identifier: MIT

package open62541

/*
#cgo LDFLAGS: -L/usr/local/include -lopen62541
#include <stdlib.h>
#include "server.h"
*/
import "C"

func CreateTag(nodeID string, nodeName string, parentNodeID string, defaultValue string, uaType string) error {
	return nil
}
func CreateObjectNode(nodeID string, nodeName string, parentNodeID string, defaultValue string) error {
	return nil
}

func CreateServer(port int, parentNodes_length int) error {
	return nil
}
func StartServer() error {
	return nil
}
