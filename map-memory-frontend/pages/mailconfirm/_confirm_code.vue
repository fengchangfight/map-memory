<template>
  <div class="confirm-page">
    <div class="confirm-content">
       <p>{{showMessage}}</p>
    </div>
    <div class="bottom">
      <!-- <app-footer></app-footer> -->
    </div>

  </div>
</template>

<script>
import {AXIOS} from '~/common/http-commons'
// import Footer from '~/components/Footer.vue'
export default {
    name: 'Confirm',
    components: {
      // 'app-footer': Footer
    },
    data() {
      return {
        confirm_code: '',
        showMessage: '验证中...',
      }
    },
    methods: {
    },
    created () {
      console.log('created');
    },
    mounted () {
      this.confirm_code = this.$route.params.confirm_code;
      if(this.confirm_code==null || this.confirm_code.length<1){
        this.$router.push("/login");
      }else{
        AXIOS.post('/register/confirm',{
          confirm_code: this.confirm_code
        }).then(response =>{
          if(response.data.ok==true){
            this.showMessage = "验证成功正在跳转";
            var count = 3;
            var storeThis = this;
            var refreshIntervalId = setInterval(function() {
                count = count - 1;
                if(count==0){
                  clearInterval(refreshIntervalId);
                  storeThis.$router.push("/login");
                }
            }, 1000);
          }else{
            this.showMessage = response.data.message;
            var count = 3;
            var storeThis = this;
            var refreshIntervalId = setInterval(function() {
                count = count - 1;
                if(count==0){
                  clearInterval(refreshIntervalId);
                  storeThis.$router.push("/login");
                }
            }, 1000);
          }
        }).catch(e => {
          console.log(e);
          this.showMessage = "error";
          var count = 3;
          var storeThis = this;
          var refreshIntervalId = setInterval(function() {
              count = count - 1;
              if(count==0){
                clearInterval(refreshIntervalId);
                storeThis.$router.push("/login");
              }
          }, 1000);
        })
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.confirm-page{
  flex-flow: column;
  width: 100%;
  height: 100vh;
  display: flex;
  left: 0;
  position:absolute;
}

.confirm-content {
  width: 100%;
  min-height: 700px;
  height: 700px;
  background-color: #f1f4f5;
}


.bottom{
  flex: 1;
  background-color: #F3F0EC;
}

</style>
