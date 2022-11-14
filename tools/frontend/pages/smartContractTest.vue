<template>
  <div class="w-screen pl-24 pt-32">
    <div>
      <h1>This is my dApp!</h1>
      <p>Here we can set or get the mood:</p>
      <label for="mood">Input Mood:</label> <br />
      <input type="text" id="mood" v-model="mood" class="border border-solid border-gray-400" />
      <button class="bg-gray-400 hover:bg-gray-700 text-xs cursor-pointer rounded-sm p-1 text-white" @click="getMood()">Get Mood</button>
      <button class="bg-gray-400 hover:bg-gray-700 text-xs cursor-pointer rounded-sm p-1 text-white" @click="setMood()">Set Mood</button>
    </div>
  </div>
</template>

<style scoped>
</style>

<script>
import { ethers } from "ethers";
import TEST_ABI_JSON from "../static/testAbi.json"

const tokenContractAddress = '0xbE448c48bA49025Fbf0d7182D4FF55DB051e0d8d'
export default {
  data() {
    return {
      abi: TEST_ABI_JSON,
      provider: null,
      accounts: [],
      balance: null,
      contract: null,
      mood: null
    }
  },
  async mounted () {
    try {
      // 1
      const provider = new ethers.providers.Web3Provider(window.ethereum, "goerli");
      this.accounts = await provider.send("eth_requestAccounts", []);
      const signer = await provider.getSigner(this.accounts[0])
      this.contract = new ethers.Contract(tokenContractAddress, this.abi, signer)

      // 2
      // await provider.send("eth_requestAccounts", []).then(async () => {
      //   await provider.listAccounts().then(async (accounts) => {
      //     console.log(accounts)
      //     const signer = await provider.getSigner(accounts[0]);
      //     this.contract = new ethers.Contract(tokenContractAddress, this.abi, signer)
      //   });
      // });
    } catch (e) {
      console.log(e)
    }
  },
  methods: {
    async getMood () {
      const getMoodPromise = await this.contract.getMood();
      const Mood = await getMoodPromise;
      console.log(Mood)
    },
    async setMood () {
      const setMoodPromise = await this.contract.setMood(this.mood);
      await setMoodPromise;
    }
  }
}
</script>
