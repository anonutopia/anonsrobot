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

			im.checkNodeAINTAHRK()

			// im.checkAINTOrders()

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

func (im *InfrastructureMonitor) checkNodeAINTAHRK() {
	abr, err := gowaves.WNC.AssetsBalance(AINTAddress, AHRKId)
	if err == nil && int(abr.Balance) < 1000*int(AHRKDec) {
		logTelegram("AHRK balance on AINT node is too small.")
	} else if err != nil {
		logTelegram("Error checking AINT node balance.")
	}
}

func (im *InfrastructureMonitor) checkAINTOrders() {
	opr, err := gowaves.WMC.OrderbookPair(AINTId, "WAVES", 10)
	if err == nil {
		if len(opr.Asks) < 2 {
			logTelegram("There are not enough AINT orders on DEX.")
		}
	} else {
		logTelegram("Error checking AINT orders.")
	}
}

func initInfrastructureMonitor() {
	im := &InfrastructureMonitor{}
	im.start()
}
