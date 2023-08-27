// Copyright (c) eden.zh.li@outlook.com, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"github.com/LZiHaN/terraform-provider-sealos/providerhelper/ssh"
	"strings"
)

type ResetCmd struct {
	Cluster string
	Force   bool
	Debug   bool
	Masters []string
	Nodes   []string
	SSH     *ssh.SSH
}

func (rc *ResetCmd) Args() []string {
	if rc.SSH == nil {
		rc.SSH = &ssh.SSH{}
	}
	var args Args = []string{}
	return args.appendFlagsWithValues("--cluster", rc.Cluster).
		appendFlagsWithValues("--debug", rc.Debug).
		appendFlagsWithValues("--force", rc.Force).
		appendFlagsWithValues("--masters", strings.Join(rc.Masters, ",")).
		appendFlagsWithValues("--nodes", strings.Join(rc.Nodes, ",")).
		appendFlagsWithValues("--user", rc.SSH.User).
		appendFlagsWithValues("--passwd", rc.SSH.Passwd).
		appendFlagsWithValues("--pk", rc.SSH.Pk).
		appendFlagsWithValues("--pk-passwd", rc.SSH.PkPasswd).
		appendFlagsWithValues("--port", rc.SSH.Port)
}
