// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"
import "strings"

import (
	"fmt"
	"github.com/Yahya-Elamri/signeitfaster/views/icons"
)

func BackgroundImage(path, rounded string) templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`width:100%;`)
	templ_7745c5c3_CSSBuilder.WriteString(`height:100%;`)
	templ_7745c5c3_CSSBuilder.WriteString(string(templ.SanitizeCSS(`background-image`, fmt.Sprintf("url('/assets%s')", path))))
	templ_7745c5c3_CSSBuilder.WriteString(`background-size:cover;`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-position:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-repeat:no-repeat;`)
	templ_7745c5c3_CSSBuilder.WriteString(`overflow:hidden;`)
	templ_7745c5c3_CSSBuilder.WriteString(string(templ.SanitizeCSS(`border-radius`, fmt.Sprintf("%s", rounded))))
	templ_7745c5c3_CSSID := templ.CSSID(`BackgroundImage`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Header(logo string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full z-10 sticky flex gap-10 h-[100px]\"><div class=\"h-full container flex mx-auto justify-between items-center px-4\"><a href=\"/\" class=\"text-3xl poppins-regular h-full flex items-center\"><img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(logo)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/header.templ`, Line: 22, Col: 96}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"h-[85%] cursor-pointer\"></a><div class=\"flex gap-[40px] w-fit items-center\"><ul class=\"flex gap-2 items-center p-2 justify-center poppins-regular w-full rounded-full bg-[#F4F2F3] h-[50px] shadow-sm\"><li class=\"poppins-regular rounded-full px-4 py-2 bg-[#ffffff] hover:bg-[#F4F2F3] cursor-pointer\"><a href=\"/\">Home</a></li><li class=\"poppins-regular rounded-full px-4 py-2 bg-[#ffffff] hover:bg-[#F4F2F3] cursor-pointer\"><a href=\"/signup\">Register</a></li><li class=\"poppins-regular rounded-full px-4 py-2 bg-[#ffffff] hover:bg-[#F4F2F3] cursor-pointer\"><a href=\"/signup\">Pricing</a></li><li class=\"poppins-regular rounded-full px-4 py-2 bg-[#ffffff] hover:bg-[#F4F2F3] cursor-pointer\"><a href=\"/signup\">About</a></li><li class=\"poppins-regular rounded-full px-4 py-2 bg-[#ffffff] hover:bg-[#F4F2F3] cursor-pointer\"><a href=\"/signup\">Contact</a></li></ul><div class=\"flex gap-5 items-center w-full\"><button class=\"text-[#80AF81]\"><a href=\"/login\">Log in</a></button> <button class=\"bg-[#508D4E] px-4 py-2 rounded-full poppins-regular text-white hover:bg-green-700\"><a href=\"/signup\">Sign up</a></button></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func HeaderLogged(users struct {
	UserId     uint
	Username   string
	ProfileUrl string
}) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full z-10 sticky flex gap-10 h-[63px] p-2 px-7 border-b dark:border-[#3F3F3F] top-0 bg-[#ffffff] dark:bg-[#0F0F0F] text-black dark:text-white\"><style>\r\n            .header-dropdown {\r\n            position: relative;\r\n            display: inline-block;\r\n            }\r\n            \r\n            .header-dropdown-content {\r\n            display: none;\r\n            position: absolute;\r\n            background-color: #f9f9f9;\r\n            min-width: 160px;\r\n            box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);\r\n            z-index: 1;\r\n            }\r\n\r\n            .show{\r\n                display:block\r\n            }\r\n        </style><div class=\"h-full w-full flex justify-between items-center px-6 md:p-0\"><a href=\"/\" class=\"poppins-regular h-full gap-2 flex items-center\"><img class=\"w-[55px] h-auto\" src=\"/assets/imgs/logo.png\" alt=\"logo\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.Logo("100px", "fill-black dark:fill-white").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a><div class=\"sm:flex items-center hidden max-w-[750px] md:w-[560px] lg:w-[580px] justify-between border-[#d5e0d5] border-[1px] dark:border-[#3F3F3F] rounded-full px-4 py-2\"><div class=\"w-[25px]\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"full\" class=\"stroke-black dark:stroke-white\" fill=\"none\" aria-hidden=\"true\" viewBox=\"0 0 24 24\" role=\"img\"><path vector-effect=\"non-scaling-stroke\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-miterlimit=\"10\" stroke-width=\"1.5\" d=\"M10.688 18.377a7.688 7.688 0 100-15.377 7.688 7.688 0 000 15.377zm5.428-2.261L21 21\"></path></svg></div><input class=\"focus:border-none focus:outline-none w-[90%] bg-transparent\" name=\"search\" type=\"text\" placeholder=\"Que Recherchez-vous\"></div><div class=\"flex justify-between w-fit items-center\"><div class=\"flex gap-2 items-center\"><div class=\"relative inline-block\"><a href=\"/poster/post\" class=\"dark:hover:bg-[#101010] rounded-full cursor-pointer border border-[#d5e0d5] dark:border-[#3F3F3F] p-2 hidden md:block poppins-regular\"><abbr title=\"Ajoutez votre anonnce\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.Plus("24px", "fill-black dark:fill-white").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</abbr></a><div class=\"hidden bg-[#f9f9f9] w-[250px] border rounded-xl border-[#d5e0d5] absolute top-[60px] -right-1/3\" id=\"myDropdown1\"><a class=\"dark:hover:bg-[#101010] block px-4 py-2 hover:bg-[#f1f1f1] rounded-xl poppins-regular\" href=\"/posts\">Gérer les publications</a> <a class=\"dark:hover:bg-[#101010] block px-4 py-2 hover:bg-[#f1f1f1] rounded-xl poppins-regular\" href=\"/newpost\">Créer un nouveau post</a></div></div><div class=\"rounded-full cursor-pointer border border-[#d5e0d5] dark:border-[#3F3F3F] p-2 hidden md:block\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.Notification("24px", "fill-black dark:fill-white").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"header-dropdown\"><div class=\"w-[45px] h-[45px] cursor-pointer bg-cover bg-center rounded-full border border-[#d5e0d5] dark:border-[#3F3F3F]\" onclick=\"toggleHeaderDropdown(&#39;myDropdown&#39;)\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 = []any{BackgroundImage(users.ProfileUrl, "9999rem")}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var4...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var4).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/header.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div><div class=\"header-dropdown-content hidden bg-[#f9f9f9] dark:bg-[#060809] border rounded-xl w-[280px] px-[16px] py-[8px] border-[#d5e0d5] dark:border-[#3F3F3F] absolute top-[60px] -right-3\" id=\"myDropdown\"><div class=\"w-full flex flex-col items-start justify-start\"><a class=\"w-full py-1 border-b dark:border-[#3F3F3F]\" href=\"/profile\"><div class=\"px-2 py-3 dark:hover:bg-[#101010] hover:bg-[#f1f1f1] rounded-xl poppins-regular text-md flex items-center gap-2 w-full\"><div class=\"w-[45px] h-[45px] cursor-pointer bg-cover bg-center rounded-full border border-[#d5e0d5] dark:border-[#3F3F3F]\" onclick=\"toggleHeaderDropdown(&#39;myDropdown&#39;)\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 = []any{BackgroundImage(users.ProfileUrl, "9999rem")}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var6...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var6).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/header.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div><div class=\"flex flex-col items-start\"><p class=\"poppins-regular text-md underline\">voir le profil</p><p class=\"poppins-regular text-sm text-grey-200\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(users.Username)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/header.templ`, Line: 99, Col: 108}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div></div></a> <a class=\"dark:hover:bg-[#101010] hover:bg-[#f1f1f1] rounded-xl my-1 px-2 py-3 poppins-regular text-md flex items-center w-full gap-1\" href=\"/\"><span class=\"flex justify-center gap-xs min-w-0 shrink w-[45px]\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.Fuel("25px", "fill-black dark:fill-white").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"flex items-center gap-xs min-w-0 shrink\">Fuel</span></a> <a class=\"dark:hover:bg-[#101010] hover:bg-[#f1f1f1] rounded-xl px-2 py-3 poppins-regular text-md flex items-center w-full justify-between\" href=\"/\"><div class=\"flex items-center gap-1\"><span class=\"flex justify-center gap-xs min-w-0 shrink w-[45px]\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.Moone("25px", "fill-black dark:fill-white").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"flex items-center gap-xs min-w-0 shrink\">Mode sombre</span></div><style>\r\n                                        .switch {\r\n                                            position: relative;\r\n                                            display: inline-block;\r\n                                            width: 60px;\r\n                                            height: 34px;\r\n                                            }\r\n\r\n                                            /* Hide default HTML checkbox */\r\n                                            .switch input {\r\n                                            opacity: 0;\r\n                                            width: 0;\r\n                                            height: 0;\r\n                                            }\r\n\r\n                                            /* The slider */\r\n                                            .slider {\r\n                                            position: absolute;\r\n                                            cursor: pointer;\r\n                                            top: 0;\r\n                                            left: 0;\r\n                                            right: 0;\r\n                                            bottom: 0;\r\n                                            background-color: #ccc;\r\n                                            -webkit-transition: .4s;\r\n                                            transition: .4s;\r\n                                            }\r\n\r\n                                            .slider:before {\r\n                                            position: absolute;\r\n                                            content: \"\";\r\n                                            height: 26px;\r\n                                            width: 26px;\r\n                                            left: 4px;\r\n                                            bottom: 4px;\r\n                                            background-color: white;\r\n                                            -webkit-transition: .4s;\r\n                                            transition: .4s;\r\n                                            }\r\n\r\n                                            input:checked + .slider {\r\n                                            background-color: #2196F3;\r\n                                            }\r\n\r\n                                            input:focus + .slider {\r\n                                            box-shadow: 0 0 1px #2196F3;\r\n                                            }\r\n\r\n                                            input:checked + .slider:before {\r\n                                            -webkit-transform: translateX(26px);\r\n                                            -ms-transform: translateX(26px);\r\n                                            transform: translateX(26px);\r\n                                            }\r\n\r\n                                            /* Rounded sliders */\r\n                                            .slider.round {\r\n                                            border-radius: 34px;\r\n                                            }\r\n\r\n                                            .slider.round:before {\r\n                                            border-radius: 50%;\r\n                                            }\r\n                                    </style><label class=\"switch\"><input onclick=\"toggleTheme()\" type=\"checkbox\"> <span class=\"slider round\"></span></label></a> <a class=\"dark:hover:bg-[#101010] hover:bg-[#f1f1f1] rounded-xl px-2 py-3 poppins-regular text-md flex items-center w-full gap-1\" href=\"/\"><span class=\"flex justify-center gap-xs min-w-0 shrink w-[45px]\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.Dashboard("25px", "stroke-black dark:stroke-white").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"flex items-center gap-xs min-w-0 shrink\">Dashboard</span></a> <a class=\"w-full py-1 border-b dark:border-[#3F3F3F]\" href=\"/\"><div class=\"dark:hover:bg-[#101010] hover:bg-[#f1f1f1] rounded-xl px-2 py-3 poppins-regular text-md flex items-center w-full gap-1\"><span class=\"flex justify-center gap-xs min-w-0 shrink w-[45px]\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.Money("25px", "fill-black dark:fill-white").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"flex items-center gap-xs min-w-0 shrink\">Monétisation</span></div></a> <a class=\"w-full py-1 border-b dark:border-[#3F3F3F]\" href=\"/parametre\"><div class=\"dark:hover:bg-[#101010] hover:bg-[#f1f1f1] rounded-xl px-2 py-3 poppins-regular text-md flex items-center w-full gap-1\"><span class=\"flex justify-center gap-xs min-w-0 shrink w-[45px]\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.Settings("25px", "fill-black dark:fill-white").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"flex items-center gap-xs min-w-0 shrink\">Parametre</span></div></a> <a class=\"dark:hover:bg-[#101010] my-1 hover:bg-[#f1f1f1] rounded-xl px-2 py-3 poppins-regular text-md flex items-center w-full gap-1\" href=\"/disconnect\"><span class=\"flex justify-center gap-xs min-w-0 shrink w-[45px]\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.LogOut("25px").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"flex items-center gap-xs min-w-0 shrink\">Se Deconnecter</span></a></div></div></div></div></div></div><script>\r\n        function toggleHeaderDropdown(DivId) {\r\n            document.getElementById(DivId).classList.toggle(\"show\");\r\n        }\r\n        function getValue() {\r\n            Value = document.getElementById(\"myDropdown2\").value;\r\n            console.log(Value)\r\n            var actionUrl = '/home/' + Value;\r\n            document.getElementById('searchFrom').action = actionUrl;\r\n        }\r\n        </script></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}