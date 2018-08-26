<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <b-alert variant="danger"
                         fade
                         dismissible
                         :show="this.$store.getters['hasError']">
                    {{this.$store.state.globalError}}
                </b-alert>
            </div>
        </div>
    </div>
</template>