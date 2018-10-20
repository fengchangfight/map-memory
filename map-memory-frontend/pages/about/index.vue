<template>
  <div class="about-page">

    <div class="about-content">
      <div id="introduction">
        <h3>本站概况</h3>
        <hr width="690px">
        <div id="article">
          <p style="text-indent: 2em;">地图记忆创建于2018年，受<a href="https://book.douban.com/subject/1032501/">《心理学与生活》</a>关于记忆的部分以及<a href="https://book.douban.com/subject/27076474/">《超好用的宫殿记忆法》</a>的启发而来。帮助大家整理自己的记忆，知识，笔记等。</p>
          <p style="text-indent: 2em;">本站既可以作为普通笔记本软件使用，也可以作为记忆术的辅助工具（空间记忆法，关联记忆法等），兼具丰富性和灵活性。</p>
          <p style="text-indent: 2em;">本站工具在使用方式上尽力保持极简的风格，让只要会电脑基本操作的人都能零门槛使用，但如果有任何使用方式上的问题，仍然可以参考 <router-link to="/demo">演示</router-link>或者<router-link to="/contact">联系本站</router-link></p>
          <p style="text-indent: 2em;">感谢百度地图API提供的优质服务</p>

          <p style="text-align: right;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; 最后更新：2018年7月</p>
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
    name: 'About',
    components: {
      'app-footer': Footer
    },
    data() {
      return {

      }
    },
    created () {

    },
    methods: {

    },
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.about-page{
  flex-flow: column;
  width: 100%;
  height: 100vh;
  display: flex;
  left: 0;
  position:absolute;
}

.about-content {
  width: 100%;
  min-height: 700px;
  height: 700px;
  background-color: #f1f4f5;
}

#introduction{
  width: 700px;
  height: 500px;
  margin: 0 auto;
  margin-top: 100px;

}

h3{
  font-size: 20px;
  font-weight: normal;
  color: #0e6aad;
}

.bottom{
  flex: 1;
  background-color: #F3F0EC;
}
</style>
