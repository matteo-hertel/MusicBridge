import gql from "graphql-tag";

const defaultState = {
  authUrl: "",
  accessToken: "",
  accessCode: "",
  expiry: null
};

export const state = () => defaultState;

export const mutations = {
  storeUrl(state, url) {
    state.authUrl = url;
  },
  storeAuthInfo(state, { accessToken, accessCode, expiry }) {
    state.accessToken = accessToken;
    state.accessCode = accessCode;
    state.expiry = expiry;
  },
  wipeState(state) {
    Object.keys(state).map(k => {
      state[k] = defaultState[k];
    });
  }
};
export const getters = {
  isExpired: state => {
    if (!state.expiry) return false;

    if (new Date() > new Date(state.expiry)) {
      return true;
    }
    return false;
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
            accessToken,
            expiry
          }
       }
      `,
      variables: {
        accessCode
      }
    });
    context.commit("storeAuthInfo", {
      ...data.spotifyAuth,
      accessCode
    });
  },
  async getAccessTokenFromUrl(context, payload) {
    context.dispatch("storeAccessToken", payload);
  }
};
