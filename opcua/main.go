// SPDX-FileCopyrightText: 2021 Kaelan Thijs Fouwels <kaelan.thijs@fouwels.com>
//
// SPDX-License-Identifier: MIT

package opcua

import (
	"fmt"
	"opc/config"
)

type OPCUA struct {
	c config.Config
}

func NewOPCUA(cfg config.Config) *OPCUA {
	return &OPCUA{
		c: cfg,
	}
}

func (o *OPCUA) ListenAndServe() error {
	return fmt.Errorf("NOT IMPLEMENTED")
}
