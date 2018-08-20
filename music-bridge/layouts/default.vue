
 <!--make this happen all around on all pages-->
<!--<b-alert variant="danger"-->
<!--fade-->
<!--dismissible-->
<!--:show="this.$store.getters['hasError']">-->
<!--{{this.$store.state.globalError}}-->
<!--</b-alert>-->

<template>
  <div>
    <nuxt />
               <b-alert variant="danger"
                fade
                 dismissible
                 :show="true">
Hello There
                </b-alert>

  </div>
</template>

<style lang="scss">
.full-height {
  height: 100vh;
}
.page-enter-active,
.page-leave-active {
  transition: all 0.25s ease-out;
}
.page-enter,
.page-leave-active {
  opacity: 0;
  transform: scale(0.95);
  transform-origin: 50% 50%;
}
.youtube-text {
  background-color: #ff0400;
  padding: 10px;
  border-radius: 20px;
  color: #fff;
}
.spotify-text {
  background-color: #1db954;
  padding: 10px;
  border-radius: 20px;
  color: #fff;
}

.spotify-pulse {
  box-shadow: 0 0 0 rgba(29, 185, 84, 0.4);
  animation: spotify-pulse 2s;
  &.pulse-repeat {
    animation: spotify-pulse 2s infinite;
   }
}

.pulse-repeat {
  animation: infinite;
}

@-webkit-keyframes spotify-pulse {
  0% {
    -webkit-box-shadow: 0 0 0 0 rgba(29, 185, 84, 0.2);
  }
  85% {
    -webkit-box-shadow: 0 0 0 500px rgba(29, 185, 84, 0);
  }
  100% {
    -webkit-box-shadow: 0 0 0 0 rgba(29, 185, 84, 0);
  }
}
@keyframes spotify-pulse {
  0% {
    -moz-box-shadow: 0 0 0 0 rgba(29, 185, 84, 0.2);
    box-shadow: 0 0 0 0 rgba(29, 185, 84, 0.4);
  }
  85% {
    -moz-box-shadow: 0 0 0 500px rgba(29, 185, 84, 0);
    box-shadow: 0 0 0 500px rgba(29, 185, 84, 0);
  }
  100% {
    -moz-box-shadow: 0 0 0 0 rgba(29, 185, 84, 0);
    box-shadow: 0 0 0 0 rgba(29, 185, 84, 0);
  }
}

.youtube-pulse {
  box-shadow: 0 0 0 rgba(255, 4, 0, 0.2);
  animation: youtube-pulse 2s;
  &.pulse-repeat {
    animation: youtube-pulse 2s infinite;
  }
}

@-webkit-keyframes youtube-pulse {
  0% {
    -webkit-box-shadow: 0 0 0 0 rgba(255, 4, 0, 0.2);
  }
  85% {
    -webkit-box-shadow: 0 0 0 500px rgba(255, 4, 0, 0);
  }
  100% {
    -webkit-box-shadow: 0 0 0 500px rgba(255, 4, 0, 0);
  }
}
@keyframes youtube-pulse {
  0% {
    -moz-box-shadow: 0 0 0 500px rgba(255, 4, 0, 0.2);
    box-shadow: 0 0 0 0 rgba(255, 4, 0, 0.2);
  }
  85% {
    -moz-box-shadow: 0 0 0 500px rgba(255, 4, 0, 0);
    box-shadow: 0 0 0 500px rgba(255, 4, 0, 0);
  }
  100% {
    -moz-box-shadow: 0 0 0 500px rgba(255, 4, 0, 0);
    box-shadow: 0 0 0 500px rgba(255, 4, 0, 0);
  }
}
</style>
