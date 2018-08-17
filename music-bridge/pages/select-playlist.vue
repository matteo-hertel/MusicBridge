<template>
    <div class="container">
        <div class="row full-height align-items-top">
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
                </div>
                <div class="row">
                    <div class="col">
<b-button @click="testQuery">Test Query</b-button>
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
                            <p class="text-center" v-if="selectedPlaylist">
                                Cool, make sure we found the right songs.<br>
                                You can remove the ones we got wrong.
                            </p>
                            <b-button>
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
<LazyCollapse :url="getVideoUrl(video.videoId)" :id="getAccordionID('accordion', index)"></LazyCollapse>
                            </div>
                            </b-card>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import LazyCollapse from "~/components/LazyCollapse";
export default {
  components: {
    LazyCollapse
  },
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
    getAccordionID(prefix, index) {
      return `${prefix}-${index}`;
    },
    getVideoUrl(videoId) {
      return `https://www.youtube.com/embed/${videoId}`;
    },
    testQuery() {
      this.$apollo
        .mutate({
          mutation: require("~/graphql/SearchSongs.gql"),
          variables: {
            songs: this.selectedSongs().map(({ artist, name }) => {
              return { artist, title: name };
            }),
            accessToken: this.$store.state.youtube.accessToken
          }
        })
        .then(
          ({ data: { youtubeSearchSongs } }) =>
            (this.searchResults = youtubeSearchSongs)
        );
    },
    selectedSongs() {
      if (!this.playlists.length) return [];
      const tracks = this.playlists.filter((_, index) => {
        return index === this.selectedPlaylist;
      })[0].tracks;
      return tracks;
    }
  },
  computed: {
    computedSongs() {
      return this.selectedSongs();
    },
    playlistTitles() {
      if (!this.playlists.length) return [];

      return this.playlists.map((item, index) => ({
        value: index,
        text: `${item.name} (${item.tracks.length})`
      }));
    }
  },
  data() {
    return {
      selectedPlaylist: false,
      searchResults: [],
      playlists: []
    };
  }
};
</script>

<style>
</style>
