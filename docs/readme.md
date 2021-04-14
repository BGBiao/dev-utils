## 相关文档


背景:通常在企业内部开发过程中经常需要进行自转换一些字符串到对应的数据结构中，来进行简单的数据排查以及结构验证，而这些服务在互联网上也均有相关的服务，比如[tool.lu](https://tool.lu/)、[tool.oschina.net](http://tool.oschina.net/)、[在线正则](https://regexper.com/) 、[toolnb](https://www.toolnb.com/) . 但是会存在一个问题，比如随着企业内部的权限收缩，越来越多的企业会直接屏蔽访问外网的权限，另外由于一些网络质量的原因，通常访问对应的工具网站会很慢。因此，该项目旨在构建一个内网常用的工具集合。



- 支持常见的Basic Auth认证
- 正则表达式解析 (regexp)
- 时间戳解析 (time|golang 和普通时间戳)
- 特殊字符串解析，unicode 解析，ascii 编解码 (strconv)
- base64 编解码 (base64)
- url 编解码 (url/encode)
- json 编解码 (如何规格化解析啊？)
- crontab 规则解析 (?)
- 进制转换 (strconv)
- MD5,hex 加解密 (几种加解密方式)
- uuid 生成器

