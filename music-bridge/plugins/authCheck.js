export default ({ store, redirect }) => {
  setInterval(() => {
    if (
      store.getters["spotify/isExpired"] ||
      store.getters["youtube/isExpired"]
    ) {
      store.dispatch("logout");
      redirect(302, { path: "/logged-out" });
    }
  }, 10 * 1000);
};
