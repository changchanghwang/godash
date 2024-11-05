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
  - [Filter](#filter)

#### Slice, Array

##### Map

```go
import "github.com/changchanghwang/godash"

// without chain
godash.Map([]int{1,2,3,4,5}, func(v int) int {
    return v * 2
}) // []int{2,4,6,8,10}

// with chain
godash.Chain([]int{1,2,3,4,5}).Map(func(v int) int {
    return v * 2
}).Value() // []int{2,4,6,8,10}
```

##### Head

```go
import "github.com/changchanghwang/godash"

// without chain
godash.Head([]int{1,2,3,4,5}) // 1, true
godash.Head([]int{}) // 0, false

// with chain
godash.Chain([]int{1,2,3,4,5}).Head() // 1, true
godash.Chain([]int{}).Head() // 0, false
```

##### Tail

```go
import "github.com/changchanghwang/godash"

// with chain
godash.Tail([]int{1,2,3,4,5}) // 5, true
godash.Tail([]int{}) // 0, false

// with chain
godash.Chain([]int{1,2,3,4,5}).Tail() // 5, true
godash.Chain([]int{}).Tail() // 0, false
```

##### DeepCopySlice

```go
import "github.com/changchanghwang/godash"

arr := []int{1,2,3,4,5}
godash.DeepCopySlice(arr) // []int{1,2,3,4,5}
```

##### Filter

```go
import "github.com/changchanghwang/godash"

// without chain
arr := []int{1,2,3,4,5}
godash.Filter(arr, func(v int) bool {
  return v > 3
}) // []int{4,5}

// with chain
arr := []int{1,2,3,4,5}
godash.Chain(arr).Filter(func(v int) bool {
  return v > 3
}).Value() // []int{4,5}
```
