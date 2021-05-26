package fung

// #dependsOn {"typeCtor":"Errors", "funcTArgs":[{"type":"kType"}]}

// #importAnchor

func (a Errors_kType) IsEmpty() bool {
	return len(a) == 0
}

func (a Errors_kType) PutAll(b map[kType]error) {
	for k, v := range b {
		a[k] = v
	}
}
