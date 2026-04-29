目录名	作用说明
conf/	配置文件目录，存放项目的各类配置（如数据库连接、服务器端口、日志级别、第三方服务密钥等），常见的配置文件格式可能是 yaml/toml/json/ini 等。
tests/	测试用例目录，存放项目的单元测试、集成测试代码，遵循 Go 测试规范（测试文件以 _test.go 结尾），用于验证代码逻辑的正确性。
controllers/	控制器目录，属于 Web 开发的业务逻辑层，负责接收路由转发的请求、处理业务逻辑（如参数校验、调用服务、组装响应），是请求处理的核心环节。
views/	视图目录，存放前端页面模板（如 HTML 模板、模板引擎文件，如 Go 的 html/template 模板），若项目是服务端渲染（SSR）架构，页面渲染逻辑的模板文件会放在这里。
routers/	路由目录，负责定义 URL 路径与控制器方法的映射关系，管理请求的路由规则（如 GET/POST 方法、路由分组、中间件挂载等），是请求分发的核心。
static/	静态资源目录，存放前端静态文件，如 CSS 样式表、JavaScript 脚本、图片（jpg/png/svg）、字体文件等，供前端页面加载使用。




type Context struct {
	Input          *BeegoInput   //请求输入封装：封装HTTP请求的所有输入数据（GET/POST参数，Cookie，Header等）
	Output         *BeegoOutput  // 2. 响应输出封装：封装 HTTP 响应的所有输出操作（写 Body、设置状态码、Cookie、Header 等）
	Request        *http.Request // 3. 原生 HTTP 请求对象：Go 标准库 http.Request，是底层的请求原始数据
	ResponseWriter *Response     // 4. 原生 HTTP 响应写入器封装：Beego 对标准库 ResponseWriter 的增强实现
	_xsrfToken     string        // 5. XSRF 防跨站请求伪造令牌：用于安全校验，防止 CSRF 攻击
}

在Web开发中，session是一种用于跟踪用户在网站上的交互状态的机制。当用户访问一个启用了session功能的网站时，服务器会为该用户创建一个唯一的会话标识（session ID）。这个ID通常会通过cookie存储在用户的浏览器中，或者通过URL参数的形式传递。后续用户在该网站上的每次请求都会携带这个session ID，服务器根据这个ID来识别用户，并关联该用户的相关会话数据。

“hook function”即钩子函数，是Windows消息处理机制的一部分。通过设置“钩子”，应用程序能在系统级对所有消息、事件进行过滤，访问正常情况下无法访问的消息。每当特定消息发出，在未到达目的窗口前，钩子函数会先捕获该消息，它可以对消息进行加工处理、继续传递或强制结束传递。

RESTful 是什么？
用大白话讲：RESTful 是一套设计 API 接口的规范风格，专门用来规范前后端、服务之间怎么通过 HTTP 接口通信，让接口统一、好理解、好维护。
核心设计原则（最常用 5 条）
用 HTTP 动词表示操作
GET：查（获取资源）
POST：增（创建资源）
PUT：全量更新
PATCH：局部更新
DELETE：删除资源

1.Ctx  *context.Context
Ctx：变量名，约定俗成都叫 ctx，随便改名也行（比如 c、ctxReq）
*：指针，传的是地址，不是拷贝
context：Go 标准库包 import "context"
Context：包里的接口类型
合起来：定义了一个名叫 Ctx、类型为 context.Context 指针的变量
2. context.Context 是干嘛的？
一句话：Go 用来做「请求上下文管控」的神器专门干三件事：
传请求超时、取消信号（超时自动终止、手动取消请求）
跨函数、协程传递请求级变量（Token、用户 ID、追踪 ID）
串联整个调用链路（接口→service→dao→数据库，全程带同一个 ctx）

ctx *context.Context
这是请求上下文，包含了这次 HTTP 请求的所有信息：
浏览器发过来了什么
你要返回什么
请求头、响应头、状态码、Body 内容……

ctx.Output
它是上下文里专门管 “给浏览器返回数据” 的模块作用 = 响应输出器
它能干这些事：
返回字符串、JSON、HTML
设置 HTTP 状态码（200、404、500）
设置响应头
输出文件、下载等

ctx.Output.Body(...)
意思是：把内容写入 HTTP 响应体，返回给浏览器

常见 Output 用法（你以后一定会用到）go运行// 返回字符串
ctx.Output.Body([]byte("hello"))
// 返回 JSON（最常用）
ctx.Output.JSON(map[string]any{
    "code": 0,
    "msg":  "success",
})
// 返回 HTML
ctx.Output.HTML("<h1>你好</h1>")
// 设置状态码
ctx.Output.SetStatus(200)

1. func(ctx *context.Context) 是什么？
它不是普通自定义函数，是 Web 框架预留的路由回调函数格式。
ctx *context.Context 就代表：当前这一次 HTTP 请求的全部上下文。

2. 什么叫「请求上下文」？
一次浏览器访问服务器：
浏览器发过来的请求数据（URL、参数、请求头）
服务器要返回给浏览器的响应（你刚才问的 Output）
还有超时、取消、传用户信息等底层控制
全部都装在这一个 ctx 里面。

RPC = 远程过程调用，跨机器直接调函数，屏蔽网络细节，专门给后端微服务内部通信用的技术。