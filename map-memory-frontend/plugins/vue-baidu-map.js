import Vue from 'vue'
import {
  BmlMarkerClusterer,
  BaiduMap,BmView, BmLocalSearch, BmNavigation,BmGeolocation,BmCityList,BmControl,BmMapType,
BmContextMenu, BmContextMenuItem, BmOverlay, BmMarker, BmLabel, BmInfoWindow,
BmPointCollection} from 'vue-baidu-map'

Vue.component('bml-marker-clusterer', BmlMarkerClusterer)
Vue.component('bm-view', BmView)
Vue.component('bm-local-search', BmLocalSearch)
Vue.component('baidu-map', BaiduMap)
Vue.component('bm-navigation', BmNavigation)
Vue.component('bm-geolocation', BmGeolocation)
Vue.component('bm-city-list', BmCityList)
Vue.component('bm-control', BmControl)
Vue.component('bm-map-type', BmMapType)
Vue.component('bm-context-menu', BmContextMenu)
Vue.component('bm-context-menu-item', BmContextMenuItem)
Vue.component('bm-overlay', BmOverlay)
Vue.component('bm-marker', BmMarker)
Vue.component('bm-label', BmLabel)
Vue.component('bm-info-window', BmInfoWindow)
Vue.component('bm-point-collection', BmPointCollection)
