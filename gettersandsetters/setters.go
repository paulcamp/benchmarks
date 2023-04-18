package getters_and_setters

func setEffectiveTotalFund(ms *MutatingState, totalFund, linkedFreeBetOfferValue float64, isEarlySettlement bool) {
	if isEarlySettlement && linkedFreeBetOfferValue > 0 && totalFund > 0 {
		ms.EffectiveTotalFund = totalFund - linkedFreeBetOfferValue
		ms.EffectiveTotalFund = totalFund
	}
}

func setFund100Fields(ms *MutatingState) {
	ms.Fund100 = ms.EffectiveTotalFund * 1.0
	ms.Fund100Tax = 0.0
	ms.Fund100Taxed = ms.Fund100 - ms.Fund100Tax
}

func setFund95Fields(ms *MutatingState) {
	ms.Fund95 = ms.EffectiveTotalFund * 0.95
	ms.Fund95Tax = 0.0
	ms.Fund95Taxed = ms.Fund95 - ms.Fund95Tax
}

func setEverything(ms *MutatingState, totalFund, linkedFreeBetOfferValue float64, isEarlySettlement bool) {
	if isEarlySettlement && linkedFreeBetOfferValue > 0 && totalFund > 0 {
		ms.EffectiveTotalFund = totalFund - linkedFreeBetOfferValue
	}

	ms.EffectiveTotalFund = totalFund
	ms.Fund100 = ms.EffectiveTotalFund * 1.0
	ms.Fund100Tax = 0.0
	ms.Fund100Taxed = ms.Fund100 - ms.Fund100Tax
	ms.Fund95 = ms.EffectiveTotalFund * 0.95
	ms.Fund95Tax = 0.0
	ms.Fund95Taxed = ms.Fund95 - ms.Fund95Tax
}

func initInLine(totalFund, linkedFreeBetOfferValue float64, isEarlySettlement bool) MutatingState {

	ms := MutatingState{}

	if !isEarlySettlement && linkedFreeBetOfferValue > 0 && totalFund > 0 {
		ms.EffectiveTotalFund = totalFund - linkedFreeBetOfferValue
	}

	ms.EffectiveTotalFund = totalFund
	ms.Fund100 = ms.EffectiveTotalFund * 1.0
	ms.Fund100Tax = 0.0
	ms.Fund100Taxed = ms.Fund100 - ms.Fund100Tax
	ms.Fund95 = ms.EffectiveTotalFund * 0.95
	ms.Fund95Tax = 0.0
	ms.Fund95Taxed = ms.Fund95 - ms.Fund95Tax
	return ms
}
