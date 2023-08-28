// Copyright (c) eden.zh.li@outlook.com, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/LZiHaN/terraform-provider-sealos/providerhelper/cmd"
	"github.com/LZiHaN/terraform-provider-sealos/providerhelper/operators"
	"github.com/LZiHaN/terraform-provider-sealos/providerhelper/ssh"
	"github.com/LZiHaN/terraform-provider-sealos/providerhelper/utils/conversions"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceClusterCreate,
		ReadContext:   resourceClusterRead,
		UpdateContext: resourceClusterUpdate,
		DeleteContext: resourceClusterDelete,
		Schema: map[string]*schema.Schema{
			"cluster_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the cluster where the operation is to be run.",
			},
			"masters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The master nodes to be run.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"nodes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The node nodes to be run.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"command": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Overwrite the CMD instruction in the image.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"config_file": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The path to the custom configuration file, used to replace resources.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"env": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The environment variables set during command execution.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"images": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The images parameter is the name and version of the Docker image you want to run in the cluster.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"transport": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Load image transport from a tar archive file. (Optional values: oci-archive, docker-archive)",
			},
			"sealos_binary": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path to sealos executable (binary).",
			},
			"ssh": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "SSH command configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The username for authentication.",
						},
						"passwd": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Authenticate using the provided password.",
						},
						"pk": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Choose the private key file from which to read the public key authentication identity.",
						},
						"pk_passwd": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The password to decrypt the PEM-encoded private key.",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The connection port of the remote host.",
						},
					},
				},
			},
		},
	}
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config, diags := readClusterConfig(d)
	if diags.HasError() {
		return diags
	}

	sealosBinary := d.Get("sealos_binary").(string)

	client := operators.NewClient(config.Cluster, sealosBinary)

	outputByte, err := client.Cluster.Run(
		&cmd.RunCmd{
			Cluster:    config.Cluster,
			Debug:      true,
			Cmd:        config.Cmd,
			ConfigFile: config.ConfigFile,
			Env:        config.Env,
			Force:      true,
			Masters:    config.Masters,
			Nodes:      config.Nodes,
			Images:     config.Images,
			SSH:        config.SSH,
			Transport:  config.Transport,
		})
	fmt.Println("client.Cluster.Run 输出日志: ", string(outputByte))

	if err != nil {
		return diag.Errorf("client.Cluster.Run 执行出错:", string(outputByte))
	}

	d.SetId("08270452")
	return resourceClusterRead(ctx, d, meta)
}

func resourceClusterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceClusterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceClusterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// TODO 当集群创建失败时，销毁操作不会做任何事情
	config, diags := readClusterConfig(d)
	if diags.HasError() {
		return diags
	}

	sealosBinary := d.Get("sealos_binary").(string)

	client := operators.NewClient(config.Cluster, sealosBinary)

	outputByte, err := client.Cluster.Reset(&cmd.ResetCmd{
		Cluster: config.Cluster,
		Force:   true,
		Debug:   true,
		Masters: config.Masters,
		Nodes:   config.Nodes,
		SSH:     config.SSH,
	})
	fmt.Println("client.Cluster.Reset 输出日志: ", string(outputByte))

	if err != nil {
		return diag.Errorf("client.Cluster.Reset 执行出错:", string(outputByte))
	}

	d.SetId("")
	return nil
}

func readClusterConfig(d *schema.ResourceData) (cmd.RunCmd, diag.Diagnostics) {
	var config cmd.RunCmd

	// 从模式数据中读取并填充配置结构体
	config.Cluster = d.Get("cluster_name").(string)

	sshConfRaw, sshConfExists := d.GetOk("ssh.0")
	if sshConfExists {
		sshConf := sshConfRaw.(map[string]interface{})
		config.SSH = &ssh.SSH{
			User:     conversions.GetStringValue(sshConf, "user"),
			Passwd:   conversions.GetStringValue(sshConf, "passwd"),
			Pk:       conversions.GetStringValue(sshConf, "pk"),
			PkPasswd: conversions.GetStringValue(sshConf, "pk_passwd"),
			Port:     conversions.GetUInt16Value(sshConf, "port"),
		}
	}

	config.Masters = conversions.ToStringSlice(d.Get("masters"))
	config.Nodes = conversions.ToStringSlice(d.Get("nodes"))
	config.Cmd = conversions.ToStringSlice(d.Get("command"))
	config.ConfigFile = conversions.ToStringSlice(d.Get("config_file"))
	config.Env = conversions.ToStringSlice(d.Get("env"))
	config.Images = conversions.ToStringSlice(d.Get("images"))
	config.Transport = d.Get("transport").(string)

	return config, nil
}
