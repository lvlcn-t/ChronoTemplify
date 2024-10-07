// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Header() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"top-0 sticky w-full z-50 text-white bg-gray-800 dark:bg-gray-900\"><header class=\"backdrop-filter backdrop-blur-md\"><div class=\"container mx-auto flex flex-wrap p-5 flex-col md:flex-row items-center\"><a class=\"flex title-font font-medium items-center text-white mb-4 md:mb-0\" href=\"/\"><img src=\"https://avatars.githubusercontent.com/u/75443136\" alt=\"Home\" class=\"w-10 h-10 rounded-full object-cover\"> <span class=\"ml-3 text-xl\">lvlcn-t</span></a><nav class=\"md:ml-auto flex flex-wrap items-center text-base justify-center\"><a class=\"mr-5 hover:text-indigo-300 transition duration-300 ease-in-out\" href=\"/\">Home</a> <a class=\"mr-5 hover:text-indigo-300 transition duration-300 ease-in-out\" href=\"/projects\">Projects</a></nav></div></header></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
