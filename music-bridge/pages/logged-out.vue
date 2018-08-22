<template>
      <div class="col">
        <div class="row">
          <div class="col">
            <h1 class="text-center">Music Bridge</h1>
          </div>
        </div>
        <div class="row">
        <div class="col">
          <p class="text-center lead">Your Session is Expired 😅</p>
          <p class="text-center">
As you know we do not store any of the information related to your or your music provider(s), the upside you don't have to worry about your data being stolen, the downside is your session is valid for only 1 hour.
          </p>
          <p class="text-center">
            <nuxt-link to="/link-spotify" class="btn btn-primary btn-lg">
              Let's go again!
            </nuxt-link>
          </p>
        </div>
      </div>
      </div>
</template>

<script>
export default {
  components: {}
};
</script>

<style>
.full-height {
  height: 100vh;
}
</style>
