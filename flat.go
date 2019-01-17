package flat

import (
	"log"
	"reflect"
	"strconv"
)

type Options struct {
	Delimiter string
	Safe      bool
	Object    bool
	MaxDepth  int
}

func Flatten(nested map[string]interface{}, opts Options) (m map[string]interface{}, err error) {
	// construct default value
	if opts.Delimiter == "" {
		opts.Delimiter = "."
	}
	m, err = flatten("", nested, 0, opts)
	return
}

func flatten(prefix string, nested interface{}, depth int, opts Options) (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	log.Println("depth", depth)
	log.Println("prefix", prefix)

	switch nested := nested.(type) {
	case map[string]interface{}:
		log.Println("map", nested)
		// map
		for k, v := range nested {
			// create new key
			newKey := k
			if prefix != "" {
				newKey = prefix + opts.Delimiter + newKey
			}
			switch v.(type) {
			case map[string]interface{}:
				temp, fErr := flatten(newKey, v, depth+1, opts)
				if fErr != nil {
					err = fErr
					return
				}
				update(m, temp, newKey)
			case []interface{}:
				if opts.Safe == true {
					m[newKey] = v
					continue
				}
				temp, fErr := flatten(newKey, v, depth+1, opts)
				if fErr != nil {
					err = fErr
					return
				}
				update(m, temp, newKey)

			default:
				m[newKey] = v
			}
		}
	case []interface{}:
		log.Println("slice", nested)
		// slice
		for i, v := range nested {
			newKey := strconv.Itoa(i)
			if prefix != "" {
				newKey = prefix + opts.Delimiter + newKey
			}
			switch v.(type) {
			case map[string]interface{}, []interface{}:
				temp, fErr := flatten(newKey, v, depth+1, opts)
				if fErr != nil {
					err = fErr
					return
				}
				update(m, temp, newKey)
			default:
				m[newKey] = v
			}
		}
	default:
		log.Println("error")
	}
	return
}

func f(prefix string, nested interface{}, opts Options) (flatmap map[string]interface{}, err error) {
	flatmap = make(map[string]interface{})

	switch nested := nested.(type) {
	case map[string]interface{}:
		for k, v := range nested {
			// create new key
			newKey := k
			if prefix != "" {
				newKey = prefix + opts.Delimiter + newKey
			}
			fm1, fe := f(newKey, v, opts)
			if fe != nil {
				err = fe
				return
			}
			update(flatmap, fm1, newKey)
		}
	case []interface{}:
		for i, v := range nested {
			newKey := strconv.Itoa(i)
			if prefix != "" {
				newKey = prefix + opts.Delimiter + newKey
			}
			fm1, fe := f(newKey, v, opts)
			if fe != nil {
				err = fe
				return
			}
			update(flatmap, fm1, newKey)
		}
	default:
		flatmap[prefix] = nested
	}
	return
}

// merge is the function that update to map with from and key
// example: from is a map
// to = {"hi": "there"}
// from = {"foo": "bar"}
// new to = {"hi": "there", "foo": "bar"}
// example: from is an empty map
// to = {"hi": "there"}
// from = {}
// key = "world"
// key = {"hi": "there", "world": {}}
func update(to map[string]interface{}, from map[string]interface{}, key string) {

	if reflect.DeepEqual(from, map[string]interface{}{}) {
		to[key] = from
		return
	}
	// merge temp to m
	for kt, vt := range from {
		to[kt] = vt
	}
}
