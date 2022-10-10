// Sirius CLI tool thats shit and unofficial
// Windows ONLY!!!!
// this is only for joke and for rando porpuses only
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	discord string
	domainx string
	info    string
}

func initialModel() model {
	return model{
		//discord link
		discord: "https://discord.gg/sirius",
		domainx: "loadstring(game:HttpGet('https://raw.githubusercontent.com/shlexware/DomainX/main/source',true))()",
		info:    "Sirius is a free utility tool that has a mass amount of high quality features with little to no compromise on ease of use. However, some of Sirius' best features are only available in Essential or Pro. We take pride in our higher tiers of Sirius and it's the best way to experience the smartest script.",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// These keys should exit the program.
		case "q":
			return m, tea.Quit
		case "1":
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Println(m.discord)
			fmt.Println("Press enter to go back")
		case "2":
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Println(m.domainx)
			fmt.Println("Press enter to go back")
		case "3":
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Println(m.info)
			fmt.Println("Press enter to go back")
		case "enter":
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
			main()
		}
	}
	return m, nil
}

func (m model) View() string {
	welcome := "Welcome to the Sirius CLI tool\nThe current time is: " + time.Now().Format("15:04") + "\nPlease choose one of the options below:\n1. Discord\n2. DomainX Loadstring\n3. Info\nq. Exit\n"
	return welcome
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
