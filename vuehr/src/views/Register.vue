<template>
    <div>
        
        <el-form
                :rules="rules"
                ref="loginForm"
                v-loading="loading"
                element-loading-text="正在登录..."
                element-loading-spinner="el-icon-loading"
                element-loading-background="rgba(0, 0, 0, 0.8)"
                :model="loginForm"
                class="loginContainer">
            <h2 class="loginTitle"><i class="el-icon-platform-eleme warpperfront">我,秦始皇.[打钱]</i></h2>
            <h2 class="loginTitle">注册</h2>
            <el-form-item prop="userId">
                <el-input size="normal" type="text" v-model="loginForm.userId" auto-complete="off"
                          placeholder="请输入用户ID"></el-input>
            </el-form-item>
            <el-form-item prop="userName">
                <el-input size="normal" type="text" v-model="loginForm.userName" auto-complete="off"
                          placeholder="请输入用户名"></el-input>
            </el-form-item>
            <el-form-item prop="userPassword">
                <el-input size="normal" type="password" v-model="loginForm.userPassword" auto-complete="off"
                          placeholder="请输入密码" @keydown.enter.native="submitLogin"></el-input>
            </el-form-item>
            <el-form-item>
                <el-link href="/#/" type="primary" >登录账号</el-link>
            </el-form-item>
            
            <el-button size="normal" type="success" round style="width: 100%;" @click="submitLogin">注册</el-button>
        </el-form>
    </div>
</template>

<script>

    export default {
        name: "Register",
        data() {
            return {
                loading: false,
                loginForm: {
                },
                checked: true,
                rules: {
                    userId: [{required: true, message: '请输入用户名', trigger: 'blur'}],
                    userPassword: [{required: true, message: '请输入密码', trigger: 'blur'}]
                }
            }
        },
        methods: {
            submitLogin() {
                this.$refs.loginForm.validate((valid) => {
                    if (valid) {
                        this.loading = true;
                        this.postKeyValueRequest('/user/register', this.loginForm).then(resp => {
                            this.loading = false;
                            if (resp) {
                                this.$store.commit('INIT_CURRENTHR', resp);
                                let path = this.$route.query.redirect;
                                this.$router.replace((path == '/' || path == undefined) ? '/' : path);
                            }
                        })
                    } else {
                        this.$message.error('请输入所有字段');
                        return false;
                    }
                });
            }
        }
    }
</script>

<style>
    .loginContainer {
        border-radius: 15px;
        background-clip: padding-box;
        margin: 180px auto;
        width: 350px;
        padding: 15px 35px 15px 35px;
        background: #fff;
        border: 1px solid #eaeaea;
        box-shadow: 0 0 25px #cac6c6;
    }

    .loginTitle {
        margin: 15px auto 20px auto;
        text-align: center;
        color: #505458;
    }

    .loginRemember {
        text-align: left;
        margin: 0px 0px 15px 0px;
    }

    .warpperfront{
        font-size: 30px;
    }
</style>