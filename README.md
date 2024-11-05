# godash

Lodash style go library

## Installation

```bash
go get github.com/changchanghwang/godash
```

## Usage

You can use with importing `godash``

```go
import "github.com/changchanghwang/godash"
```

### Quick LInks

For now, there is only supported Slice, Array mehotds

- [Slice, Array](#slice-array)
  - [Map](#map)
  - [Head](#head)
  - [Tail](#tail)
  - [DeepCopySlice](#deppcopyslice)

#### Slice, Array

##### Map

```go
import "github.com/changchanghwang/godash"

godash.Map([]int{1,2,3,4,5}, func(v int) int {
    return v * 2
}) // []int{2,4,6,8,10}
```

##### Head

```go
import "github.com/changchanghwang/godash"

godash.Head([]int{1,2,3,4,5}) // 1
```

##### Tail

```go
import "github.com/changchanghwang/godash"

godash.Tail([]int{1,2,3,4,5}) // 5
```

##### DeppCopySlice

```go
import "github.com/changchanghwang/godash"

arr := []int{1,2,3,4,5}
godash.Head(arr) // []int{1,2,3,4,5}
```
