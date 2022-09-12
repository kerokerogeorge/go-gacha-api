export const state = () => ({
  gachaId: '',
  token: ''
})

export const mutations = {
  setGachaId (state, gachaId) {
    state.gachaId = gachaId
  },
  setToken (state, token) {
    state.token = token
  }
}

export const getters = {
  gachaId (state) {
    return state.gachaId
  },
  token (state) {
    return state.token
  },
}

export const actions = {
  selectGachaId(context, params) {
    context.commit('setGachaId', params.gachaId)
  },
  setUserToken(context, params) {
    context.commit('setToken', params.token)
  },
}
