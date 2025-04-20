package command

func init() {
	const HELP = `my social media links and timezone information.`
	const TEMPLATE = `
SOCIALS:

- GitHub | https://github.com/daipham3213
- Launchpad | https://launchpad.net/~daiplg
- Gerrit | https://review.opendev.org/q/owner:daiplg
- LinkedIn | https://www.linkedin.com/in/daipham-3213
- Email | daipham.3213@gmail.com

Timezone: UTC+07 (Asia/Ho_Chi_Minh)
Please feel free to reach out to me!
`

	RegisterCommand("social", "s", HELP, TEMPLATE)
}
