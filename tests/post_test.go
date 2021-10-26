package tests

type PostTest struct {
	Request float64 `json:"request"`
	Test    string  `json:"test"`
}

var PostTests = []PostTest{
	{Request: 1, Test: "This Is Test Request One"},
}
