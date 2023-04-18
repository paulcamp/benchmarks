package getters_and_setters

func getEffectiveTotalFund(totalFund, linkedFreeBetOfferValue float64, isEarlySettlement bool) float64 {

	if !isEarlySettlement && linkedFreeBetOfferValue > 0 && totalFund > 0 {
		return totalFund - linkedFreeBetOfferValue
	}

	return totalFund
}

func getFund100Fields(effectiveTotalFund float64) (float64, float64, float64) {
	Fund100 := effectiveTotalFund * 1.0
	Fund100Tax := 0.0
	Fund100Taxed := Fund100 - Fund100Tax
	return Fund100, Fund100Tax, Fund100Taxed
}

func getFund95Fields(effectiveTotalFund float64) (float64, float64, float64) {

	Fund95 := effectiveTotalFund * 0.95
	Fund95Tax := 0.0
	Fund95Taxed := Fund95 - Fund95Tax

	return Fund95, Fund95Tax, Fund95Taxed
}
