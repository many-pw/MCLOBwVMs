resource "digitalocean_droplet" "http" {
  image = "fedora-32-x64"
  name = "jjaa.me"
  region = "nyc3"
  size = "s-1vcpu-1gb"
  private_networking = true
  ssh_keys = [
   var.ssh_fingerprint
  ]
  connection {
    host = self.ipv4_address
    user = "root"
    type = "ssh"
    private_key = file(var.pvt_key)
    timeout = "2m"
  }
  provisioner "remote-exec" {
    inline = [
     "fallocate -l 4G /swapfile",
     "chmod 600 /swapfile",
     "mkswap /swapfile",
     "swapon /swapfile",
     "echo '/swapfile   none    swap    sw    0   0' >> /etc/fstab",
     "echo '[mariadb]' > /etc/yum.repos.d/MariaDB.repo",
     "echo 'name = MariaDB' >> /etc/yum.repos.d/MariaDB.repo",
     "echo 'baseurl = http://yum.mariadb.org/10.4/fedora31-amd64' >> /etc/yum.repos.d/MariaDB.repo",
     "echo 'gpgkey=https://yum.mariadb.org/RPM-GPG-KEY-MariaDB' >> /etc/yum.repos.d/MariaDB.repo",
     "echo 'gpgcheck=1' >> /etc/yum.repos.d/MariaDB.repo",
     "dnf -y install MariaDB-server MariaDB-client",
     "systemctl start mariadb.service",
     "systemctl enable mariadb.service",
     "mysqladmin --user=root password '${var.mysql_root_password}'",
     "mysql -uroot -e 'CREATE DATABASE jjaa_me CHARACTER SET utf8 COLLATE utf8_general_ci'",
     "dnf -y install go words",
     "git clone https://github.com/many-pw/MCLOBwVMs.git",
     "mkdir /http",
     "mkdir /http/templates",
     "mkdir /http/assets",
     "cd MCLOBwVMs",
     "go build; cp mclob /",
     "cd examples/video-share",
     "mysql -uroot jjaa_me < migrations/first.sql",
     "cd http",
     "cp templates/* /http/templates",
     "cp assets/* /http/assets",
     "go build; cp http /bin/",
     "/mclob --add-service http ${var.mysql_root_password}",
    ]
  }
  provisioner "file" {
    content     =  file("../../../priv_dkim.key")
    destination = "/http/priv_dkim.key"
  }
}

resource "digitalocean_domain" "jjaa_me" {
  name = "jjaa.me"
}

resource "digitalocean_record" "a" {
  depends_on = ["digitalocean_droplet.http"]
  domain = digitalocean_domain.jjaa_me.name
  type   = "A"
  name   = "@"
  value  = digitalocean_droplet.http.ipv4_address
}
resource "digitalocean_record" "mail" {
  domain = digitalocean_domain.jjaa_me.name
  type   = "CNAME"
  name   = "mail"
  value  = "@"
}
resource "digitalocean_record" "dkim" {
  domain = digitalocean_domain.jjaa_me.name
  type   = "TXT"
  value = file("../../../pub_dkim.key")
  name = "jjaame._domainkey"
}
