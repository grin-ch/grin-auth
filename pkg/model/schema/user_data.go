package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/grin-ch/grin-auth/pkg/enum"
)

type UserData struct {
	ent.Schema
}

func (UserData) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Immutable().Comment("自增主键"),
		field.String("nickname").Unique().NotEmpty().Comment("昵称"),
		field.String("avatar_url").Default("").Comment("头像地址"),
		field.Time("birthday").Default(time.Now).Comment("生日"),
		field.Enum("sex").Values(enum.Sexs...).Default(enum.UnKnowSex).Comment("性别"),
		field.Int("level_exp").Default(0).Comment("等级经验"),
	}
}

func (UserData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_data").Unique().Required(),
	}
}
