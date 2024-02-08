package conf

import (
	"github.com/beego/beego/v2/server/web"
)

type OnlyOfficeConf struct {
	DocumentServer string
	CallbackUrl    string
	Secret         string
	DocApiUrl      string
	ConvertApiUrl  string
	CsApiUrl       string
	ApiUrl         string
}

func GetOoConfig() *OnlyOfficeConf {
	DocumentServer, _ := web.AppConfig.String("oo_document_server")
	CallbackUrl, _ := web.AppConfig.String("oo_callback_url")
	Secret, _ := web.AppConfig.String("oo_secrrt")
	DocApiUrl, _ := web.AppConfig.String("oo_doc_api_url")
	ConvertApiUrl, _ := web.AppConfig.String("oo_convert_api_url")
	CsApiUrl, _ := web.AppConfig.String("oo_cs_api_url")
	ApiUrl, _ := web.AppConfig.String("oo_api_url")
	c := &OnlyOfficeConf{
		DocumentServer: DocumentServer,
		CallbackUrl:    CallbackUrl,
		Secret:         Secret,
		DocApiUrl:      DocApiUrl,
		ConvertApiUrl:  ConvertApiUrl,
		CsApiUrl:       CsApiUrl,
		ApiUrl:         ApiUrl,
	}
	return c
}
