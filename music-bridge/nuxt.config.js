module.exports = {
  /*
  ** Headers of the page
  */
  head: {
    title: "Music Bridge",
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      {
        hid: "description",
        name: "description",
        content: "Convert Spotify playlists into YouTube Music playlists"
      }
    ],
    link: [
      { rel: "icon", type: "image/x-icon", href: "/favicon.ico" },
      {
        rel: "text/css",
        rel: "stylesheet",
        href:
          "https://stackpath.bootstrapcdn.com/bootswatch/4.1.2/flatly/bootstrap.min.css"
      },
      {
        rel: "text/css",
        rel: "stylesheet",
        href: "//unpkg.com/bootstrap-vue@latest/dist/bootstrap-vue.css"
      }
    ]
  },
  /*
  ** Customize the progress bar color
  */
  loading: false,
  /*
  ** Build configuration
  */
  plugins: [
    { src: "~/plugins/localStorage.js", ssr: false },
    { src: "~/plugins/storeInit.js", ssr: false },
    { src: "~/plugins/authCheck.js", ssr: false }
  ],
  modules: [["bootstrap-vue/nuxt", { css: false }], "@nuxtjs/apollo"],
  apollo: {
    clientConfigs: {
      default: {
        httpEndpoint: "http://localhost:3450/graphql",
        wsEndpoint: null,
        // LocalStorage token
        tokenName: "apollo-token", // optional
        // Enable Automatic Query persisting with Apollo Engine
        persisting: false, // Optional
        // Use websockets for everything (no HTTP)
        // You need to pass a `wsEndpoint` for this to work
        websocketsOnly: false // Optional
      }
    }
  },
  build: {
    vendor: ["vuex-persistedstate"],
    /*
    ** Run ESLint on save
    */
    extend(config, { isDev, isClient }) {
      if (isDev && isClient) {
        config.module.rules.push({
          enforce: "pre",
          test: /\.(js|vue)$/,
          loader: "eslint-loader",
          exclude: /(node_modules)/
        });
      }
    }
  }
};
