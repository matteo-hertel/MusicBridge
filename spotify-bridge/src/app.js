if (process.env.NODE_ENV !== "production") {
  require("dotenv").config();
}
const express = require("express");

const {
  accessTokenMiddleware,
  auth,
  authCallback,
  playlists
} = require("./handlers/spotify");

const app = express();

app.set("trust proxy", true);

app.get("/", (req, res) => res.send("Hello, World 🎉"));
app.get("/playlists", [accessTokenMiddleware], playlists);
app.get("/auth", auth);
app.get("/auth-callback", authCallback);

app.listen(process.env.PORT, () => {
  console.log(`App listening on port ${process.env.PORT}`);
});
