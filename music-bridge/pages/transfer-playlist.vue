<template>
    <b-container>
      <div class="full-height">
          <b-row class="align-items-top">
              <b-col>
                <b-row>
                    <b-col>
                        <h1 class="text-center">You're all set!</h1>
                    </b-col>
                </b-row>
                <b-row>
                    <b-col>
                        <p class="text-center lead">Use the controls below to create your playlist and transfer your selected songs.</p>
                    </b-col>
                </b-row>
                <b-row>
                    <b-col>
                      <p class="lead">Playlist Info</p>
                      <b-row>
                          <b-col>
                              <p><strong>Name:</strong> {{ playlist.name }}</p>
                          </b-col>
                      </b-row>
                        <b-row>
                            <b-col>
                                <p><strong>Number of tracks:</strong> {{ transferableSongs.length }}</p>
                            </b-col>
                        </b-row>
                    </b-col>
                </b-row>
                  <b-row>
                      <b-col>
                          <b-button @click="createPlaylist">Create Playlist</b-button>
                      </b-col>
                      <b-col v-if="playlistId">
                              <b-button @click="updatePlaylist">Transfer Your Songs</b-button>
                      </b-col>
                  </b-row>
                  <b-row>
                      <b-col>
                          <b-progress height="100%" :max="this.transferableSongs.length">
                              <b-progress-bar :value="completed" show-value :label="progressLabel">
                              </b-progress-bar>
                          </b-progress>
                      </b-col>
                  </b-row>
              </b-col>

          </b-row>
        </div>
    </b-container>
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
    },
    progressLabel() {
      return `${this.completed} / ${this.transferableSongs.length}`
    },
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
