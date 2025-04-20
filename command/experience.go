package command

func init() {
	const HELP = `tells you about my work experience, projects, and roles.`
	const TEMPLATE = `
--------------------------------------------------------------------------------
Fullstack Developer | FPT Smart Cloud | 2022 - present
--------------------------------------------------------------------------------
* FPT File Storage:
	- Designed and developed a high-availability File Storage service with a
      multi-tier storage architecture, optimizing performance and cost 
      efficiency for diverse workloads.
    - Implemented a high-performance storage tier specifically tailored for AI
      and machine learning workloads, reducing data retrieval times and 
      accelerating training processes.
    - Enabled seamless scalability and reliability, ensuring data availability 
      and integrity even under high-demand conditions.
    - Integrated with OpenStack and other open-source technologies to provide a 
      flexible and vendor-agnostic storage solution.
    - Tech stack: Python, ReactJS, VAST, DDN, Prometheus, Grafana

* FPT VPNaaS:
	- Designed and developed a new VPN as a Service (VPNaaS) solution compatible
      with OpenStack, enabling seamless, scalable, and highly available 
      networking for users.
    - Integrated an exporter for the monitoring service, improving system 
      observability and reducing incident response time by 40%.
    - Developed management tools for VPNaaS, integrating with the OpenStack CLI
      to streamline operations and reduce manual configuration time by 50%.
    - Supported the migration and configuration of VPNaaS, ensuring a smooth 
      transition with zero downtime and improved network reliability.
    - Tech stack: Python, Prometheus, Grafana, RabbitMQ, Openstack (Neutron, 
	  Cinder, Keystone, Glance,  Nova) Openstack CLI, K8s, ArgoCD.

* FPT Object Storage:
    - Refactored and enhanced the existing codebase, improving maintainability
      and reducing technical debt, leading to a 25% increase in development
      efficiency.
    - Added support for multi-backend Software-Defined Storage (SDS), enhancing
	  flexibility and optimizing storage performance across diverse workloads.
    - Improved the UX/UI of the Object Storage service on the FPT Cloud portal, 
	  streamlining user interactions and boosting user satisfaction by 30%.
    - Tech stack: Python, ReactJS, CEPH.

* FPT Backup Native:
    - Develop an agent-less backup solution for FPT Cloud Platform.
    - Develop exporter for monitoring service.
    - Integrate into Billing service with FPT Cloud Platform.
    - Tech stack: Python, Prometheus, Grafana, Kafka, RabbitMQ, Openstack
      Cinder, AWX, Ansible, K8s.

* FPT AutoScaling:
    - Integrate Openstack Senlin with FPT Cloud Platform.
    - Develop new driver for integrate Senlin with VMWare vCenter.
    - Develop new auto-scaling policy support AVI load balancer.
    - Tech stack: Python, Openstack Senlin, VMWare, AVI load balancer.

--------------------------------------------------------------------------------
Intern Backend Developer | CMC TSSG | 2021 - 2022
--------------------------------------------------------------------------------
- Learned about the software development process.
- Tech stack: Django, DRF, MySQL, Docker, Git.
`
	RegisterCommand("experience", "ex", HELP, TEMPLATE)
}
