<template>
  <div>
    <div class="footer">
      <div class="footer-row-1">
        <a 
          target="_blank" 
          @click="goPage('/')">网站首页</a> -
        <a 
          target="_blank" 
          @click="goPage('/about')">关于我们</a> -
        <a 
          target="_blank" 
          @click="goPage('/contact')">联系站长</a> -
        <a 
          target="_blank" 
          @click="goPage('/latestnews')">网站动态</a>
      </div>
      <div class="footer-row-2">
        ©CopyRight 2018-2018 DITUJIYI.COM All Rights Reserved. 地图记忆 版权所有
      </div>
      <div class="footer-row-3">
        备案号： <a href="http://www.miitbeian.gov.cn/">沪ICP备13033858号-5</a>
      </div>
    </div>
  </div>
</template>

<script>
  import {
    AXIOS
  } from '../common/http-commons'
  export default {
    name: 'Footer',
    data() {
      return {

      };
    },
    mounted() {},
    methods: {
      goPage(val) {
        this.$router.push(val)
      }
    }

  }
</script>

<style scoped>
  .footer {
    width: 100%;
    min-height: 100px;
    background-color: #2f3033;
    display: flex;
    flex-direction: column;
    flex: 0 1 100px;
    color: #d3dce6;
  }

  .footer-row-1 {
    margin: 0 auto;
    margin-top: 10px;
    height: 20px;
    width: 280px;
    font: 12px/1.5 tahoma, arial, \5b8b\4f53;
    line-height: 24px;
  }

  .footer-row-2 {
    margin: 0 auto;
    margin-top: 5px;
    height: 20px;
    width: 500px;
    font: 12px/1.5 tahoma, arial, \5b8b\4f53;
    line-height: 24px;
  }

  .footer-row-3 {
    margin: 0 auto;
    margin-top: 5px;
    height: 20px;
    width: 200px;
    font: 12px/1.5 tahoma, arial, \5b8b\4f53;
    line-height: 24px;
  }

  a:hover {
    cursor: pointer;
    text-decoration: underline;
  }
</style>
