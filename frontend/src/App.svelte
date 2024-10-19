<script lang="ts">
  import TopAppBar from "@smui/top-app-bar";
  import Drawer, { AppContent, Content } from "@smui/drawer";
</script>

<div class="app">
  <div class="top-bar" style="display: none"></div>
  <TopAppBar variant="static" color="secondary" class="top-bar">Test</TopAppBar>
  <div class="drawer-container">
    <Drawer>
      <Content class="drawer-content">Ciao</Content>
    </Drawer>
    <AppContent class="app-content">
      <main class="main-content">Hello</main>
    </AppContent>
  </div>
</div>

<style>
  .app {
    width: 100%;
    height: 100%;
  }

  .drawer-container {
    position: relative;
    display: flex;
    height: calc(100% - 60px);
    overflow: hidden;
    z-index: 0;
  }

  .app :global(.top-bar) {
    height: 60px;
    padding: 16px;
  }

  .drawer-container :global(.drawer-content) {
    padding: 16px;
  }

  .drawer-container :global(.app-content) {
    flex: auto;
    overflow: auto;
    position: relative;
    flex-grow: 1;
  }

  .main-content {
    overflow: auto;
    padding: 16px;
    height: 100%;
    box-sizing: border-box;
  }
</style>
