// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package register

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/Yahya-Elamri/signeitfaster/views/components"
)

func AddNewPost(users struct {
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
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = components.HeaderLogged(users).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"grid grid-cols-1 md:grid-cols-[272px_1fr]\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.SidePanel().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full md:max-w-[calc(100vw-272px)] md:min-h-[calc(100vh-63px)] flex flex-col items-center mx-auto justify-start p-4\"><div class=\"w-[74%] flex flex-col items-start justify-center gap-3 my-5\"><h1 class=\"poppins-semibold text-3xl\">Créer une annonce</h1><div class=\"w-full mb-4\"><ul class=\"flex gap-4\"><a href=\"/register/agence\"><li class=\"poppins-medium text-md py-1 border-b-2 border-black dark:border-white cursor-pointer\">Creer agence</li></a> <a href=\"/register/voiturelocation\"><li class=\"poppins-medium text-md py-1 hover:border-b-2 border-black dark:border-white cursor-pointer\">voiture de location</li></a> <a href=\"/register/voiturevendre\"><li class=\"poppins-medium text-md py-1 hover:border-b-2 border-black dark:border-white cursor-pointer\">voiture a vendre</li></a> <a href=\"/register/voitureparts\"><li class=\"poppins-medium text-md py-1 hover:border-b-2 border-black dark:border-white cursor-pointer\">Pièces de voiture</li></a></ul></div><div class=\"w-[78%] h-full flex flex-col justify-start items-start gap-3\"><div class=\"flex flex-col w-full items-start\"><p class=\"poppins-regular text-sm\">Remplissez le formulaire ci-dessous pour lister vos voitures de location sur notre place de marché.</p></div><div class=\"w-[78%] flex justify-start items-center\"><div class=\"w-full flex flex-col gap-2 items-start\"><form hx-post=\"/add-agency\" hx-target=\"this\" hx-swap=\"outerHTML\" class=\"w-full flex flex-col fade-in gap-5 justify-between items-start\" method=\"POST\"><div class=\"w-full flex flex-col gap-2 dark:bg-[#0F0F0F] text-black dark:text-white\"><label class=\"poppins-regular\" for=\"username\">Nom D'agence</label> <input class=\"border border-[#d5e0d5] dark:bg-[#0F0F0F] text-black dark:text-white focus:outline-none hover:shadow-pink-300 hover:shadow-sm focus:shadow-pink-300 focus:shadow-sm px-[20px] py-[18px] rounded-xl\" type=\"text\" name=\"Username\" id=\"username\" required></div><div class=\"w-full flex flex-col gap-2 dark:bg-[#0F0F0F] text-black dark:text-white\"><label class=\"poppins-regular\" for=\"email\">Email</label> <input class=\"border border-[#d5e0d5] dark:bg-[#0F0F0F] text-black dark:text-white focus:outline-none hover:shadow-pink-300 hover:shadow-sm focus:shadow-pink-300 focus:shadow-sm px-[20px] py-[18px] rounded-xl\" type=\"text\" name=\"Email\" id=\"email\" required></div><div class=\"w-full flex flex-col gap-2 dark:bg-[#0F0F0F] text-black dark:text-white\"><label class=\"poppins-regular\" for=\"Phone\">Numero de Telephone</label> <input class=\"border border-[#d5e0d5] dark:bg-[#0F0F0F] text-black dark:text-white focus:outline-none hover:shadow-pink-300 hover:shadow-sm focus:shadow-pink-300 focus:shadow-sm px-[20px] py-[18px] rounded-xl\" type=\"text\" name=\"Phone\" id=\"Phone\" required></div><div class=\"w-full flex flex-col gap-2 dark:bg-[#0F0F0F] text-black dark:text-white\"><label class=\"poppins-regular\" for=\"Address\">Address</label> <input class=\"border border-[#d5e0d5] dark:bg-[#0F0F0F] text-black dark:text-white focus:outline-none hover:shadow-pink-300 hover:shadow-sm focus:shadow-pink-300 focus:shadow-sm px-[20px] py-[18px] rounded-xl\" type=\"text\" name=\"Address\" id=\"Address\" required></div><div class=\"w-full flex flex-col items-start gap-2\"><button class=\"poppins-regular bg-green-500 px-4 py-[12px] w-1/2 rounded-xl text-white hover:bg-green-700\" type=\"submit\">Ajoutez votre agence</button></div></form></div></div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
