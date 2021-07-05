// SPDX-FileCopyrightText: 2021 Kaelan Thijs Fouwels <kaelan.thijs@fouwels.com>
//
// SPDX-License-Identifier: MIT

/* Includes */
#include <signal.h>
#include <string.h>
#include <stdio.h>
#include <open62541.h>
#include <stdint.h>

/* Definitions */
#define STRING_BUFFER_LENGTH 120
#define OPCUA_NAMESPACE 1

/* Globals */
UA_Server *server;
UA_ServerConfig *config;
UA_Boolean isServerRunning = true;

int createTag(char *nodeID, char *nodeName, char *parentNodeID, UA_Variant defaultValue, uint32_t ua_type);

int Initialise(uint32_t port, uint32_t parentNodes_length)
{
    server = UA_Server_new();
    UA_ServerConfig_setMinimal(UA_Server_getConfig(server), port, NULL);

    UA_ServerConfig *cfg = UA_Server_getConfig(server);

    /* Disable time stamp verification for Prosys OPC UA */
    cfg->verifyRequestTimestamp = UA_RULEHANDLING_ACCEPT;
    return 0;
}

int ListenAndServe()
{
    UA_Server_run(server, &isServerRunning);
    UA_Server_delete(server);
    return -1;
}


int CreateBool(char *nodeID, char *nodeName, char *parentNodeID, bool defaultValue)
{

    int ua_type = UA_TYPES_BOOLEAN;
    /* Value */
    UA_Variant ua_defaultValue;
    UA_Variant_init(&ua_defaultValue);
    UA_Variant_setScalar(&ua_defaultValue, &defaultValue, &UA_TYPES[ua_type]);

    return createTag(nodeID, nodeName, parentNodeID, ua_defaultValue, ua_type);
}

int CreateInt32(char *nodeID, char *nodeName, char *parentNodeID, int32_t defaultValue)
{

    int ua_type = UA_TYPES_INT32;
    /* Value */
    UA_Variant ua_defaultValue;
    UA_Variant_init(&ua_defaultValue);
    UA_Variant_setScalar(&ua_defaultValue, &defaultValue, &UA_TYPES[ua_type]);

    return createTag(nodeID, nodeName, parentNodeID, ua_defaultValue, ua_type);
}

int CreateUInt32(char *nodeID, char *nodeName, char *parentNodeID, uint32_t defaultValue)
{

    int ua_type = UA_TYPES_UINT32;
    /* Value */
    UA_Variant ua_defaultValue;
    UA_Variant_init(&ua_defaultValue);
    UA_Variant_setScalar(&ua_defaultValue, &defaultValue, &UA_TYPES[ua_type]);

    return createTag(nodeID, nodeName, parentNodeID, ua_defaultValue, ua_type);
}

int CreateInt16(char *nodeID, char *nodeName, char *parentNodeID, int16_t defaultValue)
{

    int ua_type = UA_TYPES_INT16;
    /* Value */
    UA_Variant ua_defaultValue;
    UA_Variant_init(&ua_defaultValue);
    UA_Variant_setScalar(&ua_defaultValue, &defaultValue, &UA_TYPES[ua_type]);

    return createTag(nodeID, nodeName, parentNodeID, ua_defaultValue, ua_type);
}

int CreateUInt16(char *nodeID, char *nodeName, char *parentNodeID, uint16_t defaultValue)
{

    int ua_type = UA_TYPES_UINT16;
    /* Value */
    UA_Variant ua_defaultValue;
    UA_Variant_init(&ua_defaultValue);
    UA_Variant_setScalar(&ua_defaultValue, &defaultValue, &UA_TYPES[ua_type]);

    return createTag(nodeID, nodeName, parentNodeID, ua_defaultValue, ua_type);
}

int CreateFloat32(char *nodeID, char *nodeName, char *parentNodeID, float defaultValue)
{

    int ua_type = UA_TYPES_FLOAT;
    /* Value */
    UA_Variant ua_defaultValue;
    UA_Variant_init(&ua_defaultValue);
    UA_Variant_setScalar(&ua_defaultValue, &defaultValue, &UA_TYPES[ua_type]);

    return createTag(nodeID, nodeName, parentNodeID, ua_defaultValue, ua_type);
}

int CreateFloat64(char *nodeID, char *nodeName, char *parentNodeID, double defaultValue)
{

    int ua_type = UA_TYPES_DOUBLE;
    /* Value */
    UA_Variant ua_defaultValue;
    UA_Variant_init(&ua_defaultValue);
    UA_Variant_setScalar(&ua_defaultValue, &defaultValue, &UA_TYPES[ua_type]);

    return createTag(nodeID, nodeName, parentNodeID, ua_defaultValue, ua_type);
}

int createTag(char *nodeID, char *nodeName, char *parentNodeID, UA_Variant defaultValue, uint32_t ua_type)
{
    /* Parent */
    UA_NodeId ua_parentNodeId = UA_NODEID_STRING(1, parentNodeID);
    UA_NodeId ua_parentReferenceNodeId = UA_NODEID_NUMERIC(0, UA_NS0ID_ORGANIZES);

    /* Node */
    UA_NodeId ua_nodeID = UA_NODEID_STRING(OPCUA_NAMESPACE, nodeID);
    UA_NodeId ua_nodeTypeDefinition = UA_NODEID_NUMERIC(0, UA_NS0ID_BASEDATAVARIABLETYPE);
    UA_QualifiedName ua_nodeName = UA_QUALIFIEDNAME(OPCUA_NAMESPACE, nodeName);

    /* DisplayName */
    UA_VariableAttributes ua_attributes = UA_VariableAttributes_default;
    ua_attributes.displayName = UA_LOCALIZEDTEXT("en-US", nodeName);
    ua_attributes.dataType = UA_TYPES[ua_type].typeId;
    ua_attributes.accessLevel = UA_ACCESSLEVELMASK_READ | UA_ACCESSLEVELMASK_WRITE;
    ua_attributes.value = defaultValue;

    int statuscode = UA_Server_addVariableNode(
        server,
        ua_nodeID,
        ua_parentNodeId,
        ua_parentReferenceNodeId,
        ua_nodeName,
        ua_nodeTypeDefinition,
        ua_attributes,
        NULL,
        NULL);

    return statuscode;
}

int CreateObjectNode(char *nodeID, char *nodeName)
{
    UA_NodeId ua_parentNodeId = UA_NODEID_NUMERIC(0, UA_NS0ID_OBJECTSFOLDER);
    UA_NodeId ua_parentReferenceTypeId = UA_NODEID_NUMERIC(0, UA_NS0ID_ORGANIZES);
    UA_QualifiedName ua_browseName = UA_QUALIFIEDNAME(1, nodeName);
    UA_NodeId ua_nodeTypeDefinition = UA_NODEID_NUMERIC(0, UA_NS0ID_BASEOBJECTTYPE);

    UA_NodeId ua_nodeID = UA_NODEID_STRING(1, nodeID);

    UA_ObjectAttributes attr = UA_ObjectAttributes_default;
    attr.displayName = UA_LOCALIZEDTEXT("en-US", nodeName);

    int statuscode = UA_Server_addObjectNode(
        server,
        ua_nodeID,
        ua_parentNodeId,
        ua_parentReferenceTypeId,
        ua_browseName,
        ua_nodeTypeDefinition,
        attr,
        NULL,
        NULL);

    return statuscode;
}