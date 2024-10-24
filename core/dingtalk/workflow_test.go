package dingtalk

import (
	"testing"

	"github.com/xops-infra/jms/app"
	"github.com/xops-infra/jms/model"
)

func init() {
	model.LoadYaml("/opt/jms/config.yaml")
	app.NewApp(true, "", "").WithDB(false).WithDingTalk()
}

func TestLoadDingtalkUsers(t *testing.T) {
	err := LoadUsers()
	if err != nil {
		t.Error(err)
	}
}
