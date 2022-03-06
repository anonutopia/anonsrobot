package main

import (
	"time"

	"github.com/anonutopia/gowaves"
)

type InfrastructureMonitor struct{}

func (im *InfrastructureMonitor) start() {
	go func() {
		for {
			im.checkSponsoredFees()

			time.Sleep(time.Minute * 15)
		}
	}()
}

func (im *InfrastructureMonitor) checkSponsoredFees() {
	abr, err := gowaves.WNC.AddressesBalance(FeeAddress)
	if err == nil && abr.Balance < int(SatInBTC) {
		logTelegram("Fee address WAVES balance is too small.")
	} else {
		logTelegram("Error checking fee balance.")
	}
}

func initInfrastructureMonitor() {
	im := &InfrastructureMonitor{}
	im.start()
}
