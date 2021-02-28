<template>
  <span class="date">{{ text }}</span>
</template>

<script>
export default {
  name: 'Date',
  props: ['text'],
}
</script>

<style>
.date {
  padding: 2px 5px;
  display: inline;
  margin: 10px 0 0px 0;
}
</style>
