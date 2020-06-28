variable "do_token" {}
variable "pvt_key" {}
variable "ssh_fingerprint" {}
variable "mysql_root_password" {}

provider "digitalocean" {
  token = var.do_token
}

/*

terraform apply -var "do_token=${DO_PAT}" -var "ssh_fingerprint=${DO_SSH_FINGERPRINT}"  -var "pvt_key=$HOME/.ssh/id_rsa" -var "mysql_root_password=${MYSQL_ROOT_PASSWORD}"

*/
