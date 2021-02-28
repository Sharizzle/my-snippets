<template>
  <a class="btn" @click="handleClick">{{ text }}</a>
</template>

<script>
export default {
  name: 'ButtonClipboard',
  props: ['text'],
  methods: {
    handleClick() {
      this.$emit('copy-clipboard')
    },
  },
}
</script>

<style>
.btn {
  display: inline-block;
  padding: 10px 30px;
  margin-right: 10px;
  margin: 10px 0;
  cursor: pointer;
  text-decoration: none;
  font-size: 14px;
  text-transform: uppercase;
  font-weight: 900;
  position: relative;
  transition: all 0.3s ease;
  text-align: center;
  line-height: 1;
  border: 2px solid;
  border-radius: 3px;
  background-color: transparent;
  color: var(--active);
  fill: var(--active);
  border-color: var(--active);
}

.btn:hover {
  color: var(--white);
  border-color: var(--white);
  text-decoration: none;
}
</style>
