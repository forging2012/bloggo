package app

import (
	"github.com/jgraham909/revmgo"
	"github.com/revel/revel"
	"strings"
)

func init() {
	revel.OnAppStart(revmgo.AppInit)
	revel.OnAppStart(AppInit)
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.ActionInvoker,           // Invoke the action.
	}

	revel.TemplateFuncs["join"] = func(a []string, sep string) string {
		return strings.Join(a, sep)
	}
}
