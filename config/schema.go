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
	Description  string
	Type         string
	DefaultValue float64 `yaml:"default_value"`
}
