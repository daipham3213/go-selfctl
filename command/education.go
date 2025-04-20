package command

func init() {
	const HELP = `my education history.`
	const TEMPLATE = `
EDUCATION:

Ho Chi Minh City University of Technology
Software Engineering | 2018 - 2022
GPA: 3.32/4.0
`
	RegisterCommand("education", "e", HELP, TEMPLATE)
}
