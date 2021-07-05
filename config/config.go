// SPDX-FileCopyrightText: 2021 Kaelan Thijs Fouwels <kaelan.thijs@fouwels.com>
//
// SPDX-License-Identifier: MIT

package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfig(taglist *os.File) (Config, error) {

	c := Config{}

	y := yaml.NewDecoder(taglist)
	err := y.Decode(&c.TagList)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
