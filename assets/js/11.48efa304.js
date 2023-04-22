(window.webpackJsonp=window.webpackJsonp||[]).push([[11],{289:function(s,t,a){"use strict";a.r(t);var n=a(10),r=Object(n.a)({},(function(){var s=this,t=s._self._c;return t("ContentSlotsDistributor",{attrs:{"slot-key":s.$parent.slotKey}},[t("h1",{attrs:{id:"grpc-restful"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#grpc-restful"}},[s._v("#")]),s._v(" grpc-restful")]),s._v(" "),t("h2",{attrs:{id:"安装工具"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#安装工具"}},[s._v("#")]),s._v(" 安装工具")]),s._v(" "),t("div",{staticClass:"language- extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[s._v("go install github.com/zeromicro/go-zero/tools/goctl@latest\ngo install github.com/golang/protobuf/protoc-gen-go@v1.3.2\n")])])]),t("h2",{attrs:{id:"编写-proto"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#编写-proto"}},[s._v("#")]),s._v(" 编写 proto")]),s._v(" "),t("div",{staticClass:"language-protobuf extra-class"},[t("pre",{pre:!0,attrs:{class:"language-protobuf"}},[t("code",[t("span",{pre:!0,attrs:{class:"token keyword"}},[s._v("syntax")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token string"}},[s._v('"proto3"')]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(";")]),s._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[s._v("option")]),s._v(" go_package "),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token string"}},[s._v('"./userpb"')]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(";")]),s._v("\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[s._v("package")]),s._v(" user"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(";")]),s._v("\n\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[s._v("import")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token string"}},[s._v('"google/api/annotations.proto"')]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(";")]),s._v("\n\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[s._v("message")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token class-name"}},[s._v("AddUserReq")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n      "),t("span",{pre:!0,attrs:{class:"token builtin"}},[s._v("string")]),s._v(" name "),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token number"}},[s._v("1")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(";")]),s._v("\n      "),t("span",{pre:!0,attrs:{class:"token builtin"}},[s._v("int32")]),s._v(" age "),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token number"}},[s._v("2")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(";")]),s._v("\n"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[s._v("message")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token class-name"}},[s._v("AddUserResp")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n      "),t("span",{pre:!0,attrs:{class:"token builtin"}},[s._v("int32")]),s._v(" id "),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token number"}},[s._v("1")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(";")]),s._v("\n"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n\n"),t("span",{pre:!0,attrs:{class:"token keyword"}},[s._v("service")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token class-name"}},[s._v("user")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n      "),t("span",{pre:!0,attrs:{class:"token keyword"}},[s._v("rpc")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token function"}},[s._v("Add")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("(")]),t("span",{pre:!0,attrs:{class:"token class-name"}},[s._v("AddUserReq")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(")")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token keyword"}},[s._v("returns")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("(")]),t("span",{pre:!0,attrs:{class:"token class-name"}},[s._v("AddUserResp")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(")")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n            "),t("span",{pre:!0,attrs:{class:"token keyword"}},[s._v("option")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("(")]),s._v("google"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(".")]),s._v("api"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(".")]),s._v("http"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(")")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n                  get"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(":")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token string"}},[s._v('"/api/v1.0/user/add"')]),s._v("\n            "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(";")]),s._v("\n      "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(";")]),s._v("\n"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n")])])]),t("h2",{attrs:{id:"编写配置文件"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#编写配置文件"}},[s._v("#")]),s._v(" 编写配置文件")]),s._v(" "),t("ul",[t("li",[s._v("modsdk.yaml")])]),s._v(" "),t("div",{staticClass:"language-yaml extra-class"},[t("pre",{pre:!0,attrs:{class:"language-yaml"}},[t("code",[t("span",{pre:!0,attrs:{class:"token key atrule"}},[s._v("scopeVersion")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(":")]),s._v(" userv1\n"),t("span",{pre:!0,attrs:{class:"token key atrule"}},[s._v("goModule")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(":")]),s._v(" github.com/jaronnie/autosdk\n"),t("span",{pre:!0,attrs:{class:"token key atrule"}},[s._v("goVersion")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(":")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token number"}},[s._v("1.18")]),s._v("\n")])])]),t("ul",[t("li",[s._v("pkgsdk.yaml")])]),s._v(" "),t("div",{staticClass:"language-shell extra-class"},[t("pre",{pre:!0,attrs:{class:"language-shell"}},[t("code",[s._v("scopeVersion: userv1\nsdkDir: pkgsdk\n")])])]),t("h2",{attrs:{id:"生成-gosdk"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#生成-gosdk"}},[s._v("#")]),s._v(" 生成 gosdk")]),s._v(" "),t("div",{staticClass:"language-shell extra-class"},[t("pre",{pre:!0,attrs:{class:"language-shell"}},[t("code",[t("span",{pre:!0,attrs:{class:"token function"}},[s._v("git")]),s._v(" clone https://github.com/jaronnie/protoc-gen-grpc-gateway-gosdk.git\n"),t("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("cd")]),s._v(" protoc-gen-grpc-gateway-gosdk\ntask "),t("span",{pre:!0,attrs:{class:"token function"}},[s._v("install")]),s._v("\n\n"),t("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("cd")]),s._v(" examples/grpc-restful\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[s._v("# 生成的 gosdk 直接在服务端")]),s._v("\n"),t("span",{pre:!0,attrs:{class:"token function"}},[s._v("mkdir")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("-p")]),s._v(" pkgsdk/pb\nprotoc -I./proto "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--go_out")]),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v("./pkgsdk/pb --grpc-gateway-gosdk_out"),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v("logtostderr"),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v("true,v"),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),t("span",{pre:!0,attrs:{class:"token number"}},[s._v("1")]),s._v(",env_file"),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v("etc/pkgsdk.yaml:pkgsdk proto/user.proto\n\n"),t("span",{pre:!0,attrs:{class:"token comment"}},[s._v("# 生成的 gosdk 独立 module")]),s._v("\n"),t("span",{pre:!0,attrs:{class:"token function"}},[s._v("mkdir")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("-p")]),s._v(" modsdk/pb\nprotoc -I./proto "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[s._v("--go_out")]),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v("./modsdk/pb --grpc-gateway-gosdk_out"),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v("logtostderr"),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v("true,v"),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),t("span",{pre:!0,attrs:{class:"token number"}},[s._v("1")]),s._v(",env_file"),t("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v("etc/modsdk.yaml:modsdk proto/user.proto\n")])])])])}),[],!1,null,null,null);t.default=r.exports}}]);