<template>
  <div class="login-wrap">
    <van-form>
      <van-field
        v-model="form.username"
        name="用户名"
        label="用户名"
        placeholder="用户名"
      />
      <van-field
        v-model="form.password"
        type="password"
        name="密码"
        label="密码"
        placeholder="密码"
      />
      <div style="margin: 16px;">
        <van-button round block type="primary" @click="onSubmit">
          登录
        </van-button>
      </div>
    </van-form>
  </div>
</template>
<script>
import {reactive} from 'vue'
import { useRouter } from 'vue-router';
export default {
  setup(){
    const router = useRouter();
    const onSubmit = ()=>{
      $.post("/api/toLogin",{
        username:form.username,
        password:form.password
      },res=>{
        if(res.status === 'success'){
          router.push("/chat")
        }
        else{
          alert(res.msg)
        }
      })
    }
    const form = reactive({
      username: "",
      password: ""
    })
    return {
      form,
      onSubmit
    }
  }
}
</script>
<style lang="less" scoped>
@import "./login.less";
</style>