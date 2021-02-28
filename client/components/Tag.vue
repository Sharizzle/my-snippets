<template>
  <span class="tag">#{{ text }}</span>
</template>

<script>
export default {
  name: 'Tag',
  props: ['text'],
}
</script>

<style>
.tag {
  background-color: var(--primary);
  color: white;
  font-size: 0.9rem;
  font-weight: bold;
  letter-spacing: 0.95px;
  border-radius: 3px;
  border: 1px var(--primary) solid;
  padding: 2px 5px;
  display: inline;
  margin: 10px 0 0px 0;
  text-transform: capitalize;
}

@media (max-width: 985px) {
  .tag {
    font-size: 0.6rem;
  }
}
</style>
