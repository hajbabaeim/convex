package dynamic_map

type DynamicMap map[string]interface{}

func New() DynamicMap {
	return make(DynamicMap)
}
