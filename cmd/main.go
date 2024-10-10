package main

import (
	"time"

	"github.com/codevault-llc/db-online/cmd/api"
	"github.com/codevault-llc/db-online/internal/updater"
)

func main() {
	go updater.StartAutoUpdate(20 * time.Minute)

	api.Start()
}
