const {spotifyApiProvider, makeSpotifyApi} = require('./../lib/bridge/spotify');

function auth(req, res) {
  const {getAuthorizeURL} = req.apiProvider;
  return res.redirect(getAuthorizeURL());
}
function getAuthUrl(req, res) {
  const {getAuthorizeURL} = req.apiProvider;
  const {body} = req;
  return res.json({url: getAuthorizeURL(body.redirect)});
}
async function authCallback(req, res) {
  const {getApiToken} = req.apiProvider;
  try {
    const callbackToken = req.query.code;
    const auth = await getApiToken(callbackToken);

    res.json(auth);
  } catch (exc) {
    handleErrors(exc, res);
  }
}

async function playlists(req, res) {
  const {getAllPlaylists} = req.apiProvider;
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
function withApiProvider(req, res, next) {
  req.apiProvider = spotifyApiProvider(makeSpotifyApi());
  next();
}
function withCustomRedirect(req, res, next) {
  if (req.query.redirect) {
    var redirect = req.query.redirect;
  } else {
    var {
      body: {redirect},
    } = req;
  }
  if (!redirect) {
    return next();
  }
  req.apiProvider.spotifyApi.setRedirectURI(redirect);
  next();
}

async function accessTokenMiddleware(req, res, next) {
  const {setAccessToken} = req.apiProvider;
  const apiToken = req.get('X-Spotify-Token');
  if (!apiToken) {
    return res.sendStatus('401');
  }
  setAccessToken(apiToken);
  next();
}

module.exports = {
  accessTokenMiddleware,
  auth,
  authCallback,
  getAuthUrl,
  playlists,
  withApiProvider,
  withCustomRedirect,
};
