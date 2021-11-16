<template>
<div>
  <div class="ui fixed  menu">
      <a href="#" class="header item">
        <img class="ui tiny image" src="https://2tjosk2rxzc21medji3nfn1g-wpengine.netdna-ssl.com/wp-content/themes/konghq/assets/img/konnect/konnect-logo-outline.svg" />
      </a>
  </div>

  <div class="ui main text container">
    <Services title="Services"/>
  </div>

  <div class="ui divider"></div>

  <div class="ui vertical footer segment">
    <div class="ui horizontal small divided link list">
      <a class="item" href="#">Site Map</a>
      <a class="item" href="#">Contact Us</a>
      <a class="item" href="#">Terms and Conditions</a>
      <a class="item" href="#">Privacy Policy</a>
    </div>
  </div>

</div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import Services from './components/Services.vue'


export default {
  name: 'App',
  components: {
    HelloWorld,
    Services
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
