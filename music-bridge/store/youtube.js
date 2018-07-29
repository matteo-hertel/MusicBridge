import gql from "graphql-tag";
export const state = () => ({
  authUrl: ""
});

export const mutations = {
  storeUrl(state, url) {
    state.authUrl = url;
  }
};

export const actions = {
  async fetchAuthUrl(context, payload) {
    const client = this.app.apolloProvider.defaultClient;
    const sleep = () => new Promise(resolve => setTimeout(resolve, 3500));
    await sleep();
    const url = await client.query({
      query: gql`
        {
          youtubeAuthUrl {
            url
          }
        }
      `
    });
    context.commit("storeUrl", url.data.youtubeAuthUrl.url);
  }
};
