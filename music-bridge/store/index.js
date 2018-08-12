export const state = () => ({});
export const actions = {
  logout: context => {
    context.commit("spotify/wipeState");
    context.dispatch("spotify/fetchAuthUrl");
    context.commit("youtube/wipeState");
    context.dispatch("youtube/fetchAuthUrl");
  }
};
