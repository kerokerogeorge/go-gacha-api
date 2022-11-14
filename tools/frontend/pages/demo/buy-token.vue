<template>
  <div class="w-full pl-24 pt-32">
    <p>tokenContractAddress: {{ addresses.tokenContractAddress }}</p>
    <p>vendorContractAddress: {{ addresses.vendorContractAddress }}</p>
    <p>myAddress: {{ addresses.myAddress }}</p>
    <p>toAddress: {{ addresses.toAddress }}</p>
    <div v-if="!loading">
      <div class="my-3 py-3 border-b border-t border-gray-400 border-solid w-10/12">
        <div class="mb-3">
          <h2 class="text-lg text-gray-600">TEST TOKENS</h2>
        </div>
        <div>
          <app-input v-model="values.buy.value" name="TOKENS TO BUY" />
        </div>
        <div class="my-3">ETHER:{{ values.buy.value / 1000000 }}</div>
        <div class="my-3">WEI:{{ (values.buy.value / 1000000 )* 10 **18 }}</div>
        <button class="button button-green" data-ripple-light="true" @click="buy">PURCHASE</button>
      </div>
      <div class="my-3 py-3 border-b border-t border-gray-400 border-solid w-10/12">
        <div class="mb-3">
          <h2 class="text-lg text-gray-600">TRANSFER TOKENS</h2>
        </div>
        <div>
          <app-input v-model="values.transfer.value" name="TOKENS TO transfer" />
        </div>
        <div class="my-3">ETHER:{{ values.transfer.value / 1000000 }}</div>
        <div class="my-3">WEI:{{ (values.transfer.value / 1000000 )* 10 **18 }}</div>
        <button class="button button-green" data-ripple-light="true" @click="getPayload">Get Payload</button>
        <button class="button button-blue" data-ripple-light="true" @click="transfer">transfer</button>
      </div>
      <div class="my-3 py-3 border-b border-t border-gray-400 border-solid w-10/12">
        <div class="mb-3">
          <h2 class="text-lg text-gray-600">GET BLOCK</h2>
        </div>
        <button class="button button-blue" data-ripple-light="true" @click="getBlock">get Block</button>
      </div>
      <div class="my-3 py-3 border-b border-t border-gray-400 border-solid w-10/12">
        <div class="mb-3">
          <h2 class="text-lg text-gray-600">GET FeeHistory</h2>
        </div>
        <button class="button button-blue" data-ripple-light="true" @click="getFeeHistory">GET FeeHistory</button>
      </div>
    </div>
    <div v-else>
      <div>
        <img src="~/assets/loading.gif" alt="">
      </div>
    </div>
    <template v-if="payload.transfer">
      <div class="text-xs text-gray-500 font-semibold p-2 bg-gray-100 rounded-md my-2">
        <p>chainID: {{payload.transfer.chainId}}</p>
        <p>data: {{payload.transfer.data}}</p>
        <p>gas: {{payload.transfer.gas}}</p>
        <p>gasFeeCap: {{payload.transfer.gasFeeCap}}</p>
        <p>gasTipCap: {{payload.transfer.gasTipCap}}</p>
        <p>nonce: {{payload.transfer.nonce}}</p>
        <p>to: {{payload.transfer.to}}</p>
      </div>
    </template>
    <template v-if="tx.data">
      <div class="text-xs text-gray-500 font-semibold p-2 bg-gray-100 rounded-md">
        <p>chainID: {{tx.data.chainId}}</p>
        <p>input: {{tx.data.input}}</p>
        <p>gas: {{tx.data.gas}}</p>
        <p>maxFeePerGas: {{tx.data.maxFeePerGas}}</p>
        <p>maxPriorityFeePerGas: {{tx.data.maxPriorityFeePerGas}}</p>
        <p>nonce: {{tx.data.nonce}}</p>
        <p>to: {{tx.data.to}}</p>
      </div>
    </template>
  </div>
</template>

<style scoped>
</style>

<script>
import { ethers } from "ethers";
import { tokenABI, vendorABI } from '../../vendorAbi';
import AppInput from '~/components/AppInput'
import Web3 from 'web3';
import contractTestRepository from '~/repositories/contractTestRepository'
const web3 = new Web3(Web3.givenProvider);

const tokenContractAddress = '0x6a7edAd9c7f49Bf215Add73e5d8F8Cb550177297'
const vendorContractAddress = '0x7dBa3cc9bDf7B3F79dcDD90B0c19768190a5aC5b'
const myAddress = '0x6941cee0e87cb8ABE7A1985bf24c4f54CFeE9785'
const toAddress = '0xec64414617F2B65bB4a7adD57e82a1c5CF53B328'
export default {
  components: {
    AppInput
  },
  data() {
    return {
      addresses: {
        tokenContractAddress: tokenContractAddress,
        vendorContractAddress: vendorContractAddress,
        myAddress: myAddress,
        toAddress: toAddress
      },
      vendorAbi: vendorABI,
      tokenAbi: tokenABI,
      provider: null,
      signer: null,
      accounts: [],
      balance: null,
      contract: null,
      loading: false,
      values: {
        buy: {
          value: null,
        },
        transfer: {
          value: null
        }
      },
      web3: web3,
      vendor: null,
      payload: {
        transfer: null
      },
      tx: {
        transfer: {
          chainId: null,
          gas: null,
          input: null,
          maxFeePerGas: null,
          maxPriorityFeePerGas: null,
          nonce: null,
          value: null,
          to: null
        },
        data: {
          chainId: null,
          gas: null,
          input: null,
          maxFeePerGas: null,
          maxPriorityFeePerGas: null,
          nonce: null,
          value: null,
          to: null
        }
      }

    }
  },
  async mounted () {
    try {
      this.provider = new ethers.providers.Web3Provider(window.ethereum, "goerli");
      this.accounts = await this.provider.send("eth_requestAccounts", []);
      this.signer = await this.provider.getSigner(this.accounts[0])
      console.log(this.accounts[0])
      console.log(vendorContractAddress)
      console.log(this.vendorAbi)
      this.contract = new this.web3.eth.Contract(this.vendorAbi, vendorContractAddress);
      // this.contract = new ethers.Contract(this.addresses.vendorContractAddress, this.vendorAbi, this.signer)
    } catch (e) {
      console.log(e)
    }
  },
  methods: {
    async getPayload () {
      const amount = this.values.transfer.value
      const { data } = await contractTestRepository.getTokenTransferTransactionPayload({
        fromAddress: this.addresses.myAddress,
        toAddress: this.addresses.toAddress,
        contractAddress: this.addresses.tokenContractAddress,
        amount: Number(this.web3.utils.toWei(amount.toString(), "ether"))
      })
      this.payload.transfer = data.Payload
      this.tx.transfer = data.Tx
      const maxFeePerGas = Number(this.web3.utils.toBN(this.tx.transfer.maxFeePerGas).toString())
      const maxPriorityFeePerGas = Number(this.web3.utils.toBN(this.tx.transfer.maxPriorityFeePerGas).toString())
      console.log(maxFeePerGas)
      console.log(maxPriorityFeePerGas)
      console.log(maxFeePerGas > maxPriorityFeePerGas)
      Object.assign(this.tx.data, {
        chainId: this.web3.utils.toBN(data.Tx.chainId).toString(),
        gas: this.web3.utils.toBN(data.Tx.gas).toString(),
        input: this.web3.utils.toBN(data.Tx.input).toString(),
        maxFeePerGas: this.web3.utils.toBN(data.Tx.maxFeePerGas).toString(),
        maxPriorityFeePerGas: this.web3.utils.toBN(data.Tx.maxPriorityFeePerGas).toString(),
        nonce: this.web3.utils.toBN(data.Tx.nonce).toString(),
        value: this.web3.utils.toBN(data.Tx.value).toString(),
        to: data.Tx.to
      })
      console.log(this.tx.data)
    },
    async transfer () {
      const accounts = await this.web3.eth.getAccounts();
      const maxFeePerGas = Number(this.web3.utils.toBN(this.tx.transfer.maxFeePerGas).toString())
      const maxPriorityFeePerGas = Number(this.web3.utils.toBN(this.tx.transfer.maxPriorityFeePerGas).toString())
      const request = {
          from: accounts[0],
          to: this.tx.transfer.to,
          gas: this.tx.transfer.gas,
          value: this.tx.transfer.value,
          maxFeePerGas:
            maxFeePerGas > maxPriorityFeePerGas ? maxFeePerGas : maxPriorityFeePerGas,
          maxPriorityFeePerGas: maxPriorityFeePerGas,
          // 参考
          // https://goerli.etherscan.io/tx/0x2c74a240ca53e6411a33a0a1def610ae4855c3d7bcb9184b243342507225e713
          nonce: this.tx.transfer.nonce,
          chainId: this.tx.transfer.chainId,
          input: this.tx.transfer.input
      }
      console.log("request: ", request)
      this.loading = true
      // const tx = await this.contract.methods.transfer().send(request)
      const tx = await this.web3.eth.sendTransaction(request)
        .on('receipt', function(receipt){
          console.log("==called==")
          console.log(receipt)
        })
        .on('confirmation', function(confirmationNumber, receipt){
          console.log("confirmationNumber: ", confirmationNumber)
          console.log("receipt: ", receipt)
        });
      this.loading = false
      console.log("SUCCESS: ", tx);
    },
    async buy () {
      try {
        // console.log("called")
        // this.loading = true
        // const wei = ethers.utils.parseEther(this.values.buy.value)
        // console.log("value: ", wei.toString())
        // const options = {value: wei.toString()}
        // const request = await this.contract.buyTokens(options)
        // alert("You have successfully purchased TEST tokens!")
        // console.log(request);
        // this.loading = false
        const amount = this.values.buy.value / 1000000
        console.log("=========")
        console.log("amount", amount)
        console.log("ethers.utils.parseEther(amount):", ethers.utils.parseEther(amount.toString()))
        console.log("this.web3.utils.toWei:", this.web3.utils.toWei(amount.toString(), "ether"))
        console.log("=========")
        const accounts = await this.web3.eth.getAccounts();
        console.log('accounts')
        console.log(accounts)
        // this.vendor = new this.web3.eth.Contract(
        //   this.vendorAbi,
        //   vendorContractAddress
        // );
        const gasPrice = await this.web3.eth.getGasPrice();
        // const gasAmount = await this.getGasAmountForContractCall(accounts[0], tokenContractAddress)
        // console.log("value: ", this.web3.utils.toWei(this.values.buy.value.toString(), "ether"))
        // console.log("gasPrice: ",gasPrice)
        // console.log("gasAmount: ", gasAmount)
        const request = await this.contract.methods.buyTokens().send({
          from: accounts[0],
          gasPrice: gasPrice,
          gas: 300000,
          value: this.web3.utils.toWei(amount.toString(), "ether")
        })
        alert("You have successfully purchased TEST tokens!");
        console.log(request);
      } catch (err) {
        console.error(err);
        alert("Error purchasing tokens");
      }
    },
    async getGasAmountForContractCall (fromAddress, toAddress) {
      const data = await this.contract.methods.buyTokens().encodeABI()
      console.log(data)
      const gasAmount = await this.web3.eth.estimateGas({to: toAddress, data: data})
      console.log(gasAmount)
      return gasAmount
    },
    async getBlock () {
      const { data } = await this.$axios.get(`http://localhost:8000/test/block`);
      console.log(data)
    },
    async getFeeHistory () {
      const historicalBlocks = 10;
      this.web3.eth.getFeeHistory(historicalBlocks, "pending", [1, 50, 99]).then((feeHistory) => {
        const blocks = this.formatFeeHistory(feeHistory, false, historicalBlocks)
        const slow    = this.avg(blocks.map(b => b.priorityFeePerGas[0]))
        const average = this.avg(blocks.map(b => b.priorityFeePerGas[1]))
        const fast    = this.avg(blocks.map(b => b.priorityFeePerGas[2]))

        console.log("blocks")
        console.log(blocks)
        this.web3.eth.getBlock("pending").then((block) => {
          const baseFeePerGas = Number(block.baseFeePerGas)
          console.log("Manual estimate:", {
            slow: slow + baseFeePerGas,
            average: average + baseFeePerGas,
            fast: fast + baseFeePerGas,
          })
        })
      })
    },
    formatFeeHistory(result, includePending, historicalBlocks) {
      let blockNum = result.oldestBlock;
      let index = 0;
      const blocks = [];
      while (blockNum < result.oldestBlock + historicalBlocks) {
        blocks.push({
          number: blockNum,
          baseFeePerGas: Number(result.baseFeePerGas[index]),
          gasUsedRatio: Number(result.gasUsedRatio[index]),
          priorityFeePerGas: result.reward[index].map(x => Number(x)),
        });
        blockNum += 1;
        index += 1;
      }
      if (includePending) {
        blocks.push({
          number: "pending",
          baseFeePerGas: Number(result.baseFeePerGas[historicalBlocks]),
          gasUsedRatio: NaN,
          priorityFeePerGas: [],
        });
      }
      return blocks;
    },
    avg(arr) {
      const sum = arr.reduce((a, v) => a + v);
      return Math.round(sum/arr.length);
    }
  }
}
</script>
