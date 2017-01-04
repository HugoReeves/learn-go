#Maps

[Code](maps.go)

Like slices and arrays, maps are another way to store data with their own benefits and drawbacks.

* Creating maps - You can create maps using a similar syntax as creating a slice. `MAPNAME := make(map[KEYDATATYPE]VALUEDATATYPE)`.

	For example: `myMap := make(map[string]bool)` will create a map with strings as keys and bools as values.

* Adding entries to maps - You can add entries to maps with the syntax `MAPNAME[KEY] = VALUE`. The key and value must conform to the datatype set out when the map is initialized. Being able to store data in this key to value method is important to creating more complex programs where values need to be stored in relation to other values.

	For example: We can add a few key value pairs,
	```go
	myMap["workDone"] = true
	myMap["projectComplete"] = true
	myMap["gitPushed"] = false
	fmt.Println("Entered values:", myMap)
	```
	Will result in `map[workDone:true projectComplete:true gitPushed:false]` being printed.

* Retrieving values from maps - Values can be retrieved from maps with the following simple syntax, `MAPNAME[KEY]`.

	For example: `myMap["gitPushed"]` will be equal to false.

* Deleting keys and values - You can also delete key value pairings with the following syntax, `delete(MAPNAME, KEY)`.

	For example: `delete(myMap, "gitPushed")` will remove the value key pair of "gitPushed".

* Declaring and initializing in one statement - Similar to how you declare and initialize slices, you can declare and initialize maps. The syntax is `MAPNAME := map[KEYDATATYPE]VALUEDATATYPE{KEY1: VALUE1, KEY2: VALUE2}`

	For example: `newMap := map[string]int{"age": 21, "children": 0, "siblings": 3}` will create a new map and add the key value pairings stated.

###Next

* [Home](../../README.md)
* [Basics Overview](../basics.md)
* [Range](../range/range.md)
