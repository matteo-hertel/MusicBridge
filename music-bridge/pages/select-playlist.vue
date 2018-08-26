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
                        <b-form-select :options="playlistTitles" @change="selectPlaylist">
                             <template slot="first">
                                     this slot appears above the options from 'options' prop
                                    <option :value="false" disabled>-- Please select an option --</option>
                                  </template>
                            </b-form-select>

                            <div v-if="selectedPlaylist !== false">
                             <b-list-group>
                              <b-list-group-item v-for="(song, index) in computedSongs" v-bind:key="index">
                                {{ song.artist}} - {{song.name}}
                                <b-button @click='removeSong(index)' variant='danger'>Remove Song</b-button>
                              </b-list-group-item>
                              </b-list-group>
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
      selectedPlaylist: {},
      playlists: []
    };
  }
};
</script>

<style>
</style>
