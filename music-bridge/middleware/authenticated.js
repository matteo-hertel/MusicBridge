export default function({ store, redirect }) {
  // If the user is not authenticated
  if (process.server) return;
  if (
    !store.state.spotify.accessToken ||
    !store.state.youtube.accessToken ||
    (store.getters["spotify/isExpired"] || store.getters["youtube/isExpired"])
  ) {
    store.dispatch("logout");
    redirect(302, { path: "/logged-out" });
  }
}
