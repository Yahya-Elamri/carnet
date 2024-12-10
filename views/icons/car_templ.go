// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Voiture(width, class string) templ.Component {
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
		var templ_7745c5c3_Var2 = []any{class}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(width)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/icons/car.templ`, Line: 4, Col: 18}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/icons/car.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" viewBox=\"0 0 48 48\" id=\"a\" xmlns=\"http://www.w3.org/2000/svg\"><path class=\"stroke-black dark:stroke-white\" d=\"M8.4,36.1457s-.1939,1.511,0,2.2286c.1347,.4981,.264,1.3371,.78,1.3371h2.8971c.516,0,.6453-.8391,.78-1.3371,.1939-.7172,0-2.2286,0-2.2286\"></path><rect class=\"stroke-black dark:stroke-white\" x=\"4.5\" y=\"21.66\" width=\"39\" height=\"14.4857\" rx=\"4.4571\" ry=\"4.4571\"></rect><path class=\"stroke-black dark:stroke-white\" d=\"M35.1429,36.1457s-.1939,1.511,0,2.2286c.1347,.4981,.264,1.3371,.78,1.3371h2.8971c.516,0,.6453-.8391,.78-1.3371,.1939-.7172,0-2.2286,0-2.2286\"></path><circle class=\"stroke-black dark:stroke-white\" cx=\"10.6286\" cy=\"29.46\" r=\"2.2286\"></circle><path class=\"stroke-black dark:stroke-white\" d=\"M8.9571,21.66l2.5629-10.0286c.438-1.7138,2.3534-3.3429,4.1229-3.3429h16.7143c1.7695,0,3.6849,1.6291,4.1229,3.3429l2.5629,10.0286\"></path><circle class=\"stroke-black dark:stroke-white\" cx=\"37.3714\" cy=\"29.46\" r=\"2.2286\"></circle></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Transmission(width, class string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var6 = []any{class}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var6...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(width)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/icons/car.templ`, Line: 8, Col: 21}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var6).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/icons/car.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M2 4C2 2.89543 2.89543 2 4 2C5.10457 2 6 2.89543 6 4C6 4.83934 5.48296 5.55793 4.75 5.85462V11.25H11.25V5.85462C10.517 5.55793 10 4.83934 10 4C10 2.89543 10.8954 2 12 2C13.1046 2 14 2.89543 14 4C14 4.83934 13.483 5.55793 12.75 5.85462V11.25H16C16.964 11.25 17.6116 11.2484 18.0946 11.1835C18.5561 11.1214 18.7536 11.0142 18.8839 10.8839C19.0142 10.7536 19.1214 10.5561 19.1835 10.0946C19.2484 9.61157 19.25 8.96401 19.25 8V5.85462C18.517 5.55793 18 4.83934 18 4C18 2.89543 18.8954 2 20 2C21.1046 2 22 2.89543 22 4C22 4.83934 21.483 5.55793 20.75 5.85462V8.05199C20.75 8.95048 20.7501 9.6997 20.6701 10.2945C20.5857 10.9223 20.4 11.4891 19.9445 11.9445C19.4891 12.4 18.9223 12.5857 18.2945 12.6701C17.6997 12.7501 16.9505 12.75 16.052 12.75L12.75 12.75L12.75 18.1454C13.483 18.4421 14 19.1607 14 20C14 21.1046 13.1046 22 12 22C10.8954 22 10 21.1046 10 20C10 19.1607 10.517 18.4421 11.25 18.1454V12.75H4.75V18.1454C5.48296 18.4421 6 19.1607 6 20C6 21.1046 5.10457 22 4 22C2.89543 22 2 21.1046 2 20C2 19.1607 2.51704 18.4421 3.25 18.1454V5.85462C2.51704 5.55793 2 4.83934 2 4Z\" class=\"fill-black dark:fill-white\"></path> <path opacity=\"0.5\" fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M17.25 15C17.25 14.5858 17.5858 14.25 18 14.25H20.2857C21.6612 14.25 22.75 15.3839 22.75 16.75C22.75 17.8285 22.0713 18.7624 21.1086 19.1077L22.6396 21.6084C22.8559 21.9616 22.7449 22.4234 22.3916 22.6396C22.0384 22.8559 21.5766 22.7449 21.3604 22.3916L19.4369 19.25H18.75V22C18.75 22.4142 18.4142 22.75 18 22.75C17.5858 22.75 17.25 22.4142 17.25 22V15ZM18.75 17.75H20.2857C20.8038 17.75 21.25 17.3169 21.25 16.75C21.25 16.1831 20.8038 15.75 20.2857 15.75H18.75V17.75Z\" class=\"fill-black dark:fill-white\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Fuel(width, class string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var10 = []any{class}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var10...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg fill=\"#000000\" width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(width)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/icons/car.templ`, Line: 15, Col: 36}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var10).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/icons/car.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" viewBox=\"0 0 56 56\" id=\"Layer_1\" version=\"1.1\" xml:space=\"preserve\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"><g><path d=\"M33.7,27.5c-0.3-0.1-7.3-1.5-11.2,2.1c-3.4,3.1-3.7,9-3.7,10.7l-1.5,1.3c-0.2,0.2-0.3,0.5-0.3,0.7c0,0.2,0.1,0.5,0.3,0.7   c0.2,0.2,0.5,0.3,0.7,0.3c0.2,0,0.5-0.1,0.7-0.3l1.5-1.3c0.5,0.1,1.4,0.1,2.5,0.1c2.6,0,6.4-0.4,8.8-2.7c3.9-3.6,3.1-10.7,3.1-11   C34.5,27.9,34.1,27.6,33.7,27.5z M30.1,37.8C30.1,37.8,30.1,37.8,30.1,37.8c-2,1.9-5.6,2.2-7.9,2.1l5.1-4.6   c0.2-0.2,0.3-0.5,0.3-0.7c0-0.2-0.1-0.5-0.3-0.7c-0.4-0.4-1-0.4-1.4-0.1l-5.1,4.6c0.2-2.2,1-5.5,3-7.3c2.5-2.3,7-2,8.7-1.8   C32.7,31.1,32.6,35.5,30.1,37.8z\"></path> <path d=\"M35,8.2H16.1c-1.4,0-2.5,1.1-2.5,2.5v7.5c0,1.4,1.1,2.5,2.5,2.5H35c1.4,0,2.5-1.1,2.5-2.5v-7.5C37.5,9.3,36.4,8.2,35,8.2z    M35.5,18.2c0,0.3-0.2,0.5-0.5,0.5H16.1c-0.3,0-0.5-0.2-0.5-0.5v-7.5c0-0.3,0.2-0.5,0.5-0.5H35c0.3,0,0.5,0.2,0.5,0.5V18.2z\"></path> <path d=\"M27.9,11.4h-1c-0.6,0-1,0.4-1,1s0.4,1,1,1h1c0.6,0,1-0.4,1-1S28.5,11.4,27.9,11.4z\"></path> <path d=\"M33.1,11.4h-1c-0.6,0-1,0.4-1,1s0.4,1,1,1h1c0.6,0,1-0.4,1-1S33.6,11.4,33.1,11.4z\"></path> <path d=\"M27.9,15.1h-1c-0.6,0-1,0.4-1,1s0.4,1,1,1h1c0.6,0,1-0.4,1-1S28.5,15.1,27.9,15.1z\"></path> <path d=\"M33.1,15.1h-1c-0.6,0-1,0.4-1,1s0.4,1,1,1h1c0.6,0,1-0.4,1-1S33.6,15.1,33.1,15.1z\"></path> <path d=\"M46.8,12.5C46.8,12.5,46.8,12.5,46.8,12.5l-1.7-1.7c-0.4-0.4-1-0.4-1.4,0s-0.4,1,0,1.4l1.4,1.4v6.1l4.4,3V38   c0,0.6-0.5,1.1-1.1,1.1c-0.6,0-1.1-0.5-1.1-1.1v-9.9c0-1.1-0.9-2-2-2h-1.9v-19c0-2.2-1.8-4.1-4.1-4.1H11.9C9.7,3,7.8,4.8,7.8,7.1   v37.1c0,0.8,0.2,1.5,0.6,2.1H6.9c-1.3,0-2.4,1.1-2.4,2.4v2c0,1.3,1.1,2.4,2.4,2.4h37.4c1.3,0,2.4-1.1,2.4-2.4v-2   c0-1.3-1.1-2.4-2.4-2.4h-1.6c0.4-0.6,0.6-1.3,0.6-2.1V28.1h1.9V38c0,1.7,1.4,3.1,3.1,3.1c1.7,0,3.1-1.4,3.1-3.1V17.2L46.8,12.5z    M44.7,48.6v2c0,0.2-0.2,0.4-0.4,0.4H6.9c-0.2,0-0.4-0.2-0.4-0.4v-2c0-0.2,0.2-0.4,0.4-0.4h5h27.3h5C44.5,48.2,44.7,48.4,44.7,48.6   z M11.9,46.2c-1.1,0-2.1-0.9-2.1-2.1V7.1c0-1.1,0.9-2.1,2.1-2.1h27.3c1.1,0,2.1,0.9,2.1,2.1v37.1c0,1.1-0.9,2.1-2.1,2.1H11.9z    M47.1,18.7v-3l2.4,2.4v2.3L47.1,18.7z\"></path></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Seats(width, class string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var13 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var13 == nil {
			templ_7745c5c3_Var13 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var14 = []any{class}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var14...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(width)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/icons/car.templ`, Line: 29, Col: 21}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var16 string
		templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var14).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/icons/car.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" version=\"1.1\" id=\"Capa_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" viewBox=\"0 0 353.926 353.926\" xml:space=\"preserve\"><path class=\"fill-black dark:fill-white\" d=\"M210.286,344.926c0,4.971-4.029,9-9,9h-48.65c-4.971,0-9-4.029-9-9s4.029-9,9-9h48.65\r\n\tC206.257,335.926,210.286,339.955,210.286,344.926z M289.677,258.958v25.928c0,19.259-15.67,34.928-34.931,34.928H99.177\r\n\tc-19.259,0-34.928-15.668-34.928-34.928v-25.928c0-4.971,4.029-9,9-9h2.394c-0.021-0.258-0.033-0.52-0.033-0.784v-24.118\r\n\tc-0.013-0.535,0.023-1.066,0.105-1.588c0.204-1.329,0.699-2.561,1.418-3.631c0.705-1.055,1.639-1.969,2.767-2.659\r\n\tc0.457-0.281,0.94-0.522,1.446-0.719c3.564-1.483,7.107-3.016,10.605-4.586V101.909c0-17.877,11.375-33.581,27.599-39.623\r\n\tc-0.019-0.492-0.028-0.984-0.028-1.48V38.578C119.521,17.306,136.827,0,158.098,0h37.725C217.095,0,234.4,17.306,234.4,38.578\r\n\tv22.229c0,0.495-0.01,0.988-0.028,1.478c6.395,2.378,12.129,6.28,16.702,11.351c0.16-0.3,0.318-0.599,0.478-0.899\r\n\tc2.318-4.396,7.761-6.081,12.16-3.76c4.396,2.319,6.079,7.764,3.76,12.16c-16.845,31.926-41.307,61.508-72.707,87.923\r\n\tc-25.063,21.083-53.512,39.294-84.813,54.313v26.586h134.02V141.64c0-4.971,4.029-9,9-9s9,4.029,9,9v108.318h18.706\r\n\tC285.647,249.958,289.677,253.987,289.677,258.958z M137.521,60.807c0,1.842,0.243,3.629,0.699,5.33\r\n\tc0.073,0.22,0.138,0.444,0.193,0.672c2.574,8.428,10.424,14.576,19.684,14.576h37.725c9.259,0,17.109-6.146,19.685-14.573\r\n\tc0.057-0.231,0.122-0.458,0.195-0.68c0.455-1.699,0.698-3.484,0.698-5.325V38.578C216.4,27.231,207.169,18,195.822,18h-37.725\r\n\tc-11.346,0-20.576,9.231-20.576,20.578V60.807z M109.951,203.272c56.184-28.521,102.335-68.15,131.162-112.739\r\n\tc-2.612-4.871-6.75-8.658-11.666-10.83c-6.622,11.738-19.213,19.681-33.625,19.681h-37.725c-14.411,0-27.002-7.944-33.624-19.682\r\n\tc-8.604,3.8-14.522,12.438-14.522,22.207V203.272z M271.677,267.958h-18.57c-0.046,0-0.091,0.001-0.136,0.001h-152.02\r\n\tc-0.045,0-0.09,0-0.136-0.001H82.249v16.928c0,9.334,7.594,16.928,16.928,16.928h155.569c9.336,0,16.931-7.594,16.931-16.928\r\n\tV267.958z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func MileAge(width, class string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var17 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var17 == nil {
			templ_7745c5c3_Var17 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var18 = []any{class}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var18...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg version=\"1.1\" id=\"Capa_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var19 string
		templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs(width)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/icons/car.templ`, Line: 53, Col: 14}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var20 string
		templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var18).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/icons/car.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" viewBox=\"0 0 501.015 501.015\" xml:space=\"preserve\"><g><path class=\"fill-black dark:fill-white\" d=\"M242.105,115.922L242.105,115.922h-0.006V96.015V56.188h0.006h16.804v59.734h-0.018v0.012l-8.391-0.012H242.105z\r\n\t\t M39.555,317.098l41.919-7.389l16.908-2.979v-0.012h0.006l-1.478-8.334l-1.442-8.193h-0.006v-0.012l-58.83,10.368L39.555,317.098z\r\n\t\t M411.266,169.234l-30.228,17.437v0.012l0,0l4.663,8.086l3.735,6.452l0,0l51.731-29.864l-8.405-14.549L411.266,169.234z\r\n\t\t M150.24,390.923l-6.623-5.556l0,0l0,0l-38.402,45.749l12.879,10.805l16.71-19.919l21.687-25.842l0,0L150.24,390.923z\r\n\t\t M432.562,386.81l8.394-14.559l-18.046-10.421l-33.673-19.458v0.012l0,0l-4.641,8.027l-3.753,6.514l0,0v0.012L432.562,386.81z\r\n\t\t M36.703,242.282l20.98,3.686l37.84,6.69v-0.012h0.006l2.5-14.215l0.405-2.329v-0.012l-58.827-10.382L36.703,242.282z\r\n\t\t M461.04,226.158l-15.989,2.813l-42.841,7.542v0.012h-0.006l2.163,12.256l0.757,4.279l0,0v0.012l58.836-10.367L461.04,226.158z\r\n\t\t M189.346,128.104v0.012l13.636-4.974l2.143-0.78h0.006l-20.422-56.14l-15.791,5.757l2.911,8.003L189.346,128.104L189.346,128.104z\r\n\t\t M67.814,386.431l51.725-29.86v-0.012l0,0l-6.469-11.201l-1.93-3.346l0,0l0,0l-51.734,29.861l8.402,14.552h0.006V386.431z\r\n\t\t M343.782,396.432l38.396,45.773L395.04,431.4v-0.012l-38.385-45.762v0.012v-0.012l-8.931,7.495L343.782,396.432L343.782,396.432z\r\n\t\t M59.625,170.992L59.625,170.992l51.731,29.876l0,0l7.276-12.613l1.12-1.936v-0.012l-51.731-29.864L59.625,170.992z\r\n\t\t M463.877,300.973L463.877,300.973l-58.824-10.379l-0.626,3.547l-2.294,13.004l58.83,10.379L463.877,300.973z M118.408,101.05\r\n\t\tl-12.868,10.802l14.109,16.81l24.291,28.948l0,0l0,0l7.601-6.384l5.26-4.412l0,0l0.006-0.012L118.408,101.05z M311.249,80.203\r\n\t\tL295.846,122.5l0.012,0.012l0,0l10.521,3.818l5.261,1.912l0,0h0.006l20.434-56.126l-15.794-5.742L311.249,80.203z M353.712,155.154\r\n\t\tl3.263,2.731l0,0l0,0l38.402-45.767l-12.873-10.79h-0.006l0,0l-38.396,45.752l0,0v0.012L353.712,155.154z M250.507,27.021\r\n\t\tC112.382,27.021,0,139.394,0,277.531c0,79.57,37.333,150.535,95.361,196.462h25.36C57.036,431.778,14.939,359.49,14.939,277.531\r\n\t\tc0-129.896,105.673-235.568,235.567-235.568c129.898,0,235.571,105.672,235.571,235.568c0,81.959-42.097,154.247-105.785,196.462\r\n\t\th25.363c58.02-45.927,95.358-116.892,95.358-196.462C501.015,139.399,388.633,27.021,250.507,27.021z M201.607,387.234\r\n\t\tc5.255,0,7.876-5.236,7.876-15.698c0-10.829-2.568-16.255-7.713-16.255c-5.423,0-8.139,5.503-8.139,16.521\r\n\t\tC193.631,382.093,196.294,387.234,201.607,387.234z M179.152,336.71h44.553v69.699h-44.553V336.71z M189.177,372.009\r\n\t\tc0,6.124,1.052,10.793,3.165,14.032c2.11,3.233,5.048,4.858,8.819,4.858c4.031,0,7.164-1.702,9.41-5.083\r\n\t\tc2.246-3.398,3.372-8.299,3.372-14.735c0-12.986-3.978-19.482-11.931-19.482c-4.158,0-7.33,1.708-9.531,5.172\r\n\t\tC190.276,360.211,189.177,365.294,189.177,372.009z M229.56,336.71h43.754v69.699H229.56V336.71z M241.053,389.327\r\n\t\tc1.738,1.04,4.359,1.572,7.874,1.572c4.17,0,7.459-1.123,9.847-3.369s3.594-5.213,3.594-8.89c0-3.547-1.111-6.348-3.334-8.394\r\n\t\tc-2.233-2.033-5.378-3.062-9.445-3.062c-0.993,0-2.065,0.035-3.233,0.106v-11.136h14.287v-3.925h-18.403V371.3\r\n\t\tc2.92-0.213,4.909-0.319,5.964-0.319c3.136,0,5.538,0.697,7.214,2.092c1.679,1.396,2.524,3.346,2.524,5.828\r\n\t\tc0,2.519-0.822,4.527-2.471,6.053c-1.646,1.525-3.807,2.281-6.493,2.281c-2.689,0-5.326-0.839-7.93-2.518v4.61H241.053z\r\n\t\t M277.848,336.71h44.006v69.699h-44.006V336.71z M288.788,372.009c0,6.124,1.053,10.793,3.168,14.032\r\n\t\tc2.11,3.233,5.049,4.858,8.819,4.858c4.031,0,7.17-1.702,9.41-5.083c2.246-3.398,3.369-8.299,3.369-14.735\r\n\t\tc0-12.986-3.984-19.482-11.935-19.482c-4.154,0-7.329,1.708-9.533,5.172C289.888,360.211,288.788,365.294,288.788,372.009z\r\n\t\t M301.225,387.234c5.249,0,7.879-5.236,7.879-15.698c0-10.829-2.571-16.255-7.714-16.255c-5.426,0-8.139,5.503-8.139,16.521\r\n\t\tC293.245,382.093,295.905,387.234,301.225,387.234z M257.815,300.725c-1.052,0-1.974,0.379-2.943,0.58L140.064,200.018\r\n\t\tL242.191,314.97c-0.068,0.579-0.358,1.1-0.358,1.673c0,8.849,7.144,16,15.982,16c8.784,0,15.924-7.151,15.924-16\r\n\t\tC273.751,307.877,266.6,300.725,257.815,300.725z\"></path></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Carnet(width, color string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var21 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var21 == nil {
			templ_7745c5c3_Var21 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg version=\"1.0\" xmlns=\"http://www.w3.org/2000/svg\" width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var22 string
		templ_7745c5c3_Var22, templ_7745c5c3_Err = templ.JoinStringErrs(width)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/icons/car.templ`, Line: 96, Col: 13}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var22))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" viewBox=\"0 0 107.000000 85.000000\" preserveAspectRatio=\"xMidYMid meet\"><g transform=\"translate(0.000000,85.000000) scale(0.100000,-0.100000)\" fill=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var23 string
		templ_7745c5c3_Var23, templ_7745c5c3_Err = templ.JoinStringErrs(color)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/icons/car.templ`, Line: 99, Col: 12}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var23))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" stroke=\"none\"><path d=\"M380 820 c-11 -11 -20 -29 -20 -40 0 -15 -7 -20 -28 -20 -34 0 -52\r\n\t-16 -52 -46 0 -14 8 -24 21 -27 16 -4 19 -9 11 -24 -14 -27 -21 -112 -14 -163\r\n\t7 -45 7 -45 47 -45 33 0 40 4 45 25 l6 25 139 0 139 0 6 -25 c5 -21 12 -25 45\r\n\t-25 40 0 40 0 47 45 7 51 0 136 -14 163 -8 15 -5 20 11 24 14 3 21 13 21 28 0\r\n\t29 -22 47 -51 43 -18 -2 -24 3 -30 27 -14 50 -30 55 -176 55 -120 0 -135 -2\r\n\t-153 -20z m293 -27 c3 -5 8 -23 12 -40 l7 -33 -157 0 -157 0 7 33 c4 17 9 35\r\n\t12 40 2 4 64 7 138 7 74 0 136 -3 138 -7z m47 -122 c8 -17 6 -20 -21 -26 -19\r\n\t-4 -33 -15 -36 -26 -5 -17 -16 -19 -129 -19 -105 0 -124 2 -124 16 0 8 -13 20\r\n\t-30 25 -28 10 -38 28 -23 42 3 4 84 7 179 7 161 0 174 -1 184 -19z m-350 -61\r\n\tc12 -8 9 -10 -12 -10 -16 0 -28 5 -28 10 0 13 20 13 40 0z m370 0 c0 -5 -12\r\n\t-10 -27 -10 -22 0 -25 2 -13 10 20 13 40 13 40 0z m0 -60 c0 -6 -75 -10 -205\r\n\t-10 -130 0 -205 4 -205 10 0 6 75 10 205 10 130 0 205 -4 205 -10z m-380 -55\r\n\tc0 -8 -7 -15 -15 -15 -8 0 -15 7 -15 15 0 8 7 15 15 15 8 0 15 -7 15 -15z\r\n\tm380 0 c0 -8 -7 -15 -15 -15 -16 0 -20 12 -8 23 11 12 23 8 23 -8z\"></path> <path d=\"M430 635 c0 -13 15 -15 108 -13 63 2 107 7 107 13 0 6 -44 11 -107\r\n\t13 -93 2 -108 0 -108 -13z\"></path> <path d=\"M118 389 c-10 -5 -20 -25 -24 -44 -5 -31 -9 -34 -34 -28 -22 4 -31 0\r\n\t-40 -17 -14 -26 -7 -50 16 -50 13 0 15 -5 8 -22 -4 -13 -8 -66 -8 -118 l-1\r\n\t-95 38 -3 c35 -3 39 0 50 27 l11 31 137 0 137 0 6 -30 c5 -26 10 -30 40 -30\r\n\t51 0 59 17 53 112 -2 46 -8 93 -12 106 -6 17 -3 22 9 22 11 0 16 9 16 30 0 23\r\n\t-5 31 -24 36 -14 3 -29 3 -35 0 -5 -3 -12 7 -16 24 -12 56 -23 60 -173 60 -75\r\n\t0 -145 -5 -154 -11z m292 -38 c0 -5 3 -23 6 -40 l7 -31 -157 0 -157 0 6 23 c4\r\n\t12 9 30 11 40 5 15 21 17 145 17 76 0 139 -4 139 -9z m44 -117 c9 -22 9 -22\r\n\t-24 -29 -21 -5 -32 -14 -36 -31 -6 -24 -8 -24 -125 -24 -118 0 -120 0 -129 25\r\n\t-5 14 -20 27 -35 31 -24 6 -33 23 -18 37 3 4 86 7 184 7 152 0 178 -2 183 -16z\r\n\tm-346 -72 c3 -9 -3 -13 -19 -10 -12 1 -24 9 -27 16 -3 9 3 13 19 10 12 -1 24\r\n\t-9 27 -16z m362 3 c0 -8 -9 -15 -19 -15 -24 0 -36 16 -19 23 24 10 38 7 38 -8z\r\n\tm10 -55 c0 -6 -77 -10 -210 -10 -133 0 -210 4 -210 10 0 6 77 10 210 10 133 0\r\n\t210 -4 210 -10z m-390 -55 c0 -8 -7 -15 -15 -15 -8 0 -15 7 -15 15 0 8 7 15\r\n\t15 15 8 0 15 -7 15 -15z m380 0 c0 -8 -7 -15 -15 -15 -8 0 -15 7 -15 15 0 8 7\r\n\t15 15 15 8 0 15 -7 15 -15z\"></path> <path d=\"M160 195 c0 -13 15 -15 108 -13 63 2 107 7 107 13 0 6 -44 11 -107\r\n\t13 -93 2 -108 0 -108 -13z\"></path> <path d=\"M647 385 c-9 -8 -18 -28 -22 -45 -4 -17 -11 -27 -16 -24 -6 3 -21 3\r\n\t-35 0 -19 -5 -24 -13 -24 -36 0 -21 5 -30 16 -30 12 0 15 -5 9 -22 -4 -13 -10\r\n\t-60 -12 -106 -6 -95 2 -112 53 -112 30 0 35 4 41 30 l6 31 141 -3 c87 -2 140\r\n\t-7 139 -13 -4 -28 19 -46 55 -43 l37 3 0 95 c0 52 -4 105 -8 118 -6 17 -5 22\r\n\t8 22 22 0 29 24 15 50 -9 17 -18 21 -40 17 -25 -6 -29 -3 -34 28 -4 19 -15 39\r\n\t-26 45 -10 5 -79 10 -154 10 -106 0 -138 -3 -149 -15z m297 -45 c3 -11 9 -29\r\n\t12 -40 5 -19 1 -20 -150 -20 -86 0 -156 4 -156 9 0 5 3 23 6 40 l7 31 138 0\r\n\tc129 0 138 -1 143 -20z m44 -112 c2 -12 -5 -19 -23 -24 -15 -4 -30 -17 -35\r\n\t-30 -9 -23 -12 -24 -129 -24 -115 0 -120 1 -123 22 -2 14 -14 25 -33 31 -38\r\n\t12 -37 11 -29 31 6 15 26 16 188 14 161 -3 181 -5 184 -20z m-354 -54 c20 -8\r\n\t11 -24 -15 -24 -10 0 -19 7 -19 15 0 16 9 19 34 9z m376 -3 c0 -6 -5 -13 -10\r\n\t-16 -15 -9 -43 3 -35 15 8 13 45 13 45 1z m0 -61 c0 -6 -77 -10 -210 -10 -133\r\n\t0 -210 4 -210 10 0 6 77 10 210 10 133 0 210 -4 210 -10z m-380 -55 c0 -8 -7\r\n\t-15 -15 -15 -8 0 -15 7 -15 15 0 8 7 15 15 15 8 0 15 -7 15 -15z m380 0 c0 -8\r\n\t-7 -15 -15 -15 -8 0 -15 7 -15 15 0 8 7 15 15 15 8 0 15 -7 15 -15z\"></path> <path d=\"M697 204 c-17 -17 13 -24 109 -24 86 0 104 3 104 15 0 12 -18 15\r\n\t-103 15 -57 0 -107 -3 -110 -6z\"></path></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
