package metrics

import "time"

type CounterRESP struct {
	Rows     int       `json:"rows"`
	Counters []Counter `json:"counters"`
}

type CreateCounter struct {
	Counter Counter `json:"counter"`
}

//type createNewCounter struct {
//	Name                  string      `json:"name"`
//	GdprAgreementAccepted int         `json:"gdpr_agreement_accepted"`
//	Site                  string      `json:"site"`
//	Webvisor              Webvisor    `json:"webvisor"`
//	Goals                 []Goal      `json:"goals"`
//	Operations            []Operation `json:"operations"`
//	CodeOptions           CodeOptions `json:"code_options"`
//}

type Counter struct {
	Id                    int         `json:"id"`
	Status                string      `json:"status"`
	OwnerLogin            string      `json:"owner_login"`
	CodeStatus            string      `json:"code_status"`
	ActivityStatus        string      `json:"activity_status"`
	Name                  string      `json:"name"`
	Type                  string      `json:"type"`
	Favorite              int         `json:"favorite"`
	HideAddress           int         `json:"hide_address"`
	Permission            string      `json:"permission"`
	Webvisor              Webvisor    `json:"webvisor"`
	Goals                 []Goal      `json:"goals"`
	Operations            []Operation `json:"operations"`
	CodeOptions           CodeOptions `json:"code_options"`
	CreateTime            time.Time   `json:"create_time"`
	TimeZoneName          string      `json:"time_zone_name"`
	TimeZoneOffset        int         `json:"time_zone_offset"`
	PartnerId             int         `json:"partner_id"`
	Site                  string      `json:"site"`
	Site2                 Site2       `json:"site2"`
	GdprAgreementAccepted int         `json:"gdpr_agreement_accepted"`
}

type Webvisor struct {
	Urls           string `json:"urls"`
	ArchEnabled    int    `json:"arch_enabled"`
	ArchType       string `json:"arch_type"`
	LoadPlayerType string `json:"load_player_type"`
	WvVersion      int    `json:"wv_version"`
	AllowWv2       bool   `json:"allow_wv2"`
	WvForms        int    `json:"wv_forms"`
}

type Operation struct {
	Id     int    `json:"id"`
	Action string `json:"action"`
	Attr   string `json:"attr"`
	Value  string `json:"value"`
	Status string `json:"status"`
}

type Goal struct {
	Id            int         `json:"id"`
	Name          string      `json:"name"`
	Type          string      `json:"type"`
	DefaultPrice  float64     `json:"default_price"`
	IsRetargeting int         `json:"is_retargeting"`
	GoalSource    string      `json:"goal_source"`
	IsFavorite    int         `json:"is_favorite"`
	PrevGoalId    int         `json:"prev_goal_id"`
	Conditions    []Condition `json:"conditions"`
}

type Condition struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type CodeOptions struct {
	Async           int      `json:"async"`
	Informer        Informer `json:"informer"`
	Visor           int      `json:"visor"`
	TrackHash       int      `json:"track_hash"`
	XmlSite         int      `json:"xml_site"`
	Clickmap        int      `json:"clickmap"`
	InOneLine       int      `json:"in_one_line"`
	Ecommerce       int      `json:"ecommerce"`
	AlternativeCdn  int      `json:"alternative_cdn"`
	EcommerceObject string   `json:"ecommerce_object"`
}

type Informer struct {
	Enabled    int    `json:"enabled"`
	Type       string `json:"type"`
	Size       int    `json:"size"`
	Indicator  string `json:"indicator"`
	ColorStart string `json:"color_start"`
	ColorEnd   string `json:"color_end"`
	ColorText  int    `json:"color_text"`
	ColorArrow int    `json:"color_arrow"`
}

type Site2 struct {
	Site   string `json:"site"`
	Domain string `json:"domain"`
}
