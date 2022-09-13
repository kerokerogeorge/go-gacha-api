<template>
  <div class="w-screen">
    <div class="flex items-center pl-56 h-28 fixed z-20 bg-white border-b border-solid border-gray-400 w-full">
      <div class="text-2xl font-bold">API動作確認</div>
      <div class="ml-auto">
        <button @click="login">Login</button>
        <div>
          アドレス: {{ address ? address : '-' }}
        </div>
        <div>
          {{ symbol ? symbol : '-' }}: {{ balance ? balance : '-' }}
        </div>
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
import ABI_JSON from "../static/abi.json"

const tokenContractAddress = '0x984a6eaecBE9e77339931E191B6bf314c6f65dab'

export default {
  data() {
    return {
      provider: null,
      accounts: [],
      balance: null,
      address: null,
      abi: ABI_JSON,
      symbol: null
    }
  },
  async mounted () {
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
    async login () {
      try {
        if(!window.ethereum) return
        const provider = new ethers.providers.Web3Provider(window.ethereum)
        provider.getBalance("0x63c7a33d940113c8d9634fff125efa564aa4cc0c").then((result)=>{
          this.balance = ethers.utils.formatEther(result)
        })

        this.accounts = await window.ethereum.request({ method: 'eth_requestAccounts' })
        const signer = await provider.getSigner()
        this.address = await signer.getAddress()
        const contract = new ethers.Contract(tokenContractAddress, this.abi, provider)
        const result = (await contract.balanceOf(this.address)).toString()
        this.symbol = await contract.symbol()
        this.balance = await ethers.utils.formatUnits(result)
      } catch (e) {
        console.log(e)
      }
    }
  }
}
</script>
