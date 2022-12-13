# Mongofunc
Mongofunc is just a wrapper (mostly for personal project) to make life easier to interact with mongo query, transaction,...

## Content:

- [Installation](#installation)
- [Quick start](#quick-start)
  - [Moper package](#moper-package)
    - [Moper queries](#you-might-using-comparision-queries-like-this)
    - [Moper update](#update-commands)
  - [Mocom package](#mocom-package)
    - [Connect with mocom](#connect-with-mocom)
    - [Mocom commands](#mocom-commands)
    - [Mocom transaction](#mocom-transaction)

## Installation

```sh
$ go get -u github.com/func25/mongofunc
```

## Quick start
You can walkaround to find out more commands, but here are some

### moper package

#### You might using comparision queries like this:
```go
filter := moper.NewD().
	Equal("damage", 10).
	EqualLess("health", 100).
	Greater("speed", 20.1)
```

You can compare time.Time like this:
```go
filter := moper.Init(
  moper.EqualGreaterTime("expired", time.Now())
)
```

#### Update commands
```go
update := moper.NewD().Set(
	moper.P{"damage", 10},
	moper.P{"health", 1},
).Inc(
	moper.P{"speed", 1},
)
```

Support simple aggregation:
```go
intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

matchStage := moper.NewD().MatchD(moper.NewD().InArray("damage", intArr))
groupStage := moper.NewD().Group(
	moper.P{K: "_id", V: nil},
	moper.P{K: "total", V: moper.NewD().Sum("damage")},
)

req := &mocom.AggregationRequest[Hero]{
	Pipeline: []moper.D{matchStage, groupStage},
	Options:  []*options.AggregateOptions{},
}
```

### mocom package
Unlike moper, mocom is designed to interact directly with mongodb; 

#### Connect with mocom
But before using it, you have to connect using mocom.Connect or if you're using your own library to manage your connection, then use mocom.Setup

```go
// using connect
client, err := mocom.Connect(ctx, "mongodb://localhost:27017/?w=majority&retryWrites=false", "defaultdb")
if err != nil {
  log.Fatal("cannot connect mongo", err)
}

// using setup if you already have a client object
db = client.Database("DbName")
mocom.Setup(db)
```

#### mocom commands

To use mocom command, your models should meet the mocom.Model interface{}
```go
type Model interface {
	CollName() string
}

// Hero is Model
filter := moper.NewD().Equal("damage", -i)
update := moper.NewD().Set(moper.P{K: "damage", V: i})

result, err := mocom.UpdateMany[Hero](ctx, filter, update)
```

### mocom transaction

Beside the CRUD, mocom also supports transaction for easier and cleaner coding style:

```go
_, err := mocom.TxOptimal(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
    // .... execute your commands with ctx context here, it's extremely important to use ctx of this function.
})

if err != nil {
  t.Error(err)
}
```

TxOptimal is used with snapshot read-concern, write majority write-concern and primary read-preference (of course)