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

                        <ConditionalText
                                :textDependency="spotifyAccessToken"
                                showIfTrue="We already have your Spotify access Token."
                                showIfFalse="First, let's log in to your Spotify account."
                        ></ConditionalText>

                        <p class="text-center">

                            <LoginButton
                                    :buttonDependency="!spotifyAccessToken"
                                    :url="spotifyUrl" buttonMessage="Log in to Spotify"
                                    waitMessage="Just a second..."
                            ></LoginButton>

                            <InternalLinkButton
                                    linkTo="/link-youtube"
                                    :buttonDependency="spotifyAccessToken"
                                    buttonMessage="Next"
                            ></InternalLinkButton>

                        </p>

                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import LoginButton from "~/components/LoginButton.vue";
    import InternalLinkButton from "~/components/InternalLinkButton.vue";
    import ConditionalText from "~/components/ConditionalText.vue";

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
        ConditionalText
    }
};
</script>