package dependencies

type ScriptObject interface {
	Type() string

	Get(name string) string
	Set(name string, value string)
	Call(name string, parameters ...string)
}
