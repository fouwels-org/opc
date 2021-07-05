// SPDX-FileCopyrightText: 2021 Kaelan Thijs Fouwels <kaelan.thijs@fouwels.com>
//
// SPDX-License-Identifier: MIT

package config

type Config struct {
	TagList TagList
}

type TagList struct {
	Meta TagListMeta
	Tags []TagListTag
}

type TagListMeta struct {
	Site    string
	Comment string
}

type TagListTag struct {
	Name         string
	Namespace    string
	Comment      string
	OpcType      string `yaml:"opc_type"`
	DefaultValue int    `yaml:"default_value"`
}
