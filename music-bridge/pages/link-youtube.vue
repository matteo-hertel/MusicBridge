<template>
    <div class="col">
        <div class="row">
            <div class="col">
                <h1 class="text-center"><span class="youtube-pulse youtube-text">YouTube</span></h1>
            </div>
        </div>
        <div class="row">
            <div class="col">
                <ConditionalBlock
                        :condition="youtubeAccessToken"
                >
                    <div slot="true">
                        <p class="text-center lead">We already have your YouTube access Token.</p>
                        <p class="text-center lead">
                            <InternalLinkButton
                                    linkTo="/select-playlist"
                                    :buttonDependency="youtubeAccessToken"
                                    buttonMessage="Skip"
                            ></InternalLinkButton>
                        </p>
                    </div>

                    <div slot="false">
                        <p class="text-center lead">Now, let's log in to your YouTube account.</p>
                        <p class="text-center lead">
                            <LoginButton
                                    :buttonDependency="!youtubeAccessToken"
                                    :url="youtubeUrl" buttonMessage="Log in to YouTube"
                                    waitMessage="Just a second..."
                            ></LoginButton>
                        </p>
                    </div>
                </ConditionalBlock>
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
        ConditionalBlock
      }
    };
</script>