package command

func init() {
	const HELP = `tells you about me and this project.`
	const TEMPLATE = `
Hi there! I am Pham Le Gia Dai and Welcome to my portfolio!

I am a software engineer with a passion for building things. I have experience
in web development, cloud computing, and DevOps. I am always looking for new
challenges and opportunities to learn and grow.

This project is a CLI tool that showcases my portfolio. You can use it to learn
more about me, my work experience, education, certificates, and social media
links. The project written in Go and uses urfave/cli for the CLI interface
and then compiled into a standalone WASM binary and can be run anywhere,
including your browser.

You can also execute some commands like "experience", "education",
"certificates"", and "social" to see more about me.
Feel free to reach out to me at daipham.3213@gmail.com if you have any
questions in mind or just want to say hi!

You can check out the source code at https://github.com/daipham3213/go-selfctl.git
`
	RegisterCommand("about", "ab", HELP, TEMPLATE)
}
