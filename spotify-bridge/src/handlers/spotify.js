const {
  getAllPlaylists,
  getApiToken,
  getAuthorizeURL,
  getInitalUserInfo,
  setAccessToken,
  spotifyApi,
} = require('./../lib/bridge/spotify');

function auth(req, res) {
  return res.redirect(getAuthorizeURL());
}
function getAuthUrl(req, res) {
  const {body} = req;
  return res.json({url: getAuthorizeURL(body.redirect)});
}
async function authCallback(req, res) {
  try {
    const callbackToken = req.query.code;
    await getApiToken(callbackToken);

    res.json(await getInitalUserInfo());
  } catch (exc) {
    handleErrors(exc, res);
  }
}

async function accessTokenMiddleware(req, res, next) {
  const apiToken = req.get('X-Spotify-Token');
  if (!apiToken) {
    return res.sendStatus('401');
  }
  setAccessToken(apiToken);
  next();
}
async function playlists(req, res) {
  try {
    const playlists = await getAllPlaylists();
    res.json(playlists);
  } catch (exc) {
    handleErrors(exc, res);
  }
}
function handleErrors(exc, res) {
  console.error(exc);
  const errorCode = exc.statusCode || 500;
  res.status(errorCode).send(exc.message || '');
}
module.exports = {
  accessTokenMiddleware,
  auth,
  authCallback,
  getAuthUrl,
  playlists,
};
