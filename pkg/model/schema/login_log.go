package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type LoginLog struct {
	ent.Schema
}

func (LoginLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Immutable().Comment("自增主键"),
		field.String("ip_addr").Immutable().Comment("登入时的ip"),
		field.Time("login_time").Default(time.Now).Immutable().Comment("登入时间"),
	}
}

func (LoginLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("login_logs").Unique(),
	}
}
