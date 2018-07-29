module.exports = {
  /*
  ** Headers of the page
  */
  head: {
    title: "Music Bridge",
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { hid: "description", name: "description", content: "It talks to Humans" }
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
  loading: { color: "#3B8070" },
  /*
  ** Build configuration
  */
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
