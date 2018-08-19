const ERROR_DELAY = 6500;

export const state = () => ({
  globalError: ""
});
export const getters = {
  hasError: state => {
    return !!state.globalError;
  }
};
export const mutations = {
  _setError(state, message) {
    state.globalError = message;
  }
};
export const actions = {
  setGlobalError: (context, message) => {
    context.commit("_setError", message);
    setTimeout(() => {
      context.commit("_setError", "");
    }, ERROR_DELAY);
  },
  logout: context => {
    context.commit("spotify/wipeState");
    context.dispatch("spotify/fetchAuthUrl");
    context.commit("youtube/wipeState");
    context.dispatch("youtube/fetchAuthUrl");
  }
};
