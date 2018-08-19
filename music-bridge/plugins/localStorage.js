import createPersistedState from "vuex-persistedstate";

export default ({ store }) => {
  createPersistedState({
    key: "music-brige",
    filter: mutation => {
      if (mutation.type.startsWith("_")) return false;
      return true;
    }
  })(store);
};
