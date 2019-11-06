package fixtureTest

import (
	"github.com/MikeSofaer/pylons/x/pylons/types"
)

type ItemCheck struct {
	StringKeys   []string                     `json:"stringKeys"`
	StringValues map[string]string            `json:"stringValues"`
	DblKeys      []string                     `json:"dblKeys"`
	DblValues    map[string]types.FloatString `json:"dblValues"`
	LongKeys     []string                     `json:"longKeys"`
	LongValues   map[string]int               `json:"longValues"`
}
type CoinCheck struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}
type TxResultCheck struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	ErrorLog string `json:"errLog"`
}
type UserPropertyCheck struct {
	Cookbooks []string    `json:"cookbooks"`
	Recipes   []string    `json:"recipes"`
	Items     []ItemCheck `json:"items"`
	Coins     []CoinCheck `json:"coins"`
}
type OutputCheck struct {
	TxResult TxResultCheck     `json:"txResult"`
	Property UserPropertyCheck `json:"property"`
}
type FixtureStep struct {
	Action               string      `json:"action"`
	BlockInterval        int64       `json:"blockInterval"`
	ParamsRef            string      `json:"paramsRef"`
	ParamsRefDescription string      `json:"paramsRefDescription"`
	Output               OutputCheck `json:"output"`
}
