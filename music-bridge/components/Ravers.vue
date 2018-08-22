<template>
<div>
    <div class="row ravers">
        <div class="col"><img class="raver1" src="~/assets/ravers/raver1.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver2" src="~/assets/ravers/raver2.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver2" src="~/assets/ravers/raver1.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver2" src="~/assets/ravers/raver2.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver1" src="~/assets/ravers/raver1.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver2" src="~/assets/ravers/raver2.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver2" src="~/assets/ravers/raver1.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver2" src="~/assets/ravers/raver2.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver1" src="~/assets/ravers/raver1.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver2" src="~/assets/ravers/raver2.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver1" src="~/assets/ravers/raver1.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver1" src="~/assets/ravers/raver2.png" alt="raver yeahh!"></div>
    </div>

    <div class="row ravers distant-ravers">
        <div class="col"><img class="raver1" src="~/assets/ravers/raver1.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver2" src="~/assets/ravers/raver2.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver2" src="~/assets/ravers/raver1.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver2" src="~/assets/ravers/raver2.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver1" src="~/assets/ravers/raver1.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver2" src="~/assets/ravers/raver2.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver2" src="~/assets/ravers/raver1.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver2" src="~/assets/ravers/raver2.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver1" src="~/assets/ravers/raver1.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver2" src="~/assets/ravers/raver2.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver1" src="~/assets/ravers/raver1.png" alt="raver yeahh!"></div>
        <div class="col"><img class="raver1" src="~/assets/ravers/raver2.png" alt="raver yeahh!"></div>
    </div>
</div>
</template>

<style lang="scss">

    .ravers {
        &.distant-ravers {
            filter: blur(20px);
        }
        position: fixed;
        bottom: -140px;
        width: 100%;
        height: 50vh;
        overflow: hidden;

        .raver1 {
            bottom: -70px;
            position: fixed;

            -webkit-animation: ravejump1 2s ease-in infinite;
            -moz-animation: ravejump1 2s ease-in infinite;
            animation: ravejump1 2s ease-in infinite;

            @-webkit-keyframes ravejump1 {
                0% {
                    bottom: -80px;
                }
                10% {
                    bottom: -90px;
                }
                60% {
                    bottom: -40px;
                }
                100% {
                    bottom: -80px;
                }
            }

            @keyframes ravejump1 {
                0% {
                    bottom: -80px;
                }
                10% {
                    bottom: -90px;
                }
                60% {
                    bottom: -40px;
                }
                100% {
                    bottom: -80px;
                }
            }
        }
        .raver2 {
            bottom: -70px;
            position: fixed;

            -webkit-animation: ravejump2 2s ease-in infinite;
            -moz-animation: ravejump2 2s ease-in infinite;
            animation: ravejump2 2s ease-in infinite;

            @-webkit-keyframes ravejump2 {
                0% {
                    bottom: -70px;
                }
                10% {
                    bottom: -75px;
                }
                60% {
                    bottom: -40px;
                }
                100% {
                    bottom: -70px;
                }
            }

            @keyframes ravejump2 {
                0% {
                    bottom: -70px;
                }
                10% {
                    bottom: -75px;
                }
                60% {
                    bottom: -40px;
                }
                100% {
                    bottom: -70px;
                }
            }
        }
    }
</style>