<template>
  <div class="register-success-page">
    <div class="register-success-content">
      <p>注册成功，正在跳转...</p>
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
    name: 'RegisterSuccess',
    components: {
      // 'app-footer': Footer
    },
    data() {
      return {

      }
    },
    created () {

    },
    mounted () {
      var count = 3;
      var storeThis = this;
      var refreshIntervalId = setInterval(function() {
          count = count - 1;
          if(count==0){
            clearInterval(refreshIntervalId);
            storeThis.$router.push("/login");
          }
      }, 1000);
    },
    methods: {
    },
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.register-success-page{
  flex-flow: column;
  width: 100%;
  height: 100vh;
  display: flex;
  left: 0;
  position:absolute;
}

.register-success-content {
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
