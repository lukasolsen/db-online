package updater

import (
	"time"
)

func StartAutoUpdate(interval time.Duration) {
	go updatePokemon() // Start updating Pokemon in a separate goroutine

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		<-ticker.C
		go updatePokemon() // Start updating Pokemon in a separate goroutine
	}
}
