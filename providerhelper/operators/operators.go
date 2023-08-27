// Copyright (c) eden.zh.li@outlook.com, Inc.
// SPDX-License-Identifier: MPL-2.0

package operators

import (
	"github.com/LZiHaN/terraform-provider-sealos/providerhelper/cmd"
	"github.com/LZiHaN/terraform-provider-sealos/providerhelper/consts"
)

func NewClient(clusterName, sealosBinPath string) *Client {
	if clusterName == "" {
		clusterName = consts.DefaultClusterName
	}
	if sealosBinPath == "" {
		sealosBinPath = consts.DefaultSealosBinPath
	}

	localCmd := cmd.NewSealosCmd(sealosBinPath, &cmd.LocalCmd{})
	return &Client{
		Cluster:      newClusterClient(localCmd, clusterName),
		CmdInterface: localCmd.Executor,
	}
}

type Client struct {
	Cluster      ClusterInterface
	CmdInterface cmd.Provider
}

type ClusterInterface interface {
	Run(*cmd.RunCmd) ([]byte, error)
	Reset(*cmd.ResetCmd) ([]byte, error)
}
