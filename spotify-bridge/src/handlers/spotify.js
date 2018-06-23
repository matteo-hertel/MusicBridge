const {
  getAllPlaylists,
  getApiToken,
  getAuthorizeURL,
  getInitalUserInfo,
  setAccessToken,
  spotifyApi
} = require("./../lib/bridge/spotify");

function auth(req, res) {
  return res.redirect(getAuthorizeURL());
}
async function authCallback(req, res) {
  try {
    const callbackToken = req.query.code;
    await getApiToken(callbackToken);

    res.json(await getInitalUserInfo());
  } catch (exc) {
    console.error(exc);
    res.sendStatus(500);
  }
}

async function accessTokenMiddleware(req, res, next) {
  const apiToken = req.get("X-Spotify-Token");
  if (!apiToken) {
    return res.sendStatus("401");
  }
  setAccessToken(apiToken);
  next();
}
async function playlists(req, res) {
  try {
    const playlists = await getAllPlaylists();
    res.json(playlists);
  } catch (exc) {
    console.error(exc);
    res.sendStatus(500);
  }
}

module.exports = {
  accessTokenMiddleware,
  auth,
  authCallback,
  playlists
};
