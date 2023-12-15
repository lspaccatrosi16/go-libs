# Assorted Go Libraries

    
- [`algorithms`](#algorithms) - Various algorithms implemented in go
- [`gbin`](#gbin) - Binary serialisation for (almost) *every* go data structure 
- [`interpolator`](#interpolator) - String interpolation from maps of data
- [`structures`](#structures) - Data structures implemented in go

---

### `algorithms`

| Package     | Description                                                                |
| ----------- | -------------------------------------------------------------------------- |
| `maths`     | Core mathematical algorithms. [Docs](./algorithms//maths/README.md)        |
| `sequences` | Algorithms relating to sequences. [Docs](./algorithms/sequences/README.md) |
---
### `gbin`

Golang binary serialisation. Infers schema from go's type system.

> [!CAUTION]
> Data will not be able to be unserialised if the underlying type changes.

Encoding

```go
encoder := gbin.NewEncoder[T]()
encoded, err := encoder.Encode(&data)
```

Decoding
```go
decoder := gbin.NewDecoder[T]()
decoded, err := decoder.Decode(encoded)
```

Where `encoded` / `decoded` are of type `[]byte`. 

`EncodeStream` & `DecodeStream` can be used alternatively, which perform the same underlying function, but work with `io.Writer` and `io.Reader` respectively.

---
### `interpolator`

Interpolate a string from a map of values.

E.g.
```go
data := interpolator.Object{
    "a": "b",
    "c": interpolator.Object{
        "foo":1
        "bar":true
    }
}

input := "$a comes after a, but $c.foo doesn't come after $c.bar"

res,err := interpolator.ParseString(input, data)
```

Will produce a result of:

```b comes after a, but 1 doesn't come after true```

---
### `structures`

| Package | Description                              |
| ------- | ---------------------------------------- |
| `set`   | Implementation of an exclusive set in go |
| `stack` | Implementation of a stack in go          |
