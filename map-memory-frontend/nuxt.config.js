const webpack = require('webpack')
module.exports = {
  /*
  ** Headers of the page
  */
  head: {
    title: '地图记忆-用空间留住时间',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'Keywords', name: 'Keywords', content: '地图记忆-用空间留住时间' },
      { hid: 'Description', name: 'Description', content: '基于地图的笔记本,用地理空间关联的方式提高记忆或笔记效率'}
    ],
    script: [
      { src: 'https://cdnjs.cloudflare.com/ajax/libs/jquery/3.1.1/jquery.min.js' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      {
        rel: 'stylesheet',
        href: 'https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css'
      }
    ]
  },
  modules: [
    'bootstrap-vue/nuxt'
  ],
  plugins: ['~/plugins/froala.js','~/plugins/v-lazy-img',{ src: "~/plugins/vue-verify", ssr: false },'~/plugins/element-ui',{ src: "~/plugins/vue-baidu-map.js", ssr: false },{ src: "~/plugins/wysiwyg.js", ssr: false },{ src: "~/plugins/vue-clipboard.js", ssr: false }],
  css: [
      'vue-wysiwyg/dist/vueWysiwyg.css',
      '~/node_modules/froala-editor/css/froala_editor.pkgd.min.css',
      '~/node_modules/froala-editor/css/froala_style.min.css'
  ],
  /*
  ** Customize the progress bar color
  */
  loading: { color: '#3B8070' },
  /*
  ** Build configuration
  */
  build: {
    plugins: [
      new webpack.ProvidePlugin({
        '$': 'jquery',
        jQuery: 'jquery'
      })
    ],
    /*
    ** Run ESLint on save
    */
    extend (config, { isDev, isClient }) {
      if (isDev && isClient) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /(node_modules)/
        })
      }
    }
  }
}
