# Settings

This is a simple example of using generics in Go.

## Example

```go
type Setting = settings.Setting[Serialize]

func WithValue(value string) Setting {
	return func(o *Serialize) {
		o.Value = value
	}
}

type Serialize struct {
	Value    string
}

var (
	defaultSerialize = Serialize{
		Value:    "hello",
}

func NewSerialize(ts ...Setting) *Serialize {
	serialize := settings.Apply(&defaultSerialize, ts)
	return serialize
}

func main() {
	serialize = NewSerialize()
	fmt.Println(serialize.Value)
	// Output: hello
	serialize := NewSerialize(WithValue("world"))
	fmt.Println(serialize.Value)
	// Output: world
}
```