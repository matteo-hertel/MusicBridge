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
    const { data } = await client.query({
      query: gql`
        {
          spotifyAuthUrl(redirect: "http://localhost:3000/spotify-callback") {
            url
          }
        }
      `
    });
    context.commit("storeUrl", data.spotifyAuthUrl.url);
  },
  async storeAccessToken(context, accessCode) {
    if (!accessCode || context.state.accessCode === accessCode) {
      return;
    }
    const client = this.app.apolloProvider.defaultClient;
    const { data } = await client.query({
      query: gql`
      {
          spotifyAuth(
            code: "${accessCode}"
            redirect: "http://localhost:3000/spotify-callback"
          ) {
            accessToken
          }
       }
      `,
      variables: {
        accessCode
      }
    });
    context.commit("storeAccessToken", data.spotifyAuth.accessToken);
    context.commit("storeAccessCode", accessCode);
  },
  async getAccessTokenFromUrl(context, payload) {
    context.dispatch("storeAccessToken", payload);
  }
};
