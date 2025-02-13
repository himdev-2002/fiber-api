package helpers

import (
	"slices"
	"tde/fiber-api/core/structs"
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
