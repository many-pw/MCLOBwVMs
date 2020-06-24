# MCLOBwVMs
Modern, container like orchestration, but with VMs only

This post on hackernews [Container technologies at Coinbase: Why Kubernetes is not part of our stack](https://news.ycombinator.com/item?id=23460066) prompted a discussion about the movie [Inception](https://en.wikipedia.org/wiki/Inception). In the movie they dream within a dream and then dream within a dream within a dream... what if we re-thought k8s with one simple rule: no dreams within dreams. If I'm running a VM, don't run another VM or container or any further virtulization because nothing good will come from it, just more and more complexity.

1. terraform to make new mediums, tiny, and large on cloud provider(s).

2. use latest fedora when possible.

3. script to make linux have dnf installs u need

4. script to change no. files. and/or add SWAP

5. script to do src installs of stuff not in dnf

Your cluster of VMs can still grow and shrink. You can orchestrate symphonies of pods and services just like before in k8s. 

Meet your new orchestra:

| provider | name of "thing" that runs | name of load balancer | more |
| --- | --- | --- | --- |
| aws | [instance](https://www.terraform.io/docs/providers/aws/r/instance.html) | [lb](https://www.terraform.io/docs/providers/aws/r/lb.html) | [bucket](https://www.terraform.io/docs/providers/aws/r/s3_bucket.html)
| DigitalOcean | [droplet](https://www.terraform.io/docs/providers/do/r/droplet.html) | [loadbalancer](https://www.terraform.io/docs/providers/do/r/loadbalancer.html) | [domain](https://www.terraform.io/docs/providers/do/r/domain.html) [record](https://www.terraform.io/docs/providers/do/r/record.html) [ssh_key](https://www.terraform.io/docs/providers/do/r/ssh_key.html) [bucket](https://www.terraform.io/docs/providers/do/r/spaces_bucket.html) |
| Vultr | [server](https://www.terraform.io/docs/providers/vultr/r/server.html) | [load_balancer](https://www.terraform.io/docs/providers/vultr/r/load_balancer.html) | [domain](https://www.terraform.io/docs/providers/vultr/r/dns_domain.html) [record](https://www.terraform.io/docs/providers/vultr/r/dns_record.html) [ssh_key](https://www.terraform.io/docs/providers/vultr/r/ssh_key.html) |
| linode | [instance](https://www.terraform.io/docs/providers/linode/r/instance.html) | [nodebalancer](https://www.terraform.io/docs/providers/linode/r/nodebalancer.html) | [domain](https://www.terraform.io/docs/providers/linode/r/domain.html) [record](https://www.terraform.io/docs/providers/linode/d/domain_record.html) [ssh_key](https://www.terraform.io/docs/providers/linode/d/sshkey.html) |
| google | [compute_instance](https://www.terraform.io/docs/providers/google/r/compute_instance.html) | [compute_target_pool](https://www.terraform.io/docs/providers/google/r/compute_target_pool.html) | [bucket](https://www.terraform.io/docs/providers/google/r/storage_bucket.html)

Examples:

1. Video Sharing App, Day One all parts running on 1 tiny VM:

| vm | thing | description |
| --- | --- | --- |
| vm1,tiny | mysql | 1 central database, everyone talks to me |
| vm1,tiny | webserver | handles incoming https requests, enqueues new jobs to worker(s) |
| vm1,tiny | ffmpeg worker | converts uploaded videos into format needed |
| --- |
| $5.00 a month |

2. Video Sharing App, Day 30 each part has its own tiny VM:

| vm | thing |
| --- | --- |
| vm1,tiny | mysql |
| vm2,tiny | webserver |
| vm3,tiny | ffmpeg worker |
| --- |
| $15.00 a month |

3. Video Sharing App, Day 90 mysql in medium VM, multiple webservers + workers:

| vm | thing | cost |
| --- | --- | --- |
| vm1,medium | mysql | $40.00 |
| vm2.1,tiny | webserver | $5.00 |
| vm2.2,tiny | webserver | $5.00 |
| vm3.1,tiny | ffmpeg worker | $5.00 |
| vm3.2,tiny | ffmpeg worker | $5.00 |
| --- |
| $60.00 a month |

Example install MariaDB:

```
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
     "mysqladmin --user=root password 'foo2bar'",
```


Other Cloud:

https://us.alibabacloud.com/

https://www.exoscale.com/

https://gridscale.io/

https://www.hetzner.com/

https://www.ncloud.com/

https://www.nutanix.com/en

https://opennebula.io/

https://us.ovhcloud.com/

https://www.scaleway.com/en/

https://www.joyent.com/triton/compute

https://cloud.yandex.com/

