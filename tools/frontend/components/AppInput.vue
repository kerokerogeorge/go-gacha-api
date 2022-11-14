<template>
  <div class="w-80 py-1">
    <label :for="name" class="mr-5 font-semibold text-gray-500">{{ name }}</label>
    <input
      type="text"
      :id="name"
      :value="value"
      @input="$emit('input', $event.target.value)"
      class=" w-full rounded-sm px-3 py-1 my-2 border border-solid border-gray-700 hover:border-gray-400"
    />
  </div>
</template>

<script>

export default ({
  name: 'app-input',
  props: ['name', 'value', 'type'],
})
</script>


<style scoped>

</style>