package dependencies

type ScriptSystem interface {
	New(objType string) *ScriptObject
	Free(obj *ScriptObject)
}
