

export const state = () => ({
  // @USAToken
  // tokenContractAddress: '0x6a7edAd9c7f49Bf215Add73e5d8F8Cb550177297',
  // vendorContractAddress: '0x7dBa3cc9bDf7B3F79dcDD90B0c19768190a5aC5b',
  tokenContractAddress: '0x8B20e8B064CB5dd72CA8990aF939CEF8c1eAB1Ba',
  vendorContractAddress: '0x305b8C435Ed44202139b146EfD854D270522dcA3',
  gachaWalletAddress: '0xec64414617F2B65bB4a7adD57e82a1c5CF53B328',
  myWalletAddress: null,
  tokenBalance: {
    me: null,
    gachaVendor: null,
    gachaWallet: null
  },
  etherBalance: {
    me: null,
    gachaVendor: null
  },
  symbol: null,
  name: null,
  tokenContract: null,
  highlight: false
})

export const mutations = {
  setMyWalletAddress (state, myWalletAddress) {
    state.myWalletAddress = myWalletAddress
  },
  setMyTokenBalance (state, tokenBalance) {
    state.tokenBalance.me = tokenBalance
  },
  setGachaVendorTokenBalance (state, tokenBalance) {
    state.tokenBalance.gachaVendor = tokenBalance
  },
  setGachaWalletTokenBalance (state, tokenBalance) {
    state.tokenBalance.gachaWallet = tokenBalance
  },
  setTokenSymbol (state, symbol) {
    state.symbol = symbol
  },
  setMyEtherBalance (state, etherBalance) {
    state.etherBalance.me = etherBalance
  },
  setGachaVendorEtherBalance (state, etherBalance) {
    state.etherBalance.gachaVendor = etherBalance
  },
  setTokenSymbol (state, symbol) {
    state.symbol = symbol
  },
  setTokenName (state, name) {
    state.name = name
  },
  setTokenContract (state, tokenContract) {
    state.tokenContract = tokenContract
  },
  setHighlight (state, highlight) {
    state.highlight = highlight
  },
}

export const getters = {
  tokenContractAddress (state) {
    return state.tokenContractAddress
  },
  vendorContractAddress (state) {
    return state.vendorContractAddress
  },
  gachaWalletAddress (state) {
    return state.gachaWalletAddress
  },
  myWalletAddress (state) {
    return state.myWalletAddress
  },
  myTokenBalance (state) {
    return state.tokenBalance.me
  },
  gachaVendorTokenBalance (state) {
    return state.tokenBalance.gachaVendor
  },
  gachaWalletTokenBalance (state) {
    return state.tokenBalance.gachaWallet
  },
  myEtherBalance (state) {
    return state.etherBalance.me
  },
  gachaVendorEtherBalance (state) {
    return state.etherBalance.gachaVendor
  },
  symbol (state) {
    return state.symbol
  },
  tokenName (state) {
    return state.name
  },
  tokenContract (state) {
    return state.tokenContract
  },
  highlight (state) {
    return state.highlight
  }
}

export const actions = {
  async getAccountsAndSetMyAddress (context, params) {
    const accounts = await this.$web3util.getAccounts()
    context.commit('setMyWalletAddress', accounts[0])
  },
  async setMyEtherBalance (context, params) {
    const myEtherBalance = await this.$web3util.getBalance(params)
    context.commit('setMyEtherBalance', await this.$web3util.fromWei(myEtherBalance, 'ether'))
  },
  async setGachaVendorEtherBalance (context, params) {
    const gachaVendorBalance = await this.$web3util.getBalance(params)
    context.commit('setGachaVendorEtherBalance', await this.$web3util.fromWei(gachaVendorBalance, 'ether'))
  },
  async setTokenContract (context, params) {
    const tokenContract = await this.$web3util.tokenContract(params)
    context.commit('setTokenContract', tokenContract.methods)
  },
  async setTokenBalance (context, params) {
    if(params.me) {
      context.commit('setMyTokenBalance', await this.$web3util.fromWei(params.balance, 'ether'))
    }
    if(params.gachaVendor) {
      context.commit('setGachaVendorTokenBalance', await this.$web3util.fromWei(params.balance, 'ether'))
    }
    if(params.gachaWallet) {
      context.commit('setGachaWalletTokenBalance', await this.$web3util.fromWei(params.balance, 'ether'))
    }
  }
}
