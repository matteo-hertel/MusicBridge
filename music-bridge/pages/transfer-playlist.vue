<template>
    <div class="container">
      <div class="full-height">
        <div class="row  align-items-top">
            <div class="col">
                <div class="row">
                    <div class="col">
                        <h1 class="text-center">You're all set!</h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <p class="text-center lead">Review the info display down below, we're ready to create the playlist!!</p>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                      <b-container >
                          <p class="lead">Playlist Info</p>
                          <b-row>
                              <b-col>Name: </b-col>
                              <b-col>{{playlist.name}}</b-col>
                          </b-row>
                          <b-row>
                              <b-col>Public: </b-col>
                              <b-col>{{playlist.public}}</b-col>
                          </b-row>
                          <b-row>
                              <b-col>Tracks: </b-col>
                              <b-col>{{transferableSongs.length}}</b-col>
                          </b-row>
                      </b-container>

                      <b-container>
                          <b-row>
                              <b-col>
                                   <b-button @click="createPlaylist">Create Playlist</b-button>
                              </b-col>
                              <b-col v-if="playlistId">
                                <b-button :href="playlistUrl">View Your Playlist</b-button>
                              </b-col>
                              <b-col v-if="playlistId">
                                <b-button @click="updatePlaylist">Hydrate Your Playlist</b-button>

{{completed}} / {{this.transferableSongs.length}}
                              </b-col>
                          </b-row>
                      </b-container>
                    </div>
                </div>
            </div>
        </div>
        </div>
    </div>
</template>

<script>
export default {
  methods: {
    async createPlaylist() {
      return this.$apollo
        .query({
          query: require("~/graphql/CreatePlaylist.gql"),
          fetchPolicy: "network-only",
          variables: {
            title: this.playlist.name,
            privacyStatus: this.playlist.public ? "public" : "private",
            accessToken: this.$store.state.youtube.accessToken
          }
        })
        .then(({ data: { youtubeCreatePlaylist } }) => {
          this.$store.dispatch(
            "core/storeCreatedPlaylistId",
            youtubeCreatePlaylist.id
          );
        })
        .catch(() => {
          this.$store.dispatch(
            "setGlobalError",
            "An error occurred while creating the youtube playlist, most likely API rating limit 😞"
          );
        });
    },

    async addToPlaylist(playlistId, videoId) {
      return this.$apollo.query({
        query: require("~/graphql/AddToPlaylist.gql"),
        fetchPolicy: "network-only",
        variables: {
          playlistId,
          videoId,
          accessToken: this.$store.state.youtube.accessToken
        }
      });
    },
    async updatePlaylist() {
      for (const { videoId, title } of this.transferableSongs) {
        if (!videoId) continue;
        try {
          await this.addToPlaylist(this.playlistId, videoId);
          this.completed = ++this.completed;
        } catch (exc) {
          console.log(exc);
          return this.$store.dispatch(
            "setGlobalError",
            `An error occurred while adding the song "${title}"`
          );
        }
      }
    }
  },
  computed: {
    playlistUrl() {
      return `https://www.youtube.com/playlist?list=${this.playlistId}`;
    },
    playlistId() {
      return this.$store.state.core.createdPlaylistId;
    },
    playlist() {
      return this.$store.state.core.playlist || {};
    },
    transferableSongs() {
      return this.$store.state.core.songs.filter(song => song);
    }
  },
  middleware: "authenticated",
  data() {
    return {
      completed: 0
    };
  }
};
</script>

<style>
</style>
