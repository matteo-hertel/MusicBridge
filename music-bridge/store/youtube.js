import gql from "graphql-tag";
export const state = () => ({
  authUrl: "",
  accessToken: "",
  accessCode: ""
});

export const mutations = {
  storeUrl(state, url) {
    state.authUrl = url;
  },
  storeAccessToken(state, accessToken) {
    state.accessToken = accessToken;
  },
  storeAccessCode(state, accessCode) {
    state.accessCode = accessCode;
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
    if (!accessCode || context.state.accessCode === accessCode) {
      return;
    }
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
    context.commit("storeAccessCode", accessCode);
  },
  async getAccessTokenFromUrl(context, payload) {
    context.dispatch("storeAccessToken", payload);
  }
};
