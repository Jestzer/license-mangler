package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
)

func main() {
	// To handle keyboard input better.
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	// Colors used across the program.
	redBackground := color.New(color.BgRed).SprintFunc()
	redText := color.New(color.FgRed).SprintFunc()
	greenBackground := color.New(color.BgGreen).SprintFunc()

	// Goodies.
	var licenseFilePath string
	var logFilePath string
	var checkForUpdatesOnLaunch bool = true

	// Setup for better Ctrl+C messaging. This is a channel to receive OS signals.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Start a Goroutine to listen for signals.
	go func() {

		// Wait for the signal.
		<-signalChan

		// Handle the signal by exiting the program.
		fmt.Print(redBackground("\nExiting from user input..."))
		os.Exit(0)
	}()

	// Determine any user-defined settings.
	currentDir, err := os.Getwd() // Get the current working directory.
	if err != nil {
		fmt.Print(redText("\nError getting current working directory while looking for user settings : ", err, " Default settings will be used instead."))
		return
	} else {
		settingsPath := filepath.Join(currentDir, "settings.txt")

		// Check if the settings file exists.
		if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
			// No settings found.
			return
		} else if err != nil {
			fmt.Print(redText("\nError checking for user settings: ", err, " Default settings will be used instead."))
		} else {
			fmt.Print("\nCustom settings found!")
			file, err := os.Open(settingsPath)
			if err != nil {
				fmt.Print(redText("\nError opening settings file: ", err, " Default settings will be used instead."))
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				line := scanner.Text()

				if !strings.HasPrefix(line, "#") {
					if strings.HasPrefix(strings.ToLower(line), "checkforupdatesonlaunch") {
						if strings.Contains(strings.ToLower(line), "false") {
							checkForUpdatesOnLaunch = false
							fmt.Print("\nUpdates for this program will not be checked per your settings.")
						}

					} else if strings.HasPrefix(line, "licenseFilePath =") || strings.HasPrefix(line, "licenseFilePath=") {
						licenseFilePath = strings.TrimPrefix(line, "licenseFilePath =")
						licenseFilePath = strings.TrimPrefix(licenseFilePath, "licenseFilePath=")
						licenseFilePath = strings.TrimSpace(licenseFilePath)
						licenseFilePath = strings.Trim(licenseFilePath, "\"")

						_, err := os.Stat(licenseFilePath) // Do you actually exist? Does anything actually exist, man?
						if err != nil {
							fmt.Print(redText("\nThe licenese file path you've specified, \"", licenseFilePath, " does not exist. Please adjust your settings accordingly."))
							os.Exit(0)
						}

						fmt.Print("\nYour log file path has been set to ", logFilePath, " per your settings.")

					} else if strings.HasPrefix(line, "logFilePath =") || strings.HasPrefix(line, "logFilePath=") {
						logFilePath = strings.TrimPrefix(line, "logFilePath =")
						logFilePath = strings.TrimPrefix(logFilePath, "logFilePath=")
						logFilePath = strings.TrimSpace(logFilePath)
						logFilePath = strings.Trim(logFilePath, "\"")

						_, err := os.Stat(logFilePath) // Do you actually exist? Does anything actually exist, man?
						if err != nil {
							fmt.Print(redText("\nThe log file path you've specified, \"", logFilePath, " does not exist. Please adjust your settings accordingly."))
							os.Exit(0)
						}

						fmt.Print("\nYour log file path has been set to ", logFilePath, " per your settings.")
					}
				}
			}

			if err := scanner.Err(); err != nil {
				fmt.Print(redText("\nError reading settings file:", err, " Default settings will be used instead."))
			}
		}
	}
}

func SetLicenseFilePath() {

}
func SetLogFilePath() {

}
func SetLicenseManagerPath() {

}
func StartLicenseManager() {

}
func StopLicenseManager() {

}
func CheckLicenseManagerStatus() {

}
