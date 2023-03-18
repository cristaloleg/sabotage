# SABOTAGE

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Collection of dirty hacks in Go.

1. Repeat `sync.Once` once more.
2. Check if `sync.Mutex`/`sync.RWMutex` is locked.
3. `[]byte <-> string` conversion without additional allocation.

![logo](https://raw.githubusercontent.com/cristaloleg/sabotage/master/sabotage.jpg)

    Cause what you see, you might not get
    And we can bet, so don't you get souped yet
    Scheming on a thing, that's a mirage
    I'm trying to tell you now, it's sabotage

    Whhhhhyyyyyy??????

    (c) Beastie Boys - Sabotage

## Examples

1. You can repeat `sync.Once` as much as you want:

```go
var once sync.Once

once.Do(myFunc)

sabotage.ResetSyncOnce(&once) 
```

or check if it was invoked earlier:

```go
if sabotage.IsOnceDone(&once) {
    println("well-well-well")
}
```

2. You can unlock any mutex from anywhere:

```go
var mu sync.Mutex

// let's yolo begins
sabotage.UnlockMutex(&mu)
```

or check if it's locked:

```go
if sabotage.IsMutexLocked(&mu) {
    for {
        // let's wait a bit
    }
}
```

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristaloleg/sabotage/workflows/build/badge.svg
[build-url]: https://github.com/cristaloleg/sabotage/actions
[pkg-img]: https://pkg.go.dev/badge/cristaloleg/sabotage
[pkg-url]: https://pkg.go.dev/github.com/cristaloleg/sabotage
[reportcard-img]: https://goreportcard.com/badge/cristaloleg/sabotage
[reportcard-url]: https://goreportcard.com/report/cristaloleg/sabotage
[coverage-img]: https://codecov.io/gh/cristaloleg/sabotage/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristaloleg/sabotage
[version-img]: https://img.shields.io/github/v/release/cristaloleg/sabotage
[version-url]: https://github.com/cristaloleg/sabotage/releases
