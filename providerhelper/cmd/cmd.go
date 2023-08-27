// Copyright (c) eden.zh.li@outlook.com, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"github.com/LZiHaN/terraform-provider-sealos/providerhelper/utils/logger"
	"os/exec"
	"strconv"
)

// Provider defines the interface for executing commands
type Provider interface {
	Exec(cmd string, args ...string) ([]byte, error)
}

// LocalCmd implements the Interface for local command execution using os/exec
type LocalCmd struct{}

// Exec executes the given command on the local machine
func (c *LocalCmd) Exec(cmd string, args ...string) ([]byte, error) {
	return exec.Command(cmd, args...).CombinedOutput()
}

type ClusterExecutor interface {
	Run(*RunCmd) error
	Reset(*ResetCmd) error
}

type CommandOptions interface {
	Args() []string
}

type Args []string

func (args Args) appendFlagsWithValues(flagName string, values interface{}) Args {
	switch vv := values.(type) {
	case []string:
		if vv == nil {
			return args
		}
		for _, v := range vv {
			if flagName != "" {
				args = append(args, flagName)
			}
			args = append(args, v)
		}
	case string:
		if vv == "" {
			return args
		}
		if flagName != "" {
			args = append(args, flagName)
		}
		args = append(args, vv)
	case bool:
		if vv {
			if flagName != "" {
				args = append(args, flagName)
			}
		}
	case uint16:
		if vv == 0 {
			return args
		}
		if flagName != "" {
			args = append(args, flagName)
		}
		args = append(args, strconv.Itoa(int(vv)))
	case int:
		if vv == 0 {
			return args
		}
		if flagName != "" {
			args = append(args, flagName)
		}
		args = append(args, strconv.Itoa(vv))
	default:
		logger.Error("Unsupported %s type %T", flagName, vv)
	}
	return args
}
