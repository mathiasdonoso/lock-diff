package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mathiasdonoso/lockdiff/internal"
)

func Execute() {
	if len(os.Args) < 3 {
		showHelp()
		return
	}

	dependenciesAdapter := internal.NewPackageLockAdapter()
	reporterAdapter := internal.NewReportAdapter(os.Stdout)
	reporterService := internal.NewService(dependenciesAdapter, reporterAdapter)

	reportHandler := internal.NewHandler(reporterService)
	err := reportHandler.HandleReport(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}
}

func showHelp() {
	fmt.Printf("Usage: lockdiff <path/to/package-lock1.json> <path/to/package-lock2.json>\n\n")

	fmt.Printf("Description:\n")
	fmt.Printf("Compare two package-lock.json files and display the version differences of shared dependencies.\n\n")

	fmt.Printf("Example:\n")
	fmt.Printf("lockdiff ./package-lock1.json ./package-lock2.json\n")
}
