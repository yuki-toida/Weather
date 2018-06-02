<template>
  <div>
    <h1>todo app</h1>
    <input type="text" v-model="todoName" v-on:keyup.enter="add">
    <ul v-if="items.length">
      <li v-for="item in items" v-bind:key="item.ID">
        {{ item.Name }}
        <button v-on:click="remove(item.ID)">×</button>
      </li>
    </ul>
    <p v-else>No todo</p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  mounted: function() {
    axios
      .get('http://localhost:8080/todo')
      .then(res => this.items = res.data)
      .catch(error => console.log(error));
  },
  data: function() {
    return {
      items: [],
      todoName: null,
    }
  },
  methods: {
    add: function() {
      this.items.push({
        ID: this.items.length + 1,
        Name: this.todoName
      })
      this.todoName = null;
    },
    remove: function(id) {
			this.items = this.items.filter(item => item.ID !== id);
    }
  }
}
</script>
