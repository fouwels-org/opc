// SPDX-FileCopyrightText: 2021 Kaelan Thijs Fouwels <kaelan.thijs@fouwels.com>
//
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"log"
	"opc/config"
	"opc/open62541"
	"os"
	"path/filepath"
	"strconv"
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

	if cTagList == "" {
		return fmt.Errorf("CONFIG_TAGLIST is not set")
	}

	if cPort == "" {
		return fmt.Errorf("PORT is not set")
	}

	iPort, err := strconv.ParseUint(cPort, 10, 32)
	if err != nil {
		return fmt.Errorf("failed too parse PORT as integer")
	}

	f, err := os.Open(filepath.Clean(cTagList))
	if err != nil {
		return fmt.Errorf("failed to open %v: %w", cTagList, err)
	}

	config, err := config.LoadConfig(f)
	e2 := f.Close()
	if e2 != nil {
		return fmt.Errorf("failed to close config file: %w", err)
	}
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	err = open62541.Initialise(uint16(iPort), 0)
	if err != nil {
		return fmt.Errorf("failed to create open62541 server: %w", err)
	}

	namespaces := map[string]bool{}
	for _, v := range config.TagList.Tags {
		namespaces[v.Namespace] = true
	}

	for k := range namespaces {
		err := open62541.CreateObjectNode(k, k)
		if err != nil {
			return fmt.Errorf("failed to create namespace: %v: %w", k, err)
		}
	}

	for _, v := range config.TagList.Tags {
		switch v.Type {
		case "bool":
			def := false
			if v.DefaultValue == 1 {
				def = true
			}
			err := open62541.CreateBool(v.Name, v.Description, v.Namespace, def)
			if err != nil {
				return err
			}
		case "uint16":
			err := open62541.CreateUInt16(v.Name, v.Description, v.Namespace, uint16(v.DefaultValue))
			if err != nil {
				return err
			}
		case "int16":
			err := open62541.CreateInt16(v.Name, v.Description, v.Namespace, int16(v.DefaultValue))
			if err != nil {
				return err
			}
		case "uint32":
			err := open62541.CreateUInt32(v.Name, v.Description, v.Namespace, uint32(v.DefaultValue))
			if err != nil {
				return err
			}
		case "int32":
			err := open62541.CreateInt32(v.Name, v.Description, v.Namespace, int32(v.DefaultValue))
			if err != nil {
				return err
			}
		case "float32":
			err := open62541.CreateFloat32(v.Name, v.Description, v.Namespace, float32(v.DefaultValue))
			if err != nil {
				return err
			}
		case "float64":
			err := open62541.CreateFloat64(v.Name, v.Description, v.Namespace, float64(v.DefaultValue))
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unrecognised type %v for %+v", v.Type, v)
		}
	}

	return open62541.ListenAndServe()
}
