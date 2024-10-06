<script setup>
import { reactive } from 'vue'
import { GetMethods, SendRequest } from '../../wailsjs/go/api/HttpV1'


const http_request_handler = reactive({
  selected_method: "GET",
  methods: ["GET"],
  url: "https://google.com",
  headers: {},
  response: {
    headers: {},
    body: {},
  },
  renderResp: {},
})


function getMethods() {
  GetMethods().then(result => {
    http_request_handler.methods = result
    console.log(result)
  })
}

function sendRequest() {
  const req = {
    method: http_request_handler.selected_method,
    url: http_request_handler.url,
  }

  console.log("request", req)
  SendRequest(req).then(result => {
    http_request_handler.response.body = result.Body
    http_request_handler.response.headers = result.Headers
    console.log(result)

    renderResponse()
  })
}

function renderResponse() {
  console.log("header ct", http_request_handler.response.headers["content-type"])
  const ct = http_request_handler.response.headers["content-type"] || ""

  if (ct == "application/json") {
    const jsonRes = JSON.parse(http_request_handler.response.body)
    console.log("###", jsonRes)
    http_request_handler.renderResp = JSON.stringify(jsonRes)
  } else {
    http_request_handler.renderResp = http_request_handler.response.body
  }

}



</script>

<template>
  <main>
    <div id="result" class="result">{{ http_request_handler.methods }}</div>
    <select v-model="http_request_handler.selected_method">
      <option v-for="method in http_request_handler.methods">
        {{ method }}
      </option>
      </ select>
      <input type="url" placeholder="Enter url" v-model="http_request_handler.url">
      <button class="btn" @click="sendRequest">Send</button>
      <button class="btn" @click="getMethods">Get Methods</button>

      <p>{{ http_request_handler.renderResp }}</p>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
