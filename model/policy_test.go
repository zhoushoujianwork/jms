package model_test

import (
	"testing"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
	"github.com/xops-infra/multi-cloud-sdk/pkg/model"
	"github.com/xops-infra/noop/log"

	. "github.com/xops-infra/jms/model"
)

func init() {
	// config.LoadYaml("/opt/jms/config.yaml")
	log.Default().WithLevel(log.DebugLevel).WithFilename("/tmp/test.log").Init()
}

func TestMatchServer(t *testing.T) {
	server := Server{
		Name: "test-server",
		Host: "127.0.0.1",
		Tags: model.Tags{
			{
				Key:   "EnvType",
				Value: "prod",
			}, {
				Key:   "Team",
				Value: "ops",
			},
		},
	}
	filter := ServerFilterV1{
		EnvType: []string{"prod"},
	}
	assert.True(t, MatchServerByFilter(filter, server))

	filter.Name = []string{"!test*"}
	assert.False(t, MatchServerByFilter(filter, server))

	filter.Name = []string{"test-server"}
	filter.IpAddr = []string{"!127.0.1.*"}
	assert.True(t, MatchServerByFilter(filter, server))

	filter.IpAddr = []string{"!127.0.0.*"}
	assert.False(t, MatchServerByFilter(filter, server))

	filter.IpAddr = []string{"!127.0.1.*"}
	filter.Team = []string{"ops"}
	assert.True(t, MatchServerByFilter(filter, server))

	filter.Team = []string{"!ops"}
	assert.False(t, MatchServerByFilter(filter, server))

	// 过滤条件有一个满足就满足。
	filter.Team = []string{"!ops", "*"}
	assert.True(t, MatchServerByFilter(filter, server))
}

// TEST MatchPolicy
func TestMatchPolicy(t *testing.T) {
	Conf.WithDB.Enable = true

	user := User{
		Username: tea.String("zhoushoujian"),
		Groups:   ArrayString{},
	}
	inPutAction := Connect
	server := Server{}
	policy := Policy{
		Name:      "test policy",
		IsEnabled: true,
		Users:     ArrayString{"zhoushoujian"},
		Actions:   ArrayString{"connect"},
		ExpiresAt: time.Now().Add(ExpireTimes[OneWeek]),
	}

	user.Groups = ArrayString{"admin"}
	{
		// 测试 admin 组
		assert.True(t, MatchPolicy(user, inPutAction, server, []Policy{
			policy,
		}))
	}

	user.Groups = ArrayString{}
	{
		// 测试普通用户,IP 匹配
		policy.ServerFilterV1.IpAddr = []string{"127.0.0.1"}
		policy.ServerFilterV1.Name = nil
		policy.ServerFilterV1.EnvType = nil
		policy.ServerFilterV1.Team = nil
		server.Host = "127.0.0.1"
		assert.True(t, MatchPolicy(user, inPutAction, server, []Policy{
			policy,
		}))
		server.Host = "89.0.142.86"
		assert.False(t, MatchPolicy(user, inPutAction, server, []Policy{
			policy,
		}))
	}
	{
		// 普通用户，Name匹配
		policy.ServerFilterV1.Name = []string{"test"}
		policy.ServerFilterV1.EnvType = nil
		policy.ServerFilterV1.Team = nil
		policy.ServerFilterV1.IpAddr = nil
		server.Name = "test"
		assert.True(t, MatchPolicy(user, inPutAction, server, []Policy{
			policy,
		}))
		server.Name = "test2"
		assert.False(t, MatchPolicy(user, inPutAction, server, []Policy{
			policy,
		}))
	}
	{
		// 普通用户，EnvType匹配
		policy.ServerFilterV1.Team = nil
		policy.ServerFilterV1.Name = nil
		policy.ServerFilterV1.IpAddr = nil
		policy.ServerFilterV1.EnvType = []string{"prod"}
		server.Tags = model.Tags{
			{
				Key:   "EnvType",
				Value: "prod",
			},
		}
		assert.True(t, MatchPolicy(user, inPutAction, server, []Policy{
			policy,
		}))
		server.Tags = model.Tags{
			{
				Key:   "EnvType",
				Value: "dev",
			},
		}
		assert.False(t, MatchPolicy(user, inPutAction, server, []Policy{
			policy,
		}))
	}
	{
		// 普通用户，Team匹配
		policy.ServerFilterV1.Team = []string{"ops"}
		policy.ServerFilterV1.Name = nil
		policy.ServerFilterV1.EnvType = nil
		policy.ServerFilterV1.IpAddr = nil

		server.Tags = model.Tags{
			{
				Key:   "Team",
				Value: "ops",
			},
		}
		assert.True(t, MatchPolicy(user, inPutAction, server, []Policy{
			policy,
		}))
		server.Tags = model.Tags{
			{
				Key:   "Team",
				Value: "others",
			},
		}
		assert.False(t, MatchPolicy(user, inPutAction, server, []Policy{
			policy,
		}))
	}
	{
		// 普通用户，Owner匹配
		policy.ServerFilterV1.Team = nil
		policy.ServerFilterV1.Name = nil
		policy.ServerFilterV1.EnvType = nil
		policy.ServerFilterV1.IpAddr = nil

		server.Tags = model.Tags{
			{
				Key:   "Owner",
				Value: "zhoushoujian",
			},
		}
		assert.True(t, MatchPolicy(user, inPutAction, server, []Policy{
			policy,
		}))
		server.Tags = model.Tags{
			{
				Key:   "Owner",
				Value: "xxx",
			},
		}
		assert.False(t, MatchPolicy(user, inPutAction, server, []Policy{
			policy,
		}))
	}

}

func TestMultipolicy(t *testing.T) {
	Conf.WithDB.Enable = true

	user := User{
		Username: tea.String("zhoushoujian"),
		Groups:   ArrayString{},
	}
	inPutAction := Connect

	defaultPolicy := Policy{
		IsEnabled: true,
		Users:     ArrayString{"zhoushoujian"},
		Actions:   ArrayString{"connect"},
		ExpiresAt: time.Now().Add(ExpireTimes[OneWeek]),
		ServerFilterV1: &ServerFilterV1{
			IpAddr: []string{"127.0.0.1"},
		},
	}
	{
		server := Server{
			Host: "127.0.0.1",
		}
		// 测试 deny 匹配
		assert.True(t, MatchPolicy(user, inPutAction, server, []Policy{
			defaultPolicy,
		}))
		assert.False(t, MatchPolicy(user, inPutAction, server, []Policy{
			defaultPolicy,
			{
				IsEnabled: true,
				Users:     ArrayString{"zhoushoujian"},
				Actions:   ArrayString{string(DenyConnect)},
				ExpiresAt: time.Now().Add(ExpireTimes[OneWeek]),
				ServerFilterV1: &ServerFilterV1{
					Name: []string{"*"},
				},
			},
		}))

		// 测试 ! 匹配
		server.Tags = model.Tags{
			{
				Key:   "EnvType",
				Value: "prod",
			},
		}
		assert.False(t, MatchPolicy(user, inPutAction, server, []Policy{
			{
				IsEnabled: true,
				Users:     ArrayString{"zhoushoujian"},
				Actions:   ArrayString{string((Connect))},
				ExpiresAt: time.Now().Add(ExpireTimes[OneWeek]),
				ServerFilterV1: &ServerFilterV1{
					EnvType: []string{"!prod"},
				},
			},
		}))
		assert.True(t, MatchPolicy(user, inPutAction, server, []Policy{
			{
				IsEnabled: true,
				Users:     ArrayString{"zhoushoujian"},
				Actions:   ArrayString{string((Connect))},
				ExpiresAt: time.Now().Add(ExpireTimes[OneWeek]),
				ServerFilterV1: &ServerFilterV1{
					EnvType: []string{"!dev"},
				},
			},
		}))

		// 测试 * 匹配
		server.Tags = model.Tags{
			{
				Key:   "Team",
				Value: "ops",
			},
		}
		assert.True(t, MatchPolicy(user, inPutAction, server, []Policy{
			{
				IsEnabled: true,
				Users:     ArrayString{"zhoushoujian"},
				Actions:   ArrayString{string((Connect))},
				ExpiresAt: time.Now().Add(ExpireTimes[OneWeek]),
				ServerFilterV1: &ServerFilterV1{
					Team: []string{"*"},
				},
			},
		}))
		assert.False(t, MatchPolicy(user, inPutAction, server, []Policy{
			{
				IsEnabled: true,
				Users:     ArrayString{"zhoushoujian"},
				Actions:   ArrayString{string((Connect))},
				ExpiresAt: time.Now().Add(ExpireTimes[OneWeek]),
				ServerFilterV1: &ServerFilterV1{
					Team: []string{"data"},
				},
			},
		}))
	}
}
