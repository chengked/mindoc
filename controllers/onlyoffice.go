package controllers

import (
	"encoding/json"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/mindoc-org/mindoc/models"
	"github.com/mindoc-org/mindoc/utils"
)

type OnlyController struct {
	web.Controller
}

// CommonCallback 通用回调状态
type CommonCallback struct {
	Key    string `json:"key"`
	Status int    `json:"status"`
}

// EditHasPrepareSave 2- 文档已准备好保存，
type EditHasPrepareSave struct {
	Key        string `json:"key"`
	Status     int    `json:"status"`
	Url        string `json:"url"`
	ChangesUrl string `json:"changesurl"`
	History    struct {
		ServerVersion string `json:"serverVersion"`
		Changes       []struct {
			Created string `json:"created"`
			User    struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"user"`
		} `json:"changes"`
	} `json:"history"`
	Users   []string `json:"users"`
	Actions []struct {
		Type   int    `json:"type"`
		Userid string `json:"userid"`
	} `json:"actions"`
	LastSave    time.Time `json:"lastsave"`
	NotModified bool      `json:"notmodified"`
	Filetype    string    `json:"filetype"`
}

// EditHasSaved 6 文档正在编辑，但当前文档状态已保存，
type EditHasSaved struct {
	Key        string `json:"key"`
	Status     int    `json:"status"`
	Url        string `json:"url"`
	ChangesUrl string `json:"changesurl"`
	History    struct {
		ServerVersion string `json:"serverVersion"`
		Changes       []struct {
			Created string `json:"created"`
			User    struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"user"`
		} `json:"changes"`
	} `json:"history"`
	Users   []string `json:"users"`
	Actions []struct {
		Type   int    `json:"type"`
		Userid string `json:"userid"`
	} `json:"actions"`
	LastSave      time.Time `json:"lastsave"`
	ForceSaveType int       `json:"forcesavetype"`
	Filetype      string    `json:"filetype"`
}

// savePath 保存路径
const savePath = "./"

var serverUrl = ""
var wsServer = ""
var documentServer = ""
var innerDocumentServer = "undefined"
var callBackUrl = ""

func init() {
	//取得客户端用户名
	err := os.MkdirAll(savePath, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error("[文件存储目录初始化异常 %s]", err)
	}
	val, err := config.String("serverUrl")
	if err != nil {
		logs.Error("[获取文件预览服务外部访问地址 %s]", val)
	} else {
		serverUrl = val
		logs.Info("[获取文件预览服务外部访问地址 %s]", val)
	}
	val, err = config.String("wsServer")
	if err != nil {
		logs.Error("[ws 访问地址  %s]")
	} else {
		wsServer = val
		logs.Info("[ws 访问地址  %s]", val)
	}
	val, err = config.String("documentServer")
	if err != nil {
		logs.Error("[文档服务器访问地址 %s]")
	} else {
		documentServer = val
		logs.Info("[文档服务器访问地址 %s]", val)
	}
	val, err = config.String("callBackUrl")
	if err != nil {
		logs.Error("[文档服务器回调地址 %s]")
	} else {
		callBackUrl = val
		logs.Info("[文档服务器内网回调地址 %s]", val)
	}
	val, err = config.String("innerDocumentServer")
	if err != nil {
		logs.Error("[内网文档服务器地址 %s]")
	} else {
		if val == "undefined" {
			innerDocumentServer = documentServer
		} else {
			innerDocumentServer = val
		}
		logs.Info("[内网文档服务器地址 %s]", val)
	}
}

func (c *OnlyController) OnlyOffice() {
	//pid转成64为
	// docId, err := c.GetInt64(":id")
	// if err != nil {
	// 	logs.Error(err)
	// 	c.Data["json"] = base.BuildResult(-1, "文件id参数存在问题")
	// 	_ = c.ServeJSON()
	// 	return
	// }
	//根据附件id取得附件的详细信息
	// attachment, err := models.GetOnlyAttachmentById(docId)
	// if err != nil {
	// 	logs.Error(err)
	// 	c.Data["json"] = base.BuildResult(-1, "根据文档id未查询到当前文档")
	// 	_ = c.ServeJSON()
	// 	return
	// }

	c.Data["Username"] = "打工人"

	c.Data["Mode"] = "edit"
	c.Data["Edit"] = true
	c.Data["Review"] = false

	c.Data["Doc"] = ""          //attachment
	c.Data["Key"] = ""          //strconv.FormatInt(attachment.Updated.UnixNano(), 10)
	c.Data["Doc.FileName"] = "" // attachment.FileName
	c.Data["serverUrl"] = serverUrl
	c.Data["wsServer"] = wsServer
	c.Data["documentServer"] = documentServer
	c.Data["callBackUrl"] = callBackUrl

	// if path.Ext(attachment.FileName) == ".docx" || path.Ext(attachment.FileName) == ".DOCX" {
	// 	c.Data["fileType"] = "docx"
	// 	c.Data["documentType"] = "text"
	// } else if path.Ext(attachment.FileName) == ".XLSX" || path.Ext(attachment.FileName) == ".xlsx" {
	// 	c.Data["fileType"] = "xlsx"
	// 	c.Data["documentType"] = "spreadsheet"
	// } else if path.Ext(attachment.FileName) == ".pptx" || path.Ext(attachment.FileName) == ".PPTX" {
	// 	c.Data["fileType"] = "pptx"
	// 	c.Data["documentType"] = "presentation"
	// } else if path.Ext(attachment.FileName) == ".doc" || path.Ext(attachment.FileName) == ".DOC" {
	// 	c.Data["fileType"] = "doc"
	// 	c.Data["documentType"] = "text"
	// } else if path.Ext(attachment.FileName) == ".txt" || path.Ext(attachment.FileName) == ".TXT" {
	// 	c.Data["fileType"] = "txt"
	// 	c.Data["documentType"] = "text"
	// } else if path.Ext(attachment.FileName) == ".XLS" || path.Ext(attachment.FileName) == ".xls" {
	// 	c.Data["fileType"] = "xls"
	// 	c.Data["documentType"] = "spreadsheet"
	// } else if path.Ext(attachment.FileName) == ".csv" || path.Ext(attachment.FileName) == ".CSV" {
	// 	c.Data["fileType"] = "csv"
	// 	c.Data["documentType"] = "spreadsheet"
	// } else if path.Ext(attachment.FileName) == ".ppt" || path.Ext(attachment.FileName) == ".PPT" {
	// 	c.Data["fileType"] = "ppt"
	// 	c.Data["documentType"] = "presentation"
	// } else if path.Ext(attachment.FileName) == ".pdf" || path.Ext(attachment.FileName) == ".PDF" {
	// 	c.Data["fileType"] = "pdf"
	// 	c.Data["documentType"] = "text"
	// 	c.Data["Mode"] = "view"
	// }

	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		c.Data["Type"] = "mobile"
	} else {
		c.Data["Type"] = "desktop"
	}
}

// getExpireTime 从文件url中提取过期时间
func getExpireTime(url string) time.Time {
	//获取连接到下载的预期过期事件
	paramsMap, err := utils.GetParams(url)
	if err != nil {
		logs.Error(err)
	}
	var expiresTime int64
	expires, ok := paramsMap["expires"]
	if ok {
		expiresTime, err = strconv.ParseInt(expires, 10, 64)
		if err != nil {
			expiresTime = time.Now().Unix()
		}
	}
	//时间戳转日期
	dataTimeStr := time.Unix(expiresTime, 0)
	return dataTimeStr
}

// DocCallback 协作页面的保存和回调
// 关闭浏览器标签后获取最新文档保存到文件夹
func (c *OnlyController) DocCallback() {
	c.Prepare()
	docId, err := c.GetInt("id")
	if err != nil {
		logs.Error(err)
	}
	//根据doc_id取得附件的详细信息
	doc, err := models.NewDocument().Find(docId)
	if err != nil {
		logs.Error("[通用-回调函数 解析回调异常] 查询不到该文档 %v", docId)
		c.Data["json"] = map[string]interface{}{"error": 0}
		_ = c.ServeJSON()
		return
	}

	var callback CommonCallback
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &callback)
	if err != nil {
		logs.Error("[通用-回调函数 解析回调异常]： %v", err)
	} else {
		logs.Info("[通用-回调函数 解析回调] :%v", callback)
	}
	logs.Debug("{}", callback.Status)
	if callback.Status == 1 || callback.Status == 4 {
		c.Data["json"] = map[string]interface{}{"error": 0}
		_ = c.ServeJSON()
		return
	} else if callback.Status == 6 {
		var editHasSaved EditHasSaved
		_ = json.Unmarshal(c.Ctx.Input.RequestBody, &editHasSaved)
		downloadUrl := strings.Replace(editHasSaved.Url, documentServer, innerDocumentServer, -1)
		//下载文件到本地
		err := utils.DownloadFile(downloadUrl, savePath+doc.FilePath)
		if err != nil {
			logs.Error(err)
		} else {
		}
		c.Data["json"] = map[string]interface{}{"error": 0}
		_ = c.ServeJSON()
	} else if callback.Status == 2 {
		var editHasPrepareSave EditHasPrepareSave
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &editHasPrepareSave)
		if err != nil {
			logs.Error("[状态2-回调函数 解析回调异常]： %v", err)
		}
		downloadUrl := strings.Replace(editHasPrepareSave.Url, documentServer, innerDocumentServer, -1)
		//下载文件到本地
		err := utils.DownloadFile(downloadUrl, savePath+doc.FilePath)
		if err != nil {
			logs.Error(err)
		} else {
		}
		c.Data["json"] = map[string]interface{}{"error": 0}
		_ = c.ServeJSON()
		return
	} else if callback.Status == 3 {
		c.Data["json"] = map[string]interface{}{"error": 0}
		_ = c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"error": 0}
		_ = c.ServeJSON()
	}
}
