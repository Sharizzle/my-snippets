package code

import "github.com/sharizzle/my-snippets/server/router"

var Routes = router.RoutePrefix{
	"/codes",
	[]router.Route{
		router.Route{
			"CodesIndex",
			"GET",
			"",
			IndexHandler,
		},
		router.Route{
			"CodesShow",
			"GET",
			"/{codeId}",
			ShowHandler,
		},
		router.Route{
			"CodesCreate",
			"POST",
			"",
			CreateHandler,
		},
		router.Route{
			"DeleteHandler",
			"DELETE",
			"/{codeId}",
			DeleteHandler,
		},
		router.Route{
			"UpdateHandler",
			"PUT",
			"/{codeId}",
			UpdateHandler,
		},
		router.Route{
			"OptionHandler",
			"OPTIONS",
			"/{codeId}",
			OptionHandler,
		},
		router.Route{
			"OptionHandler",
			"OPTIONS",
			"",
			OptionHandler,
		},
	},
}