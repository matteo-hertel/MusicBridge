const stepMapping = {
  0: "/select-playlist",
  1: "/select-best-match",
  2: "/transfer-playlist"
};

const defaultState = {
  step: 0,
  playlist: {},
  songs: [],
  createdPlaylistId: ""
};

export const state = () => defaultState;

export const getters = {
  stepUrl: state => stepMapping[state.step]
};

export const mutations = {
  decrementStep(state, step) {
    state.step--;
    if (state.step < 0) state.step = 0;
  },
  incrementStep(state, step) {
    state.step++;
  },
  setCreatedPlaylistId(state, id) {
    state.createdPlaylistId = id;
  },
  setSongs(state, songs) {
    state.songs = songs;
  },
  setPlaylist(state, playlist) {
    state.playlist = playlist;
  }
};
export const actions = {
  async storeCreatedPlaylistId(context, payload) {
    context.commit("setCreatedPlaylistId", payload);
  },
  async storeSongs(context, songs) {
    context.commit("setSongs", songs);
  },
  async storePlaylist(context, payload) {
    context.commit("setPlaylist", payload);
  }
};
