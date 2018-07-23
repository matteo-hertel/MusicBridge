const typeDefs = `
type SpotifyUserInfo{
  accessToken: String!,
  display_name: String
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
  spotifyAuthUrl: AuthUrl!
  spotifyAuth(code: String!): SpotifyUserInfo
`;
module.exports = {
  typeDefs,
  queryDefs,
};
