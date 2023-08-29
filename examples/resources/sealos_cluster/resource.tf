resource "sealos_cluster" "example" {
  cluster_name = "test-cluster"
  masters      = ["ip1", "ip2", "ip3"]
  nodes        = ["ip4", "ip5", "ip6"]
  command      = ["echo", "hello", "sealos"]
  config_file  = ["test-config-file"]
  env          = ["test-env"]
  images       = ["testimage1", "testimage2", "testimage3"]
  transport    = "test-transport"
  ssh {
    user      = "test-user"
    passwd    = "test-passwd"
    pk        = "test-pk"
    pk_passwd = "test-pk-passwd"
    port      = 22
  }
  sealos_binary = "/usr/bin/sealos"
}