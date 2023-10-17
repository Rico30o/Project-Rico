package main

import (
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/generate-kill-script", func(c *fiber.Ctx) error {
		script := `#!/bin/bash
# Replace with your process name or PID
process_name="your_process_name"

# Find the process ID
pid=$(pgrep "$process_name")

if [ -z "$pid" ]; then
 echo "Process not found: $process_name"
else
 # Kill the process
 kill -9 $pid
 echo "Killed process: $process_name (PID: $pid)"
fi
`

		// Save the script to a file
		err := ioutil.WriteFile("kill.sh", []byte(script), 0755)
		if err != nil {
			return err
		}

		// Send the script file as a response
		return c.SendFile("kill.sh")
	})

	app.Listen(":3000")
}
