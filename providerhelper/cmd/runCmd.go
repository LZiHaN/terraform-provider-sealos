// Copyright (c) eden.zh.li@outlook.com, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"github.com/LZiHaN/terraform-provider-sealos/providerhelper/ssh"
	"strings"
)

type RunCmd struct {
	Cluster    string
	Debug      bool
	Cmd        []string
	ConfigFile []string
	Env        []string
	Force      bool
	Masters    []string
	Nodes      []string
	Images     []string
	SSH        *ssh.SSH
	Transport  string
}

func (rc *RunCmd) Args() []string {
	if rc.SSH == nil {
		rc.SSH = &ssh.SSH{}
	}
	var args Args = []string{}
	return args.appendFlagsWithValues("--cluster", rc.Cluster).
		appendFlagsWithValues("--debug", rc.Debug).
		appendFlagsWithValues("--masters", strings.Join(rc.Masters, ",")).
		appendFlagsWithValues("--nodes", strings.Join(rc.Nodes, ",")).
		appendFlagsWithValues("", rc.Images).
		appendFlagsWithValues("--cmd", rc.Cmd).
		appendFlagsWithValues("--env", rc.Env).
		appendFlagsWithValues("--config-file", rc.ConfigFile).
		appendFlagsWithValues("--user", rc.SSH.User).
		appendFlagsWithValues("--passwd", rc.SSH.Passwd).
		appendFlagsWithValues("--pk", rc.SSH.Pk).
		appendFlagsWithValues("--pk-passwd", rc.SSH.PkPasswd).
		appendFlagsWithValues("--port", rc.SSH.Port).
		appendFlagsWithValues("--transport", rc.Transport)
}
