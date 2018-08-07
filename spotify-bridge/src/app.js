if (process.env.NODE_ENV !== 'production') {
  require('dotenv').config();
}
const express = require('express');
const bodyParser = require('body-parser');

const {
  accessTokenMiddleware,
  auth,
  authCallback,
  getAuthUrl,
  playlists,
  withApiProvider,
  withCustomRedirect,
} = require('./handlers/spotify');

const app = express();

app.set('trust proxy', true);
app.use(bodyParser.json());

app.get('/', (req, res) => res.send('Hello, World 🎉'));
app.get('/playlists', [withApiProvider, accessTokenMiddleware], playlists);
app.post('/auth-url', [withApiProvider, withCustomRedirect], getAuthUrl);
app.get('/auth', [withApiProvider, withCustomRedirect], auth);
app.get('/auth-callback', [withApiProvider, withCustomRedirect], authCallback);

app.listen(process.env.PORT, () => {
  console.log(`App listening on port ${process.env.PORT}`);
});
