<template>
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
                <div class="col">
                    <p class="text-center lead">Check the songs</p>
                </div>
                <div class="col" v-if='chosenSongs.length'>
                    <p class="text-center lead">Confirm YourSelection</p>
                </div>
            </div>
            <div class="row">
                <div class="col">
                    <b-form-select :options="playlistTitles" v-model="selectedPlaylist" id="playlistSelect">
                         <template slot="first">
                                 this slot appears above the options from 'options' prop
                                <option :value="false" disabled>-- Please select an option --</option>
                              </template>
                        </b-form-select>
                        <p class="text-center" v-if="selectedPlaylist !== false">
                        <ul>
                          <li v-for="(song, index) in computedSongs" v-bind:key="index">
                            {{ song.artist}} - {{song.name}}
                          </li>
                        </ul>
                        </p>
                        <b-button @click='makeSearch' v-if="selectedPlaylist">
                            Transfer Playlist
                        </b-button>
                </div>
                <div class="col" >
                    <div role="tablist" class="shadow-lg">
                        <b-card v-for="(song, index) in searchResults" v-bind:key="index" no-body class="mb-1">
                            <b-card-header header-tag="header" class="p-1" role="tab">
                                <b-btn block  v-b-toggle="getAccordionID('accordion', index)" variant="info">{{ song.results[0].title }} - {{ song.results[0].artist }}</b-btn>
                            </b-card-header>
                        <div v-for="(video, i) in song.results" v-bind:key="i" no-body class="mb-1">
                          <LazyCollapse :url="getVideoUrl(video.videoId)" :id="getAccordionID('accordion', index)">
                              <b-button @click='useMe(index, i)'>Select this</b-button>
                          </LazyCollapse>
                        </div>
                       </b-card>
                    </div>
                </div>
                <div class="col" v-if='chosenSongs.length'>
                        <ul>
                          <li v-for="(song, index) in chosenSongs" v-bind:key="index">
                            {{ song.title}} - {{song.videoId}}
                          </li>
                        </ul>
                            <p class='center'>Transfered  {{this.transferCompleted}} / {{this.chosenSongs.length}}</p>
                        <b-button @click='transferPlaylist' >
                            Make it so!
                        </b-button>
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
