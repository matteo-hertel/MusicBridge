<template>
    <div class="container">
        <div class="row full-height align-items-center">
            <div class="col">
                <div class="row">
                    <div class="col">
                        <h1 class="text-center">Welcome Back <span class="spotify-pulse spotify-text">Spotify</span> Person!</h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <p class="text-center lead">We missed you!</p>
                        <!--<p class="text-center">-->
                            <!--<nuxt-link to="" class="btn btn-primary btn-lg">-->
                                <!--Log in to Spotify-->
                            <!--</nuxt-link>-->
                            <!--<nuxt-link to="link-youtube" class="btn btn-primary btn-lg">-->
                                <!--Next-->
                            <!--</nuxt-link>-->
                        <!--</p>-->
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        mounted () {
            this.$store.dispatch('spotify/getAccessTokenFromUrl', this.$route.query.code);
        }
    };
</script>

<style>

</style>
