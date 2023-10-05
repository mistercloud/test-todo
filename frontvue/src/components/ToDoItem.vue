<script setup>
import SupportIcon from './icons/IconSupport.vue'

defineProps(['value'])
defineEmits(['remove'])
</script>

<template>
  <div class="item">
    <i @click="$emit('remove')">
      <SupportIcon />
    </i>
    <div class="details">
      {{ value }}
    </div>
  </div>
</template>

<style scoped>
  .item {
    display: flex;
    position: relative;
  }

  .details {
    flex: 1;
    margin-left: 1rem;
    font-size: 1.2rem;
    font-weight: 500;
    margin-bottom: 0.4rem;
    color: var(--color-heading);
  }

  i {
    display: flex;
    place-items: center;
    place-content: center;
    width: 32px;
    height: 32px;

    color: var(--color-text);

    &:hover {
      color: red;
    }
  }
</style>