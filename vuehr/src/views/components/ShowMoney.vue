<template>
  <div>
    <el-button type="text" @click="dialogFormVisible = true">查询余额</el-button>

    <el-dialog title="查询余额" :visible.sync="dialogFormVisible">
      <el-form :model="form">
        <el-form-item label="余额" :label-width="formLabelWidth">
          <el-input disabled="true" v-model="form.balance" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitLogin">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      dialogFormVisible: false,
      form: {
         userId:JSON.parse(window.sessionStorage.getItem("user")).userId,
         balance: 0.0
      },  
      formLabelWidth: "120px"
    };
  },
  updated () {
    this.form.userId= JSON.parse(window.sessionStorage.getItem("user")).userId
    this.postRequest("/user/inquiry", this.form).then(resp => {
        if (resp) {
          console.log(resp);
          this.form.balance = resp.balance;
        }
      });
  },
  methods: {
    submitLogin() {
        this.dialogFormVisible = false;
    }
  }
};
</script>