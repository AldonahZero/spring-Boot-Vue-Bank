<template>
  <div>
    <el-container>
      <el-header class="homeHeader">
        <div class="title">我,秦始皇.[打钱]</div>
        <div>
          <el-button
            icon="el-icon-bell"
            type="text"
            style="margin-right: 8px;color: #000000;"
            size="normal"
            @click="goChat"
          ></el-button>
          <el-dropdown class="userInfo" @command="commandHandler">
            <span class="el-dropdown-link">
              {{user.userName}}
              <i>
                <img :src="userface" alt />
              </i>
            </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item command="userinfo">个人中心</el-dropdown-item>
              <el-dropdown-item command="setting">设置</el-dropdown-item>
              <el-dropdown-item command="logout" divided>注销登录</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </div>
      </el-header>
      <el-container>
        <el-aside width="230px" height="100%">
          <el-menu
            width="auto"
            default-active="1-4-1"
            class="el-menu-vertical-demo"
            @mouseenter.native="collapseOpen"
            @mouseleave.native="collapseClose"
            :collapse="isCollapse"
          >
            <el-submenu index="1">
              <template slot="title">
                <i class="el-icon-location"></i>
                <span slot="title">存取款业务</span>
              </template>
              <el-menu-item-group>
                <span slot="title">存取款</span>
                <el-menu-item index="1-1"><AddMoney/></el-menu-item>
                <el-menu-item index="1-2"><SubMoney/></el-menu-item>
              </el-menu-item-group>
            </el-submenu>
            <el-submenu index="2">
              <template slot="title">
                <i class="el-icon-menu"></i>
                <span slot="title">查询业务</span>
              </template>
              <el-menu-item-group>
                <span slot="title">查钱</span>
                <el-menu-item index="1-1"><ShowMoney/></el-menu-item>
              </el-menu-item-group>
            </el-submenu>
            <el-submenu index="3">
              <template slot="title">
                <i class="el-icon-document"></i>
                <span slot="title">转账业务</span>
              </template>
              <el-menu-item-group>
                <span slot="title">转钱</span>
                <el-menu-item index="1-1"><TransferMoney/></el-menu-item>
              </el-menu-item-group>
            </el-submenu>
          </el-menu>
        </el-aside>
        <el-main>
          <el-breadcrumb
            separator-class="el-icon-arrow-right"
            v-if="this.$router.currentRoute.path!='/home'"
          >
            <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>{{this.$router.currentRoute.name}}</el-breadcrumb-item>
          </el-breadcrumb>
          <div class="homeWelcome" v-if="this.$router.currentRoute.path=='/home'">欢迎给我打钱,</div>
          <div
            class="homeWelcome"
            v-if="this.$router.currentRoute.path=='/home'"
          >支付宝 139 - XXXX - 0307.</div>
          <div class="homeWelcome" v-if="this.$router.currentRoute.path=='/home'">多多益善好人一生平安.</div>
          <router-view class="homeRouterView" />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import image from "./../assets/29779367.jpeg";
import AddMoney from "./components/AddMoney";
import SubMoney from "./components/SubMoney";
import ShowMoney from "./components/ShowMoney";
import TransferMoney from "./components/TransferMoney";

export default {
  name: "Home",
  data() {
    return {
      user: JSON.parse(window.sessionStorage.getItem("user")),
      userface: image,
      isCollapse: true
    };
  },
  components: {
      AddMoney,
      SubMoney,
      ShowMoney,
      TransferMoney
  },
  computed: {
    routes() {
      return this.$store.state.routes;
    }
  },
  methods: {
    goChat() {
      this.$router.push("/chat");
    },
    collapseStatus() {
      this.collapseBtnClick = this.isCollapse;
      this.isCollapse = !this.isCollapse;
    },
    collapseOpen() {
      if (this.collapseBtnClick) return;
      this.isCollapse = false;
    },
    collapseClose() {
      if (this.collapseBtnClick) return;
      this.isCollapse = true;
    },
    commandHandler(cmd) {
      if (cmd == "logout") {
        this.$confirm("此操作将注销登录, 是否继续?", "提示", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        })
          .then(() => {
            this.getRequest("/logout");
            window.sessionStorage.removeItem("user");
            this.$store.commit("initRoutes", []);
            this.$router.replace("/");
          })
          .catch(() => {
            this.$message({
              type: "info",
              message: "已取消操作"
            });
          });
      }
    }
  }
};
</script>

<style>
.homeRouterView {
  margin-top: 10px;
}

.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}

.homeWelcome {
  text-align: center;
  font-size: 30px;
  font-family: 华文行楷;
  color: #409eff;
  padding-top: 50px;
}

.homeHeader {
  background-color: #409eff;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0px 15px;
  box-sizing: border-box;
}

.homeHeader .title {
  font-size: 30px;
  font-family: 华文行楷;
  color: #ffffff;
}

.homeHeader .userInfo {
  cursor: pointer;
}

.el-dropdown-link img {
  width: 48px;
  height: 48px;
  border-radius: 24px;
  margin-left: 8px;
}

.el-dropdown-link {
  display: flex;
  align-items: center;
}
</style>