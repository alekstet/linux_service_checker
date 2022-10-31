<template>
  <div id="app" class="container">
    <h2 class="display-5">My Services</h2>
    <div>
      <b-dropdown class="m-md-2">
        <template #button-content>
          show: <strong>{{show}}</strong>
        </template>
        <b-dropdown-item @click="change_show('active')">active</b-dropdown-item>
        <b-dropdown-item @click="change_show('inactive')">inactive</b-dropdown-item>
        <b-dropdown-item @click="change_show('all')">all</b-dropdown-item>
      </b-dropdown>
      <div>
        <b-form inline>
          <label class="sr-only" for="inline-form-input-name">Name</label>
          <b-form-input
            id="inline-form-input-name"
            class="mb-2 mr-sm-2 mb-sm-0"
            placeholder="Service name"
            v-model="service"
          ></b-form-input>
          <div class="mt-2">Value: {{ service }}</div>
          <b-button variant="primary" @click="search()">Search</b-button>
          <h5 v-if="error">Error occured</h5>
        </b-form>
      </div>
    </div>
    <b-list-group>
      <b-list-group-item v-for="(index, value, num) in tasks" :key="value" active class="flex-column align-items-start">
        <div class="d-flex w-100 justify-content-sm-between">
          <h4 class="mb-1">{{value}}</h4>
          <h4 class="mb-1">{{num}}</h4>
          
          <small>
            <b-button-group>
              <h5 class="h5" :class="[{active : index.Active == 'active(running)'}, {activeEx : index.Active == 'active(exited)'}]">{{index.Active}}</h5>
              <b-button variant="danger" v-if="index.Active == 'active(running)'" @click="make(value, 'stop')">stop</b-button>
              <b-button variant="success" v-if="index.Active == 'active(exited)' || index.Active == 'inactive(dead)'" @click="make(value, 'start')">start</b-button>
              <b-button variant="secondary" @click="journal(num)">journal</b-button>
            </b-button-group>
          </small>
        </div>
        <p v-if="arr[num] == 1" class="mb-1">{{index.Journal}}</p> 
        <hr id="hr">
      </b-list-group-item>   
    </b-list-group>
    <h5>ARR {{this.arr}}</h5>
  </div>
</template>

<script>

export default {
  name: 'App',
  data: function () {
    return {
      err : false,
      service: "",
      show: "all",
      tasks: [],
      arr : [],
    }
  },
  created: function() {
    setInterval(() => {
      this.collect()
    }, 1000);
  },
  methods: {
    collect: function() {
      fetch("/collect?show=" + this.show)
      .then(resp => resp.json())
      .then(data => {
        this.tasks = data
        console.log(data)
        for (let index = 0; index < data.length; index++) {
          console.log(index, data[index])
        }
        var len = Object.keys(data).length
        this.arr = new Array(len).fill(0)
      })
      .catch(resp => console.error(resp))
    },
    make: function(name, command) {
      var res = {
        name: name,
        command: command,
      }
      fetch("/make", {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json;charset=utf-8'
        },
        body: JSON.stringify(res)
      })
      .then(resp => console.error(resp))
      .catch(resp => console.error(resp))
      this.collect()
    },
    change_show: function(n) {
      this.show = n
    },
    search: function() {
      console.log("hello")
    },
    journal: function(index) {
      console.log(index)
      var a = this.arr[index]
      if (a==1) {
        this.arr[index] = 0
      } else {
        this.arr[index] = 1
      }
      console.log(this.arr)
    }
  }
} 
</script>

<style>
body {
  background-color: #b0f2ff;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #5b7794;
  margin-top: 60px;
}
#hr {
  margin: 3px 0px;
}
.mb-1 {
  text-align: left;
  white-space: pre-wrap;
}
.h5 {
  margin-top: 0.5rem;
  margin-right: 0.5rem;
  color: #ff5f5f;
}
.activeEx {
  color: #f8b155;
}
.active {
  color: #1bff00;
}
.test{
  white-space: pre-wrap;
}
</style>
