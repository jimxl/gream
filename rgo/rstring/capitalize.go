package rstring

import (
	"github.com/huandu/xstrings"
)

// TODO: 添加rails的string函数支持 https://api.rubyonrails.org/classes/String.html
func Capitalize(str string) string {
	return xstrings.ToCamelCase(str)
}
