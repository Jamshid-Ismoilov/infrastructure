package localize

// localizers for corp_api
var SmsService map[string]map[string]string = map[string]map[string]string{
	"test_message": {
		"en": "report",
		"ru": "отчет",
		"uz": "hisobot",
	},
	"balance": {
		"en":       "SMS balance is %d",
		"ru":       "СМС баланс %d",
		"uz":       "SMS balansi %d",
		"alphabet": "A",
	},
}
