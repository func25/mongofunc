# Mongofunc
Just a wrapper to make life easier to interact with mongo query

### Quick start
You can walkaround to find out more commands, but here are some

You might using comparision queries like this:
```go
filter := moper.Init(
  moper.Equal("damage", 10),
  moper.EqualLess("health", 100),
  moper.Greater("speed", 20),
  moper.EqualGreaterTime("expired", time.Now())
)
```

Update commands:
```go
filter := moper.Init(
  moper.Equal("damage", 10),
)
update := moper.Init(
  moper.Set(moper.PairSetter{
    FieldName: "damage",
    Value:     -10,
  }),
  moper.IncInt(moper.PairSetterInt{
    FieldName: "health",
    Value:     3,
  }),
)
```

## Extra
You can compare time like this:
```go
filter := moper.Init(
  moper.EqualGreaterTime("expired", time.Now())
)
```

## Notes:
If you are using your own mongo connection library, then ignore mogorely package.
