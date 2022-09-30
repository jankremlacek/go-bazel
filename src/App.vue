<script setup lang="ts">
import { ref } from 'vue'

const counter = ref(0)

const onInc = () => {
    counter.value++
}
</script>

<template>
    <div>
        <p>
            Counter: <strong>{{counter}}</strong>
        </p>
        <button @click.prevent="onInc">
            Inc
        </button>
    </div>
</template>
