package config

import "go-devops-admin/pkg/config"

func init() {
	config.Add("paging", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认分页条数
			"page_size": 10,

			// URL 中分页参数字段
			// 此值若修改,需要同时修改 requests 中的验证规则
			"url_query_page": "page",

			// URL 中排序的参数字段
			// 此值若修改,需要同时修改 requests 中的验证规则
			"url_query_sort": "sort",

			// URL 中排序规则的参数字段
			// 此值若修改,需要同时修改 requests 中的验证规则
			"url_query_order": "order",

			// URL 中每页页数的参数字段
			// 此值若修改,需要同时修改 requests 中的验证规则
			"url_query_page_size": "10",
		}
	})
}
