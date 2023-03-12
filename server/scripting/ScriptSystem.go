package scripting

type ScriptSystem struct {
	objects map[string]*ScriptObject
}

func (s *ScriptSystem) Create(typeName string, id string) *ScriptObject {
	return nil
}

func (s *ScriptSystem) Find(id string) *ScriptObject {
	if obj, ok := s.objects[id]; ok {
		return obj
	} else {
		return nil
	}
}

func (s *ScriptSystem) Release(id string) {
	if o, ok := s.objects[id]; ok {
		o.Close()
		delete(s.objects, id)
	}
}
