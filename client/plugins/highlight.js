import Vue from 'vue'
import VueHighlightJS from 'vue-highlight.js'

import javascript from 'highlight.js/lib/languages/javascript'
import go from 'highlight.js/lib/languages/go'
import python from 'highlight.js/lib/languages/python'
import matlab from 'highlight.js/lib/languages/matlab'
import bash from 'highlight.js/lib/languages/bash'
import typescript from 'highlight.js/lib/languages/typescript'

import 'highlight.js/styles/dracula.css'

Vue.use(VueHighlightJS, {
  languages: {
    javascript,
    go,
    python,
    matlab,
    bash,
    typescript,
  },
})
