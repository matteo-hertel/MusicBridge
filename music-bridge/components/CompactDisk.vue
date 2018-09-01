<template>
    <div class="row">
        <div class="compact-disk" :class="{ 'spinning':  spinning }" :style="{ backgroundImage: 'url(https://img.youtube.com/vi/' + videoId + '/0.jpg)' }"></div>
    </div>
</template>

<script>
    export default {
        props: ["videoId", "spinning"],
    };
</script>

<style lang="scss">
    .compact-disk {
        width:200px;
        height:200px;
        background-color: black;
        border-radius: 100px;
        transition: background-image 0.3s ease;
        background-position: center;
        border: 2px solid white;
        &.spinning {
            animation: spin 3s linear infinite;
        }
    }

    @keyframes spin { 100% { -webkit-transform: rotate(360deg); transform:rotate(360deg); } }
</style>