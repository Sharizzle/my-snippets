<template>
  <div
    :class="{ active: isActive }"
    class="tab-item"
    @click="clickHandler(title)"
  >
    <h3 class="tab-item__heading">{{ title }}</h3>
    <p class="tab-item__subheading">{{ description }}</p>
    <div class="tab-item__footer">
      <Tag :text="category" />
      <Date :text="date" />
    </div>
  </div>
</template>

<script>
export default {
  name: 'Tab',
  props: ['title', 'description', 'isActive', 'category', 'date'],
  methods: {
    clickHandler() {
      this.$emit('clickHandler')
    },
  },
}
</script>

<style>
.tab-item {
  background: var(--white);
  border: 1px #d5dadf solid;
  border-left: 5px solid #d5dadf;
  box-shadow: 0 0.5px 0.5px rgba(213, 218, 223, 0.35);
  padding: 15px;
  border-radius: 3px;
  cursor: pointer;
  transition: all 0.2s ease;
  margin-bottom: 15px;
}

.tab-item:hover {
  border-left: 10px solid var(--primary);
}

.tab-item.active {
  box-shadow: 0px 0.5px 0.5px 0.5px rgba(213, 218, 223, 0.35);
  border-left: 10px solid var(--active);
}

.tab-item__heading {
  font-weight: bold;
  font-size: 20px;
  line-height: 1.3em;
  letter-spacing: 0.02em;
  color: var(--primary);
  margin: 0px;
  margin-bottom: 4px;
}

.tab-item__subheading {
  font-size: 14px;
  color: #333;
  margin: 0px;
}

.tab-item__footer {
  display: flex;
  justify-content: space-between;
}

@media (max-width: 985px) {
  .tab-item {
    min-width: 80%;
  }

  .tab-item__heading {
    font-size: 15px;
  }

  .tab-item__subheading {
    font-size: 12px;
  }
}
</style>
