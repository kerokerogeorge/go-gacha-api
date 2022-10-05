<template>
  <div class="loading-container">
    <span class="loading" />
  </div>
</template>

<style scoped>
.loading-container {
  z-index: 3000;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color:rgba(0, 0, 0, 0.6);
  position: fixed;
  display: flex;
  justify-content: center;
  align-items: center;
}

.loading {
  z-index: 2;
  border: 5px solid #EC6B3E;
  border-top: 5px solid #ffffff;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  margin: auto;
  animation: spin 0.9s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style>