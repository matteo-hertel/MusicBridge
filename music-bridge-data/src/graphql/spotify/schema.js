const typeDefs = `
type SpotifyUserAuth{
  accessToken: String!,
  expiry: String!
}

type SpotifyPlaylist{
  id: ID!,
  name: String!,
  owner: String!,
  public: Boolean,
  tracks: [SpotifyTrack]
}

type SpotifyTrack{
  name: String!,
  artist: String!
}`;

const queryDefs = `
  spotifyPlaylists(accessToken: String!): [SpotifyPlaylist],
  spotifyAuthUrl(redirect: String): AuthUrl!
  spotifyAuth(code: String!,redirect: String): SpotifyUserAuth
`;
module.exports = {
  typeDefs,
  queryDefs,
};
