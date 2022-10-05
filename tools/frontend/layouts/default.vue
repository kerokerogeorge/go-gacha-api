<template>
  <div class="w-screen">
    <UtilLoading v-if="loading" />
    <div class="flex items-center pl-56 h-28 fixed z-20 bg-white border-b border-solid border-gray-400 w-full">
      <div>
        <div class="text-2xl font-bold">API動作確認</div>
        <div class="text-xs text-gray-600">
          <div>
          アドレス: {{ address ? address : '-' }}
          </div>
          <div>
            ETH: {{ etherBalance ? etherBalance : '-' }}
          </div>
          <div>
            {{ symbol ? symbol : '-' }}トークン: {{ balance ? balance : '-' }}
          </div>
          <div>
            選択中のユーザー: {{ token ? token : '-' }}
          </div>
          <div>
            選択中のガチャ: {{ gachaId ? gachaId : '-' }}
          </div>
        </div>
      </div>
      <div v-if="!isConnected" class="ml-auto w-auto pr-5">
        <button class="bg-gray-400 hover:bg-gray-700 text-xs cursor-pointer rounded-sm p-1 text-white" @click="connectToMetamask">Metamaskに接続</button>
      </div>
    </div>
    <div class="relative">
      <CommonSidebar />
    </div>
    <main class="pl-48">
      <Nuxt />
    </main>
  </div>
</template>

<style scoped>
</style>

<script>
import { ethers } from "ethers";
import { mapGetters, mapActions } from 'vuex'
// import ABI_JSON from "../static/abi.json"
import CONTRACT_ABI_JSON from "../static/contractAbi.json"

const tokenContractAddress = '0x984a6eaecBE9e77339931E191B6bf314c6f65dab'

export default {
  data() {
    return {
      abi: CONTRACT_ABI_JSON,
      provider: null,
      accounts: [],
      balance: null,
      address: null,
      symbol: null,
      etherBalance: null,
      loading: false
    }
  },
  computed: {
    ...mapGetters('token', [
      'isConnected'
    ]),
    ...mapGetters('gacha', [
      'gachaId',
      'token'
    ]),
  },
  async mounted () {
    await this.changeConnectionStatus({ isConnected: false })
    try {
      if (typeof window.ethereum !== 'undefined') {
        console.log('MetaMask is installed!')
        console.log(window.ethereum.selectedAddress)
      }
    } catch (e) {
      console.log(e)
    }
  },
  methods: {
    ...mapActions('token', [
      'changeConnectionStatus'
    ]),
    async connectToMetamask () {
      try {
        if(!window.ethereum) return
        this.accounts = await window.ethereum.request({ method: 'eth_requestAccounts' })
        const provider = new ethers.providers.Web3Provider(window.ethereum)
        provider.getBalance("0x63c7a33d940113c8d9634fff125efa564aa4cc0c").then((result)=>{
          this.etherBalance = ethers.utils.formatEther(result)
        })
        const signer = await provider.getSigner()
        this.address = await signer.getAddress()
        const contract = new ethers.Contract(tokenContractAddress, this.abi, provider)
        const result = (await contract.balanceOf(this.address)).toString()
        this.symbol = await contract.symbol()
        this.balance = await ethers.utils.formatUnits(result)

        await contract.on("Transfer", async (from, to, value, event) => {
          this.loading = true
          let info = {
            from: from,
            to: to,
            value: ethers.utils.formatUnits(value),
            data: event,
          };
          const b = (await contract.balanceOf(this.address)).toString()
          this.balance = ethers.utils.formatUnits(b)
          this.loading = false
        })
        await this.changeConnectionStatus({ isConnected: true })
      } catch (e) {
        console.log(e)
      }
    }
  }
}
</script>
