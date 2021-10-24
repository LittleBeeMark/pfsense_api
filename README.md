# pfsense_api
这是我使用 Pfsense 的 [API文档]（https://opensourcelibs.com/lib/pfsense-api） 及未实现 API 的使用页面表单请求对接实现的 SDK
主要使用它来实现 haproxy 模块的自动化证书更新的功能
代码主要实现了页面表单登陆获取cookie, haproxy 模块的证书更新，证书模块的证书导入等功能的对接。
由于 haproxy 模块是插件形式，没有 API 所以采用模仿手动调用方式的表单请求对接，进行了 cookie 认证，crsf 防止重复攻击的破解等认证操作。
