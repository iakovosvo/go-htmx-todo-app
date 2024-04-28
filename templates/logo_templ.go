// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Logo() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col items-center justify-center mt-8\"><h1 class=\"text-4xl font-extrabold text-gray-800 flex items-center gap-2\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"50\" height=\"50\" fill=\"currentColor\" class=\"bi bi-person-lines-fill text-blue-500 rounded-full\" viewBox=\"0 0 16 16\" style=\"clip-path: circle();\"><path d=\"M1 1.5A.5.5 0 0 1 1.5 1h13a.5.5 0 0 1 .5.5v13a.5.5 0 0 1-.5.5h-13a.5.5 0 0 1-.5-.5v-13Zm8 0v13h3V1.5h-3ZM1.5 10a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-5ZM2 7 a1 1 0 1 1 2 0 1 1 0 0 1-2 0Zm3-4 a1 1 0 1 0 0-2 1 1 0 0 0 0 2Zm-2.496 6 a.5.5 0 1 0 0-1 .5.5 0 0 0 0 1Z\"></path></svg> <span class=\"text-blue-500\">Connectly</span></h1><p class=\"text-gray-300 text-lg\">A minimal app for adding contacts.</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}