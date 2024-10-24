package pui

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/elfgzp/ssh"
	"github.com/manifoldco/promptui"
	"github.com/xops-infra/noop/log"

	"github.com/xops-infra/jms/app"
	"github.com/xops-infra/jms/core/sshd"
	. "github.com/xops-infra/jms/model"
)

// PUI pui
type PUI struct {
	sess       *ssh.Session
	timeOut    time.Duration
	lastActive time.Time
}

func NewPui(s *ssh.Session, timeout time.Duration) *PUI {
	return &PUI{
		sess:       s,
		timeOut:    timeout,
		lastActive: time.Now(),
	}
}

func (ui *PUI) SessionWrite(msg string) error {
	_, err := (*ui.sess).Write([]byte(msg))
	return err
}

// exit
func (ui *PUI) Exit() {
	ui.SessionWrite(fmt.Sprintf(BybLabel, time.Now().Local().Format("2006-01-02 15:04:05")))
	err := (*ui.sess).Close()
	if err == nil {
		log.Infof("User %s form %s exit", (*ui.sess).User(), (*ui.sess).RemoteAddr().String())
	}
}

func (ui *PUI) IsTimeout() bool {
	_, found := app.App.Cache.Get((*ui.sess).RemoteAddr().String())
	if found {
		ui.FlashTimeout()
		return false
	}
	return time.Since(ui.lastActive) > ui.timeOut
}

// getTimeout
func (ui *PUI) GetTimeout() string {
	return fmt.Sprint(ui.timeOut)
}

// FlashTimeout flash timeout
func (ui *PUI) FlashTimeout() {
	ui.lastActive = time.Now()
}

// ShowMenu show menu
func (ui *PUI) ShowMenu(label string, menu []*MenuItem, BackOptionLabel string, selectedChain []*MenuItem) {
	user := User{
		Username: tea.String((*ui.sess).User()),
	}
	var broadcast *Broadcast

	if app.App.Config.WithDB.Enable {
		_user, err := app.App.JmsDBService.DescribeUser((*ui.sess).User())
		if err != nil {
			log.Errorf("DescribeUser error: %s", err)
			sshd.ErrorInfo(err, ui.sess)
		}
		user = _user
		_broadcast, err := app.App.JmsDBService.GetBroadcast()
		if err != nil {
			log.Errorf("GetBroadcast error: %s", err)
		} else {
			broadcast = _broadcast
		}
	}

loopMenu:
	for {
		menuLabels := make([]string, 0) // 菜单，用于显示
		menuItems := make([]*MenuItem, 0)
		if menu == nil {
			log.Debugf("menu is nil, label: %s", label)
			// break
		}
		// 返回顶级菜单
		log.Debugf("label: %s MainLabel:%s", label, MainLabel)
		switch label {
		case MainLabel:
			// 顶级菜单，如果有审批则主页支持选择审批或者服务器
			menu = make([]*MenuItem, 0)

			if app.App.Config.WithDB.Enable && !app.App.Config.WithDingtalk.Enable {
				// 没有审批策略时候，会在 admin 服务器选择列表里面显示审批菜单
				policies, err := app.App.JmsDBService.NeedApprove((*ui.sess).User())
				if err != nil {
					log.Errorf("Get need approve policy for admin error: %s", err)
				}
				if len(policies) > 0 {
					sshd.Info(fmt.Sprintf("作为管理员，有新的审批工单（%d）待处理。", len(policies)), ui.sess)
					menu = append(menu, GetApproveMenu(policies)...)
				}
			}
			_menus, err := GetServersMenuV2(ui.sess, user, ui.GetTimeout())
			if err != nil {
				sshd.ErrorInfo(err, ui.sess)
				break loopMenu
			}
			menu = append(menu, _menus...)

			filter, err := ui.inputFilter(broadcast)
			if err != nil {
				break loopMenu
			}
			if filter == "^C" {

				continue
			}
			for index, menuItem := range menu {
				if menuItem.IsShow == nil || menuItem.IsShow(index, menuItem, ui.sess, selectedChain) {
					if !strings.Contains(menuItem.Label, filter) {
						continue
					}
					menuLabels = append(menuLabels, menuItem.Label)
					menuItems = append(menuItems, menuItem)
				}
			}
		default:
			for index, menuItem := range menu {
				if menuItem.IsShow == nil || menuItem.IsShow(index, menuItem, ui.sess, selectedChain) {
					log.Debugf("index: %d label: %s", index, menuItem.Label)
					menuLabels = append(menuLabels, menuItem.Label)
					menuItems = append(menuItems, menuItem)
				}
			}
		}

		log.Debugf("menuLabels: %s", tea.Prettify(menuLabels))
		if len(menuLabels) == 0 {
			continue
		}
		menuLabels = append(menuLabels, BackOptionLabel) // 添加返回选项
		backIndex := len(menuLabels) - 1                 // 返回选项的索引
		menuPui := promptui.Select{
			Label:  label,
			Size:   15, // 菜单栏服务器数最大为15
			Items:  menuLabels,
			Stdin:  *ui.sess,
			Stdout: *ui.sess,
		}

		// get sub menu label
		index, subMenuLabel, err := menuPui.Run()
		if err != nil {
			// ^C ^D is not error
			if strings.Contains(err.Error(), "^C") {
				log.Debugf(label, MainLabel)
				// 返回主菜单
				ui.ShowMenu(MainLabel, menu, BackOptionLabel, selectedChain)

			} else if strings.Contains(err.Error(), "^D") {
				app.App.Cache.Delete((*ui.sess).User())
				ui.Exit()
				break
			}
			log.Errorf("Select menu error %s\n", err)
		}
		if index == backIndex {
			// 返回上一级菜单，如果主菜单了则退无可退。
			if label == MainLabel {
				log.Debugf("main menu, no back")
				continue
			} else {
				break
			}
		}

		// get sub menu
		selected := menuItems[index]
		if selected.GetSubMenu != nil {
			getSubMenu := selected.GetSubMenu
			subMenu := getSubMenu(index, selected, ui.sess, selectedChain)

			if len(subMenu) > 0 {
				back := "back"
				if selected.BackOptionLabel != "" {
					back = selected.BackOptionLabel
				}
				if selected.SubMenuTitle != "" {
					subMenuLabel = selected.SubMenuTitle
				}
				ui.ShowMenu(subMenuLabel, subMenu, back, append(selectedChain, selected))
			} else {
				noSubMenuInfo := "No options under this menu ..."
				if selected.NoSubMenuInfo != "" {
					noSubMenuInfo = selected.NoSubMenuInfo
				}
				sshd.ErrorInfo(errors.New(noSubMenuInfo), ui.sess)
			}
		}

		// run selected func
		if selected.SelectedFunc != nil {
			selectedFunc := selected.SelectedFunc
			isTop, err := selectedFunc(index, selected, ui.sess, selectedChain)
			if err != nil {
				sshd.ErrorInfo(err, ui.sess)
			}
			if isTop {
				label = MainLabel
			}
		}
	}
	log.Debugf("pui exit")
}

// inputFilter input filter
func (ui *PUI) inputFilter(broadcast *Broadcast) (string, error) {
	ui.FlashTimeout()
	defer ui.SessionWrite("\033c") // clear
	servers := app.GetServers()
	servers.SortByName()
	// 发送屏幕清理指令
	// 发送当前时间
	ui.SessionWrite(fmt.Sprintf("Last connect time: %s\t OnLineUser: %d\t AllServerCount: %d\n",
		time.Now().Format("2006-01-02 15:04:05"), app.App.Cache.ItemCount()-1, len(servers),
	))
	// 发送欢迎信息
	if broadcast != nil {
		ui.SessionWrite(fmt.Sprintf(InfoLabel, app.App.Version, "\n"+broadcast.Message))
	} else {
		ui.SessionWrite(fmt.Sprintf(InfoLabel, app.App.Version, ""))
	}
	prompt := promptui.Prompt{
		Label:  "Filter",
		Stdin:  *ui.sess,
		Stdout: *ui.sess,
	}
	filter, err := prompt.Run()
	if err != nil {
		// ^C ^D is not error
		if err.Error() == "^C" {
			return "^C", nil
		} else if err.Error() == "^D" {
			ui.Exit()
			return "", fmt.Errorf("exit")
		}
		log.Errorf("Prompt error: %s", err)
		return "", err
	}
	log.Debugf("Filter: %s", filter)
	return filter, nil
}

// ShowMainMenu show main menu
func (ui *PUI) ShowMainMenu() {
	MainMenu := make([]*MenuItem, 0)
	selectedChain := make([]*MenuItem, 0)
	ui.ShowMenu(MainLabel, MainMenu, "Quit", selectedChain)
}
