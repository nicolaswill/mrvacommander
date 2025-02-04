// Copyright © 2024 github
// Licensed under the Apache License, Version 2.0 (the "License").

package main

import (
	"flag"
	"log"
	"log/slog"
	"os"

	"mrvacommander/config/mcc"

	"mrvacommander/pkg/agent"
	"mrvacommander/pkg/logger"
	"mrvacommander/pkg/qldbstore"
	"mrvacommander/pkg/qpstore"
	"mrvacommander/pkg/queue"
	"mrvacommander/pkg/server"
	"mrvacommander/pkg/storage"
)

func main() {
	// Define flags
	helpFlag := flag.Bool("help", false, "Display help message")
	logLevel := flag.String("loglevel", "info", "Set log level: debug, info, warn, error")
	mode := flag.String("mode", "standalone", "Set mode: standalone, container, cluster")

	// Custom usage function for the help flag
	flag.Usage = func() {
		log.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		log.Println("\nExamples:")
		log.Println("  go run main.go --loglevel=debug --mode=container")
	}

	// Parse the flags
	flag.Parse()

	// Handle the help flag
	if *helpFlag {
		flag.Usage()
		return
	}

	// Apply 'loglevel' flag
	switch *logLevel {
	case "debug":
		slog.SetLogLoggerLevel(slog.LevelDebug)
	case "info":
		slog.SetLogLoggerLevel(slog.LevelInfo)
	case "warn":
		slog.SetLogLoggerLevel(slog.LevelWarn)
	case "error":
		slog.SetLogLoggerLevel(slog.LevelError)
	default:
		log.Printf("Invalid logging verbosity level: %s", *logLevel)
		os.Exit(1)
	}

	// Read configuration
	config := mcc.LoadConfig("mcconfig.toml")

	// Output configuration summary
	log.Printf("Help: %t\n", *helpFlag)
	log.Printf("Log Level: %s\n", *logLevel)
	log.Printf("Mode: %s\n", *mode)

	// Apply 'mode' flag
	switch *mode {
	case "standalone":
		// Assemble single-process version

		sl := logger.NewLoggerSingle(&logger.Visibles{})

		// FIXME take value from configuration
		sq := queue.NewQueueSingle(2, &queue.Visibles{
			Logger: sl,
		})

		ss := storage.NewStorageSingle(config.Storage.StartingID, &storage.Visibles{})

		qp, err := qpstore.NewStore(&qpstore.Visibles{})
		if err != nil {
			slog.Error("Unable to initialize query pack storage")
			os.Exit(1)
		}

		ql, err := qldbstore.NewStore(&qldbstore.Visibles{})
		if err != nil {
			slog.Error("Unable to initialize ql database storage")
			os.Exit(1)
		}

		server.NewCommanderSingle(&server.Visibles{
			Logger:         sl,
			Queue:          sq,
			ServerStore:    ss,
			QueryPackStore: qp,
			QLDBStore:      ql,
		})

		// FIXME take value from configuration
		agent.NewAgentSingle(2, &agent.Visibles{
			Logger:         sl,
			Queue:          sq,
			QueryPackStore: qp,
			QLDBStore:      ql,
		})

	case "container":
		// Assemble container version
		sl := logger.NewLoggerSingle(&logger.Visibles{})

		// FIXME take value from configuration
		sq := queue.NewQueueSingle(2, &queue.Visibles{
			Logger: sl,
		})

		ss := storage.NewStorageSingle(config.Storage.StartingID, &storage.Visibles{})

		qp, err := qpstore.NewStore(&qpstore.Visibles{})
		if err != nil {
			slog.Error("Unable to initialize query pack storage")
			os.Exit(1)
		}

		ql, err := qldbstore.NewStore(&qldbstore.Visibles{})
		if err != nil {
			slog.Error("Unable to initialize ql database storage")
			os.Exit(1)
		}

		agent.NewAgentSingle(2, &agent.Visibles{
			Logger:         sl,
			Queue:          sq,
			QueryPackStore: qp,
			QLDBStore:      ql,
		})

		server.NewCommanderSingle(&server.Visibles{
			Logger:         sl,
			Queue:          sq,
			ServerStore:    ss,
			QueryPackStore: qp,
			QLDBStore:      ql,
		})

	case "cluster":
		// Assemble cluster version
	default:
		slog.Error("Invalid value for --mode. Allowed values are: standalone, container, cluster\n")
		os.Exit(1)
	}

}
