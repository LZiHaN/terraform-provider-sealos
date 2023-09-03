# Terraform Provider for Sealos

The Terraform Provider for Sealos provides a more straightforward and robust means of executing Sealos automation from Terraform than local-exec. Users can run Sealos on infrastructure provisioned by Terraform.

This provider can be [found in the Terraform Registry here](https://registry.terraform.io/providers/LZiHaN/sealos/latest).


## Requirements

- install Go: [official installation guide](https://go.dev/doc/install)
- install Terraform: [official installation guide](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)
- install Sealos: [official installation guide](https://sealos.io/docs/lifecycle-management/quick-start/installation)


## Getting Started

This is a small example of how to install the Kubernetes cluster. Please read the [documentation](https://registry.terraform.io/providers/LZiHaN/sealos/latest/docs) for more
information.

```hcl
provider "sealos" {}

resource "sealos_cluster" "example" {
  cluster_name = "test-cluster"
  masters      = ["ip1", "ip2", "ip3"]
  nodes        = ["ip4", "ip5", "ip6"]
  images       = ["testimage1", "testimage2", "testimage3"]
  ssh {
    user      = "test-user"
    passwd    = "test-passwd"
    port      = 22
  }
}
```

### Examples
The [examples](./examples/) subdirectory contains a usage example for this provider.

## Contributing

The Sealos Provider for Terraform is the work of many contributors. We appreciate your help!

