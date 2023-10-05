<script setup>
import { ref, onMounted } from 'vue'
import ToDoItem from "@/components/ToDoItem.vue";


onMounted(() => {
  if (localStorage.getItem('items')) {
    try {
      todos.value = JSON.parse(localStorage.getItem('items'));
    } catch(e) {
      localStorage.removeItem('items');
    }
  }
})

const newTodoText = ref('')
const todos = ref([])

let nextTodoId = 4

function addNewTodo() {
  todos.value.push({
    id: nextTodoId++,
    value: newTodoText.value
  })
  newTodoText.value = ''

  saveToStorage()
}

function saveToStorage() {
  const parsed = JSON.stringify(todos.value);
  localStorage.setItem('items', parsed);
}

function remove(index) {
  todos.value.splice(index, 1)
  saveToStorage()
}

</script>

<template>
  <ToDoItem v-for="(item,index) in todos"
            :index="item.id"
            :value="item.value"
            :key="item.id"
            @remove="remove(index)"
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