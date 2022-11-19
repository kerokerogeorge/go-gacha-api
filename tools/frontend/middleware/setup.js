export default async ({store}) => {
  await store.dispatch('gacha/fetchGachas')
  await store.dispatch('user/fetchUsers')
  await store.dispatch('web3/getAccountsAndSetMyAddress')
  await store.dispatch('web3/setMyEtherBalance', await store.state.web3.myWalletAddress)
  await store.dispatch('web3/setGachaVendorEtherBalance', await store.state.web3.vendorContractAddress)
  await store.dispatch('web3/setTokenContract', await store.state.web3.tokenContractAddress)
  await store.commit('web3/setTokenSymbol', await store.state.web3.tokenContract.symbol().call())
  await store.commit('web3/setTokenName', await store.state.web3.tokenContract.name().call())
  const myTokenBalance = await store.state.web3.tokenContract.balanceOf(store.state.web3.myWalletAddress).call()
  const vendorContractBalance = await store.state.web3.tokenContract.balanceOf(store.state.web3.vendorContractAddress).call()
  const gachaWalletBalance = await store.state.web3.tokenContract.balanceOf(store.state.web3.gachaWalletAddress).call()
  await store.dispatch('web3/setTokenBalance', { me: true, balance: myTokenBalance })
  await store.dispatch('web3/setTokenBalance', { gachaVendor: true, balance: vendorContractBalance })
  await store.dispatch('web3/setTokenBalance', { gachaWallet: true, balance: gachaWalletBalance })

}