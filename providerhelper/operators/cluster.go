// Copyright (c) eden.zh.li@outlook.com, Inc.
// SPDX-License-Identifier: MPL-2.0

package operators

import "github.com/LZiHaN/terraform-provider-sealos/providerhelper/cmd"

func newClusterClient(sealosCmd *cmd.SealosCmd, clusterName string) ClusterInterface {
	return &ClusterClient{
		SealosCmd:   sealosCmd,
		clusterName: clusterName,
	}
}

type ClusterClient struct {
	*cmd.SealosCmd
	clusterName string
}

func (c *ClusterClient) Run(args *cmd.RunCmd) ([]byte, error) {
	return c.SealosCmd.Run(&cmd.RunCmd{
		Cluster:    c.clusterName,
		Debug:      args.Debug,
		Cmd:        args.Cmd,
		ConfigFile: args.ConfigFile,
		Env:        args.Env,
		Force:      args.Force,
		Masters:    args.Masters,
		Nodes:      args.Nodes,
		Images:     args.Images,
		SSH:        args.SSH,
		Transport:  args.Transport,
	})
}

func (c *ClusterClient) Reset(args *cmd.ResetCmd) ([]byte, error) {
	return c.SealosCmd.Reset(&cmd.ResetCmd{
		Cluster: c.clusterName,
		Force:   args.Force,
		Debug:   args.Debug,
		Masters: args.Masters,
		Nodes:   args.Nodes,
		SSH:     args.SSH,
	})
}
