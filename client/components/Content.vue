<template>
  <div class="container-content">
    <div class="tab-content__heading">
      <h3 class="tab-content__header">
        {{ title }} <br />
        <div class="button-container">
          <ButtonClipboard
            text="Copy to Clipboard"
            @copy-clipboard="copyCode(content)"
          />
        </div>
      </h3>
      <h5 class="tab-content__date">{{ date }}</h5>
    </div>
    <highlight-code lang="js">
      {{ content }}
    </highlight-code>
  </div>
</template>

<script>
export default {
  name: 'Content',
  props: ['title', 'content', 'date'],
  methods: {
    copyCode(code) {
      navigator.clipboard
        .writeText(code)
        .then(() => alert('Code Copied to Clipboard!'))
    },
  },
}
</script>

<style>
.container-content {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: var(--secondary);
}

.button-container {
  margin: 10px 0;
}

.hljs {
  padding-left: 0;
}

.tab-content__date {
  color: var(--active);
  letter-spacing: 2px;
}

.tab-content__heading {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.tab-content__header {
  color: var(--secondary);
  font-weight: bold;
  margin: 0px 0px 0px;
  font-size: 30px;
  line-height: 1.3em;
  letter-spacing: 0.02em;
  align-self: flex-start;
}

.tab-content__text {
  margin: 0px 0px 30px;
  font-size: 1.1em;
}
</style>
