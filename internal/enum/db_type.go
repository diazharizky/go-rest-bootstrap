package enum

type DbType string

const (
	DbTypePostgres DbType = "postgres"
)

func (t DbType) Validate() bool {
	switch t {
	case DbTypePostgres:
		return true
	}

	return false
}

func (t DbType) String() string {
	return string(t)
}
