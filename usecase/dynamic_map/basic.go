package dynamic_map

func (d *DynamicMap) Set(key string, value interface{}) {
	(*d)[key] = value
}

func (d *DynamicMap) Get(key string) interface{} {
	return (*d)[key]
}

func (d *DynamicMap) Del(key string) {
	delete(*d, key)

}

func (d *DynamicMap) Keys() []string {
	keys := []string{}
	for k, _ := range *d {
		keys = append(keys, k)

	}
	return keys
}

func (d *DynamicMap) Values() []interface{} {
	values := []interface{}{}
	for k, _ := range *d {
		values = append(values, k)

	}
	return values
}

func (d *DynamicMap) Clear() {
	*d = DynamicMap{}

}
