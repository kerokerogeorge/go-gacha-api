<template>
  <div class="w-full pl-24 pt-32">
    <p>tokenContractAddress: {{ addresses.contractAddress }}</p>
    <p>myAddress: {{ addresses.myAddress }}</p>
    <p>toAddress: {{ addresses.toAddress }}</p>
    <div v-if="!loading">
       <div class="my-3 py-3 border-b border-t border-gray-400 border-solid w-10/12">
        <button class="button button-lime" data-ripple-light="true" @click="buy">Create transaction</button>
      </div>

      <div class="my-3 py-3 border-b border-t border-gray-400 border-solid w-10/12">
        <div class="mb-3">
          <h2 class="text-lg text-gray-600">APPROVE</h2>
        </div>
        <div>
          <app-input v-model="values.approve.address" name="Address" />
          <app-input v-model="values.approve.amount" name="Amount" type="number" />
        </div>
        <button class="button button-green" data-ripple-light="true" @click="contractApprove">approve</button>
      </div>

      <div class="my-3 py-3 border-b border-t border-gray-400 border-solid w-10/12">
        <div class="mb-3">
          <h2 class="text-lg text-gray-600">TransferFrom</h2>
        </div>
        <div>
          <app-input v-model="values.transferFrom.owner" name="Owner" />
          <app-input v-model="values.transferFrom.buyer" name="Buyer" />
          <app-input v-model="values.transferFrom.amount" name="Amount" type="number" />
        </div>
        <button class="button button-green" data-ripple-light="true" @click="contractTransferFrom">transfer From</button>
      </div>

      <div class="my-3 py-3 border-b border-t border-gray-400 border-solid w-10/12">
        <div class="mb-3">
          <h2 class="text-lg text-gray-600">Buy</h2>
        </div>
        <div>
          <app-input v-model="values.transferFrom.owner" name="Owner" />
          <app-input v-model="values.transferFrom.buyer" name="Buyer" />
          <app-input v-model="values.transferFrom.amount" name="Amount" type="number" />
        </div>
        <button class="button button-green" data-ripple-light="true" @click="contractTransferFrom">transfer From</button>
      </div>

      <div class="my-3 py-3 border-b border-t border-gray-400 border-solid w-10/12">
        <button class="button button-green" data-ripple-light="true" @click="getTotalSupply">get total</button>
      </div>
    </div>
    <div v-else>
      <div>
        <img src="~/assets/loading.gif" alt="">
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>

<script>
import { ethers } from "ethers";
import TEST_ABI_JSON from "../../static/gachaAbi.json"
import AppInput from '~/components/AppInput'
import contractTestRepository from '~/repositories/contractTestRepository'

const tokenContractAddress = '0x9baCF53718c5137Ddf8152406b86F4920e13e1aC'
const myAddress = '0x6941cee0e87cb8ABE7A1985bf24c4f54CFeE9785'
const toAddress = '0xec64414617F2B65bB4a7adD57e82a1c5CF53B328'
export default {
  components: {
    AppInput
  },
  data() {
    return {
      addresses: {
        contractAddress: tokenContractAddress,
        myAddress: myAddress,
        toAddress: toAddress
      },
      abi: TEST_ABI_JSON,
      provider: null,
      signer: null,
      accounts: [],
      balance: null,
      contract: null,
      loading: false,
      values: {
        approve: {
          address: null,
          amount: null
        },
        transferFrom: {
          owner: null,
          buyer: null,
          amount: null
        },
        buy: {
          value: null,
          buyer: null,
          amount: null
        }
      }
    }
  },
  async mounted () {
    try {
      // 1
      this.provider = new ethers.providers.Web3Provider(window.ethereum, "goerli");
      console.log(this.provider)
      this.accounts = await this.provider.send("eth_requestAccounts", []);
      this.signer = await this.provider.getSigner(this.accounts[0])
      console.log('this.signer')
      // console.log(this.signer.signTransaction)
      this.contract = new ethers.Contract(tokenContractAddress, this.abi, this.signer)
      console.log("=====this.contract====")
      console.log(this.contract)



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
    async contractApprove () {
      const setApprovePromise = await this.contract.approve(this.values.approve.address, Number(this.values.approve.amount));
      console.log('setting...')
      this.loading = true
      await setApprovePromise.wait();
      this.loading = false
      console.log('setting finished')
      await setApprovePromise;
      console.log(await setApprovePromise)
    },
    async contractTransferFrom () {
      const setTransferFromPromise = await this.contract.transferFrom(this.values.transferFrom.owner, this.values.transferFrom.buyer, Number(this.values.transferFrom.amount));
      console.log('setting...')
      this.loading = true
      await setTransferFromPromise.wait();
      this.loading = false
      console.log('setting finished')
    },
    async buy () {
      // const { data } = await contractTestRepository.createBuyTokenTransaction({
      //   fromAddress: this.addresses.myAddress,
      //   toAddress: this.addresses.toAddress,
      //   contractAddress: this.addresses.contractAddress
      // })
      // console.log(data)
      const options = {value: ethers.utils.parseUnits("100.0", "gwei")}
      const setTransferFromPromise = await this.contract.buy(options);
      console.log('buying...')
      this.loading = true
      await setTransferFromPromise.wait();
      this.loading = false
      console.log('buying finished')
    },
    async getTotalSupply () {
      const getTotalSupply = await this.contract.totalSupply()
      // const amount = ethers.BigNumber.from()
      // console.log(amount)
      let res = getTotalSupply.toHexString()
      // console.log(res)
      // res = ethers.utils.parseUnits(res, 18)
      // console.log(res)
      res = getTotalSupply.toString()
      console.log(res)
      // res = ethers.BigNumber.from(res)
      // console.log(res)
    },
    async getMood () {
      const getMoodPromise = await this.contract.getMood();
      const Mood = await getMoodPromise;
      this.currentMood = Mood
      console.log(Mood)
    },
    async setMood () {
      const setMoodPromise = await this.contract.setMood(this.mood);
      console.log('setting mood ...')
      this.loading = true
      await setMoodPromise.wait();
      this.loading = false
      console.log('setting mood finished')
      await setMoodPromise;
      console.log(await setMoodPromise)
    }
  }
}
</script>
