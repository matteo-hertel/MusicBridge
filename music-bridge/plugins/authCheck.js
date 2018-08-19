export default ({ store, redirect }) => {
  const DELAY = 40 * 1000;
  autchCheck(DELAY);

  function autchCheck(delay) {
    if (
      store.getters["spotify/isExpired"] ||
      store.getters["youtube/isExpired"]
    ) {
      store.dispatch("logout");
      redirect(302, { path: "/logged-out" });
    }
    setInterval(autchCheck, delay);
  }
};
