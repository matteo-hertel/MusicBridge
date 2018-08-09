export default ({ store }) => {
    store.dispatch("youtube/fetchAuthUrl");
    store.dispatch("spotify/fetchAuthUrl");
};
