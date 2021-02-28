<template>
  <div class="index">
    <Header />
    <h1 class="pen-heading">Useful Code Snippets</h1>
    <div class="container">
      <div class="row">
        <div class="col-md-3 col-md">
          <Tab
            v-for="code in codes"
            :key="code.id"
            :is-active="selectedTab === code.id"
            :description="code.description"
            :title="code.title"
            :date="code.date"
            :category="code.category"
            @clickHandler="selectTab(code.id)"
          />
        </div>
        <div class="col-md-7 col-md">
          <Content
            :date="codes.find((code) => code.id === selectedTab).date"
            :title="codes.find((code) => code.id === selectedTab).title"
            :content="codes.find((code) => code.id === selectedTab).content"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  async asyncData({ $axios }) {
    const codes = await $axios.$get('http://localhost:3001/codes')
    return { codes }
  },
  data() {
    return {
      selectedTab: null,
    }
  },
  head() {
    return {
      title: 'My Snippets - Sharif',
      meta: [
        {
          hid: 'home',
          name: 'My Snippets',
          content:
            'Useful Web Development Code Snippets that I use on a daily bases for my personal projects',
        },
      ],
    }
  },
  created() {
    this.selectedTab = this.codes[0].id
  },
  methods: {
    selectTab(id) {
      this.selectedTab = id
    },
  },
}
</script>

<style>
.index {
  font-family: 'SourceSansPro';
  padding: 10px;
  height: 100vh;
}

.col-md {
  width: 100%;
  float: left;
  position: relative;
  min-height: 1px;
  padding-right: 30px;
  padding-left: 0px;
}

@media (max-width: 850px) {
  .col-md-3 {
    width: 100%;
    overflow: auto;
    float: left;
    height: 30vh;
    padding: 10px;
    margin-bottom: 30px;
  }

  .col-md-7 {
    padding: 18px;
  }
}

@media (min-width: 850px) {
  .col-md-7 {
    width: 74%;
    padding-left: 30px;
  }

  .col-md-3 {
    width: 26%;
    overflow: auto;
    float: left;
    height: 70vh;
    padding: 10px;
  }
}

.pen-heading {
  font-family: 'DancingScript';
  letter-spacing: 6px;
  font-weight: bold;
  font-size: 2.8em;
  text-align: center;
  margin: 20px 0 30px 0;
  color: var(--white);
  -webkit-text-stroke-width: 0.5px;
  -webkit-text-stroke-color: var(--primary);
}

.tab-content__header {
  color: var(--secondary);
  font-weight: bold;
  margin: 0px 0px 30px;
  font-size: 30px;
  line-height: 1.3em;
  letter-spacing: 0.02em;
}

.tab-content__text {
  margin: 0px 0px 30px;
  font-size: 1.1em;
}

@media (max-width: 985px) {
  .col-md-3 {
    font-size: 12px !important;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
  }
}
</style>
