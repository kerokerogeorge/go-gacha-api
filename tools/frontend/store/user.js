export const state = () => ({
  address: '',
})

export const mutations = {
  setAddress (state, address) {
    state.address = address
  },
}

export const getters = {
  address (state) {
    return state.address
  },
}

export const actions = {
  setAddress(context, params) {
    context.commit('setAddress', params.address)
  },
}
