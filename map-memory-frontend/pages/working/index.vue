<template>
  <div style="width:100%;position:relative;height:100%;">
    <el-dialog :visible.sync="inputReadCodeVisible" :fullscreen="false" title="输入阅读密码" width="30%">
      <div @keyup.enter="submitReadCode">
        <el-form autocomplete="off">
          <el-input
            v-model="read_code"
            placeholder="阅读密码(初始为123)"
            autocomplete="off"
            type="password"
          />
          <input type="password" style="visibility:hidden" />
          <el-button @click="submitReadCode">提交</el-button>
        </el-form>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="deleteFavLocVisible" title="删除常用位置" width="200px">
      <div>
        <p>{{ selected_favname }}</p>
        <el-button @click="deleteFavloc">删除收藏位置</el-button>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="addFavlocVisible" title="添加常用位置" width="30%">
      <div>
        <el-input v-model="fav_loc_name" placeholder="收藏名" />
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="addFavlocVisible = false">取消</el-button>
        <el-button type="primary" @click="add2Faviloc">添加</el-button>
      </span>
    </el-dialog>

    <el-dialog
      :visible.sync="addMemoryVisible"
      :fullscreen="true"
      :before-close="handleClose"
      title="添加记忆点"
      width="100%"
    >
      <el-form ref="form" :model="form" label-width="80px">
        <el-form-item label="图标">
          <el-select v-model="selected_icon" filterable placeholder="记忆图标">
            <el-option
              v-for="item in available_icons"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            >
              <span style="float: left">
                <img :src="'imgs/'+item.id" width="20" height="20" />
              </span>
              <span style="float: right; color: #8492a6; font-size: 13px">{{ item.name }}</span>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="记忆标题">
          <el-input v-model="form.title" maxlength="300" />
        </el-form-item>
        <el-form-item label="记忆详情">
          <vue-editor :editorToolbar="customToolbar" v-model="form.memory_content"></vue-editor>
        </el-form-item>
        <el-form-item>
          <el-switch v-model="form.is_public" active-text="公开" inactive-text="仅本人可见" />
        </el-form-item>
      </el-form>

      <span slot="footer" class="dialog-footer">
        <el-button @click="addMemoryVisible = false">取消</el-button>
        <el-button type="primary" @click="createMemoryPoint">创建</el-button>
      </span>
    </el-dialog>

    <el-dialog :visible.sync="memoryDetailBoxVisible" :fullscreen="true" title="记忆详情" width="100%">
      <div style="display:flex;flex-direction:column">
        <div style="display:flex;position:relative;">
          <button
            v-if="detailMode=='view' && memDetail.i_am_owner && readonlydetail==false"
            class="delete-mem-point"
            title="删除"
            @click="deleteMemPoint(memDetail.id)"
          />
          <button
            v-clipboard:copy="mem_url"
            v-clipboard:success="onCopy"
            v-clipboard:error="onCopyError"
            v-if="memDetail.is_public"
            id="copy-lng-lat"
            title="分享本记忆"
          />
          <button
            v-if="detailMode=='view' && memDetail.i_am_owner && readonlydetail==false"
            class="edit-mem-point"
            title="编辑"
            @click="editMemPoint(memDetail.id)"
          />
          <h3 v-if="detailMode=='view'" style="margin: 0 auto;padding:0px 30px 0px 40px;">
            <span style="float: left">
              <img :src="'imgs/'+memDetail.icon" width="40" height="40" />
            </span>
            &nbsp;&nbsp;{{ memDetail.title }}
            <span
              v-if="!memDetail.is_public"
              class="small-notes"
            >(仅本人可见)</span>
            <span v-if="memDetail.is_public" class="small-notes">(公开)</span>
          </h3>
          <div v-if="detailMode=='edit'" style="display:flex;width:100%;">
            <el-select v-model="memDetail.icon" filterable placeholder="记忆图标">
              <el-option
                v-for="item in available_icons"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              >
                <span style="float: left">
                  <img :src="'imgs/'+item.id" width="20" height="20" />
                </span>
                <span style="float: right; color: #8492a6; font-size: 13px">{{ item.name }}</span>
              </el-option>
            </el-select>
            <el-input v-model="memDetail.title" style="margin: 0 auto;padding:0px 1px 0px 1px;" />
          </div>
        </div>
        <div style="display:block;margin-top: 10px;">
          <div v-if="readonlydetail==false" style="margin-left: auto; font-size:70%;color:gray">
            <label style="color:#003300;">{{ memDetail.nickname }}</label>
            &nbsp;发布于：{{ memDetail.created_at }} &nbsp;
            <label
              v-if="autosaveflag"
            >(已自动保存)</label>
          </div>
        </div>
        <div v-if="detailMode=='edit'">
          <vue-editor
            @text-change="autosave"
            :editorToolbar="customToolbar"
            v-model="memDetail.content"
          ></vue-editor>
          <el-switch
            v-model="memDetail.is_public"
            active-text="公开"
            inactive-text="仅本人可见"
            @change="changeAccessibility"
          />

          <!-- <froala :config="option" v-model="memDetail.content"></froala> -->
          <div class="button-container">
            <el-button
              style="float: right; margin-right:10px;"
              type="primary"
              size="small"
              @click="saveEdit"
            >保存</el-button>
            <el-button style="float: right; margin-right:10px;" size="small" @click="cancelEdit">取消</el-button>
          </div>
        </div>
        <div class="ql-editor ql-snow">
          <div
            class="ql-editor ql-snow"
            v-loading="loading_content"
            v-if="detailMode=='view'"
            v-html="memDetail.content"
          />
        </div>
      </div>
    </el-dialog>

    <div class="search-box">
      <div style="display: flex;margin: 0 auto;">
        <div class="search-input-container-row-keyword">
          <el-input v-model="keyword" placeholder="地名关键词">
            <template slot="append">
              <el-button @click="clear()">清空搜索</el-button>
            </template>
          </el-input>
        </div>

        <div class="search-input-container-row">
          <el-select v-model="selected_center" filterable placeholder="常用位置" @change="changeCenter">
            <el-option
              v-for="item in fav_loc_list"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </div>
        <div class="search-input-container-row-ll">
          <el-input v-model="locstring" placeholder="(经,纬),例:116.328427,39.990929">
            <template slot="append">
              <el-button @click="gotoloc">定位</el-button>
            </template>
          </el-input>
        </div>
        <div class="search-input-container-row-switch">
          <el-switch
            v-model="view_public"
            active-text="显示公共笔记"
            inactive-text="仅自己的"
            @change="changeView"
          />
        </div>
      </div>
    </div>
    <div class="bm-view">
      <baidu-map
        :max-zoom="20"
        :scroll-wheel-zoom="true"
        :center="center"
        :zoom="zoom"
        class="bm-view"
        ak="badBwvN3hjlfXUCFSpcGzVLuG0oqpTNR"
        @moveend="syncCenterAndZoom"
        @zoomend="syncCenterAndZoom"
        @ready="mapready"
      >
        <bm-navigation anchor="BMAP_ANCHOR_TOP_RIGHT" />
        <bm-geolocation
          :show-address-bar="true"
          :auto-location="true"
          anchor="BMAP_ANCHOR_BOTTOM_LEFT"
        />
        <bm-local-search
          :keyword="keyword"
          :auto-viewport="true"
          :location="location"
          style="position:absolute;left:10px;top:40px;"
        />
        <bm-city-list style="z-index:10;" anchor="BMAP_ANCHOR_TOP_LEFT" />
        <bm-map-type
          :map-types="['BMAP_NORMAL_MAP', 'BMAP_PERSPECTIVE_MAP', 'BMAP_SATELLITE_MAP', 'BMAP_HYBRID_MAP']"
          anchor="BMAP_ANCHOR_BOTTOM_RIGHT"
        />
        <bm-context-menu>
          <bm-context-menu-item :callback="addMemory" text="在此添加记忆" />
          <bm-context-menu-item :callback="popAdd2FavoriteLocation" text="设为常用位置" />
          <bm-context-menu-item :callback="logAndLat" text="拷贝此经纬度" />
        </bm-context-menu>

        <bml-marker-clusterer :average-center="true" :maxZoom="16">
          <bm-marker
            v-for="item in my_mem_data"
            :key="item.id"
            :title="limitStringLength(item.title)"
            :position="{lng: item.longitude, lat: item.latitude}"
            :dragging="true"
            :icon="{url: 'imgs/'+item.icon, size: {width: 40, height: 40}}"
            animation="BMAP_ANIMATION_DROP"
            @mouseover="infoWindowOpen(item.id)"
            @mouseout="infoWindowClose(item.id)"
            @dragend="moveMemoryPos($event, item.id)"
            @click="showMemDetailWin(item.id, item.locked)"
          >
            <bm-info-window :show="showInfoWin[item.id]">{{ item.title }}</bm-info-window>
          </bm-marker>
        </bml-marker-clusterer>

        <bm-marker
          v-if="is_newbie"
          :position="{lng: center.lng, lat: center.lat}"
          :dragging="false"
          :icon="{url: 'imgs/'+tutorialItem.icon, size: {width: 60, height: 60}}"
          animation="BMAP_ANIMATION_BOUNCE"
          @mouseover="infoWindowOpen(tutorialItem.id)"
          @mouseout="infoWindowClose(tutorialItem.id)"
          @click="showMemDetailWin(tutorialItem.id)"
        >
          <bm-info-window :show="showInfoWin[tutorialItem.id]">{{ tutorialItem.title }}</bm-info-window>
        </bm-marker>

        <bm-marker
          v-for="item in fav_loc_list"
          :key="item.id"
          :title="limitStringLength(item.name)"
          :position="{lng: item.longitude, lat: item.latitude}"
          :dragging="false"
          :icon="{url: 'imgs/redstar.png', size: {width: 40, height: 40}}"
          @click="showFavlocDetail(item.id, item.name)"
        />
      </baidu-map>
    </div>
  </div>
</template>

<script>
import swal from "sweetalert";
import { AXIOS } from "~/common/http-commons";
import "~/css/map.public.css";
import state from "~/common/state";
import base from "~/mixins/base";
import _ from "lodash";
import Qs from "qs";

export default {
  components: {},
  mixins: [base],
  data() {
    return {
      map_ready: false,
      url_longitude: "",
      url_latitude: "",
      url_showpublic: "",
      loading_content: false,
      locstring: "",
      inputReadCodeVisible: false,
      current_id: "",
      read_code: "",
      tutorialItem: {
        id: "tutorial",
        title: "点我看说明哦(づ￣ 3￣)づ",
        icon: "tutorial.png"
      },
      autosaveflag: false,
      state,
      isLoggedIn: false,
      editorOption: {
        // some quill options
        modules: {
          toolbar: [
            ["bold", "italic", "underline", "strike"],
            ["blockquote", "code-block"]
          ]
        }
      },
      deleteFavLocVisible: false,
      selected_center: "",
      fav_loc_name: "",
      addFavlocVisible: false,
      customToolbar: [
        [{ font: [] }],
        [{ header: [false, 1, 2, 3, 4, 5, 6] }],
        [{ size: ["small", false, "large", "huge"] }],
        ["bold", "italic", "underline", "strike"],
        [
          { align: "" },
          { align: "center" },
          { align: "right" },
          { align: "justify" }
        ],
        [{ header: 1 }, { header: 2 }],
        ["blockquote", "code-block"],
        [{ list: "ordered" }, { list: "bullet" }, { list: "check" }],
        [{ script: "sub" }, { script: "super" }],
        [{ indent: "-1" }, { indent: "+1" }],
        [{ color: [] }, { background: [] }],
        ["link", "formula"],
        [{ direction: "rtl" }],
        ["clean"]
      ],
      memoryDetailBoxVisible: false,
      my_mem_data: [],

      form: {
        title: "",
        memory_content: "",
        is_public: false
      },
      load_scope: {
        south_west_x: 0.0,
        south_west_y: 0.0,
        north_east_x: 0.0,
        north_east_y: 0.0
      },
      current_bound: {
        south_west_x: 0.0,
        south_west_y: 0.0,
        north_east_x: 0.0,
        north_east_y: 0.0
      },
      is_newbie: false,
      selected_longitude: 0.0,
      selected_latitude: 0.0,
      addMemoryVisible: false,
      center: { lng: 0, lat: 0 },
      zoom: 19,
      keyword: "",
      location: "",
      selected_icon: "",
      showInfoWin: {},
      memDetail: {},
      detailMode: "view",
      readonlydetail: false,
      selected_id: "",
      fav_loc_list: [],
      selected_favname: "",
      base_service_url: "",
      physicLongitude: "",
      physicLatitude: "",
      view_public: false
    };
  },
  computed: {
    mem_url() {
      //return this.base_url+'/working'+"?longitude="+this.memDetail.longitude+"&latitude="+this.memDetail.latitude+"&showpublic=true";
      return (
        this.base_url +
        "/working" +
        "?memid=" +
        this.memDetail.id +
        "&showpublic=true"
      );
    },
    fallbackLongitude() {
      return this.physicLongitude != "" ? this.physicLongitude : 116.404;
    },
    fallbackLatitude() {
      return this.physicLatitude != "" ? this.physicLatitude : 39.915;
    },
    optionNoAutoSave() {
      return {
        imageUpload: false,
        heightMin: 350
      };
    },
    option() {
      return {
        imageUpload: false,
        heightMin: 350,
        saveInterval: 2500,
        saveMethod: "POST",
        saveParam: "memDetail.content",
        saveParams: {
          id: this.memDetail.id
        },
        requestWithCredentials: true,
        requestWithCORS: true,
        saveURL: this.base_service_url + "/api/v1/memory-content"
      };
    }
  },
  mounted() {
    this.checklogin();

    this.url_longitude = this.$route.query.longitude;
    if (
      this.url_longitude == null ||
      this.url_longitude.length < 1 ||
      isNaN(this.url_longitude)
    ) {
      this.url_longitude = null;
    }

    this.url_latitude = this.$route.query.latitude;
    if (
      this.url_latitude == null ||
      this.url_latitude.length < 1 ||
      isNaN(this.url_latitude)
    ) {
      this.url_latitude = null;
    }
    this.url_showpublic = this.$route.query.showpublic;
    if (this.url_showpublic != "true") {
      this.url_showpublic = null;
    } else {
      this.url_showpublic = true;
      this.view_public = this.url_showpublic;
    }

    var mem_id = this.$route.query.memid;
    if (mem_id != null && mem_id.length > 0) {
      this.getLocById(mem_id);
    }

    this.getBaseUrl();
    this.getBaseServiceUrl();
    this.isFirstDayUser();
    this.loadAllFavloc();
  },
  methods: {
    getLocById(memid) {
      console.log("reach here 10");
      AXIOS.get("/api/v1/mem-location?id=" + memid)
        .then(response => {
          if (response.data.ok == true) {
            this.url_latitude = response.data.data.latitude;
            this.url_longitude = response.data.data.longitude;
            console.log("reach here 11");
            if (this.map_ready) {
              this.center.lng = Number(this.url_longitude);
              this.center.lat = Number(this.url_latitude);
            }
            console.log("reach here 12");
            console.log(this.center);
            this.zoom = 19;
          } else {
            if (this.isLoggedIn) {
              this.$notify.error({
                title: "错误",
                message: response.data.message
              });
            }
          }
        })
        .catch(e => {
          console.log("根据id获取坐标失败");
          this.goPage("/login");
        });
    },
    onCopyError() {
      swal("提示", "拷贝失败", "info");
    },
    onCopy() {
      //var url = this.base_service_url+'/'+"?longitude="+this.memDetail.longitude+"&latitude="+this.memDetail.latitude+"&showpublic=true"
      //var st=this.memDetail.longitude+","+this.memDetail.latitude
      swal("提示", "已复制本记忆地址:" + this.mem_url + "到剪贴板", "info");
    },
    gotoloc() {
      if (this.locstring == null || this.locstring.length < 1) {
        swal("提示", "(经度,纬度)组合不能为空哦(づ￣ 3￣)づ", "info");
        return;
      }
      var arr = this.locstring.split(",", 2);
      var arrchinese = this.locstring.split("，", 2);

      if (arr.length < 2 && arrchinese.length == 2) {
        arr = arrchinese;
      }

      if (arr.length < 2) {
        swal("提示", '输入需要按照:"经度,纬度"的格式哦(づ￣ 3￣)づ', "info");
        return;
      }
      var latitude = arr[1];
      var longitude = arr[0];
      if (isNaN(latitude) || isNaN(longitude)) {
        swal("提示", "输入的经纬度都必须是数字哦(づ￣ 3￣)づ", "info");
        return;
      }

      if (!(longitude < 136 && longitude > 72)) {
        swal(
          "提示",
          "经度范围不在中华人民共和国，不予查询哦(づ￣ 3￣)づ",
          "info"
        );
        return;
      }
      if (!(latitude <= 54 && latitude >= 3)) {
        swal(
          "提示",
          "纬度范围不在中华人民共和国，不予查询哦(づ￣ 3￣)づ",
          "info"
        );
        return;
      }

      this.center.lng = longitude;
      this.center.lat = latitude;
      this.zoom = 19;
    },
    changeAccessibility(val) {
      if (val == true) {
        this.memDetail.is_public == true;
      } else {
        this.memDetail.is_public == false;
      }
    },
    submitReadCode() {
      if (this.read_code == null || this.read_code.length < 1) {
        swal("提示", "不能提交空的阅读密码哦(づ￣ 3￣)づ", "info");
        return;
      }
      this.inputReadCodeVisible = false;
      this.memDetail = {};
      this.memoryDetailBoxVisible = true;
      this.loadMemoryDetailById(this.current_id, this.read_code);
    },
    isFirstDayUser() {
      AXIOS.get("/api/v1/first-day-user")
        .then(response => {
          if (response.data == true) {
            // first day registered user, give instructions
            this.is_newbie = true;
            console.log(true);
          } else {
            // old user, do nothing
            this.is_newbie = false;
            console.log(false);
          }
        })
        .catch(e => {
          console.log(e);
        });
    },
    autosave: _.debounce(function() {
      var data = { id: this.memDetail.id, content: this.memDetail.content };
      AXIOS.put("/api/v1/memory-content", Qs.stringify(data), {
        headers: { "Content-Type": "application/x-www-form-urlencoded" }
      })
        .then(response => {
          if (response.data.ok == true) {
            this.autosaveflag = true;
            this.$notify({
              title: "自动保存",
              type: "success",
              message: "已自动保存"
            });
            var count = 2;
            var storeThis = this;
            var refreshIntervalId = setInterval(function() {
              count = count - 1;
              if (count == 0) {
                clearInterval(refreshIntervalId);
                storeThis.autosaveflag = false;
              }
            }, 1000);
          }
        })
        .catch(e => {});
    }, 2500),
    onEditMemoryEditorChange({ editor, html, text }) {
      this.memDetail.content = html;
    },
    onAddMemoryEditorChange({ editor, html, text }) {
      //console.log('editor change!', editor, html, text)
      this.form.memory_content = html;
    },
    showFavlocDetail(id, name) {
      this.selected_id = id;
      this.selected_favname = name;
      this.deleteFavLocVisible = true;
    },
    deleteFavloc() {
      AXIOS.delete("/api/v1/favorite-location/" + this.selected_id)
        .then(response => {
          if (response.data.ok == true) {
            this.$notify({
              title: "成功",
              type: "success",
              message: response.data.message
            });
            this.loadAllFavloc();
            this.deleteFavLocVisible = false;
          } else {
            this.$notify.error({
              title: "错误",
              message: response.data.message
            });
          }
        })
        .catch(e => {
          this.$notify.error({
            title: "错误",
            message: "未知错误"
          });
        });
    },

    updateFavLocUse(id) {
      AXIOS.put("/api/v1/favorite-location/used/" + id)
        .then(response => {
          if (response.data.ok == true) {
            console.log(response.data.message);
            this.loadAllFavloc();
          }
        })
        .catch(e => {});
    },
    changeCenter() {
      for (var i in this.fav_loc_list) {
        var loc = this.fav_loc_list[i];
        if (loc.id == this.selected_center) {
          // do stuff and return
          this.center.lng = loc.longitude;
          this.center.lat = loc.latitude;
          this.zoom = 19;
          this.updateFavLocUse(loc.id);
          return;
        }
      }
    },
    moveMemoryPos(e, id) {
      //when drag end, decide weather to update its longitude and latitude for id

      var longitude = e.target.point.lng;
      var latitude = e.target.point.lat;

      AXIOS.put("/api/v1/memory-pos/" + id, {
        longitude: longitude,
        latitude: latitude
      })
        .then(response => {
          if (response.data.ok == true) {
            this.$notify({
              title: "成功",
              type: "success",
              message: response.data.message
            });
            this.my_mem_data = [];
            this.loadMemPoints();
          } else {
            this.$notify.error({
              title: "错误",
              message: response.data.message
            });
            this.my_mem_data = [];
            this.loadMemPoints();
          }
        })
        .catch(e => {
          this.$notify.error({
            title: "错误",
            message: e.response.statusText
          });
          this.my_mem_data = [];
          this.loadMemPoints();
        });
    },
    saveEdit() {
      // update with: id, title, content, that's it
      AXIOS.put("/api/v1/memory/" + this.memDetail.id, {
        title: this.memDetail.title,
        content: this.memDetail.content,
        icon: this.memDetail.icon,
        is_public: this.memDetail.is_public
      })
        .then(response => {
          if (response.data.ok == true) {
            this.$notify({
              title: "成功",
              type: "success",
              message: response.data.message
            });
            this.cancelEdit();
            this.loadMemPoints();
          } else {
            this.$notify.error({
              title: "错误",
              message: response.data.message
            });
          }
        })
        .catch(e => {
          this.$notify.error({
            title: "错误",
            message: "未知错误"
          });
        });
    },
    cancelEdit() {
      this.detailMode = "view";
      //this.loadMemoryDetailById(this.memDetail.id);
    },
    loadMemoryDetailById(id, read_code) {
      this.loading_content = true;
      var params = {};
      if (read_code != null && read_code.length > 0) {
        params.read_code = read_code;
      }
      AXIOS.get("/api/v1/memory/" + id, {
        params: params
      })
        .then(response => {
          if (response.data.ok == true) {
            this.memDetail = response.data.data;
            this.loading_content = false;
          } else {
            this.$notify.error({
              title: "错误",
              message: response.data.message
            });
            this.memoryDetailBoxVisible = false;
          }
        })
        .catch(e => {
          this.$notify.error({
            title: "错误",
            message: "未知错误"
          });
        });
    },
    showMemDetailWin(id, locked) {
      this.detailMode = "view";
      if (locked == true) {
        this.read_code = "";
        this.inputReadCodeVisible = true;
        this.current_id = id;
      } else {
        this.memDetail = {};
        this.memoryDetailBoxVisible = true;
        if (id == "tutorial") {
          this.readonlydetail = true;
          this.memDetail.title = "操作说明";
          this.memDetail.content = `<p>
    在地图上任意位置右键弹出菜单
</p>
<p>
</p>
<img src="http://www.fengchang.cc/imageprocessing/fit?width=700&height=600&type=jpeg&file=NYVv2NiktGC4OR1e.png">
<p>
    菜单选项
</p>
<ol class=" list-paddingleft-2" style="list-style-type: decimal;">
    <li>
        <p>
            可查看该位置<strong>经纬度</strong><br/>
        </p>
    </li>
    <li>
        <p>
            可在该位置<strong>创建笔记本</strong>
        </p>
    </li>
    <li>
        <p>
            可收藏该位置为<strong>常用位置</strong>
        </p>
    </li>
</ol>
<p>
    <strong>拖拽图标</strong>可以调整笔记本位置
</p>
<p>
    按<strong>Escape键</strong>可<strong>退出</strong>记忆详情页面
</p>`;
          this.memDetail.icon = "tutorial.png";
        } else {
          this.readonlydetail = false;
          this.detailMode = "view";
          this.loadMemoryDetailById(id);
        }
      }
    },
    editMemPoint(id) {
      // make title and content edit mode
      this.detailMode = "edit";
    },
    deleteMemPoint(id) {
      this.$confirm("确认删除该记忆点？")
        .then(_ => {
          AXIOS.delete("/api/v1/memory/" + id)
            .then(response => {
              if (response.data.ok == true) {
                this.$notify({
                  title: "成功",
                  type: "success",
                  message: response.data.message
                });
                this.memoryDetailBoxVisible = false;
                this.loadMemPoints();
              } else {
                this.$notify.error({
                  title: "错误",
                  message: response.data.message
                });
              }
            })
            .catch(e => {
              this.$notify.error({
                title: "错误",
                message: "未知错误"
              });
            });
          done();
        })
        .catch(_ => {});
    },
    limitStringLength(st) {
      if (st == null) {
        return "";
      }
      var maxlength = 20;
      if (st.length < maxlength) {
        return st;
      } else {
        return st.substring(0, maxlength) + "...";
      }
    },
    infoWindowClose(id) {
      this.showInfoWin = {};
      this.showInfoWin[id] = false;
    },
    infoWindowOpen(id) {
      this.showInfoWin = {};
      this.showInfoWin[id] = true;
    },
    syncCenterAndZoom(e) {
      this.zoom = e.target.getZoom();
      const { lng, lat } = e.target.getCenter();
      if (
        lng == 0 &&
        lat == 0 &&
        e.target.getBounds().getNorthEast().lng -
          e.target.getBounds().getSouthWest().lng >
          360
      ) {
        // skip the whole world view at first load second, it's too big, it's gonna load every memory data in the whole world
        return;
      }

      this.current_bound.south_west_x = e.target.getBounds().getSouthWest().lng;
      this.current_bound.south_west_y = e.target.getBounds().getSouthWest().lat;
      this.current_bound.north_east_x = e.target.getBounds().getNorthEast().lng;
      this.current_bound.north_east_y = e.target.getBounds().getNorthEast().lat;

      function boundInScope(cb, ls) {
        if (
          cb.south_west_x > ls.south_west_x &&
          cb.south_west_y > ls.south_west_y &&
          cb.north_east_x < ls.north_east_x &&
          cb.north_east_y < ls.north_east_y
        ) {
          return true;
        } else {
          return false;
        }
      }
      //bound not fully in scope,  update scope and load new data
      if (!boundInScope(this.current_bound, this.load_scope)) {
        // update new loadscope
        this.load_scope.south_west_x =
          this.current_bound.south_west_x -
          (this.current_bound.north_east_x - this.current_bound.south_west_x) /
            2;
        this.load_scope.south_west_y =
          this.current_bound.south_west_y -
          (this.current_bound.north_east_y - this.current_bound.south_west_y) /
            2;
        this.load_scope.north_east_x =
          this.current_bound.north_east_x +
          (this.current_bound.north_east_x - this.current_bound.south_west_x) /
            2;
        this.load_scope.north_east_y =
          this.current_bound.north_east_y +
          (this.current_bound.north_east_y - this.current_bound.south_west_y) /
            2;
        // load data with in updated loadscope, do the real important stuff
        this.loadMemPoints();
      }

      // check if current view is totally within load scope, if not, reload within new scope with new center, and update load scope itself
    },
    loadAllFavloc() {
      // this is loading all fav loc list for the drop down
      AXIOS.get("/api/v1/favorite-location")
        .then(response => {
          if (response.data.ok == true) {
            this.fav_loc_list = response.data.data;
          } else {
            this.$notify.error({
              title: "错误",
              message: response.data.message
            });
          }
        })
        .catch(e => {
          console.log("无法载入收藏地点列表");
        });
    },
    loadMemPoints() {
      AXIOS.post("/api/v1/person-memory-inbound", {
        south_west_x: this.load_scope.south_west_x,
        south_west_y: this.load_scope.south_west_y,
        north_east_x: this.load_scope.north_east_x,
        north_east_y: this.load_scope.north_east_y,
        view_public: this.view_public ? "1" : "0"
      })
        .then(response => {
          if (response.data.ok == true) {
            this.my_mem_data = response.data.data;
          } else {
            this.$notify.error({
              title: "错误",
              message: response.data.message
            });
          }
        })
        .catch(e => {
          console.log("载入地图记忆点点失败");
          this.goPage("/login");
        });
    },
    createMemoryPoint() {
      if (!Boolean(this.selected_icon)) {
        swal("提示", "您必须选中一个记忆图标哦(づ￣ 3￣)づ", "info");
        return;
      }

      if (!Boolean(this.form.title)) {
        swal("提示", "标题不能为空哦(づ￣ 3￣)づ", "info");
        return;
      }
      // create memory with
      AXIOS.post("/api/v1/memorypoint", {
        longitude: this.selected_longitude,
        latitude: this.selected_latitude,
        title: this.form.title,
        content: this.form.memory_content,
        icon: this.selected_icon,
        is_public: this.form.is_public
      })
        .then(response => {
          if (response.data.ok == true) {
            this.$notify({
              title: "成功",
              type: "success",
              message: response.data.message
            });
            this.addMemoryVisible = false;
            this.loadMemPoints();
          } else {
            this.$notify.error({
              title: "错误",
              message: response.data.message
            });
          }
        })
        .catch(e => {
          this.$notify.error({
            title: "错误",
            message: "未知错误"
          });
        });
    },
    handleClose(done) {
      this.$confirm("确认关闭？")
        .then(_ => {
          done();
        })
        .catch(_ => {});
    },
    logAndLat(e) {
      console.log(e);
      console.log(e.map.cs);
      let container = this.$refs.container;
      var val = e.point.lng + "," + e.point.lat;
      this.$copyText(val, container);
      swal("信息", "已复制经纬度:" + val + "到剪贴板", "info");
    },
    add2Faviloc() {
      if (!Boolean(this.fav_loc_name)) {
        swal("提示", "名字不能为空哦(づ￣ 3￣)づ", "info");
        return;
      }
      AXIOS.post("/api/v1/favorite-location", {
        name: this.fav_loc_name,
        longitude: this.selected_longitude,
        latitude: this.selected_latitude
      })
        .then(response => {
          if (response.data.ok == true) {
            this.$notify({
              title: "成功",
              type: "success",
              message: response.data.message
            });
            this.addFavlocVisible = false;
            this.loadAllFavloc();
          } else {
            this.$notify.error({
              title: "错误",
              message: response.data.message
            });
          }
        })
        .catch(e => {
          this.$notify.error({
            title: "错误",
            message: "未知错误"
          });
        });
    },
    popAdd2FavoriteLocation(e) {
      this.selected_longitude = e.point.lng;
      this.selected_latitude = e.point.lat;
      this.addFavlocVisible = true;
      this.fav_loc_name = "";
    },
    addMemory(e) {
      this.selected_longitude = e.point.lng;
      this.selected_latitude = e.point.lat;
      // step1 pop up a window to enter message about title and content, at the same time record longitude and latitude

      // clear previous value
      this.selected_icon = "";
      this.form["title"] = "";
      this.form["memory_content"] = "";
      this.form["is_public"] = false;

      this.addMemoryVisible = true;

      // step 2 handle the add memory point request(post)
    },
    clear() {
      this.keyword = "";
      this.location = "";
    },
    checkCurrentLoc() {
      var storeThis = this;
      $.ajax({
        url: "https://api.map.baidu.com/location/ip",
        data: {
          coor: "bd09ll",
          ak: "Er8iGG4UMfSd3Ckuc6w8C56peI4ge1Ih"
        },
        type: "get",
        dataType: "jsonp",
        async: true,
        success: function(data) {
          console.log(data.content);
          var locInfo = data.content;
          if (locInfo != null && locInfo.point != null) {
            storeThis.physicLongitude = locInfo.point.x;
            storeThis.physicLatitude = locInfo.point.y;
          }

          console.log(storeThis.url_longitude);
          console.log(storeThis.url_latitude);
          if (
            storeThis.url_latitude != null &&
            storeThis.url_latitude.toString().length > 0 &&
            storeThis.url_longitude != null &&
            storeThis.url_longitude.toString().length > 0 &&
            !isNaN(storeThis.url_latitude) &&
            !isNaN(storeThis.url_longitude)
          ) {
            storeThis.center.lng = storeThis.url_longitude;
            storeThis.center.lat = storeThis.url_latitude;
          } else {
            storeThis.center.lng =
              storeThis.state.longitude == null
                ? storeThis.fallbackLongitude
                : storeThis.state.longitude;
            storeThis.center.lat =
              storeThis.state.latitude == null
                ? storeThis.fallbackLatitude
                : storeThis.state.latitude;
          }
          //storeThis.loadMemPoints();
        },
        error: function(data) {
          console.log("error getting my location");
        }
      });
    },
    changeView() {
      this.loadMemPoints();
    },
    mapready({ BMap, map }) {
      //console.log(BMap, map)
      this.map_ready = true;
      this.checkCurrentLoc();

      //this.center.lng = this.state.longitude==null?this.fallbackLongitude:this.state.longitude;
      //this.center.lat = this.state.latitude==null?this.fallbackLatitude:this.state.latitude;
      this.zoom = 19;

      //this.loadMemPoints();
    }
  }
};
</script>

<style scoped>
.container {
  width: 100%;
  height: 100%;
  background-color: yellow;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.mapborder {
  background-color: gray;
  width: 100%;
  height: 100%;
  flex: 1;
}
.bm-view {
  width: 100%;
  height: 100%;
}
.search-box {
  z-index: 1;
  width: 600px;
  height: 38px;
  position: absolute;
  top: 0;

  margin-left: 100px;
  margin-top: 5px;
  display: flex;
}
.search-result-holder {
  width: 400px;
  height: 400px;
  background-color: gray;
  position: absolute;
  left: 15px;
  top: 50px;
  z-index: 0;
}
.search-input-container-row-keyword {
  width: 250px;
  margin-left: 5px;
}
.search-input-container-row {
  width: 150px;
  margin-left: 5px;
}
.search-input-container-row-ll {
  width: 320px;
  margin-left: 25px;
}
.search-input-container-row-switch {
  margin-top: 5px;
  margin-left: 100px;
  width: 200px;
}

.quill-editor {
  min-height: 200px;
  max-height: 400px;
  overflow-y: auto;
}
</style>
