import gql from "graphql-tag";
export const state = () => ({
    authUrl: "",
    accessToken: ""
});

export const mutations = {
    storeUrl(state, url) {
        state.authUrl = url;
    },
    async storeAccessToken(state, accessCode) {
        const client = this.app.apolloProvider.defaultClient;
        const accessToken = await client.query({
            query: gql`
        {
          spotifyPapoi {
            accessCode
          }
        }
      `
        });
    }
};

export const actions = {
    async fetchAuthUrl(context, payload) {
        const client = this.app.apolloProvider.defaultClient;
        const url = await client.query({
            query: gql`
        {
          spotifyAuthUrl {
            url
          }
        }
      `
        });
        context.commit("storeUrl", url.data.spotifyAuthUrl.url);
    },
    async getAccessTokenFromUrl(context, payload) {
        context.commit("storeAccessToken", payload);
    },
};
