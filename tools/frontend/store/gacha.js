import gachaRepository from '~/repositories/gachaRepository'
export const state = () => ({
  gachaId: '',
  token: '',
  gachas: null
})

export const mutations = {
  setGachaId (state, gachaId) {
    state.gachaId = gachaId
  },
  setToken (state, token) {
    state.token = token
  },
  setGachas (state, gachas) {
    state.gachas = gachas
  }
}

export const getters = {
  gachaId (state) {
    return state.gachaId
  },
  token (state) {
    return state.token
  },
  gachas (state) {
    return state.gachas
  }
}

export const actions = {
  selectGachaId(context, params) {
    context.commit('setGachaId', params.gachaId)
  },
  setUserToken(context, params) {
    context.commit('setToken', params.token)
  },
  async fetchGachas(context, params) {
    const { data } = await gachaRepository.getGachas()
    context.commit('setGachas', data.data)
  }
}
