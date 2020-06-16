# MCLOBwVMs
Modern, container like orchestration, but with VMs only

This post on hackernews [Container technologies at Coinbase: Why Kubernetes is not part of our stack](https://news.ycombinator.com/item?id=23460066) prompted a discussion about the movie [Inception](https://en.wikipedia.org/wiki/Inception). In the movie they dream within a dream and then dream within a dream within a dream... what if we re-thought k8s with one simple rule: no dreams within dreams. If I'm running a VM, don't run another VM or container or any further virtulization because nothing good will come from it, just more and more complexity.

1. terraform to make new mediums, tiny, and large on cloud provider(s).

2. use default tiny linux available

3. script to make linux have dnf/yum installs u need

4. script to change no. files. and/or add SWAP

5. script to do src installs of stuff not in dnf/yum

Your cluster of VMs can still grow and shrink. You can orchestrate symphonies of pods and services just like before in k8s. 

Meet your new orchestra:

| provider | name of "thing" that runs | name of load balancer | more |
| --- | --- | --- | --- |
| aws | [instance](https://www.terraform.io/docs/providers/aws/r/instance.html) | [lb](https://www.terraform.io/docs/providers/aws/r/lb.html) |
| DigitalOcean | [droplet](https://www.terraform.io/docs/providers/do/r/droplet.html) | [loadbalancer](https://www.terraform.io/docs/providers/do/r/loadbalancer.html) | [domain](https://www.terraform.io/docs/providers/do/r/domain.html) [record](https://www.terraform.io/docs/providers/do/r/record.html) [ssh_key](https://www.terraform.io/docs/providers/do/r/ssh_key.html) |
| Vultr | [server](https://www.terraform.io/docs/providers/vultr/r/server.html) | [load_balancer](https://www.terraform.io/docs/providers/vultr/r/load_balancer.html) | [domain](https://www.terraform.io/docs/providers/vultr/r/dns_domain.html) [record](https://www.terraform.io/docs/providers/vultr/r/dns_record.html) [ssh_key](https://www.terraform.io/docs/providers/vultr/r/ssh_key.html) |
| linode | [instance](https://www.terraform.io/docs/providers/linode/r/instance.html) | [nodebalancer](https://www.terraform.io/docs/providers/linode/r/nodebalancer.html) | [domain](https://www.terraform.io/docs/providers/linode/r/domain.html) [record](https://www.terraform.io/docs/providers/linode/d/domain_record.html) [ssh_key](https://www.terraform.io/docs/providers/linode/d/sshkey.html) |
| google | [compute_instance](https://www.terraform.io/docs/providers/google/r/compute_instance.html) | [compute_target_pool](https://www.terraform.io/docs/providers/google/r/compute_target_pool.html) |

Example Dockerfile:

https://github.com/bitnami/bitnami-docker-wordpress/blob/master/5/debian-10/Dockerfile

Example Chart:

https://github.com/bitnami/charts/tree/master/bitnami/wordpress

Example Symphony A:

1. One load balancer with 3 targets:

  1. nginx with php on small
  2. nginx with php on small
  3. nginx with php on small

2. Two MariaDB on medium:

  1. main 
  2. replica

Example Symphony B:

1. One <a href="https://github.com/andrewarrow/feedbacks">feedbacks</a> + MariaDB on medium


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

