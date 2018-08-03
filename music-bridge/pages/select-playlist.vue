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
                        <b-form-select :options="playlists" v-model="selectedPlaylist" id="playlistSelect"></b-form-select>
                            <p class="text-center" v-if="selectedPlaylist">
                                Cool, make sure we found the right songs.<br>
                                You can remove the ones we got wrong.
                            </p>
                            <b-button>
                                Transfer Playlist
                            </b-button>
                    </div>
                    <div class="col">
                        <div role="tablist" class="shadow-lg">
                            <!--<b-card v-for="(song, index) in songs" b-bind:key="index" no-body class="mb-1">-->
                                <!--<b-card-header header-tag="header" class="p-1" role="tab">-->
                                    <!--<b-btn block href="#" v-b-toggle.accordion.index variant="info">{{ song.title }} - {{ song.artist }}</b-btn>-->
                                <!--</b-card-header>-->
                                <!--<b-collapse id="accordion{{ index }}" accordion="my-accordion" role="tabpanel">-->
                                    <!--<b-card-body>-->
                                        <!--<p class="card-text">-->
                                            <!--<b-embed type="iframe" aspect="16by9" controls poster="poster.png">-->
                                                <!--<source src="https://www.youtube.com/embed/{{ song.videoId }}"-->
                                                        <!--type='iframe' />-->
                                            <!--</b-embed>-->
                                        <!--</p>-->
                                    <!--</b-card-body>-->
                                <!--</b-collapse>-->
                            <!--</b-card>-->
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>

    export default {
        data() {
            return {
                playlists: [
                    { text: "Select your playlist", value: null },
                    { text: 'Playlist One', value: 1 },
                    { text: 'Playlist Two', value: 2 },
                    { text: 'Playlist Three', value: 3 }
                ],
                selectedPlaylist: null,
                songs: [
                    { title: "The blahh of papoi", artist: "Lorem", videoId: "aghBrGfF2Bw" },
                    { title: "The art of the papoi", artist: "King Bob", videoId: "Tc2EvaiY0s4" }
                ]
            }
        }
    };

</script>

<style>

</style>
