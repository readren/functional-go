package typectors

type setOfTemplateArgs []TemplateArguments

func (taSetPtr *setOfTemplateArgs) add(taPtr *TemplateArguments) bool {
	if taSetPtr.contains(taPtr) {
		return false
	} else {
		*taSetPtr = append(*taSetPtr, *taPtr)
		return true
	}
}

func (taSet setOfTemplateArgs) contains(taPtr *TemplateArguments) bool {
	for i := range taSet {
		if taSet[i].IsEqual(taPtr) {
			return true
		}
	}
	return false
}

// Gives a new `requirementSet` that contains all the requirements that are contained by this `requirementSet` and not contained by the other `requirementSet`.
func (taSet setOfTemplateArgs) diff(other setOfTemplateArgs) setOfTemplateArgs {
	newSet := make(setOfTemplateArgs, 0, len(taSet))
	for i := range taSet {
		p := &taSet[i]
		if !other.contains(p) {
			newSet.add(p)
		}
	}
	return newSet
}
