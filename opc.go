// SPDX-FileCopyrightText: 2021 Kaelan Thijs Fouwels <kaelan.thijs@fouwels.com>
//
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"log"
	"opc/config"
	"opc/opcua"
	"os"
)

func main() {
	err := run()
	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("exit without error?: %v", err)
	}

	os.Exit(1)
}

func run() error {

	cTagList := os.Getenv("CONFIG_TAGLIST")
	cPort := os.Getenv("PORT")
	cAddress := os.Getenv("ADDRESS")

	if cTagList == "" {
		return fmt.Errorf("CONFIG_TAGLIST is not set")
	}

	if cPort == "" {
		return fmt.Errorf("PORT is not set")
	}

	if cAddress == "" {
		return fmt.Errorf("ADDRESS is not set")
	}

	f, err := os.Open(cTagList)
	if err != nil {
		return fmt.Errorf("failed to open %v: %w", cTagList, err)
	}
	config, err := config.LoadConfig(f)
	f.Close()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	op := opcua.NewOPCUA(config)

	err = op.ListenAndServe()
	return err
}
