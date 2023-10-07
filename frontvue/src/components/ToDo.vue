<script setup lang="ts">
import { ref, onMounted } from 'vue'
import ToDoItem from "@/components/ToDoItem.vue";
import TodoService from "@/services/TodoService";

const newTodoText = ref('')

interface Item {
  id : number
  title: string
}

const todos = ref([])

onMounted(() => {
  updateItems()
})

function updateItems() {
  TodoService.getList().then(response => {
    todos.value = response.data.items
  })
      .catch(error => {
        console.log(error)
      })
}



function addNewTodo() {
  TodoService.addItem(newTodoText.value).then(response => {
    todos.value.push(response.data)
  })
}


function remove(index:number) {
  TodoService.removeItem(index).then(response => {
    updateItems()
  })
}

</script>

<template>
  <ToDoItem v-for="(item,index) in todos"
            :index="item.id"
            :value="item.title"
            :key="item.id"
            @remove="remove(item.id)"
  />

  <form v-on:submit.prevent="addNewTodo">
    <label for="new-todo">Add a todo</label>
    <input
        v-model="newTodoText"
        id="new-todo"
        placeholder="Добавить элемент в список"
    />
    <button>Add</button>
  </form>
</template>

<style scoped>

</style>