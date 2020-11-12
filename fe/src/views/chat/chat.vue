<template>
  <div class="chat-wrap">
    <van-field type="textarea" v-model="msg" placeholder="输入聊天内容" />
    <van-button type="primary" @click="send">发送</van-button>
  </div>
</template>
<script>
import {ref} from "vue"
export default {
  setup(){
    const msg = ref("");
    let ws = new WebSocket("ws://127.0.0.1:8000/ws");
    ws.onopen = function(){
      console.log("连接")
    }
    ws.onmessage = function(e){
      console.log("收到消息：",e.data)
    }
    const send = ()=>{
      ws.send(msg.value);
    }
    return {
      msg,
      send
    }
  }
}
</script>
<style lang="less" scoped>
@import "./chat.less";
</style>