package command

func init() {
	const HELP = `my awards and recognitions.`
	const TEMPLATE = `
CERTIFICATES:

- Certified Kubernetes for application developers (CKAD) | 2024

NOMINATIONS:

- OpenStack Caracal Contributor - Pham Le Gia Dai | https://www.openstack.org/software/openstack-caracal/
- OpenStack Dalmatian Contributor - Pham Le Gia Dai | https://www.openstack.org/software/openstack-dalmatian-2/

MEETUPS:
- Distributed tracing for microservices in OpenStack | https://superuser.openinfra.org/articles/viet-openinfra-user-group-holds-its-30th-meetup/
`
	RegisterCommand("awards", "a", HELP, TEMPLATE)
}
