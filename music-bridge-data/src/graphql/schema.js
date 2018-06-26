const { makeExecutableSchema } = require("graphql-tools");
const resolvers = require("./resolvers");

const typeDefs = `
type UserInfo{
accessToken: String!,
display_name: Int
}

type Playlist{
id: ID!,
name: String!,
owner: String!,
public: Boolean,
tracks: [Track]

}
type Track{
name: String!,
artist: String!
}

type authUrl{
url: String!
}

type Query{
playlists(accessToken: String!): [Playlist],
authUrl: authUrl!
auth(code: String!): UserInfo
}
`;
module.exports = makeExecutableSchema({
  typeDefs,
  resolvers
});
