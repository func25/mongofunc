# Mongofunc
Just a wrapper to make life easier to interact with mongo query

### Quick start
You can walkaround to find out more commands, but here are some

You might using comparision queries like this:
```go
filter := mongoquery.Init(
  mongoquery.Equal("damage", 10),
  mongoquery.EqualLess("health", 100),
  mongoquery.Greater("speed", 20),
  mongoquery.EqualGreaterTime("expired", time.Now())
)
```

Update commands:
```go
filter := mongoquery.Init(
  mongoquery.Equal("damage", 10),
)
update := mongoquery.Init(
  mongoquery.Set(mongoquery.PairSetter{
    FieldName: "damage",
    Value:     -10,
  }),
  mongoquery.IncInt(mongoquery.PairSetterInt{
    FieldName: "health",
    Value:     3,
  }),
)
```

## Extra
You can compare time like this:
```go
filter := mongoquery.Init(
  mongoquery.EqualGreaterTime("expired", time.Now())
)
```

## Notes:
If you are using your own mongo connection library, then ignore mogorely package.
