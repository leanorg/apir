<script setup>
import Httpv1ApiRequest from './components/Httpv1ApiRequest.vue'</script>

<template>
  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png" />
  <Httpv1ApiRequest />
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
