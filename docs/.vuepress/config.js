module.exports = {
  head: [
    [
        'link', // 设置 favicon.ico
        { rel: 'icon', href: 'favicon.ico' }
    ]
],
  base: '/',
  locales: {
    // 键名是该语言所属的子路径
    // 作为特例，默认语言可以使用 '/' 作为其路径。
    '/': {
      lang: 'zh-CN', // 将会被设置为 <html> 的 lang 属性
      title: 'protoc-gen-grpc-gateway-go',
      description: 'protoc-gen-grpc-gateway-go docs'
    },
  },
  plugins: [
    '@vuepress/back-to-top',
    '@vuepress/last-updated',
  ],
  themeConfig: {
    docsRepo: 'https://github.com/jaronnie/protoc-gen-grpc-gateway-go',
    docsBranch: 'main',
    docsDir: 'docs',
    editLinks: true, // 底部增加编辑此页
    lastUpdated: 'Last Updated', // 最近更新时间
    logo: 'https://oss.jaronnie.com/logo.jpeg',
    locales: {
      '/': {
        // 多语言下拉菜单的标题
        selectText: '选择语言',
        // 该语言在下拉菜单中的标签
        label: '简体中文',
        // Service Worker 的配置
        serviceWorker: {
          updatePopup: {
            message: "发现新内容可用.",
            buttonText: "刷新"
          } 
        },
        // 当前 locale 的 algolia docsearch 选项
        algolia: {},
        nav: [
          { text: 'Github', link: 'https://github.com/jaronnie/protoc-gen-grpc-gateway-go' },
        ],
        sidebar: {
          '/guide/': [
          {
            title: '1. Introduction',
            collapsable: false,
            children: [
              { title: '1.1 什么是 protoc-gen-grpc-gateway-go?', path: '/guide/introduction/what-is-protoc-gen-grpc-gateway-go' },
              { title: '1.2 它解决了什么问题?', path: '/guide/introduction/what-problems-can-be-solved' },
            ]
          },
          {
            title: '2. Quick start',
            collapsable: false,
            children: [
              { title: '2.1 下载 protoc-gen-grpc-gateway-go', path: '/guide/quickstart/install' },
            ]
          },
          ],
        }
      },
    }
  },
}
