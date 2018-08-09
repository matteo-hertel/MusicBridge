import gql from "graphql-tag";
export const state = () => ({
  authUrl: "",
  accessToken: ""
});

export const mutations = {
  storeUrl(state, url) {
    state.authUrl = url;
  },
  storeAccessToken(state, accessToken) {
    state.accessToken = accessToken;
  }
};

export const actions = {
  async fetchAuthUrl(context, payload) {
    const client = this.app.apolloProvider.defaultClient;
    const url = await client.query({
      query: gql`
        {
          youtubeAuthUrl(redirect: "http://localhost:3000/youtube-callback") {
            url
          }
        }
      `
    });
    context.commit("storeUrl", url.data.youtubeAuthUrl.url);
  },
  async storeAccessToken(context, accessCode) {
    const client = this.app.apolloProvider.defaultClient;
    const { data } = await client.query({
      query: gql`
      {
          youtubeAuth(
            code: "${accessCode}"
            redirect: "http://localhost:3000/youtube-callback"
          ) {
            accessToken
          }
       }
      `,
      variables: {
        accessCode
      }
    });
    context.commit("storeAccessToken", data.youtubeAuth.accessToken);
  },
  async getAccessTokenFromUrl(context, payload) {
    context.dispatch("storeAccessToken", payload);
  }
};
