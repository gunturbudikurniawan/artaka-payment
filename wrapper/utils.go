package wrapper

import "fmt"

func ParamToString(queries map[string][]string) (bool, string) {
	base := "?"
	i := 0
	for k, v := range queries {
		var concat = ""
		if i < len(queries)-1 {
			concat = "&"
		}
		base += fmt.Sprintf("%s=%s%s", k, v[0], concat)
	}
	isQuery := len(queries) > 0
	return isQuery, base
}
