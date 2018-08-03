<template>
  <div class="container">
    <div class="row full-height align-items-center">
      <div class="col">
        <div class="row">
          <div class="col">
            <h1 class="text-center">Music Bridge</h1>
          </div>
        </div>
        <div class="row">
        <div class="col">
          <p class="text-center lead">Spotify playlists to YouTube Music - The fast way.</p>
          <p class="text-center">
            Music Bridge stores none of you data. We use short duration tokens to access your accounts.<br/>
            These tokens last one hour and then expire. Think of it like a parking ticket for your account!
          </p>
          <p class="text-center">
            <nuxt-link to="/link-spotify" class="btn btn-primary btn-lg">
              Neat! Let's Go!
            </nuxt-link>
          </p>
        </div>
      </div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  components: {
  }
};
// import AppLogo from "~/components/AppLogo.vue";

// export default {
//   components: {
//     AppLogo
//   }
// };
</script>

<style>
  .full-height {
    height: 100vh;
  }
</style>
