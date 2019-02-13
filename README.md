# SABOTAGE

![logo](https://raw.githubusercontent.com/cristaloleg/sabotage/master/sabotage.jpg)

    Cause what you see, you might not get
    And we can bet, so don't you get souped yet
    Scheming on a thing, that's a mirage
    I'm trying to tell you now, it's sabotage

    Whhhhhyyyyyy??????

    (c) Beastie Boys - Sabotage

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