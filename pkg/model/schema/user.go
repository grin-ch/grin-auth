package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Immutable().Comment("自增主键"),
		field.String("email").Unique().Comment("邮箱"),
		field.String("phone").Unique().Comment("手机号"),
		field.String("password").NotEmpty().Comment("密码"),
		field.Time("reg_time").Default(time.Now).Immutable().Comment("注册时间"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_data", UserData.Type).Unique(),
		edge.To("login_logs", LoginLog.Type),
	}
}
