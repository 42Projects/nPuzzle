import Vue from 'vue'
import App from './App.vue'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.config.productionTip = false;
Vue.use(BootstrapVue);

new Vue({
  render: h => h(App),
}).$mount('#app');

/*
import { Handshake, Problem } from './npuzzle_pb';
import { NpuzzleClient } from './npuzzle_grpc_web_pb'

let client = new NpuzzleClient('http://' + window.location.hostname + ':8080',
    null, null);

let handshake = new Handshake();
handshake.setMessage('ping');

client.greets(handshake, {}, (err, res) => {
  console.log(res.getMessage());
});

let problem = new Problem();
problem.setHeuristic("manhattan + linear conflicts");
problem.setSearch("uniform-cost");
problem.setMatrix("4\n1 5 9 13\n2 6 10 14\n3 7 11 15\n4 8 12 0");

client.solve(problem, {}, (err, res) => {
  if (res.getSuccess()) {
    console.log(res.getMoves());
    console.log(res.getMaxstates());
    console.log(res.getTotalstates());
  } else {
    console.log(err);
  }
});*/
