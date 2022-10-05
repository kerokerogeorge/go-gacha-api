export const state = () => ({
  isConnected: false,
})

export const mutations = {
  setConnectionStatus (state, isConnected) {
    state.isConnected = isConnected
  }
}

export const getters = {
  isConnected (state) {
    return state.isConnected
  }
}

export const actions = {
  changeConnectionStatus(context, params) {
    context.commit('setConnectionStatus', params.isConnected)
  }
}