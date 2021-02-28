<template>
  <span>
    <a class="btn" @click="clickHandler">{{ text }}</a>
  </span>
</template>

<script>
export default {
  name: 'ButtonShare',
  props: ['text'],
  data() {
    return {
      clicked: false,
    }
  },
  methods: {
    clickHandler() {
      console.log('hi')
    },
  },
}
</script>

<style>
.btn {
  display: inline-block;
  padding: 10px 30px;
  margin-right: 10px;
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
