package enums

type RouterPrefix string

const (
	AuthPrefix RouterPrefix = "auth"
)

func (prefix RouterPrefix) ToString() string {
	return string(prefix)
}
