package metrics

// todo вынести в когфиг настройки?
func newCounter(counterName string) *CreateCounter {
	// Webvisor
	wv := Webvisor{
		Urls:           "regexp:.*",
		ArchEnabled:    1,
		ArchType:       "none",
		LoadPlayerType: "proxy",
		WvVersion:      2,
		AllowWv2:       true,
		WvForms:        1,
	}

	// Operations
	var ops []Operation
	op := Operation{
		Action: "cut_all_parameters",
		Attr:   "url",
		Status: "active",
	}
	ops = append(ops, op)

	// Goals
	var goals []Goal
	var conditions []Condition
	condition := Condition{
		Type: "contain",
		Url:  "confirm.html",
	}

	conditions = append(conditions, condition)

	goal1 := Goal{
		Name:       "L",
		Type:       "url",
		GoalSource: "user",
		Conditions: conditions,
	}
	goals = append(goals, goal1)

	//informer
	informer := Informer{
		Enabled:    1,
		Type:       "ext",
		Indicator:  "pageviews",
		Size:       3,
		ColorStart: "FFFFFFFF",
		ColorEnd:   "EFEFEFFF",
	}

	codeOptions := CodeOptions{
		Visor:    1,
		Clickmap: 1,
		Informer: informer,
	}

	// Site 2
	site2 := Site2{
		Site: counterName,
	}

	counter := Counter{
		Name:                  counterName,
		GdprAgreementAccepted: 1,
		Site:                  counterName,
		Status:                "Active",
		ActivityStatus:        "low",
		Type:                  "simple",
		Permission:            "own",
		TimeZoneName:          "Europe/Moscow",
		Webvisor:              wv,
		Operations:            ops,
		Goals:                 goals,
		CodeOptions:           codeOptions,
		Site2:                 site2,
	}

	var createCounter CreateCounter
	createCounter.Counter = counter

	return &createCounter
}
