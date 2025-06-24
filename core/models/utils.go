package models

import (
	"reflect"
)

func CompareStructs(oldVal, newVal interface{}) map[string][2]interface{} {
	changes := make(map[string][2]interface{})

	oldValRef := reflect.ValueOf(oldVal)
	newValRef := reflect.ValueOf(newVal)

	// Dereference pointer if needed
	if oldValRef.Kind() == reflect.Ptr {
		oldValRef = oldValRef.Elem()
	}
	if newValRef.Kind() == reflect.Ptr {
		newValRef = newValRef.Elem()
	}

	oldType := oldValRef.Type()

	for i := 0; i < oldValRef.NumField(); i++ {
		field := oldType.Field(i)

		// Skip unexported fields
		if !oldValRef.Field(i).CanInterface() {
			continue
		}

		oldField := oldValRef.Field(i).Interface()
		newField := newValRef.Field(i).Interface()

		if !reflect.DeepEqual(oldField, newField) {
			changes[field.Name] = [2]interface{}{oldField, newField}
		}
	}

	return changes
}

func compareFields(oldVal, newVal interface{}, include []string, exclude []string) map[string][2]interface{} {
	changes := make(map[string][2]interface{})

	oldRef := reflect.ValueOf(oldVal).Elem()
	newRef := reflect.ValueOf(newVal).Elem()
	typ := oldRef.Type()

	// Buat map untuk pencarian cepat
	includeMap := make(map[string]bool)
	excludeMap := make(map[string]bool)

	for _, f := range include {
		includeMap[f] = true
	}
	for _, f := range exclude {
		excludeMap[f] = true
	}

	for i := 0; i < oldRef.NumField(); i++ {
		field := typ.Field(i)
		fieldName := field.Name

		// ❌ Lewati jika tidak di-include (jika include list tidak kosong)
		if len(includeMap) > 0 && !includeMap[fieldName] {
			continue
		}
		// ❌ Lewati jika ada di exclude
		if excludeMap[fieldName] {
			continue
		}

		// Lewati field unexported
		if !oldRef.Field(i).CanInterface() {
			continue
		}

		oldField := oldRef.Field(i).Interface()
		newField := newRef.Field(i).Interface()

		if !reflect.DeepEqual(oldField, newField) {
			changes[fieldName] = [2]interface{}{oldField, newField}
		}
	}
	return changes
}
