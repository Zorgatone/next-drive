import "./style.css"

// import "svelte-material-ui/themes/svelte-dark.css"
import "svelte-material-ui/themes/svelte.css"

import App from './App.svelte'

const app = new App({
  target: document.getElementById('app')
})

export default app
