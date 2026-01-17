package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/yhiraki/wakeonlan-webapp/backend/config"
	"github.com/yhiraki/wakeonlan-webapp/backend/server"
	"github.com/yhiraki/wakeonlan-webapp/backend/wol"
)

//go:embed frontend/dist/*
var frontendFS embed.FS

var Version = "dev"

func main() {
	versionFlag := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if *versionFlag {
		fmt.Println(Version)
		return
	}

	// 1. Parse Config
	// Use flag.Args() instead of os.Args because flag.Parse() consumes flags
	args := flag.Args()
	targets, err := config.ParseTargets(args)
	if err != nil {
		log.Fatalf("Error parsing targets: %v\nUsage: %s [options] name=mac [name=mac ...]", err, os.Args[0])
	}
	if len(targets) == 0 {
		log.Println("Warning: No targets configured. Start with name=mac arguments.")
	}

	// 2. Initialize Services
	wolSvc := wol.NewService()

	// 3. Initialize Server
	srv := server.NewServer(targets, wolSvc, Version)

	// 4. Serve Static Files
	// Get the subdirectory of the embedded FS
	distFS, err := fs.Sub(frontendFS, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}
	
	srv.MountStatic(distFS)

	// 5. Start Server
	port := "8080"
	fmt.Printf("Starting server on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, srv); err != nil {
		log.Fatal(err)
	}
}
