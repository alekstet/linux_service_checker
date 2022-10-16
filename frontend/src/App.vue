<template>
  <div id="app" class="container">
    <h2 class="display-5">My Services</h2>
    <div>
      <b-dropdown class="m-md-2">
        <template #button-content>
          sort by <strong>{{sort}}</strong>
        </template>
        <b-dropdown-item></b-dropdown-item>
        <b-dropdown-item @click="change_sort('config')">by config</b-dropdown-item>
        <b-dropdown-item @click="change_sort('work')">by work</b-dropdown-item>
        <b-dropdown-item @click="change_sort('stop')">by stop</b-dropdown-item>
      </b-dropdown>
    </div>
    <b-list-group>
      <b-list-group-item v-for="(index, value) in tasks" :key="value" active class="flex-column align-items-start">
        <div class="d-flex w-100 justify-content-sm-between">
          <h4 class="mb-1">{{value}}</h4>
          
          <small>
            <b-button-group>
              <h5 class="h5" :class="{active : index.Active == 'active(running)'}">{{index.Active}}</h5>
              <b-button variant="danger" v-if="index.Active == 'active(running)'" @click="make(value, 'stop')">stop</b-button>
              <b-button variant="success" v-if="index.Active == 'active(exited)' || index.Active == 'inactive(dead)'" @click="make(value, 'start')">start</b-button>
              <b-button variant="secondary" @click="journal(value)">journal</b-button>
              
            </b-button-group>
          </small>
        </div>
        <p v-if="show_journal[value] == 1" class="mb-1">{{index[3]}}</p> 
        <hr id="hr">
      </b-list-group-item>   
    </b-list-group>
  </div>
</template>

<script>

export default {
  name: 'App',
  data: function () {
    return {
      tasks: [],
      sort: "config",
      show_journal: [0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
    }
  },
  created: function() {
    setInterval(() => {
      this.collect()
    }, 4000);
  },
  methods: {
    collect: function() {
      fetch("/collect")
      .then(resp => resp.json())
      .then(data => {
        this.tasks = data
      })
      //this.sorting(services, this.sort)
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
      this.datas()
    },
    sorting: function(arr, type) { 
      if (type=="config") {
        this.tasks = arr
      }
      if (type=="work") {
        arr.sort(function(a, b) {
        var A = a[2].length
        var B = b[2].length
        return B - A
        })
        this.tasks = arr
      }
      if (type=="stop") {
        arr.sort(function(a, b) {
        var A = a[2].length
        var B = b[2].length
        return A - B
        })
        this.tasks = arr
      }   
    },
    change_sort: function(n) {
      this.sort = n
    },
    journal: function(index) {
      var a = this.show_journal[index]
      if (a==1) {
        this.show_journal[index] = 0
      } else {
        this.show_journal[index] = 1
      }
    }
  }
} 
</script>

<style>
body {
  background-color: #6bdef5;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
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
.active {
  color: #1bff00;
}
.test{
  white-space: pre-wrap;
}
</style>
