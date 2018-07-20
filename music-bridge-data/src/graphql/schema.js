const {makeExecutableSchema} = require('graphql-tools');
const resolvers = require('./resolvers');

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
}

type SpotifyAuthUrl{
url: String!
}
type YoutubeAuthUrl{
url: String!
}

type Query{
spotifyPlaylists(accessToken: String!): [SpotifyPlaylist],
spotifyAuthUrl: SpotifyAuthUrl!
spotifyAuth(code: String!): SpotifyUserInfo
youtubeAuthUrl: YoutubeAuthUrl!
youtubeCreatePlaylist(accessToken: String!, title: String!, description: String, privacyStatus: String): String!
}
`;
module.exports = makeExecutableSchema({
  typeDefs,
  resolvers,
});
