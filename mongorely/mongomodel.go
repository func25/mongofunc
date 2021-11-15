package mongorely

import "fmt"

type MongoModel interface {
	GetMongoCollName() string
}

type DbConfig struct {
	DbName   string
	UserName string
	Password string
	Host     string
	Port     string
	// IsReplica  bool
	// ReplicaSet string
}

type ObjectId struct {
	ID interface{} `json:"id" bson:"_id,omitempty"`
}

func (cfg *DbConfig) ToConnectionString() string {
	link := fmt.Sprintf("%s:%s/?w=majority", cfg.Host, cfg.Port)

	// if cfg.IsReplica {
	// 	link = fmt.Sprintf("%s", cfg.ReplicaSet)
	// }

	var uri string
	if cfg.UserName == "" && cfg.Password == "" {
		uri = fmt.Sprintf("mongodb://%s", link)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s@%s", cfg.UserName, cfg.Password, link)
	}

	return uri
}
