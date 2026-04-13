<template>
  <div>
    <div style="display: flex; justify-content: space-between; margin-bottom: 10px">
      <div>
        <el-input
          placeholder="请输入员工姓名搜索..."
          prefix-icon="el-icon-search"
          clearable
          style="width: 300px; margin-right: 10px"
          v-model="keyword"
          @keydown.enter.native="initResignation"
        ></el-input>
        <el-button type="primary" icon="el-icon-search" @click="initResignation">搜索</el-button>
      </div>
      <div>
        <el-button type="primary" icon="el-icon-plus" @click="showAddDialog">发起离职申请</el-button>
      </div>
    </div>

    <el-table :data="resignationList" border style="width: 100%" v-loading="loading">
      <el-table-column prop="employeeNo" label="工号" width="120"></el-table-column>
      <el-table-column prop="employeeName" label="员工姓名" width="100"></el-table-column>
      <el-table-column prop="department" label="部门" width="100"></el-table-column>
      <el-table-column prop="position" label="职位" width="100"></el-table-column>
      <el-table-column prop="applyDate" label="申请日期" width="120">
        <template slot-scope="scope">
          {{ formatDate(scope.row.applyDate) }}
        </template>
      </el-table-column>
      <el-table-column prop="lastWorkingDate" label="最后工作日" width="120">
        <template slot-scope="scope">
          {{ formatDate(scope.row.lastWorkingDate) }}
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态" width="100">
        <template slot-scope="scope">
          <el-tag :type="getStatusType(scope.row.status)">{{ getStatusText(scope.row.status) }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="currentStep" label="当前步骤" width="100">
        <template slot-scope="scope">
          <el-steps :active="scope.row.currentStep" simple size="mini">
            <el-step title="申请"></el-step>
            <el-step title="审批"></el-step>
            <el-step title="交接"></el-step>
            <el-step title="归还"></el-step>
            <el-step title="结算"></el-step>
          </el-steps>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="350">
        <template slot-scope="scope">
          <el-button size="mini" @click="showDetail(scope.row)">详情</el-button>
          <el-button size="mini" type="success" @click="approveResignation(scope.row)" v-if="scope.row.status === 'pending'">审批</el-button>
          <el-button size="mini" type="warning" @click="rejectResignation(scope.row)" v-if="scope.row.status === 'pending'">拒绝</el-button>
          <el-button size="mini" type="info" @click="generateCertificate(scope.row)" v-if="scope.row.currentStep >= 4">生成证明</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog title="发起离职申请" :visible.sync="addDialogVisible" width="600px">
      <el-form :model="resignationForm" label-width="100px">
        <el-form-item label="员工姓名">
          <el-input v-model="resignationForm.employeeName"></el-input>
        </el-form-item>
        <el-form-item label="工号">
          <el-input v-model="resignationForm.employeeNo"></el-input>
        </el-form-item>
        <el-form-item label="部门">
          <el-input v-model="resignationForm.department"></el-input>
        </el-form-item>
        <el-form-item label="职位">
          <el-input v-model="resignationForm.position"></el-input>
        </el-form-item>
        <el-form-item label="申请日期">
          <el-date-picker v-model="resignationForm.applyDate" type="date" placeholder="选择日期"></el-date-picker>
        </el-form-item>
        <el-form-item label="最后工作日">
          <el-date-picker v-model="resignationForm.lastWorkingDate" type="date" placeholder="选择日期"></el-date-picker>
        </el-form-item>
        <el-form-item label="离职原因">
          <el-input type="textarea" v-model="resignationForm.resignationReason"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="addDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createResignation">确定</el-button>
      </span>
    </el-dialog>

    <el-dialog title="离职流程详情" :visible.sync="detailDialogVisible" width="900px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="工号">{{ currentProcess.employeeNo }}</el-descriptions-item>
        <el-descriptions-item label="员工姓名">{{ currentProcess.employeeName }}</el-descriptions-item>
        <el-descriptions-item label="部门">{{ currentProcess.department }}</el-descriptions-item>
        <el-descriptions-item label="职位">{{ currentProcess.position }}</el-descriptions-item>
        <el-descriptions-item label="申请日期">{{ formatDate(currentProcess.applyDate) }}</el-descriptions-item>
        <el-descriptions-item label="最后工作日">{{ formatDate(currentProcess.lastWorkingDate) }}</el-descriptions-item>
        <el-descriptions-item label="离职原因" :span="2">{{ currentProcess.resignationReason }}</el-descriptions-item>
      </el-descriptions>

      <el-divider content-position="left">工作交接清单</el-divider>
      <div style="margin-bottom: 10px">
        <el-button size="small" type="primary" @click="showAddHandoverDialog">添加交接项</el-button>
      </div>
      <el-table :data="currentProcess.handoverItems" border size="small">
        <el-table-column prop="itemName" label="交接项目" width="150"></el-table-column>
        <el-table-column prop="itemType" label="类型" width="100">
          <template slot-scope="scope">
            <el-tag size="small">{{ getItemTypeText(scope.row.itemType) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="itemDescription" label="描述" width="200"></el-table-column>
        <el-table-column prop="receiverName" label="接收人" width="100"></el-table-column>
        <el-table-column prop="status" label="状态" width="80">
          <template slot-scope="scope">
            <el-tag :type="scope.row.status === 'completed' ? 'success' : 'info'" size="small">
              {{ scope.row.status === 'completed' ? '已完成' : '待完成' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="80">
          <template slot-scope="scope">
            <el-button
              size="mini"
              type="success"
              @click="completeHandover(scope.row)"
              :disabled="scope.row.status === 'completed'"
            >完成</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-divider content-position="left">资产归还清单</el-divider>
      <div style="margin-bottom: 10px">
        <el-button size="small" type="primary" @click="showAddAssetDialog">添加资产</el-button>
      </div>
      <el-table :data="currentProcess.assetReturns" border size="small">
        <el-table-column prop="assetName" label="资产名称" width="150"></el-table-column>
        <el-table-column prop="assetType" label="资产类型" width="100">
          <template slot-scope="scope">
            <el-tag size="small">{{ getAssetTypeText(scope.row.assetType) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="assetNo" label="资产编号" width="120"></el-table-column>
        <el-table-column prop="status" label="状态" width="80">
          <template slot-scope="scope">
            <el-tag :type="getAssetStatusType(scope.row.status)" size="small">
              {{ getAssetStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="receiverName" label="接收人" width="100"></el-table-column>
        <el-table-column label="操作" width="80">
          <template slot-scope="scope">
            <el-button
              size="mini"
              type="success"
              @click="completeAssetReturn(scope.row)"
              :disabled="scope.row.status === 'returned'"
            >归还</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-divider content-position="left">薪资结算</el-divider>
      <div style="margin-bottom: 10px">
        <el-button size="small" type="primary" @click="calculateSalary">计算薪资</el-button>
      </div>
      <el-descriptions :column="4" border v-if="salarySettlement.id">
        <el-descriptions-item label="基本工资">{{ salarySettlement.baseSalary }}</el-descriptions-item>
        <el-descriptions-item label="工作天数">{{ salarySettlement.workDays }}</el-descriptions-item>
        <el-descriptions-item label="奖金">{{ salarySettlement.bonus }}</el-descriptions-item>
        <el-descriptions-item label="扣款">{{ salarySettlement.deduction }}</el-descriptions-item>
        <el-descriptions-item label="最终金额" :span="4">
          <span style="font-size: 18px; font-weight: bold; color: #409EFF">{{ salarySettlement.finalAmount }}</span>
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>

    <el-dialog title="添加交接项" :visible.sync="handoverDialogVisible" width="500px">
      <el-form :model="handoverForm" label-width="100px">
        <el-form-item label="交接项目">
          <el-input v-model="handoverForm.itemName"></el-input>
        </el-form-item>
        <el-form-item label="类型">
          <el-select v-model="handoverForm.itemType" placeholder="请选择">
            <el-option label="项目" value="project"></el-option>
            <el-option label="文档" value="document"></el-option>
            <el-option label="权限" value="permission"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="描述">
          <el-input type="textarea" v-model="handoverForm.itemDescription"></el-input>
        </el-form-item>
        <el-form-item label="接收人">
          <el-input v-model="handoverForm.receiverName"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handoverDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="addHandoverItem">确定</el-button>
      </span>
    </el-dialog>

    <el-dialog title="添加资产" :visible.sync="assetDialogVisible" width="500px">
      <el-form :model="assetForm" label-width="100px">
        <el-form-item label="资产名称">
          <el-input v-model="assetForm.assetName"></el-input>
        </el-form-item>
        <el-form-item label="资产类型">
          <el-select v-model="assetForm.assetType" placeholder="请选择">
            <el-option label="电脑" value="computer"></el-option>
            <el-option label="办公用品" value="office"></el-option>
            <el-option label="门禁卡" value="card"></el-option>
            <el-option label="其他" value="other"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="资产编号">
          <el-input v-model="assetForm.assetNo"></el-input>
        </el-form-item>
        <el-form-item label="接收人">
          <el-input v-model="assetForm.receiverName"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="assetDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="addAssetReturn">确定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'ResignationProcess',
  data() {
    return {
      loading: false,
      keyword: '',
      resignationList: [],
      addDialogVisible: false,
      detailDialogVisible: false,
      handoverDialogVisible: false,
      assetDialogVisible: false,
      resignationForm: {
        employeeName: '',
        employeeNo: '',
        department: '',
        position: '',
        applyDate: '',
        lastWorkingDate: '',
        resignationReason: ''
      },
      currentProcess: {
        handoverItems: [],
        assetReturns: []
      },
      handoverForm: {
        itemName: '',
        itemType: '',
        itemDescription: '',
        receiverName: ''
      },
      assetForm: {
        assetName: '',
        assetType: '',
        assetNo: '',
        receiverName: ''
      },
      salarySettlement: {}
    }
  },
  mounted() {
    this.initResignation()
  },
  methods: {
    initResignation() {
      this.loading = true
      this.postRequest('/resignation/list', { keyword: this.keyword }).then(resp => {
        this.loading = false
        if (resp) {
          this.resignationList = resp.data.list || []
        }
      })
    },
    showAddDialog() {
      this.addDialogVisible = true
      this.resignationForm = {
        employeeName: '',
        employeeNo: '',
        department: '',
        position: '',
        applyDate: '',
        lastWorkingDate: '',
        resignationReason: ''
      }
    },
    createResignation() {
      this.postRequest('/resignation/create', this.resignationForm).then(resp => {
        if (resp) {
          this.addDialogVisible = false
          this.initResignation()
        }
      })
    },
    showDetail(row) {
      this.postRequest('/resignation/detail', { id: row.id }).then(resp => {
        if (resp) {
          this.currentProcess = resp.data.process
          this.detailDialogVisible = true
        }
      })
    },
    approveResignation(row) {
      this.$prompt('请输入审批意见', '审批通过', {
        confirmButtonText: '确定',
        cancelButtonText: '取消'
      }).then(({ value }) => {
        this.postRequest('/resignation/approve', { id: row.id, approvalRemark: value }).then(resp => {
          if (resp) {
            this.initResignation()
          }
        })
      })
    },
    rejectResignation(row) {
      this.$prompt('请输入拒绝原因', '拒绝申请', {
        confirmButtonText: '确定',
        cancelButtonText: '取消'
      }).then(({ value }) => {
        this.postRequest('/resignation/reject', { id: row.id, approvalRemark: value }).then(resp => {
          if (resp) {
            this.initResignation()
          }
        })
      })
    },
    showAddHandoverDialog() {
      this.handoverDialogVisible = true
      this.handoverForm = {
        itemName: '',
        itemType: '',
        itemDescription: '',
        receiverName: ''
      }
    },
    addHandoverItem() {
      this.handoverForm.resignationProcessId = this.currentProcess.id
      this.postRequest('/resignation/addHandoverItem', this.handoverForm).then(resp => {
        if (resp) {
          this.handoverDialogVisible = false
          this.showDetail(this.currentProcess)
        }
      })
    },
    completeHandover(item) {
      this.postRequest('/resignation/completeHandoverItem', { id: item.id }).then(resp => {
        if (resp) {
          this.showDetail(this.currentProcess)
        }
      })
    },
    showAddAssetDialog() {
      this.assetDialogVisible = true
      this.assetForm = {
        assetName: '',
        assetType: '',
        assetNo: '',
        receiverName: ''
      }
    },
    addAssetReturn() {
      this.assetForm.resignationProcessId = this.currentProcess.id
      this.postRequest('/resignation/addAssetReturn', this.assetForm).then(resp => {
        if (resp) {
          this.assetDialogVisible = false
          this.showDetail(this.currentProcess)
        }
      })
    },
    completeAssetReturn(asset) {
      this.postRequest('/resignation/completeAssetReturn', { id: asset.id }).then(resp => {
        if (resp) {
          this.showDetail(this.currentProcess)
        }
      })
    },
    calculateSalary() {
      this.$prompt('请输入基本工资', '薪资计算', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputPattern: /^[0-9]+(\.[0-9]+)?$/,
        inputErrorMessage: '请输入有效金额'
      }).then(({ value }) => {
        this.postRequest('/resignation/calculateSalary', {
          resignationProcessId: this.currentProcess.id,
          baseSalary: parseFloat(value),
          workDays: 22,
          bonus: 0,
          deduction: 0
        }).then(resp => {
          if (resp) {
            this.salarySettlement = resp.data.settlement
          }
        })
      })
    },
    generateCertificate(row) {
      this.postRequest('/resignation/generateCertificate', { id: row.id }).then(resp => {
        if (resp) {
          this.$message.success('离职证明已生成')
          this.initResignation()
        }
      })
    },
    formatDate(date) {
      if (!date) return ''
      return new Date(date).toLocaleDateString()
    },
    getStatusType(status) {
      const map = {
        pending: 'info',
        approved: 'success',
        rejected: 'danger',
        processing: 'warning',
        completed: 'success'
      }
      return map[status] || 'info'
    },
    getStatusText(status) {
      const map = {
        pending: '待审批',
        approved: '已通过',
        rejected: '已拒绝',
        processing: '进行中',
        completed: '已完成'
      }
      return map[status] || status
    },
    getItemTypeText(type) {
      const map = {
        project: '项目',
        document: '文档',
        permission: '权限'
      }
      return map[type] || type
    },
    getAssetTypeText(type) {
      const map = {
        computer: '电脑',
        office: '办公用品',
        card: '门禁卡',
        other: '其他'
      }
      return map[type] || type
    },
    getAssetStatusType(status) {
      const map = {
        pending: 'info',
        returned: 'success',
        damaged: 'warning',
        lost: 'danger'
      }
      return map[status] || 'info'
    },
    getAssetStatusText(status) {
      const map = {
        pending: '待归还',
        returned: '已归还',
        damaged: '损坏',
        lost: '丢失'
      }
      return map[status] || status
    }
  }
}
</script>

<style scoped>
</style>
