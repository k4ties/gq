# gq
Go map utilities

Thanks, AI for commenting the code.

Usage:

gq.Map
```go
v := make(gq.Map[string, any])

v.Set("test", 1)
v.Set("test 2", 2)

v.Iterate(func(key string, value any) bool {
	fmt.Printf("key: %s, value: %v\n", key, value)
	return true
})

v.Clear()
```

gq.Set
```go
v := make(gq.Set[string]) // type must be comparable

v.Add("item 1")
v.Add("item 2")

v.Iterate(f func(value string) bool {
	fmt.Printf("value: %v\n", value)
	return true
})
```