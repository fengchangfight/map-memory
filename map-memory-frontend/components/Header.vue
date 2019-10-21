<template>
  <div class="header">
    <b-navbar toggleable="md" type="dark" variant="dark">
      <b-navbar-toggle target="nav_collapse" />
      <b-navbar-brand href="#">
        <img
          src="~/assets/banner_logo.png"
          class="d-inline-block align-top"
          alt="LOGO"
          @click="goPage('/')"
        />
      </b-navbar-brand>
      <b-collapse id="nav_collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item />
          <b-nav-item />
          <b-nav-item @click="goPage('/working')">地图视图</b-nav-item>
          <b-nav-item />
          <b-nav-item />
          <b-nav-item @click="goPage('/memory-list')">列表视图</b-nav-item>
          <b-nav-item />
          <b-nav-item />
          <b-nav-item @click="goPage('/about')">关于我们</b-nav-item>
          <b-nav-item />
          <b-nav-item />
          <b-nav-item @click="goPage('/contact')">联系我们</b-nav-item>
          <b-nav-item />
          <b-nav-item />
          <b-nav-item @click="goPage('/demo')">帮助</b-nav-item>
        </b-navbar-nav>
        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto">
          <b-nav-item-dropdown right>
            <!-- Using button-content slot -->
            <template slot="button-content">
              <em v-if="whoami!=null && whoami.length>0">{{ whoami }}</em>
              <em v-if="whoami==null || whoami.length<1">用户</em>
            </template>

            <b-dropdown-item v-if="whoami.length>0" @click="goPage('/profile')">个人信息</b-dropdown-item>
            <b-dropdown-item v-if="whoami.length>0" @click="logout">登出</b-dropdown-item>

            <b-dropdown-item v-if="whoami==null || whoami.length<1" @click="goPage('/login')">登录</b-dropdown-item>
            <b-dropdown-item v-if="whoami==null || whoami.length<1" @click="goRegister">注册</b-dropdown-item>
          </b-nav-item-dropdown>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>

<script>
import { AXIOS } from "~/common/http-commons";

export default {
  name: "Header",
  components: {},
  data() {
    return {
      profile_image: "",
      activeIndex: "1",
      activeIndex2: "1",
      whoami: "",
      no_auth_page: [
        "/",
        "/public-resource",
        "/demo",
        "/public-zone",
        "/about",
        "/contact",
        "/feedback",
        "/latestnews",
        "/register",
        "/register-success",
        "/regulation",
        "/forget-password",
        "/reset-password-success"
      ]
    };
  },
  mounted() {
    this.whoamifoo();
  },
  created() {},
  updated() {
    this.whoamifoo();
  },
  methods: {
    goRegister() {
      this.$router.push("/register");
    },
    goDemo() {
      this.$router.push("/demo");
    },
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
    goPage(val) {
      this.$router.push(val);
    },

    whoamifoo() {
      AXIOS.get("/api/v1/user/whoami")
        .then(response => {
          // JSON responses are automatically parsed.
          this.whoami = response.data;
        })
        .catch(e => {
          var noLoginRequiredPath = this.no_auth_page;
          if (
            noLoginRequiredPath.indexOf(this.$route.fullPath) != -1 ||
            this.$route.fullPath.startsWith("/mailconfirm") ||
            this.$route.fullPath.startsWith("/public-tree-detail")
          ) {
          } else {
            this.$router.push("/login");
          }
        });
    },
    logout() {
      AXIOS.post(
        "/logout",
        {},
        { headers: { "Content-Type": "application/x-www-form-urlencoded" } }
      )
        .then(response => {
          window.location.href = "/login";
        })
        .catch(e => {
          window.location.href = "/login";
        });
    }
  }
};
</script>

<style scoped>
.header {
  flex: 0 1 auto;
}
</style>
