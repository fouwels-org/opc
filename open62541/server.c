// SPDX-FileCopyrightText: 2021 Belcan Advanced Solutions
//
// SPDX-License-Identifier: MIT

/* Includes */
#include <signal.h>
#include <string.h>
#include <stdio.h>
#include <open62541.h>
#include "types.h"

/* Definitions */
#define STRING_BUFFER_LENGTH 120
#define OPCUA_NAMESPACE 1

/* Globals */
UA_Server *server;
UA_ServerConfig *config;
UA_Boolean isServerRunning = true;

int CreateServer(uint32_T port, uint32_T parentNodes_length)
{
    server = UA_Server_new();
    UA_ServerConfig_setMinimal(UA_Server_getConfig(server), port, NULL);

    UA_ServerConfig *cfg = UA_Server_getConfig(server);

    /* Disable time stamp verification for Prosys OPC UA */
    cfg->verifyRequestTimestamp = UA_RULEHANDLING_ACCEPT;

    printf("Created server on port %i\n", port);

    return 0;
}

int StartServer()
{
    printf("Starting server");
    UA_Server_run(server, &isServerRunning);

    printf("Server exit");
    UA_Server_delete(server);
    return -1;
}

int CreateTag(char *nodeID, char *nodeName, char *parentNodeID, UA_Variant defaultValue, uint32_T ua_type)
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

    if (statuscode == UA_STATUSCODE_GOOD)
    {
        printf("OK (%i): creating tag id %s on parent %s\n", statuscode, nodeID, parentNodeID);
    }
    else
    {
        printf("FAILED (%i): creating tag id %s on parent %s\n", statuscode, nodeID, parentNodeID);
    }
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

    if (statuscode == UA_STATUSCODE_GOOD)
    {
        printf("OK (%i): creating object node %s\n", statuscode, nodeID);
    }
    else
    {
        printf("FAILED (%i): creating object node %s\n", statuscode, nodeID);
    }

    return statuscode;
}