// SPDX-FileCopyrightText: 2021 Kaelan Thijs Fouwels <kaelan.thijs@fouwels.com>
//
// SPDX-License-Identifier: MIT

package config

import (
	"log"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	f, err := os.Open("taglist.yml")
	if err != nil {
		t.Fatalf("failed to open: %v", err)
	}
	defer f.Close()

	config, err := LoadConfig(f)
	if err != nil {
		t.Fatalf("failed to load: %v", err)
	}

	for _, v := range config.TagList.Tags {
		log.Printf("%+v", v)
	}
}
