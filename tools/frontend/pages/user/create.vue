<template>
  <div class="pt-28 mx-10">
    <div class="">
      <div class="my-3">
        <button class="text-sm cursor-pointer rounded-md py-1 px-3 text-white" :class="{'bg-green-300': !user.name, 'bg-green-500 hover:bg-green-700': user.name }" :disabled="!user.name" @click="createUser">ユーザーを作成する</button>
      </div>
      <input class="p-1 rounded-sm text-gray-600 border border-solid border-gray-400" type="text" v-model="user.name">
      <div class="text-xs text-gray-400 mt-2">NAME: {{ user.name }}</div>
    </div>
    <div class="border border-gray-400 border-solid mt-10 h-28">
      <div class="mx-4 my-3 text-gray-600 font-semibold text-lg">Response</div>
      <div v-if="res" class="m-4 text-xs">{{ res }}</div>
    </div>
    <template v-if="users">
      <div class="mt-10">
        <p>Users</p>
      </div>
      <div v-for="(u, index) in users" :key="index" class="m-3 flex">
        <div class="mr-3">ID: {{ u.id }}</div>
        <div>Name: {{ u.name }}</div>
      </div>
    </template>
  </div>
</template>

<script>
import userRepository from '~/repositories/userRepository'

export default {
  data () {
    return {
      user: {
        name: ''
      },
      users: null,
      res: null
    }
  },
  async mounted () {
    try {
      await this.fetchUsers()
    } catch (e) {
      console.log(e)
    }
  },
  methods: {
    async createUser () {
      try {
        const { data } = await userRepository.createUser(this.user.name)
        this.res = data
      } catch (e) {
        console.log(e)
      } finally {
        this.user.name = ''
        await this.fetchUsers()
      }
    },
    async fetchUsers () {
      try {
        const { data } = await userRepository.getUsers()
        this.users = data.users
      } catch (e) {
        console.log(e)
      } finally {}
    }
  }
}
</script>

<style scoped>

</style>