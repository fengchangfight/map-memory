<template>
  <div class="home-page">
    <div class="main-image"  v-lazy-load-bg="'/imgs/home-back-grid-icon.jpg'">
      <div id="title">
        <h1 style="color: white;">基于地图的极简记忆本</h1>
      </div>

      <div class="button-middle-holder">
        <a v-on:click="goPage('/login')" class="home-btn">
          <span>进入</span>
        </a>
        <a v-on:click="goPage('/demo')" class="home-btn">
          <span>演示</span>
        </a>
        <!-- <button v-on:click="goDemo" class="demo-button button"></button>
        <button v-on:click="enter" class="enter-button button"></button> -->
      </div>
    </div>


    <div class="page-con" style="background: rgb(255, 255, 255); min-height: 582px;">

        <el-carousel :interval="3000" arrow="always" height="582px">
            <el-carousel-item v-for="item in intro_items" :key="item.title">
               <div class="intro-container">
                      <div class="h2-div">{{item.title}}</div>
                      <div class="h3-div">{{item.description}}</div>
                      <div class="des-item" v-lazy-load-bg="'/imgs/'+item.image" ></div>
                </div>
            </el-carousel-item>
          </el-carousel>

    </div>

    <div class="bottom">
      <app-footer></app-footer>
    </div>
  </div>
</template>

<script>
import {AXIOS} from '~/common/http-commons'
import Footer from '~/components/Footer.vue'
export default {
    name: 'HomePage',
    components: {
       'app-footer': Footer
    },
    data() {
      return {
        rest_base: '.',
        intro_items: [
           {
              title: '让记忆同位置关联起来',
              description: '听说过明朝来华传教士利玛窦的宫殿记忆法吗？让记忆对象跟已知事物关联起来，极大激发你的记忆潜能！',
              image: 'intro_1.png'
           },{
              title: '读万卷书，行万里路',
              description: '书路合一，知行合一',
              image: 'intro_2.png'
           }
        ]
      }
    },
    methods: {
      goPage (val) {
        this.$router.push(val)
      }

    },
    created () {

    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.home-page{
  flex-flow: column;
  width: 100%;
  height: 100vh;
  display: flex;
  left: 0;
  position:absolute;
}

.main-image {
  width: 100%;
  min-height: 700px;
  height: 700px;
  background-color: #f1f4f5;
  background-size: 100% 100%;
  display: flex;
  flex-direction: column;
}

.center-content {
  display: flex;
  max-width: 890px;
  width: 890px;
  height: 270px;
  margin: 0 auto;
  margin-top: 20px;
  min-height: 270px;
}

.center-intro {
  max-width: 1050px;
  height: 100%;
  margin: 0 auto;
  display: flex;
}

.introduction{
  min-height: 310px;
  height: 310px;
  width: 100%;
  background-color: #f1f4f5;
}

.space {
  width: 100%;
  background-color: #f1f4f5;
  min-height: 430px;
  height: 330px;
  display: flex;
  flex-direction: column;

}
#qr-container{
  margin: 0 auto;
  display: flex;
  height: 250px;
  width: 520px;
}
.money-holder{
  width: 260px;
  height: 250px;
  background-size: 100% 100%;
}

.intro-1 {
  width: 33%;
  height: 100%;
}
.intro-2 {
  width: 34%;
  height: 100%;
}
.intro-3 {
  width: 33%;
  height: 100%;
}
.sample-jiapu-container {

  height: 400px;
  width: 400px;
  margin: 0 auto;
  margin-top: 80px;
  background-size: 100% 100%;
}

.demo-button{
  background:url('../assets/demo.png') no-repeat;
  border:none;
  width:155px;
  height:50px;
}

.demo-button:active  /* use Dot here */
{
    font-size:0;
    background:url('../assets/demo.png') no-repeat;
    box-shadow: 0 1px 0 #00823F;
    bottom:-4px;
}

.enter-button{
    background:url('../assets/enter.png') no-repeat;
    border:none;
    width:153px;
    height:50px;
    margin-left: 10px;
}

.enter-button:active  /* use Dot here */
{
    background:url('../assets/enter.png') no-repeat;
    box-shadow: 0 1px 0 #00823F;
    bottom:-4px;
}

button:hover {
  cursor:pointer;
  margin-top: 3px;
}

.button-container {
    background-color: yellow;
    height: 120px;
}

.button-middle-holder {
  display: flex;
  margin: 0 auto;
  margin-top:170px;
  width: 321px;
  height: 200px;

}

.inner {

  height: 90%;
  padding: 0 20px;
  -webkit-box-shadow: 3px 3px 5px 6px #ccc;  /* Safari 3-4, iOS 4.0.2 - 4.2, Android 2.3+ */
  -moz-box-shadow:    3px 3px 5px 6px #ccc;  /* Firefox 3.5 - 3.6 */
  box-shadow:         3px 3px 5px 6px #ccc;
}

h3{
  padding-top: 20px;
  font-size: 28px;
  color: #76bd1d;
  font-family: Montserrat,sans-serif;
  font-weight: 400;
  line-height: 1.2em;
}

h4{
  font-family: Montserrat,sans-serif;
  font-weight: 400;
  line-height: 1.5em;
  text-transform: uppercase;
}

p.feature{
  font-weight: 400;
  font-family: "Varela Round",sans-serif;
}

h1{
  font-size: 2.9rem;
  line-height: 80px;
  font-weight: 400;
  text-align: center;
  letter-spacing: 1px;
}

#title{
  margin: 0 auto;
  margin-top: 100px;
}


.h2-div{
  margin: 0 auto;
  font-size: 2rem;
    line-height: 80px;
    font-weight: 400;
    text-align: center;
    letter-spacing: 1px;
}

.h3-div{
  margin: 0 auto;
  color: #444;
  font-size: 16px;
  line-height: 28px;
  font-weight: 400;
  text-align: center;
  max-width: 600px;
  margin-bottom: 10px;
}

.page-con{
  display: flex;
  flex-direction: column;
  width: 100%;
}
.page-con-img{
  max-width: 480px;
  width: 480px;
  background-color: red;
  margin: 0 auto;
}

.des-item{
  height:380px;
  max-width: 520px;
  width: 520px;
  background-size: 100% 100%;
  display: table;
  margin: 0 auto;
  margin-bottom: 10px;
}





.intro-container{
  width: 700px;
  height: 100%;
  margin: 0 auto;
}

.bottom{
  flex: 1;
  background-color: #F3F0EC;
}

.el-carousel__item h3 {
    color: #475669;
    font-size: 14px;
    opacity: 0.75;
    line-height: 200px;
    margin: 0;
  }

  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }

  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  }


  .home-btn {
  	background-image: linear-gradient(to right, #006175 0%, #00a950 100%);
  	border-radius: 40px;
    box-sizing: border-box;
  	color: #00a84f;
  	display: block;
  	font: 1.525rem 'Oswald', Arial, sans-serif; /*18*/
  	height: 50px;
  	letter-spacing: 1px;
  	margin: 0 auto;
  	padding: 4px;
  	position: relative;
    text-decoration: none;
  	width: 214px;
  	z-index: 2;
  }

  .home-btn:hover {
  	color: #fff;
  }

  .home-btn span {
  	align-items: center;
  	background: #e7e8e9;
  	border-radius: 40px;
  	display: flex;
  	justify-content: center;
  	height: 100%;
  	transition: background .5s ease;
  	width: 100%;
  }

  .home-btn:hover span {
  	background: transparent;
  }
</style>
