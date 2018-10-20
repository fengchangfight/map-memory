<template>
  <div 
    v-loading="wholePageLoading" 
    class="register-page">
    <div class="main-content">
      <div class="register-form">
        <div class="register-header">注册新帐号 加入地图记忆</div>
        <div class="register-body">
          <b-tabs>
            <!-- <b-tab title="手机注册" active>
                     <div class="inner-body">
                     <el-form ref="form" :model="form" label-width="87px">
                       <el-form-item label="手机号码:">
                         <el-input class="phoneIcon fl" v-verify="form.phone" v-model="form.phone" @change="phoneValidate"></el-input>
                         <label v-if="phoneError" class="phone-invalid">*{{phoneErrorMessage}}*</label>
                       </el-form-item>
                       <el-form-item label="创建密码:">
                         <el-input type="password" v-verify="form.password" v-model="form.password" @blur="passwordLostFocus"></el-input>
                         <label v-if="passwordError" class="password-invalid">*{{passwordErrorMessage}}*</label>
                       </el-form-item>
                       <el-form-item label="短信验证码:">
                         <el-col :span="11"><el-input v-verify="form.authcode" v-model="form.authcode"></el-input></el-col>
                         <el-col class="line" :span="3"><label v-if="authcodeError" class="authcode-invalid">*{{authcodeErrorMessage}}*</label></el-col>
                         <el-col  :span="8"><el-button v-on:click="sendCode" :disabled="disableAuthButton || countingDown">发送验证码{{remainingTime}}</el-button></el-col>
                       </el-form-item>
                       <el-button v-on:click="register('phone')" type="primary">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;完&nbsp;成&nbsp;注&nbsp;册&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</el-button>
                       <div class="agreement" style="width:420px;height:30px;float:left;margin-top:10px;text-align:right;">
                         点击完成注册表示您同意并遵守
                         <router-link to="regulation">地图记忆服务条款</router-link>
                       </div>
                     </el-form>
                     </div>
                   </b-tab> -->
            <b-tab title="邮箱注册">
              <div class="inner-body">
                <el-form 
                  ref="form" 
                  :model="form" 
                  label-width="87px">
                  <input 
                    type="text" 
                    style="display:none">
                  <el-form-item label="邮箱:">
                    <el-input 
                      v-verify="form.email" 
                      v-model="form.email" 
                      class="phoneIcon fl" 
                      autocomplete="off" 
                      @change="emailValidate"/>
                    <label 
                      v-if="emailError" 
                      class="email-invalid">*{{ emailErrorMessage }}*</label>
                  </el-form-item>
                  <el-form-item label="创建密码:">
                    <el-input 
                      v-verify="form.password" 
                      v-model="form.password" 
                      type="password" 
                      @blur="passwordLostFocus"/>
                    <label 
                      v-if="passwordError" 
                      class="password-invalid">*{{ passwordErrorMessage }}*</label>
                  </el-form-item>
                  <el-form-item label="重复密码:">
                    <el-input 
                      v-model="form.passwordre" 
                      type="password" 
                      @blur="passwordreLostFocus"/>
                    <label 
                      v-if="passwordreError" 
                      class="password-invalid">*{{ passwordreErrorMessage }}*</label>
                  </el-form-item>

                  <el-button 
                    type="primary" 
                    @click="register('email')">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;完&nbsp;成&nbsp;注&nbsp;册&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</el-button>
                  <div 
                    class="agreement" 
                    style="width:420px;height:30px;float:left;margin-top:10px;text-align:right;">
                    点击完成注册表示您同意并遵守
                    <router-link to="regulation">地图记忆服务条款</router-link>
                  </div>
                </el-form>
              </div>
            </b-tab>
          </b-tabs>


        </div>
      </div>
    </div>
    <div class="bottom">
      <app-footer/>
    </div>

  </div>
</template>

<script>
import {AXIOS} from '~/common/http-commons'
import Footer from '~/components/Footer.vue'

export default {
    name: 'Register',
    components: {
      'app-footer': Footer
    },
    verify: {
            age:"required",
            form: {
                phone: ["required","mobile"],
                email: ["required", "email"],
                password: {
                   minLength:6,
                   message: "密码不得小于6位"
                },
                authcode: ["required"]
            }
    },
    data() {
      return {
          wholePageLoading: false,
          countingDown: false,
          disableAuthButton : true,
          phoneError: false,
          emailError: false,
          passwordError: false,
          passwordreError: false,
          authcodeError: false,
          phoneErrorMessage: '',
          emailErrorMessage: '',
          passwordErrorMessage: '',
          passwordreErrorMessage: '',
          authcodeErrorMessage: '',
          remainingTime : '',
          form: {
            phone: '',
            email: '',
            password: '',
            authcode: ''
          }
      }
    },
    created () {

    },
    methods: {
      goPage (which) {
        this.$route.push(which);
      },

      authCodeValid(){
        this.$verify.check();
        if(this.$verify.$errors.form.authcode.length>0){
          this.authcodeError = true;
          this.authcodeErrorMessage = this.$verify.$errors.form.authcode[0];
        }else{
          this.authcodeError = false;
          this.authcodeErrorMessage = '';
        }
        return !this.authcodeError;
      },

      authLostFocus () {
        this.authCodeValid();
      },
      passwordreLostFocus(){
        if(this.form.password!=this.form.passwordre){
          this.passwordreError = true;
          this.passwordreErrorMessage = '两次输入密码不一致';
        }else{
          this.passwordreError = false;
          this.passwordreErrorMessage = '';
        }
      },
      passwordValid(){
        this.$verify.check();
        if(this.$verify.$errors.form.password.length>0){
          this.passwordError = true;
          this.passwordErrorMessage = this.$verify.$errors.form.password[0];
        }else{
          this.passwordError = false;
          this.passwordErrorMessage = '';
        }
        return !this.passwordError;
      },
      passwordLostFocus () {
        this.passwordValid();
      },
      emailValidate(){
        this.$verify.check()
        if(this.$verify.$errors.form.email.length>0){
          this.emailError = true;
          this.emailErrorMessage = this.$verify.$errors.form.email[0];
        }else{
          this.emailError = false;
          this.emailErrorMessage = '';
        }
        return !this.emailError;
      },
      phoneValidate(){
        this.$verify.check()
        if(this.$verify.$errors.form.phone.length>0){
          this.phoneError = true;
          this.phoneErrorMessage = this.$verify.$errors.form.phone[0];
          this.disableAuthButton = true;
        }else{
          this.phoneError = false;
          this.phoneErrorMessage = '';
          this.disableAuthButton = false;
        }
        return !this.disableAuthButton;
      },
      sendCode() {
        this.phoneValidate();
        if(this.disableAuthButton){
          return;
        }
        this.countingDown = true;
        var storeThis = this;
        var count = 61;
        var refreshIntervalId = setInterval(function() {
            count = count - 1;
            storeThis.remainingTime = "("+count+"秒)"
            if(count==0){
              storeThis.countingDown = false;
              storeThis.disableAuthButton = false;
              storeThis.remainingTime = ""
              clearInterval(refreshIntervalId);
            }
        }, 1000);


        AXIOS.post('/register/authcode', {
          "phone":  this.form.phone
        }).then(response => {
          console.log(response);
        }).catch(e => {
          console.log(e);
        })
      },
      register (type) {
        if(type=='phone'){
          // phone register
          if(!this.phoneValidate() || !this.passwordValid() || !this.authCodeValid()){
            console.log("Phone register info not valid");
          }else{
            this.wholePageLoading=true;
            AXIOS.post('/register', {
              "phone":  this.form.phone,
              "password": this.form.password,
              "authcode": this.form.authcode
            }).then(response => {
              this.$notify({
                title: '成功',
                type: 'success',
                message: '注册成功'
              });
              this.wholePageLoading=false;
              this.$router.push("/register-success");
            }).catch(e => {
              this.$notify.error({
                title: '错误',
                message: e.response.data.message
              });
              this.wholePageLoading=false;
            })
          }
        }else{
          // email register
          if(!this.emailValidate() || !this.passwordValid() || this.form.password!=this.form.passwordre){
            console.log('Email register info not valid');
          }else{
            this.wholePageLoading=true;
            AXIOS.post('/register/email', {
              "email":  this.form.email,
              "password": this.form.password
            }).then(response => {
              if(response.data.ok==true){
                this.$notify({
                  title: '成功',
                  type: 'success',
                  message: response.data.message
                });
                this.wholePageLoading=false;
                this.$router.push("/email-register-success");
              }else{
                this.$notify.error({
                  title: '错误',
                  message: response.data.message
                });
                this.wholePageLoading=false;
              }
            }).catch(e => {
              this.$notify.error({
                title: '错误',
                message: e.response.data.message
              });
              this.wholePageLoading=false;
            })
          }
        }

      }
    },
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.register-page{
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 100%;
}

.main-content{
  background-color: #f1f4f5;
  width: 520px;
  min-height: 700px;
  height: 700px;
  width: 100%;
  margin-left: 0 auto;
  flex: 1;
}

.register-form{
  margin: 0 auto;
  margin-top: 100px;
  width: 520px;
  height: 478px;
  background-color: gray;
  display: flex;
  flex-direction: column;
  border:1px solid #cfd6de;
}
.register-header{
  width: 100%;
  background-color: white;
  height: 74px;
  font-size: 24px;
  color: #333;
  text-align: center;
  width: 100%;
  line-height: 74px;
  font-family: 'Microsoft Yahei',Tahoma,Arial,Helvetica,Simsun,STHeiti;
}
.register-body{
  width: 100%;
  background-color: #f8f8f8;
  flex: 1 1 auto;
  display: flex;
}

.inner-body{
  width: 360px;
  margin-left: 50px;
  height: 200px;
  margin-top: 30px;
}

.agreement{
      font: 12px/18px STHeiti,'Microsoft YaHei',arial,\5b8b\4f53;
}

.phone-invalid, .password-invalid, .authcode-invalid, .email-invalid{
  color: red;
  font: 12px/18px STHeiti,'Microsoft YaHei',arial,\5b8b\4f53;
}

.bottom{
  flex: 1;
  background-color: #F3F0EC;
}

</style>
