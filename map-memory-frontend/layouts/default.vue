<template>
  <div class="whole-framework">
    <div id="whole-header">
      <app-header/>
    </div>
    <div id="whole-body">
      <nuxt/>
    </div>
  </div>

</template>

<script>
import Header from '~/components/Header.vue'
export default {
  components: {
    'app-header': Header
  }
}
</script>

<style>


.whole-framework{
  width: 100%;
  height: 100vh;
  display: flex;
  flex-direction: column;
}

#whole-header{
  min-height: 55px;
  height: 55px;
}
#whole-body{
  width: 100%;
  flex: 1;
}

</style>
