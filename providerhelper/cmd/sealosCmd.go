// Copyright (c) eden.zh.li@outlook.com, Inc.
// SPDX-License-Identifier: MPL-2.0

package cmd

func NewSealosCmd(binPath string, executor Provider) *SealosCmd {
	return &SealosCmd{
		BinPath:  binPath,
		Executor: executor,
	}
}

type SealosCmd struct {
	BinPath  string
	Executor Provider
	ClusterExecutor
}

func (s *SealosCmd) Run(args *RunCmd) ([]byte, error) {
	return s.Executor.Exec(s.BinPath, append([]string{"run"}, args.Args()...)...)
}

func (s *SealosCmd) Reset(args *ResetCmd) ([]byte, error) {
	return s.Executor.Exec(s.BinPath, append([]string{"reset"}, args.Args()...)...)
}
