package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

const dtracePath = "/usr/sbin/dtrace"
const defaultScript = "../dtrace-scripts/events.d"
const defaultTrace = "trace.json"
const defaultLog = "eq.log"

func main() {
	// Parse flags
	script := flag.String("script", defaultScript, "Script to run")
	traceName := flag.String("trace", defaultTrace, "Trace output filename")
	logName := flag.String("log", defaultLog, "Error log filename")
	flag.Parse()
	fmt.Printf("Starting up eq with script %s\n", *script)

	// Open trace and error log files
	traceFile, err := os.Create(*traceName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "eq: %s\n", err)
		os.Exit(1)
	}
	logFile, err := os.Create(*logName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "eq: %s\n", err)
		os.Exit(1)
	}

	// Run DTrace
	cmd := exec.Command(dtracePath, "-s", *script)
	cmd.Stdout = traceFile
	cmd.Stderr = logFile
	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "eq: Failed to run dtrace: %s\n", err)
		os.Exit(1)
	}
}
