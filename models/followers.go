package models

type Follower struct {
	Title          string           `json:"title"`
	MediaListData  []MediaListData  `json:"media_list_data"`
	StringListData []StringListData `json:"string_list_data"`
}
type StringListData struct {
	Href      string `json:"href"`
	Value     string `json:"value"`
	Timestamp int64  `json:"timestamp"`
}

type MediaListData struct {
}

func (f Follower) Equals(o Follower) bool {
	return f.StringListData[0].Value == o.StringListData[0].Value
}

func (f Follower) GetValue() string {
	if f.StringListData == nil || len(f.StringListData) < 1 {
		return ""
	}
	return f.StringListData[0].Value
}
