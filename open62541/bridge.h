// SPDX-FileCopyrightText: 2021 Kaelan Thijs Fouwels <kaelan.thijs@fouwels.com>
//
// SPDX-License-Identifier: MIT
#include <stdint.h>
#include <stdbool.h>

int Initialise(uint32_t port, uint32_t parentNodes_length);
int ListenAndServe();

int CreateObjectNode(char *nodeID, char *nodeName);
int CreateBool(char *nodeID, char *nodeName, char *parentNodeID, bool defaultValue);
int CreateInt32(char *nodeID, char *nodeName, char *parentNodeID, int32_t defaultValue);
int CreateUInt32(char *nodeID, char *nodeName, char *parentNodeID, uint32_t defaultValue);
int CreateInt16(char *nodeID, char *nodeName, char *parentNodeID, int16_t defaultValue);
int CreateUInt16(char *nodeID, char *nodeName, char *parentNodeID, uint16_t defaultValue);
int CreateFloat32(char *nodeID, char *nodeName, char *parentNodeID, float defaultValue);
int CreateFloat64(char *nodeID, char *nodeName, char *parentNodeID, double defaultValue);