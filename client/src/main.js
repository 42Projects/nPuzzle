import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false;

import { Handshake } from './npuzzle_pb';
import { NpuzzleClient } from './npuzzle_grpc_web_pb'

let client = new NpuzzleClient('http://' + window.location.hostname + ':8080',
    null, null);

// simple unary call
let handshake = new Handshake();
handshake.setMessage('ping');

client.greets(handshake, {}, (err, response) => {
  console.log(response.getMessage());
});

new Vue({
  render: h => h(App),
}).$mount('#app');
