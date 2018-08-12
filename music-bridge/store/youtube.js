import gql from "graphql-tag";
const defaultState = {
  authUrl: "",
  accessToken: "",
  accessCode: "",
  expiry: null
};

export const state = () => defaultState;

export const getters = {
  isExpired: state => {
    if (!state.expiry) return false;

    if (new Date() > new Date(state.expiry)) {
      return true;
    }
  }
};
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
            accessToken,
              expiry
          }
       }
      `,
      variables: {
        accessCode
      }
    });
    context.commit("storeAuthInfo", { ...data.youtubeAuth, accessCode });
  },
  async getAccessTokenFromUrl(context, payload) {
    context.dispatch("storeAccessToken", payload);
  }
};
