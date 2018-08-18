<template>
    <div class="container">
        <div class="row full-height align-items-center">
            <div class="col">
                <div class="row">
                    <div class="col">
                        <h1 class="text-center"><span class="youtube-pulse youtube-text">YouTube</span></h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col">

                        <ConditionalText
                                :textDependency="youtubeAccessToken"
                                showIfTrue="We already have your YouTube access Token."
                                showIfFalse="Now, let's log in to your YouTube account."
                        ></ConditionalText>

                        <p class="text-center">

                            <LoginButton
                                    :buttonDependency="!youtubeAccessToken"
                                    :url="youtubeUrl" buttonMessage="Log in to YouTube"
                                    waitMessage="Just a second..."
                            ></LoginButton>

                            <InternalLinkButton
                                    linkTo="/select-playlist"
                                    :buttonDependency="youtubeAccessToken"
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
            youtubeUrl: function() {
                return this.$store.state.youtube.authUrl;
            },
            youtubeAccessToken: function() {
                return this.$store.state.youtube.accessToken;
            }
        },
      components: {
        LoginButton,
        InternalLinkButton,
        ConditionalText
      }
    };
</script>