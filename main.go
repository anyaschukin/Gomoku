package main

import (
	gomoku "Gomoku/src"
	// "flag"
	// "github.com/pkg/profile"
	// "runtime"
	// "log"
	// "os"
	// "runtime/pprof"
)

// var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	// flag.Parse()
	// if *cpuprofile != "" {
	// f, err := os.Create(*cpuprofile)
	// if err != nil {
	// log.Fatal(err)
	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()
	// }
	// runtime.SetCPUProfileRate(200)
	// profile.CPUProfileRate(1),
	// profile.NoShutdownHook
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	gomoku.Play()
}

// ## To run enter either command:
// go run main.go
// go build; ./Gomoku
