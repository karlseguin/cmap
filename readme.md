# Concurrent Map

Wraps a map[string]interface{} inside some locks

```go
m := cmap.New()

m.Set("power", 9000)
value, ok := m.Get("power")
m.Delete("power")
m.Len()
```
