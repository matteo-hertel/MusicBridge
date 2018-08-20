<template>
    <div class="container">
        <div class="row full-height align-items-center">
            <div class="col">
                <div class="row">
                    <div class="col">
                        <h1 class="text-center"><span class="spotify-pulse spotify-text">Spotify</span></h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <ConditionalBlock
                                :condition="spotifyAccessToken"
                        >
                            <div slot="true">
                                <p class="text-center lead">We already have your Spotify access Token.</p>
                                <p class="text-center">
                                    <InternalLinkButton
                                            linkTo="/link-youtube"
                                            :buttonDependency="spotifyAccessToken"
                                            buttonMessage="Skip"
                                    ></InternalLinkButton>
                                </p>
                            </div>
                            <div slot="false">
                                <p class="text-center lead">First, let's log in to your Spotify account.</p>
                                <p class="text-center">
                                    <LoginButton
                                            :buttonDependency="!spotifyAccessToken"
                                            :url="spotifyUrl" buttonMessage="Log in to Spotify"
                                            waitMessage="Just a second..."
                                    ></LoginButton>
                                </p>
                            </div>
                        </ConditionalBlock>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import LoginButton from "~/components/LoginButton.vue";
    import InternalLinkButton from "~/components/InternalLinkButton.vue";
    import ConditionalBlock from "~/components/ConditionalBlock.vue";

export default {
  computed: {
    spotifyUrl: function() {
      return this.$store.state.spotify.authUrl;
    },
    spotifyAccessToken: function() {
        return this.$store.state.spotify.accessToken;
    }
  },
    components: {
        LoginButton,
        InternalLinkButton,
        ConditionalBlock
    }
};
</script>