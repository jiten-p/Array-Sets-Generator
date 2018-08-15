// arrsets generates sets out of an array
package arrsets

// ArraySetsGenerator is a set that generatss subsets
func ArraySetsGenerator(arr []interface{}) [][]interface{} {
	if arr == nil {
		return nil
	}
	if len(arr) == 1 {
		return [][]interface{}{{arr[0]}}
	}
	rArr := [][]interface{}{}
	j := 0
	for i := 0; i < len(arr)*(len(arr)-1); i++ {
		if j == len(arr)-1 {
			j = 0
		}
		t := arr[j+1]
		arr[j+1] = arr[j]
		arr[j] = t
		tmp := make([]interface{}, len(arr))
		copy(tmp, arr)
		rArr = append(rArr, tmp)
		j++
	}
	return rArr
}
