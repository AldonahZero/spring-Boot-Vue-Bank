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
          @keydown.enter.native="initOnboarding"
        ></el-input>
        <el-button type="primary" icon="el-icon-search" @click="initOnboarding">搜索</el-button>
      </div>
      <div>
        <el-button type="primary" icon="el-icon-plus" @click="showAddDialog">新增入职流程</el-button>
      </div>
    </div>

    <el-table :data="onboardingList" border style="width: 100%" v-loading="loading">
      <el-table-column prop="employeeNo" label="工号" width="120"></el-table-column>
      <el-table-column prop="employeeName" label="员工姓名" width="100"></el-table-column>
      <el-table-column prop="email" label="邮箱" width="180"></el-table-column>
      <el-table-column prop="department" label="部门" width="100"></el-table-column>
      <el-table-column prop="position" label="职位" width="100"></el-table-column>
      <el-table-column prop="onboardingDate" label="入职日期" width="120">
        <template slot-scope="scope">
          {{ formatDate(scope.row.onboardingDate) }}
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
            <el-step title="预约"></el-step>
            <el-step title="IT准备"></el-step>
            <el-step title="行政准备"></el-step>
            <el-step title="入职引导"></el-step>
          </el-steps>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="280">
        <template slot-scope="scope">
          <el-button size="mini" @click="showDetail(scope.row)">详情</el-button>
          <el-button size="mini" type="success" @click="advanceStep(scope.row)" :disabled="scope.row.status === 'completed'">推进</el-button>
          <el-button size="mini" type="danger" @click="deleteOnboarding(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog title="新增入职流程" :visible.sync="addDialogVisible" width="600px">
      <el-form :model="onboardingForm" label-width="100px">
        <el-form-item label="员工姓名">
          <el-input v-model="onboardingForm.employeeName"></el-input>
        </el-form-item>
        <el-form-item label="部门">
          <el-input v-model="onboardingForm.department"></el-input>
        </el-form-item>
        <el-form-item label="职位">
          <el-input v-model="onboardingForm.position"></el-input>
        </el-form-item>
        <el-form-item label="入职日期">
          <el-date-picker v-model="onboardingForm.onboardingDate" type="date" placeholder="选择日期"></el-date-picker>
        </el-form-item>
        <el-form-item label="联系电话">
          <el-input v-model="onboardingForm.phone"></el-input>
        </el-form-item>
        <el-form-item label="紧急联系人">
          <el-input v-model="onboardingForm.emergencyContact"></el-input>
        </el-form-item>
        <el-form-item label="紧急联系电话">
          <el-input v-model="onboardingForm.emergencyPhone"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="addDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createOnboarding">确定</el-button>
      </span>
    </el-dialog>

    <el-dialog title="入职流程详情" :visible.sync="detailDialogVisible" width="800px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="工号">{{ currentProcess.employeeNo }}</el-descriptions-item>
        <el-descriptions-item label="员工姓名">{{ currentProcess.employeeName }}</el-descriptions-item>
        <el-descriptions-item label="邮箱">{{ currentProcess.email }}</el-descriptions-item>
        <el-descriptions-item label="部门">{{ currentProcess.department }}</el-descriptions-item>
        <el-descriptions-item label="职位">{{ currentProcess.position }}</el-descriptions-item>
        <el-descriptions-item label="入职日期">{{ formatDate(currentProcess.onboardingDate) }}</el-descriptions-item>
      </el-descriptions>

      <el-divider content-position="left">入职检查清单</el-divider>
      <el-table :data="currentProcess.checklistItems" border size="small">
        <el-table-column prop="itemName" label="检查项" width="200"></el-table-column>
        <el-table-column prop="itemCategory" label="分类" width="100">
          <template slot-scope="scope">
            <el-tag size="small">{{ getCategoryText(scope.row.itemCategory) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="isCompleted" label="状态" width="100">
          <template slot-scope="scope">
            <el-tag :type="scope.row.isCompleted ? 'success' : 'info'" size="small">
              {{ scope.row.isCompleted ? '已完成' : '待完成' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="completedBy" label="完成人" width="100"></el-table-column>
        <el-table-column label="操作" width="100">
          <template slot-scope="scope">
            <el-button
              size="mini"
              type="success"
              @click="completeChecklist(scope.row)"
              :disabled="scope.row.isCompleted"
            >完成</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'OnboardingProcess',
  data() {
    return {
      loading: false,
      keyword: '',
      onboardingList: [],
      addDialogVisible: false,
      detailDialogVisible: false,
      onboardingForm: {
        employeeName: '',
        department: '',
        position: '',
        onboardingDate: '',
        phone: '',
        emergencyContact: '',
        emergencyPhone: ''
      },
      currentProcess: {
        checklistItems: []
      }
    }
  },
  mounted() {
    this.initOnboarding()
  },
  methods: {
    initOnboarding() {
      this.loading = true
      this.postRequest('/onboarding/list', { keyword: this.keyword }).then(resp => {
        this.loading = false
        if (resp) {
          this.onboardingList = resp.data.list || []
        }
      })
    },
    showAddDialog() {
      this.addDialogVisible = true
      this.onboardingForm = {
        employeeName: '',
        department: '',
        position: '',
        onboardingDate: '',
        phone: '',
        emergencyContact: '',
        emergencyPhone: ''
      }
    },
    createOnboarding() {
      this.postRequest('/onboarding/create', this.onboardingForm).then(resp => {
        if (resp) {
          this.addDialogVisible = false
          this.initOnboarding()
        }
      })
    },
    showDetail(row) {
      this.postRequest('/onboarding/detail', { id: row.id }).then(resp => {
        if (resp) {
          this.currentProcess = resp.data.process
          this.detailDialogVisible = true
        }
      })
    },
    advanceStep(row) {
      this.postRequest('/onboarding/advanceStep', { id: row.id }).then(resp => {
        if (resp) {
          this.initOnboarding()
        }
      })
    },
    deleteOnboarding(row) {
      this.$confirm('确定删除该入职流程?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.postRequest('/onboarding/delete', { id: row.id }).then(resp => {
          if (resp) {
            this.initOnboarding()
          }
        })
      })
    },
    completeChecklist(item) {
      this.postRequest('/onboarding/completeChecklist', { id: item.id, completedBy: 'admin' }).then(resp => {
        if (resp) {
          this.showDetail(this.currentProcess)
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
        processing: 'warning',
        completed: 'success'
      }
      return map[status] || 'info'
    },
    getStatusText(status) {
      const map = {
        pending: '待处理',
        processing: '进行中',
        completed: '已完成'
      }
      return map[status] || status
    },
    getCategoryText(category) {
      const map = {
        it: 'IT部门',
        admin: '行政部门',
        mentor: '导师'
      }
      return map[category] || category
    }
  }
}
</script>

<style scoped>
</style>
