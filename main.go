package main

import (
	"embed"
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

func main() {
	// 1. Parse Config
	targets, err := config.ParseTargets(os.Args[1:])
	if err != nil {
		log.Fatalf("Error parsing targets: %v\nUsage: %s name=mac [name=mac ...]", err, os.Args[0])
	}
	if len(targets) == 0 {
		log.Println("Warning: No targets configured. Start with name=mac arguments.")
	}

	// 2. Initialize Services
	wolSvc := wol.NewService()

	// 3. Initialize Server
	srv := server.NewServer(targets, wolSvc)

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
