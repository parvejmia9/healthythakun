<script setup lang="ts">
import { ref, onMounted } from 'vue';
import HelloWorld from './components/HelloWorld.vue';

// Use ref to store reactive variables
const message = ref('Loading...');

// Fetch data from Go backend on component mount
onMounted(() => {
  fetch('http://localhost:8181/ping')
    .then(response => {
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    })
    .then(data => {
      message.value = data.status; // Update the message with the response
    })
    .catch(error => {
      console.error('There was a problem with the fetch operation:', error);
      message.value = 'Error fetching data';
    });
});


// Reactive variables
const tasks = ref([]);      // To store the tasks fetched from the backend
const showModal = ref(false); // To control the visibility of the modal

// Fetch tasks from Go backend when the button is clicked
const fetchTasks = () => {
  fetch('http://localhost:8181/tasks')
    .then(response => {
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    })
    .then(data => {
      tasks.value = data.tasks; // Store the tasks in the reactive variable
      showModal.value = true;   // Show the modal once tasks are fetched
    })
    .catch(error => {
      console.error('There was a problem with the fetch operation:', error);
    });
};

// Close modal
const closeModal = () => {
  showModal.value = false; // Hide the modal
};
</script>

<template>
  <div>
    <h1>{{ message }}</h1>
  </div>

  <div>
    <!-- Button to fetch tasks and show the modal -->
    <button @click="fetchTasks">Show Tasks</button>

    <!-- Modal Popup -->
    <div v-if="showModal" class="modal-overlay" @click="closeModal">
      <div class="modal-content" @click.stop>
        <h2>Task List</h2>
        <table>
          <thead>
          <tr>
            <th>ID</th>
            <th>Title</th>
            <th>Description</th>
            <th>Due Date</th>
            <th>Status</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="task in tasks" :key="task.id">
            <td class="color: green">{{ task.id }}</td>
            <td :style="{ color: 'red' }">{{ task.title }}</td>
            <td>{{ task.description }}</td>
            <td>{{ new Date(task.due_date).toLocaleDateString() }}</td>
            <td>{{ task.status }}</td>
          </tr>
          </tbody>
        </table>
        <button @click="closeModal">Close</button>
      </div>
    </div>
  </div>

</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}


nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
table {
  width: 100%;
  border-collapse: collapse;
  margin: 20px 0;
}

table, th, td {
  border: 1px solid black;
}

th, td {
  padding: 10px;
  text-align: left;
}

th {
  background-color: #f2f2f2;
}
</style>
