<!--
SPDX-FileCopyrightText: 2021 Kaelan Thijs Fouwels <kaelan.thijs@fouwels.com>

SPDX-License-Identifier: MIT
-->

# OPC/UA

Go defined tag configurable OPC/UA server, built on [Open62541](https://open62541.org) 

See `config/taglist.yml` for an example configuration input, see `compose.yml` for containerised use.

Tags are configured declaratively, and provided as the path defined by the `CONFIG_TAGLIST` environment variable. 

This file is defined as per `config/schema.go`, see `config/taglist.yml` for an example declaration.

Supports the full range of OPC/UA types, as defined in `open62541/types.go`

See `Dockerfile` for `./configured` Open62541 build flags.

## Security

OPC/UA certificate implementations within Open62541 have not been tested, and no credit is assumed. The connection should be upgraded to (m)TLS externally in the deployment (nginx) with effective crypto, if this is required.

## License
MIT and/or MIT compatible

Licensing tracked via SPDX, see file level tags for specific attribution