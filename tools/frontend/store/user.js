import userRepository from '~/repositories/userRepository'

export const state = () => ({
  address: '',
  users: null
})

export const mutations = {
  setAddress (state, address) {
    state.address = address
  },
  setUsers (state, users) {
    state.users = users
  }
}

export const getters = {
  address (state) {
    return state.address
  },
  users (state) {
    return state.users
  }
}

export const actions = {
  setAddress(context, params) {
    context.commit('setAddress', params.address)
  },
  async fetchUsers(context, params) {
    const { data } = await userRepository.getUsers()
    context.commit('setUsers', data.users)
  }
}
