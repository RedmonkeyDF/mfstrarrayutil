### MfStrArrayutils 1.0.0

##### String array slice utilities

Get the index of the string needle in array haystack

`func StringArrayIndex(haystack []string, needle string) int`

Does the hastack contain the string needle

`func StringArrayContains(haystack []string, needle string) bool`

Perform a function on each item in the slice, return the result 

`func StringArrayMap(haystack []string, f func(string) string) []string`

Return a slice which contains only items from haystack that passes the test func 

`func StringArrayFilter(haystack []string, f func(string) bool) []string`

Return a count of items in haystack that passes the test func 

`func StringArrayCount(haystack []string, f func(string) bool) int`

Does any item in the haystack pass the test func

`func StringArrayAny(haystack []string,  f func(string) bool) bool`

Do all the items in the haystack pass the test func 

`func StringArrayAll(haystack []string,  f func(string) bool) bool`