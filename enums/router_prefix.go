package enums

type RouterPrefix string

const (
	AuthPrefix    RouterPrefix = "auth"
	DepositPrefix RouterPrefix = "deposits"
)

func (prefix RouterPrefix) ToString() string {
	return string(prefix)
}
