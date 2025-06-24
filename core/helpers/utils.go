package helpers

import (
	"crypto/rand"
	"him/fiber-api/core/structs"
	"math/big"
	"reflect"
	"slices"
)

func ConvertToDataResponse(data *[]map[string]interface{}, exclude *[]string) (*structs.DataResponse, error) {
	resp := structs.DataResponse{
		Total: 0,
	}

	if len(*data) > 0 {
		tmp := *data
		f := tmp[0]
		// fmt.Println(f)

		resp.Schema = make(map[string]int, len(f))
		idx := 0
		for k := range f {
			if (len(*exclude) > 0 && !slices.Contains(*exclude, k)) || len(*exclude) == 0 {
				resp.Schema[k] = idx
				idx++
			}
		}
		var tmp2 map[int]any
		resp.Data = make([]map[int]any, 0, len(*data))
		for _, s := range *data {
			tmp2 = make(map[int]any, len(f))
			for k, v := range s {
				if (len(*exclude) > 0 && !slices.Contains(*exclude, k)) || len(*exclude) == 0 {
					tmp2[resp.Schema[k]] = v
					idx++
				}
			}
			resp.Data = append(resp.Data, tmp2)
			resp.Total++
		}
	}

	return &resp, nil
}

func generateSecret(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789" +
		"!@#$%^&*()-_=+[]{}|;:',.<>?/`~"
	secret := make([]byte, length)
	for i := 0; i < length; i++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		secret[i] = charset[nBig.Int64()]
	}
	return string(secret), nil
}

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

func compareFields(oldVal, newVal interface{}, fields []string) map[string][2]interface{} {
	changes := make(map[string][2]interface{})

	oldRef := reflect.ValueOf(oldVal).Elem()
	newRef := reflect.ValueOf(newVal).Elem()
	typ := oldRef.Type()

	fieldMap := make(map[string]bool)
	for _, f := range fields {
		fieldMap[f] = true
	}

	for i := 0; i < oldRef.NumField(); i++ {
		field := typ.Field(i)
		if !fieldMap[field.Name] {
			continue // hanya bandingkan field whitelist
		}
		if !oldRef.Field(i).CanInterface() {
			continue
		}

		oldField := oldRef.Field(i).Interface()
		newField := newRef.Field(i).Interface()

		if !reflect.DeepEqual(oldField, newField) {
			changes[field.Name] = [2]interface{}{oldField, newField}
		}
	}
	return changes
}
