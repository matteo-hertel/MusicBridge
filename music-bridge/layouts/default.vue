<template>
  <div>
    <nuxt/>
  </div>
</template>

<style>
  .full-height {
    height: 100vh;
  }
    .page-enter-active, .page-leave-active {
        transition: all .25s ease-out;
    }
    .page-enter, .page-leave-active {
        opacity: 0;
        transform: scale(0.95);
        transform-origin: 50% 50%;
    }
</style>

