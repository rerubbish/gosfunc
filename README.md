# gosfunc

在golang实现类似 kotlin 的 Scope functions 风格。

# how to use

```shell
go get -u github.com/rerubbish/gosfunc
```

```golang
type Person struct {
	Name string
	Age  uint
}
var result = gosfunc.Let([]Person{
	{
		Name: "233",
		Age:  5,
	},
	{
		Name: "wdf",
		Age:  10,
	},
}, func(it []Person) Person {
	for _, item := range it {
		if item.Age > 5 {
			return item
		}
	}
	return Person{}
})
```
