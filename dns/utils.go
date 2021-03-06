package dns

import "strings"

var rootDomain string

// SetRootDomain 设置根域名
func SetRootDomain(domain string) {
	rootDomain = domain
}

// ParentDomain 获取父域名
func ParentDomain(domain string) string {
	idx := strings.Index(domain, ".")
	return string([]byte(domain)[idx+1:])
}
