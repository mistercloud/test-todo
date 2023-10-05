<script setup>
defineProps(['value'])
defineEmits(['remove'])
</script>

<template>
  <li>
    {{ value }}
    <button @click="$emit('remove')">Remove</button>
  </li>
</template>

<style scoped>

</style>