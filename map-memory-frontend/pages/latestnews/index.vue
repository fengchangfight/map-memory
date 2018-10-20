<template>
  <div class="latestnews-page">
    <div class="latestnews-content">
      <div class="center-content">
        <div class="place-taker"/>
        <timeline>
          <timeline-title>2018</timeline-title>
          <timeline-item color="#9dd8e0">8月2日ditujiyi.com域名上线</timeline-item>
        </timeline>

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
import { Timeline, TimelineItem, TimelineTitle } from 'vue-cute-timeline'
export default {
    name: 'LatestNews',
    components: {
      'app-footer': Footer,
      Timeline,
      TimelineItem,
      TimelineTitle
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
.latestnews-page{
  flex-flow: column;
  width: 100%;
  display: flex;
  left: 0;
  position:absolute;
}

.latestnews-content {
  width: 100%;
  min-height: 700px;
  background-color: #f1f4f5;
}

.center-content {
  max-width: 800px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
}

.place-taker{
  width: 100%;
  height: 50px;

}

.bottom{
  flex: 1;
  background-color: #F3F0EC;
}
</style>
