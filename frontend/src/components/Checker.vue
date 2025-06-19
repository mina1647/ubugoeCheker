<script setup lang="ts">
import { ref } from 'vue'

const data = ref<string[]>()
const query = ref<string>('')

const getUbugoe = async (name: string) => {
  const responce = await fetch(`https://ubugoechecker.trap.show/ubugoe/${name}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  })
  if (!responce.ok) {
    data.value = ['いないよ']
    return
  }

  data.value = await responce.json()
}
</script>

<template>
  <div class="checker">
    <h1>うぶごえチェッカー</h1>
    <!-- <p>ここにうぶごえチェッカーの内容が入ります。</p> -->
    <input type="text" v-model="query" placeholder="traQIDを入力してください" class="nameinput"/>
    <button @click="getUbugoe(query)">検索</button>

    <div v-for="item in data" :key="item">
      {{ item }}
    </div>
  </div>
</template>

<style>
.checker {
  background-color: #fff;
  color: #000;
  width: 100%;
  height: 100vh;
  top: 10vh;
  left: 0;
  position: fixed;
  padding: 1rem;
}

.nameinput {
  width: 30%;
}
</style>
