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

			im.checkNodeAHRK()

			time.Sleep(time.Minute * 15)
		}
	}()
}

func (im *InfrastructureMonitor) checkSponsoredFees() {
	abr, err := gowaves.WNC.AddressesBalance(FeeAddress)
	if err == nil && abr.Balance < int(SatInBTC) {
		logTelegram("Fee address WAVES balance is too small.")
	} else if err != nil {
		logTelegram("Error checking fee balance.")
	}
}

func (im *InfrastructureMonitor) checkNodeAHRK() {
	abr, err := gowaves.WNC.AssetsBalance(AHRKAddress, AHRKId)
	if err == nil && int(abr.Balance) < 20000*int(AHRKDec) {
		logTelegram("AHRK balance on AHRK node is too small.")
	} else if err != nil {
		logTelegram("Error checking AHRK node balance.")
	}
}

func initInfrastructureMonitor() {
	im := &InfrastructureMonitor{}
	im.start()
}
