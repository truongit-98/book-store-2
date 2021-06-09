package appconfig

import "github.com/astaxie/beego/session"

var globalSessions *session.Manager
func init() {
	globalSessions, _ = session.NewManager("memory", &session.ManagerConfig{
		CookieName: "gosessionid",
		EnableSetCookie: true,
		Gclifetime: 3600,
		Maxlifetime: 3600,
		Secure: true,
		CookieLifeTime: 3600,
		ProviderConfig: "",

	})
	go globalSessions.GC()
}

func GetSessionManager() *session.Manager {
	return globalSessions
}