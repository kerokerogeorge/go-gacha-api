<template>
  <div class="min-h-screen pt-28">
    <div class="flex w-full h-full">
      <template v-if="users">
        <div class="w-4/12 h-auto pt-5 pl-5 border-r border-solid border-gray-400">
          <div class="">
            <p>Users</p>
            <div v-for="(u, index) in users" :key="index" class="mt-3">
              <div class="flex">
                <div class="text-sm mr-3">{{ u.id }}</div>
                <div>
                  <button class="bg-gray-400 hover:bg-gray-700 text-xs cursor-pointer rounded-sm p-1 text-white"  @click="fetchUser(u.token)">ユーザーを取得</button>
                </div>
              </div>
              <div class="text-xs text-gray-500">Token: {{ u.token }}</div>
            </div>
          </div>
        </div>
        <div class="w-8/12">
          <div class="bg-white z-20 fixed top-28 h-32 py-3 px-5 w-full border-b border-gray-400 border-solid">
            <template v-if="user">
              <div class="w-full">
                {{ user.name }}
              </div>
              <div class="mt-3 text-gray-500 text-xs">
                Token: {{ token }}
              </div>
              <div class="mt-1 text-gray-500 text-xs">
                Address: {{ user.address }}
              </div>
              <div class="mt-3">
                <button v-if="!isFetched" class="bg-gray-400 hover:bg-gray-700 text-xs cursor-pointer rounded-sm p-1 text-white"  @click="fetchUserCharacters(token)">キャラクターを取得</button>
              </div>
            </template>
            <template v-if="!user">
              <div class="w-full text-sm text-gray-600">ユーザーは選択されていません。</div>
            </template>
          </div>
          <div class="h-custom overflow-scroll">
            <template v-if="isFetched && characters.length > 0 ">
              <div class="px-5 pt-36 pb-10">
                <div class="border border-solid border-gray-500 text-gray-700 text-sm">
                  <div class="flex py-2 border-b border-solid border-gray-500">
                    <div class="pl-3 w-10">ID</div>
                    <div class="w-40">Name</div>
                    <div class="w-20">char id</div>
                    <div class="w-40">emmission rate</div>
                    <!-- <div class="w-40">img</div> -->
                  </div>
                  <div v-for="(c, index) in characters" :key="index" class="">
                    <div class="flex py-2">
                      <div class="ml-3 w-10">{{ c.characterId }}</div>
                      <div class="w-40">{{ c.name }}</div>
                      <div class="w-20">{{ c.userCharacterId }}</div>
                      <div class="w-20">{{ c.emissionRate }}</div>
                      <!-- <div class="relative border border-solid border-gray-400">
                        <img :src="c.imgUrl" alt="pokemon" class="w-20 h-20" />
                        <div class="absolute text-xs bottom-2 px-2 w-full bg-gray-400 text-white bg-opacity-80">{{ c.name }}</div>
                      </div> -->
                    </div>
                  </div>
                </div>
              </div>
            </template>
            <template v-if="isFetched && !characters.length > 0">
              <div class="pt-36 text-gray-500 px-8">
                no character
              </div>
            </template>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import userRepository from '~/repositories/userRepository'

export default {
  data () {
    return {
      users: null,
      res: null,
      characters: [],
      user: null,
      isFetched: false
    }
  },
  computed: {
    ...mapGetters('gacha', [
      'token'
    ]),
  },
  async mounted () {
    try {
      await this.fetchUsers()
    } catch (e) {
      console.log(e)
    }
  },
  methods: {
    ...mapActions('gacha', [
      'setUserToken',
    ]),
    async fetchUsers () {
      this.isFetched = false
      try {
        const { data } = await userRepository.getUsers()
        this.users = data.users
      } catch (e) {
        console.log(e)
      } finally {}
    },
    async fetchUser (token) {
      this.isFetched = false
      try {
        const { data } = await userRepository.getUser(token)
        this.user = data
        await this.setUserToken({ token: token })
      } catch (e) {
        console.log(e)
      } finally {}
    },
    async fetchUserCharacters (token) {
      try {
        this.isFetched = true
        const { data } = await userRepository.getUserCharacters(token)
        this.characters = data.characters
      } catch (e) {
        console.log(e)
      } finally {}
    }
  }
}
</script>

<style scoped>
.h-custom {
  height: calc(100vh);
}
</style>