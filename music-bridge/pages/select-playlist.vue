<template>
    <div class="container">
      <div class="full-height">
        <div class="row  align-items-top">
            <div class="col">
                <div class="row">
                    <div class="col">
                        <h1 class="text-center">Transfer your Playlists</h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <p class="text-center lead">Pick the playlist</p>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <b-form-select v-model="selected"  :options="playlistTitles" @change="selectPlaylist">
                             <template slot="first">
                                    <option :value="false" disabled>-- Please select a playlist--</option>
                                  </template>
                            </b-form-select>

                            <div v-if="selected !== false">
                             <b-list-group>
                              <b-list-group-item v-for="(song, index) in computedSongs" v-bind:key="index">
                                {{ song.artist}} - {{song.name}}
                                <b-button @click='removeSong(index)' variant='danger'>Remove Song</b-button>
                              </b-list-group-item>
                              </b-list-group>
                              <b-button @click="commitSongs">Loooks good, next step</b-button>
                            </div>
                    </div>
                </div>
            </div>
        </div>
        </div>
    </div>
</template>

<script>
export default {
  mounted() {
    this.$apollo.addSmartQuery("playlists", {
      query: require("~/graphql/SpotifyPlaylists.gql"),
      update: ({ spotifyPlaylists }) => spotifyPlaylists,
      variables: {
        accessToken: this.$store.state.spotify.accessToken
      }
    });
  },
  methods: {
    removeSong(index) {
      this.selectedPlaylist.tracks.splice(index, 1);
    },
    async commitSongs() {
      this.$store.dispatch("core/storePlaylist", this.selectedPlaylist);
      await this.$store.commit("core/incrementStep");
      const redirectUrl = this.$store.getters["core/stepUrl"];
      console.log(redirectUrl);
      this.$router.push({ path: redirectUrl });
    },
    selectPlaylist(id) {
      const dataCopy = { ...this.playlists[id] };
      dataCopy.tracks = [...dataCopy.tracks];
      this.selectedPlaylist = dataCopy;
    }
  },
  computed: {
    computedSongs() {
      return this.selectedPlaylist.tracks;
    },
    playlistTitles() {
      if (!this.playlists.length) return [];

      return this.playlists.map((item, index) => ({
        value: index,
        text: `${item.name} (${item.tracks.length})`
      }));
    }
  },
  middleware: "authenticated",
  data() {
    return {
      selected: false,
      selectedPlaylist: {},
      playlists: []
    };
  }
};
</script>

<style>
</style>
