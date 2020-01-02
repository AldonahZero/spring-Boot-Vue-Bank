<template>
  <div>
    <el-button type="text" @click="dialogFormVisible = true">转账</el-button>

    <el-dialog title="转账" :visible.sync="dialogFormVisible">
      <el-form :model="form">
        <el-form-item label="对方户头ID" :label-width="formLabelWidth">
          <el-input v-model="form.toId" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="转账金额" :label-width="formLabelWidth">
          <el-input v-model="form.transferNumber" autocomplete="off"></el-input>
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
        fromId: JSON.parse(window.sessionStorage.getItem("user")).userId,
        toId: JSON.parse(window.sessionStorage.getItem("user")).userId,
        transferNumber: 0.0
      },
      formLabelWidth: "120px"
    };
  },
  methods: {
    submitLogin() {
      this.postRequest("/user/transfer", this.form).then(resp => {
        if (resp) {
          console.log(resp);
          this.dialogFormVisible = false;
        }
      });
    }
  }
};
</script>