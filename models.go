package godynamo

type Model struct {
	Id      string `json:"id" godynamo:"S,hash"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated"`
	Deleted int64  `json:"deleted"`
}

type aaa struct {
	Model

	Aa string            `json:"aaa"`
	Ab string            `json:"aab"`
	Ac []bbb             `json:"aac"`
	Ad []string          `json:"aad"`
	Ae map[string]bbb    `json:"aae"`
	Af map[string]string `json:"aaf"`
}

type bbb struct {
	Model

	Ba string   `json:"bba"`
	Bb []ddd    `json:"bbb"`
	Bc []string `json:"bbc"`

	CId string `json:"cId"`

	Bd int64 `json:"bbd" godynamo:"N,range"`
}

type ccc struct {
	Model

	Ca string `json:"cca"`

	DId string `json:"dId"`
}

type ddd struct {
	Model

	Da string `json:"dda"`
	Db string `json:"ddb"`
}
