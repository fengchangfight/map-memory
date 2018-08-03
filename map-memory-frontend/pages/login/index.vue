<template>
  <div class="login-main" v-loading="wholepageloading">
      <div class="login-form">
          <div class="input-grouper">
              <el-input style="" v-model="username" placeholder="用户名"></el-input>
              <el-input style="margin-top: 17px;" v-model="password" type="password" placeholder="密码"></el-input>
              <a class="si-login green-button" @click="login" >登录</a>
              <a class="si-login blue-button"  @click="goPage('/register')">注册</a>
          </div>

      </div>
  </div>
</template>

<script>

import Qs from 'qs'
import {AXIOS} from '~/common/http-commons'
import swal from 'sweetalert'

export default {
  name: 'Login',
  components: {
  },
  methods: {
      goPage(val){
          this.$router.push(val);
      },
      whoamifoo(){
        AXIOS.get('/api/v1/user/whoami').then(response => {
              // JSON responses are automatically parsed.
              this.whoami = response.data;
              if(Boolean(this.whoami)){
                this.$router.push('/working');
              }
            }).catch(e => {

            })
      },
      login () {
          if(this.username==null || this.username.trim().length<1 || this.password==null || this.password.trim().length<1){
              swal ( "提示" ,  "用户名或密码不能为空" ,  "info" );
              return;
          }
        this.wholepageloading = true;
        var data = {'username': this.username, 'password':this.password, 'timeout':1000};
        AXIOS.post('/login', Qs.stringify(data), {headers:{'Content-Type':'application/x-www-form-urlencoded'}}).then(response => {
            //location.reload();
            window.location.href = '/working';
            //this.$router.push('/working');
        }).catch(e => {
          this.wholepageloading = false;
          this.$notify.error({
            title: '错误',
            message: '用户名或密码错误'
          });
          this.$router.push('/login')
          console.log(e)
        })
      }
  },
  mounted(){
    this.whoamifoo();
  },
  data(){
    return {
        wholepageloading: false,
      username: '',
      password: ''
    }
  }
}
</script>

<style scoped>
.login-main {
  background-image: url('~/assets/yu.jpeg');
  background-size: cover;
  width: 100%;
  height: 100vh;
  position: relative;
}

.login-form{
    position: absolute;
    width: 380px;
    height: 260px;
    background-color: rgba(181, 164, 204, 0.5);
    top: 50%;
    left: 50%;
    transform: translate(-50%,-50%);
    -moz-box-shadow: 1px 1px 25px #202020; /* 老的 Firefox */
    box-shadow: 1px 1px 25px #202020;
    border-radius: 20px;
}
.input-grouper{
    width: 80%;
    height: 80%;
    /* background-color: aliceblue; */
    margin: 0 auto;
    margin-top: 20px;
    display: block;
}
.si-login,.si-login:visited{
    padding:5px 132px;
    border-radius: 10px;
    font-size: 19px;
    text-decoration: none;
    color: white !important;
    display: inline-block;
    margin: 0 auto;
    margin-top: 17px;
}

.green-button{
    background-color: rgb(39, 199, 33);
}
.green-button:hover{
    cursor: pointer;
    background-color: rgba(44, 33, 199,0.5);
    transform: translateY(-2px);
}

.blue-button{
    background-color: rgb(44, 33, 199);
}
.blue-button:hover{
    cursor: pointer;
    background-color: rgba(39, 199, 33,0.5);
    transform: translateY(-2px);
}



</style>
