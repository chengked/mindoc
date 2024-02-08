package models

import (
	"strconv"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mindoc-org/mindoc/conf"
)

type InfactConfig struct {
	Config OoConfig `json:"config"`
	Token  string   `json:"token"`
}

// onlyoffice api配置
type OoConfig struct {
	/**
	 * 类型 embedded/desktop
	 * 默认为desktop
	 */
	Type string `json:"type"`
	/**
	 * 文档类型 word/cell/slide
	 * open a word document (.doc, .docm, .docx, .dot, .dotm, .dotx, .epub, .fodt, .htm, .html, .mht, .odt, .ott, .pdf, .rtf, .txt, .djvu, .xps)
	 * open a cell (.csv, .fods, .ods, .ots, .xls, .xlsm, .xlsx, .xlt, .xltm, .xltx)
	 * open a slide (.fodp, .odp, .otp, .pot, .potm, .potx, .pps, .ppsm, .ppsx, .ppt, .pptm, .pptx)
	 */
	DocumentType string `json:"documentType"`
	/**
	 * 自定义签名
	 */
	//Token string `json:"token" default:"123456"`
	/**
	 * 文档配置信息
	 */
	Document DocumentConfig `json:"document"`
	/**
	 * 编辑配置
	 */
	EditorConfig EditorConfig `json:"editorConfig"`

	/**
	 * 访问api 路径
	 */
	DocServiceApiUrl string `json:"docServiceApiUrl"`
}
type DocumentConfig struct {
	/**
	 * 文件类型 如 docx
	 * 只需要文件的扩展名
	 */
	FileType string `json:"fileType"`

	/**
	 * 文件名称
	 */
	Title string `json:"title"`

	/**
	 * 文件访问的url
	 */
	Url string `json:"url"`

	/**
	 * 定义用于服务的文档识别的唯一文档标识符。在发送已知密钥的情况下，文档将从高速缓存中取出。
	 * 每次编辑和保存文档时，必须重新生成密钥。
	 * 文档url可以用作键，但是没有特殊字符，长度限制为20个符号。
	 */
	Key string `json:"key"`

	/**
	 * 文件作者信息
	 */
	Info DocumentInfo `json:"info"`

	/**
	 * 权限
	 */
	Permissions DocumentPermission `json:"permissions"`
}
type EditorConfig struct {
	/**
	 * 回调地址
	 * 指定文档存储服务的绝对URL（必须由在自己的服务器上使用ONLYOFFICE Document Server的软件集成器实现）。
	 * 必填*
	 */
	CallbackUrl string `json:"callbackUrl"`

	/**
	 * 定义文档创建后的绝对URL，并在创建后可用。如果没有指定，将没有创建按钮
	 */
	CreateUrl string `json:"createUrl"`

	/**
	 * 语言 "zh-CN"
	 * 定义编辑器接口语言（如果存在除英语之外的其他语言）。
	 * 使用两个字母（DE、RU、IT等）或四个字母（EN美国、FR FR等）来设置语言代码。不填默认值是"EN US"。
	 */
	Lang string `json:"lang"`

	/**
	 * 定义编辑器打开模式。
	 * 可以是打开用于查看的文档的视图，也可以是在允许对文档数据应用更改的编辑模式下打开文档的编辑。
	 * 默认值是edit
	 * view 视图
	 * edit 编辑
	 */
	Mode string `json:"mode"`

	/**
	 * 最近打开历史
	 */
	Recent []FileRecent `json:"recent"`

	/**
	 * 用户信息
	 */
	User FileUser `json:"user"`

	/**
	 * 自定义信息
	 */
	Customization FileCustomization `json:"customization"`

	/**
	 * 共同编辑
	 */
	CoEditing string `json:"coEditing"`

	Embedded FileEmbedded `json:"embedded"`

	/**
	 * 请查看官方文档
	 * 自定义插件
	 * https://api.onlyoffice.com/editors/config/editor/plugins
	 */
	Plugins Plugins `json:"plugins"`
}
type FileEmbedded struct {
	/**
	 * 文件url
	 * "https://example.com/embedded?doc=exampledocument1.docx"
	 */
	EmbedUrl string `json:"embedUrl"`

	/**
	 * "https://example.com/embedded?doc=exampledocument1.docx#fullscreen"
	 */
	FullscreenUrl string `json:"fullscreenUrl"`

	/**
	 * 保存的url
	 * "https://example.com/download?doc=exampledocument1.docx"
	 */
	SaveUrl string `json:"saveUrl"`

	/**
	 * "https://example.com/view?doc=exampledocument1.docx"
	 */
	ShareUrl string `json:"shareUrl"`

	/**
	 * 定义嵌入式浏览器工具栏的位置，可以是顶部或底部
	 * 默认top
	 * bottom/top
	 */
	ToolbarDocked string `json:"toolbarDocked"`
}
type Plugins struct {
	/**
	 * 插件的 guid  asc.{4FF5B2DB-BDDA-CC2A-5A36-0087719EB455}
	 */
	Autostart []string `json:"autostart"`
	/**
	 * 插件地址  服务器地址+guid+/config.json  这的guid没有前缀 {4FF5B2DB-BDDA-CC2A-5A36-0087719EB455}
	 */
	PluginsData []string `json:"pluginsData"`
}
type FileCustomization struct {

	/**
	 * {
	 * "request": true,
	 * "label": "Guest"
	 * },
	 * 添加对匿名名称的请求：
	 * 请求 - 定义请求是否发送。 默认值为 true，类型：布尔值，
	 * 标签 - 添加到用户名的后缀。 默认值为来宾，类型：字符串，
	 */
	Anonymous Anonymous `json:"anonymous" default:"true" `

	/**
	 * 定义是否启用或禁用“自动保存”菜单选项。
	 * 如果设置为false，则只能选择Strict协同编辑模式，因为Fast在没有自动保存的情况下无法工作。
	 * 请注意，如果您在菜单中更改此选项，它将保存到您的浏览器localStorage。默认值为true
	 */
	Autosave bool `json:"autosave"  default:"true" `

	/**
	 * 定义是显示还是隐藏“注释”菜单按钮。
	 * 请注意，如果您隐藏了“评论”按钮，则相应的评论功能将仅供查看，添加和编辑评论将不可用。 默认值为true。
	 */
	Comments bool `json:"comments" default:"true" `
	/**
	 * 定义其他操作按钮是显示在编辑器窗口标题的上部徽标旁边 （false） 还是显示在工具栏 （true） 中，从而使标题更紧凑。 默认值为false。
	 */
	CompactHeader bool `json:"compactHeader" default:"true" `
	/**
	 * 定义显示的顶部工具栏类型是完整 （假） 还是紧凑 （真）。 默认值为false。
	 */
	CompactToolbar bool `json:"compactToolbar" default:"false" `
	/**
	 * 定义仅与 OOXML 格式兼容的功能的使用。 例如，不要对整个文档使用注释。 默认值为false。
	 */
	CompatibleFeatures bool `json:"compatibleFeatures"  default:"false" `
	/**
	 * "address": "My City, 123a-45",有权访问编辑或编辑作者的公司或个人的邮政地址
	 * "info": "Some additional information",有关您希望其他人认识的公司或个人的一些其他信息，
	 * "logo": "https://example.com/logo-big.png",图片徽标的路径
	 * "logoDark": "https://example.com/dark-logo-big.png",深色主题的图像徽标的路径
	 * "mail": "john@example.com",有权访问编辑者或编辑者的公司或个人的电子邮件地址
	 * "name": "John Smith and Co.",名称
	 * "phone": "123456789",电话
	 * "www": "example.com"以上公司或个人的家庭网站地址
	 */
	Customer Customer `json:"customer"`
	/**
	 * {"spellcheck":  true}
	 * 定义用户可以禁用或自定义的参数（如果可能）：
	 * 拼写检查 - 定义在加载编辑器时是自动打开还是关闭拼写检查器。 如果此参数是布尔值，则将其设置为初始拼写检查器值，并且不会隐藏拼写检查器设置。
	 * 默认值为 true，类型：对象或布尔值，
	 * {"spellcheck": {"mode": true}}
	 * 拼写检查模式 - 定义在加载编辑器时拼写检查器是自动打开还是关闭。
	 * 此参数仅适用于文档编辑器和演示文稿编辑器，类型：布尔值，
	 * 示例：true`json:"plugins"`
	 */
	Features Features `json:"features"`

	/**
	 * {"url": "https://example.com",
	 * "visible": true}
	 * 定义“反馈和支持”菜单按钮的设置。
	 * 可以是布尔值（仅显示或隐藏“反馈和支持”菜单按钮）或对象。
	 */
	Feedback Feedback `json:"feedback"`
	/**
	 * 在文档编辑服务中保存文档时（例如，单击“保存”按钮等），
	 * 将文件强制保存的请求添加到回调处理程序中。 默认值为false。
	 */
	Forcesave bool `json:"forcesave"  default:"true" `

	/**
	 * {
	 * "blank": true,
	 * "requestClose": false,
	 * "text": "Open file location",
	 * "url": "https://example.com"
	 * }
	 * 定义“打开文件位置”菜单按钮和右上角按钮的设置。 该对象具有以下参数：
	 * 空白 - 单击“打开文件位置”按钮时，在新的浏览器选项卡/窗口中打开网站（如果值设置为true）或当前选项卡（如果值设置为false）。
	 * 默认值为 true，类型：布尔值，
	 * <p>
	 * requestClose- 定义如果单击“打开文件位置”按钮，则会调用 events.onRequestClose事件，而不是打开浏览器选项卡或窗口。
	 * 默认值为假，类型：布尔值，
	 * <p>
	 * <p>
	 * text-将为“打开文件位置”菜单按钮和右上角按钮（即而不是“转到文档”）显示的文本，键入：字符串，
	 * <p>
	 * url-单击“打开文件位置”菜单按钮时将打开的网站地址的绝对URL，键入：字符串，
	 */
	Goback Goback `json:"goback"`
	/**
	 * 定义是显示还是隐藏“帮助”菜单按钮。 默认值为true。
	 */
	Help bool `json:"help"  default:"true" `
	/**
	 * 定义在首次加载时是显示还是隐藏注释面板。 默认值为false。此参数仅适用于演示文稿编辑器ppt。
	 */
	HideNotes bool `json:"hideNotes"  default:"false" `
	/**
	 * 定义首次加载时是显示还是隐藏右侧菜单。 默认值为false。
	 */
	HideRightMenu bool `json:"hideRightMenu"  default:"false" `
	/**
	 * 定义是显示还是隐藏编辑器标尺。 此参数可用于文档和演示文稿编辑器。文档编辑器的默认值为false，演示文稿的默认值为true。
	 */
	HideRulers bool `json:"hideRulers"  default:"true" `

	/**
	 * 定义将编辑器嵌入网页的模式。 嵌入值在加载编辑器框架时禁用滚动到编辑器框架，因为未捕获焦点。
	 */
	IntegrationMode string `json:"integrationMode"`

	/**
	 * "image": "https://example.com/logo.png",
	 * "imageDark": "https://example.com/dark-logo.png",
	 * "url": "https://www.onlyoffice.com"
	 * 更改编辑器标题左上角的图像文件。 建议的图像高度为 20 像素。 该对象具有以下参数：
	 * image- 用于在通用工作模式（即所有编辑器的查看和编辑模式下）或嵌入模式下显示的图像文件的路径（请参阅配置部分以了解如何定义嵌入式文档类型）。
	 * 图像必须具有以下大小：172x40，类型：字符串，
	 * <p>
	 * imageDark - 用于深色主题的图像文件的路径。 图像必须具有以下大小：172x40，类型：字符串，
	 * <p>
	 * url - 当有人点击徽标图像时将使用的绝对URL（可用于访问您的网站等）。
	 * 保留为空字符串或null以使徽标不可点击，键入：字符串，
	 */
	Logo Logo `json:"logo"`
	/**
	 * 定义在编辑器打开时是否自动运行文档宏。 默认值为true。false值对用户隐藏宏设置。
	 */
	Macros bool `json:"macros"  default:"true" `
	/**
	 * disable
	 * enable
	 * warn 默认
	 */
	MacrosMode string `json:"macrosMode"`
	/**
	 * 定义在注释中提及后描述事件的提示。
	 * 如果为 true，则提示指示用户将收到通知并访问文档。
	 * 如果为 false，则提示指示用户将仅收到提及通知。 默认值为true。
	 * 请注意，只有在设置了onRequestSendNotify事件的情况下，它才可用于注释。
	 */
	MentionShare bool `json:"mentionShare"  default:"true" `
	/**
	 * 定义插件是否启动并可用。 默认值为true。
	 */
	Plugins bool `json:"plugins"  default:"true" `

	/**
	 * "hideReviewDisplay": false, 定义在“协作”选项卡上显示还是隐藏“显示模式”按钮。 默认值为假，类型：布尔值，
	 * "hoverMode": false         定义审阅显示模式：通过悬停更改在工具提示中显示审阅 （true） 或通过单击更改（假）在气球中。 默认值为false。
	 */
	Review Review `json:"review"`
	/**
	 * 定义文档标题是在顶部工具栏上可见 （假） 还是隐藏 （真）。 默认值为false。
	 */
	ToolbarHideFileName bool `json:"toolbarHideFileName"  default:"false" `

	/**
	 * 定义顶部工具栏选项卡是清晰显示 （false） 还是仅突出显示以查看选择了哪个选项卡 （true）。 默认值为false。
	 */
	ToolbarNoTabs bool `json:"toolbarNoTabs"  default:"false" `

	/**
	 * 定义编辑器主题设置。 可以通过两种方式进行设置：
	 * 主题 ID - 用户通过其 ID 设置主题参数（主题-浅色、主题-经典-浅色、主题-深色、主题-对比度-深），
	 * 默认主题 - 将设置默认的深色或浅色主题值（默认深色，默认浅色）。 默认浅色主题为主题经典浅色。
	 * 第一个选项具有更高的优先级。除了可用的编辑器主题外，用户还可以为应用程序界面自定义自己的颜色主题。
	 * 颜色可以以十六进制或 RGBA 格式呈现。
	 */
	UiTheme string `json:"uiTheme"`

	/**
	 * 定义标尺和对话框中使用的度量单位。 可以采用以下值：
	 * 厘米，
	 * pt-点，
	 * 英寸 - 英寸。
	 * 默认值为厘米 （cm）。
	 */
	Unit string `json:"unit"`

	/**
	 * 定义以百分比度量的文档显示缩放值。 可以采用大于0 的值。
	 * 对于文本文档和演示文稿，可以将此参数设置为 -1（使文档适合页面选项）或 to-2（使文档页面宽度适合编辑器页面）。
	 * 默认值为100。
	 */
	Zoom int `json:"zoom"`

	SubmitForm bool `json:"submitForm"  default:"true" `
}

type Anonymous struct {
	Request bool `json:"request"  default:"true" `
}
type Features struct {
	Spellcheck bool `json:"spellcheck"  default:"true" `
}
type Customer struct {
	Address bool `json:"address"`
}
type Feedback struct {
	Url string `json:"url"`
}
type Goback struct {
	Blank bool `json:"blank"  default:"true" `
}
type Logo struct {
	Image string `json:"image"`
}
type Review struct {
	HideReviewDisplay bool `json:"hideReviewDisplay"  default:"true" `
}
type FileUser struct {
	/**
	 * 用户唯一标识
	 */
	Id string `json:"id"`

	/**
	 * 用户 全称
	 */
	Name string `json:"name"`

	/**
	 * 组
	 */
	Group []string `json:"group"`
}
type FileRecent struct {
	/**
	 * 文件夹
	 */
	Folder string `json:"folder"`

	/**
	 * 名称
	 */
	Title string `json:"title"`

	/**
	 * url 绝对路径
	 */
	Url string `json:"url"`
}
type DocumentInfo struct {
	/**
	 * 定义收藏夹图标的突出显示状态。 当用户单击该图标时，将调用onMetaChange事件
	 */
	Favorite string `json:"favorite"`

	/**
	 * 创建时间（格式化后数据）
	 */
	Owner string `json:"owner"`
	/**
	 * 创建时间（格式化后数据）
	 */
	Created string `json:"created"`

	/**
	 * 存储文件夹可以为空
	 */
	Folder string `json:"folder"`

	/**
	 * 分享
	 * Defines the settings which will allow to share the document with other users:
	 * permissions - the access rights for the user with the name above. Can be Full Access, Read Only or Deny Access
	 * type: string
	 * example: "Full Access"
	 * user - the name of the user the document will be shared with
	 * type: string
	 * example: "John Smith".
	 */
	SharingSettings []SharingSettings `json:"sharingSettings"`
}
type SharingSettings struct {
	IsLink      bool     `json:"isLink"  default:"false" ` //将用户图标更改为链接图标
	Permissions []string `json:"permissions"`              //完全访问，只读或拒绝访问  Full Access, Read Only , Deny Access
	User        string   `json:"user"`                     //共享文档的用户的名称
}
type DocumentPermission struct {
	/**
	 * 定义是否在文档中启用聊天功能。 如果聊天权限设置为true，将显示聊天菜单按钮。 默认值为true。
	 */
	Chat bool `json:"chat"`

	/**
	 * 定义文档是否可以被注释。
	 * 在注释权限设置为“true”的情况下，文档侧栏将包含“注释”菜单选项；
	 * 如果该模式参数设置为“编辑”，则文档注释仅可用于文档编辑器。
	 * 默认值与编辑参数的值一致。
	 * 默认为true
	 */
	Comment bool `json:"comment"`
	/***
	  "edit": ["Group2", ""],
	  "remove": [""],
	  "view": "" ,
	  定义用户可以编辑、删除和/或查看其注释的组。 该对象具有以下参数：
	  编辑 - 用户可以编辑其他用户的评论，类型：列表，例如：[“Group2”，“”];
	  删除 - 用户可以删除其他用户的评论，类型：列表，
	  示例：[];
	  查看 - 用户可以查看其他用户的评论，类型：列表，
	  [“”]值表示用户可以编辑/删除/查看不属于这些组的人员所做的评论（例如，如果在第三方编辑器中审阅文档）。
	  如果值为 []，则用户无法编辑/删除/查看任何组所做的注释。
	  如果编辑、删除和查看参数为“”或未指定，则用户可以查看/编辑/删除任何用户所做的评论。
	*/
	CommentGroups string `json:"commentGroups"`

	/**
	 * 定义是否可以将内容复制到剪贴板。
	 * 如果参数设置为false，则粘贴内容将仅在当前文档编辑器中可用。 默认值为true。
	 */
	Copy bool `json:"copy"`

	/**
	 * 定义用户是否只能删除其注释。 默认值为false。
	 */
	DeleteCommentAuthorOnly bool `json:"deleteCommentAuthorOnly"`
	/**
	 * 是否允许下载 默认为true
	 */
	Download bool `json:"download"`
	/**
	 * 是否允许编辑 默认true
	 */
	Edit bool `json:"edit"`

	/**
	 * 定义用户是否只能编辑其注释。 默认值为false。
	 */
	EditCommentAuthorOnly bool `json:"editCommentAuthorOnly"`

	/**
	 * 是否可以填写表单 默认true
	 * 如果编辑设置为“true”或审阅设置为“true”，则不考虑fillForms值，并且可以填写表单。
	 * 如果编辑设置为“false”，审阅设置为“false”，并且fillForms也设置为“true”，则用户只能在文档中填写表单。
	 * 如果编辑设置为“false”并且审阅设置为“false”并且fillForms设置为“true”则不考虑注释值，并且注释不可用。
	 * 仅表单填写模式目前仅适用于文档编辑器。
	 */
	FillForms bool `json:"fillForms"`
	/**
	 * 定义是否可以更改内容控件设置。 仅当mode参数设置为编辑时，内容控件修改才可用于文档编辑器。 默认值为true。
	 */
	ModifyContentControl bool `json:"modifyContentControl"`
	/**
	 * 定义筛选器是可以全局应用（true）影响所有其他用户，还是本地应用（false），即仅适用于当前用户。
	 * 仅当模式参数设置为编辑时，过滤器修改才可用于电子表格编辑器。 默认值为true。
	 */
	ModifyFilter bool `json:"modifyFilter"`
	/**
	 * 定义文档是否可以打印。 如果打印权限设置为“false”打印菜单选项将不在“文件”菜单中。 默认值为true。
	 */
	Print bool `json:"print"`
	/**
	 * 定义是显示工具栏上的“保护”选项卡和左侧菜单中的“保护”按钮（真）还是隐藏（假）。 默认值为true。
	 */
	Protect bool `json:"protect"`
	/**
	 * 定义文档是否可以被审核。
	 * 如果审阅权限设置为“true”，则文档状态栏将包含审阅菜单选项；
	 * 如果模式参数设置为编辑，则文档审阅将仅对文档编辑器可用。
	 * 默认值与编辑参数的值一致。
	 * 如果编辑设置为“true”并且审阅也设置为“true”，则用户将能够编辑文档，接受/拒绝所做的更改并自己切换到审阅模式。
	 * 如果编辑设置为“true”并且审阅设置为“false”，则用户将只能进行编辑。 如果编辑设置为“false”并且审阅设置为“true”，则文档将仅在审阅模式下可用。
	 * 默认：true
	 */
	Review bool `json:"review"`
	/**
	 * 定义用户可以接受/拒绝其更改的组。 [“”]值表示用户可以审阅不属于这些组的人员所做的更改（例如，如果在第三方编辑器中审阅文档）。
	 * 如果值为 []，则用户无法查看任何组所做的更改。 如果值为“”或未指定，则用户可以查看任何用户所做的更改。
	 */
	ReviewGroups []string `json:"reviewGroups"`
	/**
	 * 定义其信息显示在编辑器中的用户组：
	 * 用户名显示在编辑器标题的编辑用户列表中，
	 * 键入文本时，将显示用户光标和工具提示及其名称，
	 * 在严格协同编辑模式下锁定对象时，将显示用户名。
	 * [“组 1”， “”] 表示显示有关组 1 中的用户和不属于任何组的用户的信息。 [] 表示根本不显示任何用户信息。 未定义的或 “” 值表示显示有关所有用户的信息。
	 */
	UserInfoGroups []string `json:"userInfoGroups"`
}

//jwt加密
type ConfigClaims struct {
	Type             string         `json:"type"`
	DocumentType     string         `json:"documentType"`
	Document         DocumentConfig `json:"document"`
	EditorConfig     EditorConfig   `json:"editorConfig"`
	DocServiceApiUrl string         `json:"docServiceApiUrl"`
	jwt.RegisteredClaims
}

// 前端配置
func GetDocConfig(item *Document, book_identify string, mode string) *InfactConfig {
	cf := conf.GetOoConfig()
	// 文档配置
	document := new(DocumentConfig)
	document.FileType = "docx"
	document.Url = cf.ApiUrl + "/api/" + book_identify + "/file/" + strconv.Itoa(item.DocumentId)
	document.Key = strconv.Itoa(int(time.Now().Unix()))
	document.Permissions = DocumentPermission{
		Edit: true,
	}
	if mode == "view" {
		document.Permissions.Edit = false
	}
	// 编辑器配置
	editCong := new(EditorConfig)
	editCong.Mode = mode
	editCong.CallbackUrl = cf.CallbackUrl + "?id=" + strconv.Itoa(item.DocumentId)
	editCong.Lang = "zh-CN"
	editCong.Customization = FileCustomization{
		Forcesave: true,
	}
	c := &OoConfig{
		Type:             "desktop",
		DocServiceApiUrl: cf.DocumentServer + cf.DocApiUrl,
		DocumentType:     "word",
		Document:         *document,
		EditorConfig:     *editCong,
	}
	c2 := &InfactConfig{
		Config: *c,
		Token:  "",
	}
	configClaims := ConfigClaims{
		Type:             "desktop",
		DocServiceApiUrl: cf.DocumentServer + cf.DocApiUrl,
		DocumentType:     "word",
		Document:         *document,
		EditorConfig:     *editCong,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, configClaims)

	// 2 把token加密
	mySigningKey := []byte(cf.Secret)
	ss, _ := token.SignedString(mySigningKey)
	c2.Token = ss
	return c2
}

// NewWithClaims creates a new Token with the specified signing method and claims.
func NewWithClaims(method jwt.SigningMethod, claims jwt.Claims) *jwt.Token {
	return &jwt.Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": method.Alg(),
		},
		Claims: claims,
		Method: method,
	}
}
func TestHs256(t *testing.T) {
	type User struct {
		Id   int64
		Name string
	}
	type UserClaims struct {
		User User
		jwt.RegisteredClaims
	}
	// 1 jwt.NewWithClaims生成token
	user := User{
		Id:   101,
		Name: "hisheng",
	}
	userClaims := UserClaims{
		User:             user,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	// 2 把token加密
	mySigningKey := []byte("ushjlwmwnwht")
	ss, err := token.SignedString(mySigningKey)
	t.Log(ss, err)
}
