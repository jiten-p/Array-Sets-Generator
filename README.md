### Generating Array Sets from given array

Generating sets of arrays from a given array is always a daunting task.This simple GoLang package, takes a given array and generate all random available sets from it.
Sets are unique and are generated from the same array.

```
    arr := []interface{}{"A", "B", "C"}
	rArr := ArraySetGenerator(arr)
	fmt.Println(rArr)

Output:
[[B A C] [B C A] [C B A] [C A B] [A C B] [A B C]]
```

```
    arr := []interface{}{1,2,3}
	rArr := ArraySetGenerator(arr)
	fmt.Println(rArr)

Output:
[[2 1 3] [2 3 1] [3 2 1] [3 1 2] [1 3 2] [1 2 3]]
```