<template>
  <div class="demo-page" id="demoanchor">
    <div class="demo-content" >
      <div class="center-content">
        <div class="card">
          <div id="desc" >
            <div style="margin-top:22px;"><p>推荐本站最佳体验浏览器:</p></div>
            <div
              v-lazy-load-bg="'/imgs/browser.png'"
              id="browser-rec" />
          </div>
        </div>
        <el-button @click="back2topscroll" id="back2top" title="Go to top">回顶部</el-button>
        <el-button @click="back2bottomscroll" id="back2bottom" title="Go to top">至底部</el-button>

        <div class="card">
          <h3>如何在地图上添加记忆笔记?</h3>
          <p style="text-indent: 2em; margin-left: 5px;">
            在地图上任一点右键,弹出菜单选中"在此添加记忆"
          </p>
          <div
            v-lazy-load-bg="'/imgs/tianjiajiyi.jpg'"
            class="desc-img" />
          <div
            v-lazy-load-bg="'/imgs/tianjiajiyidian.jpg'"
            class="desc-img" />
          <div
            v-lazy-load-bg="'/imgs/xianshifangshi.jpg'"
            class="desc-img" />
        </div>

        <div class="card">
          <h3>如何显示某处经纬度？</h3>
          <p style="text-indent: 2em; margin-left: 5px;">
            在地图上任一点右键,弹出菜单选中"此处经纬度"
          </p>
          <div
            v-lazy-load-bg="'/imgs/cichujingweidu.jpg'"
            class="desc-img" />
          <div
            v-lazy-load-bg="'/imgs/jingweidu.jpg'"
            class="desc-img" />
        </div>

        <div class="card">
          <h3>如何添加并使用常用位置？</h3>
          <p style="text-indent: 2em; margin-left: 5px;">
            在地图上任一点右键,弹出菜单选中"设为常用位置"
          </p>
          <div
            v-lazy-load-bg="'/imgs/sheweichangyong.jpg'"
            class="desc-img" />
          <div
            v-lazy-load-bg="'/imgs/tianjiachangyong.jpg'"
            class="desc-img" />
          <div
            v-lazy-load-bg="'/imgs/changyongliebiao.jpg'"
            class="desc-img" />
        </div>

        <div id="place-bottom"/>
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
    name: 'Demo',
    components: {
      'app-footer': Footer
    },
    data() {
      return {
        static_base: '.',
        rest_base: '.'
      }
    },
    computed: {

   },
    created () {

    },
    methods: {
      back2bottomscroll(){
        var objDiv = document.getElementById("demoanchor");
        console.log(objDiv.scrollHeight);
        //objDiv.scrollTop = objDiv.scrollHeight;
        window.scrollTo(0,objDiv.scrollHeight);
      },
      back2topscroll(){
        window.scrollTo(0,0);
      },
      goPage (val) {
        this.$router.push(val)
      }

    },
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.demo-page{
  flex-flow: column;
  width: 100%;
  display: flex;
  left: 0;
  position:absolute;
}
.bottom{
  flex: 1;
  background-color: #F3F0EC;
}

.demo-content {
  width: 100%;
  min-height: 700px;
  background-color: #f1f4f5;
  flex: 1;
}

.center-content {
  max-width: 800px;
  height: 100%;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
}

.video-out-holder{
  background-size: 100% 100%;
  margin: 0 auto;
  width: 800px;
  min-height: 432px;
  height: 432px;
  margin-top: 20px;
  display: flex;
}

.video-inner-holder{
  margin: 0 auto;
  width: 657px;
  height: 432px;
  min-height: 432px;

  margin-top: 46px;
}
#desc{
  width: 440px;
  height: 65px;
  margin: 0 auto;
  margin-top: 50px;
  display: flex;
  font-family:sans-serif;

}
#browser-rec{
  width: 230px;
  height: 65px;
  background-size: 100% 100%;
  margin-left: 20px;
}
.desc-img{
  width: 730px;
  height: 365px;
  background-size: 100% 100%;
  margin-left: 20px;
}
.desc-img-short{
  width: 730px;
  height: 165px;
  background-size: 100% 100%;
  margin-left: 20px;
}
.desc-img-tall{
  width: 730px;
  height: 565px;
  background-size: 100% 100%;
  margin-left: 20px;
}
#place-bottom{
  width: 100%;
  height: 200px;
}
.card{
  margin-top: 20px;
  min-height: 150px;
  width: 100%;
  background-color: #FFFAF0;
}
h3{
  font-size: 14px;
  margin: 0 auto;
  margin-top: 20px;
  margin-bottom: 10px;
}
#back2top {
    /*display: none; /* Hidden by default */
    position: fixed; /* Fixed/sticky position */
    top: 70px; /* Place the button at the bottom of the page */
    right: 30px; /* Place the button 30px from the right */
    z-index: 99; /* Make sure it does not overlap */
    border: none; /* Remove borders */
    outline: none; /* Remove outline */
    background-color: #00FF00; /* Set a background color */
    color: white; /* Text color */
    cursor: pointer; /* Add a mouse pointer on hover */
    padding: 15px; /* Some padding */
    border-radius: 10px; /* Rounded corners */
    font-size: 12px; /* Increase font size */
    opacity: 0.5;
}

#back2top:hover {
    background-color: #555; /* Add a dark-grey background on hover */
}

#back2bottom {
    /*display: none; /* Hidden by default */
    position: fixed; /* Fixed/sticky position */
    top: 120px; /* Place the button at the bottom of the page */
    right: 30px; /* Place the button 30px from the right */
    z-index: 99; /* Make sure it does not overlap */
    border: none; /* Remove borders */
    outline: none; /* Remove outline */
    background-color: red; /* Set a background color */
    color: white; /* Text color */
    cursor: pointer; /* Add a mouse pointer on hover */
    padding: 15px; /* Some padding */
    border-radius: 10px; /* Rounded corners */
    font-size: 12px; /* Increase font size */
    opacity: 0.5;
}

#back2bottom:hover {
    background-color: #C0C0C0; /* Add a dark-grey background on hover */
}
</style>
