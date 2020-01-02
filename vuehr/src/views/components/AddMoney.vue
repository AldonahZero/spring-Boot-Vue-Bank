<template>
  <div>
    <el-button type="text" @click="dialogFormVisible = true">存款</el-button>

    <el-dialog title="存款" :visible.sync="dialogFormVisible">
      <el-form :model="form">
        <el-form-item label="存款金额" :label-width="formLabelWidth">
          <el-input v-model="form.balance" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
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
        userId: JSON.parse(window.sessionStorage.getItem("user")).userId,
        balance: 0.0
      },
      formLabelWidth: "120px"
    };
  },
  methods: {
    submitLogin() {
      this.putRequest("/user/deposit", this.form).then(resp => {
        if (resp) {
          console.log(resp);
          this.dialogFormVisible = false;
        }
      });
    }
  }
};
</script>