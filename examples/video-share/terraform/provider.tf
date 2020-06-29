variable "do_token" {}
variable "pvt_key" {}
variable "ssh_fingerprint" {}
variable "mysql_root_password" {}
variable "access_id" {}
variable "secret_key" {}

provider "digitalocean" {
  token = var.do_token
  spaces_access_id  = var.access_id
  spaces_secret_key = var.secret_key
}

/*

terraform apply -var "do_token=${DO_PAT}" -var "ssh_fingerprint=${DO_SSH_FINGERPRINT}"  -var "pvt_key=$HOME/.ssh/id_rsa" -var "mysql_root_password=${MYSQL_ROOT_PASSWORD}" -var "access_id=${DO_ID}" -var "secret_key=${DO_SECRET}"

*/
