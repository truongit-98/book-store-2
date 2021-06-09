package test

import (
	"strings"
)

func ConfigHostPort(hosts, ports string) []string {
	hosts_temp := strings.Split(hosts, ",")
	ports_temp := strings.Split(ports, ",")
	var results []string

	for i := 0; i < len(hosts_temp); i++ {
		results = append(results, strings.TrimSpace(hosts_temp[i])+":"+strings.TrimSpace(ports_temp[i]))
	}

	return results
}
