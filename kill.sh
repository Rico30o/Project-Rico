#!/bin/bash
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
