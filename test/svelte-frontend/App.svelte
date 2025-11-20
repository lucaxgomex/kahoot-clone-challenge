<script lang="ts">
  let message: string = '';

  async function fetchGreeting() {
    const response = await fetch('http://localhost:8081/api/greet');
    if (response.ok) {
      const data = await response.json();
      message = data.message;
    } else {
      message = 'Failed to load greeting';
    }
  }

  fetchGreeting();
</script>

<main>
  <h1>{message}</h1>
</main>

<style>
  main {
    text-align: center;
    margin-top: 50px;
  }
</style>
