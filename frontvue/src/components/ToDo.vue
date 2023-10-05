<script setup>
import { ref, onMounted } from 'vue'
import ToDoItem from "@/components/ToDoItem.vue";


onMounted(() => {
  //console.log(`The initial count is ${count.value}.`)
})

let items = [
  { id : 1, value : 'ToDo 1'},
  { id : 2, value : 'ToDo 2'},
  { id : 3, value : 'ToDo 3'}
]
const newTodoText = ref('')
const todos = ref(items)

let nextTodoId = 4

function addNewTodo() {
  todos.value.push({
    id: nextTodoId++,
    value: newTodoText.value
  })
  newTodoText.value = ''
}

</script>

<template>
  <ToDoItem v-for="(item,index) in items"
            :index="item.id"
            :value="item.value"
            :key="item.id"
            @remove="todos.splice(index, 1)"
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