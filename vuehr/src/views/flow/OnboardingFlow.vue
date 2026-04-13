<template>
    <div>
        <el-card>
            <div slot="header" class="clearfix">
                <span>入职流程管理</span>
                <el-button style="float: right; padding: 3px 0" type="text" @click="showAddDialog">发起入职流程</el-button>
            </div>
            
            <el-row :gutter="20" style="margin-bottom: 20px;">
                <el-col :span="6">
                    <el-statistic title="待入职" :value="statistics.pending" value-style="color: #E6A23C">
                        <template slot="prefix">
                            <i class="el-icon-time"></i>
                        </template>
                    </el-statistic>
                </el-col>
                <el-col :span="6">
                    <el-statistic title="进行中" :value="statistics.inProgress" value-style="color: #409EFF">
                        <template slot="prefix">
                            <i class="el-icon-loading"></i>
                        </template>
                    </el-statistic>
                </el-col>
                <el-col :span="6">
                    <el-statistic title="本月已完成" :value="statistics.completed" value-style="color: #67C23A">
                        <template slot="prefix">
                            <i class="el-icon-circle-check"></i>
                        </template>
                    </el-statistic>
                </el-col>
                <el-col :span="6">
                    <el-statistic title="总流程数" :value="statistics.total">
                        <template slot="prefix">
                            <i class="el-icon-document"></i>
                        </template>
                    </el-statistic>
                </el-col>
            </el-row>

            <el-tabs v-model="activeTab">
                <el-tab-pane label="入职流程列表" name="list">
                    <el-table :data="onboardingList" border stripe v-loading="loading">
                        <el-table-column prop="processNo" label="流程编号" width="120"></el-table-column>
                        <el-table-column prop="employeeName" label="员工姓名" width="100"></el-table-column>
                        <el-table-column prop="department.name" label="入职部门" width="120"></el-table-column>
                        <el-table-column prop="position.name" label="职位" width="100"></el-table-column>
                        <el-table-column prop="expectedDate" label="预计入职日期" width="120"></el-table-column>
                        <el-table-column prop="currentStep" label="当前节点" width="120">
                            <template slot-scope="scope">
                                <el-tag :type="getStepType(scope.row.currentStep)">{{scope.row.currentStep}}</el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column prop="status" label="状态" width="100">
                            <template slot-scope="scope">
                                <el-tag :type="scope.row.status === '已完成' ? 'success' : scope.row.status === '进行中' ? 'primary' : 'warning'">{{scope.row.status}}</el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column label="操作" width="250" fixed="right">
                            <template slot-scope="scope">
                                <el-button size="mini" type="primary" @click="viewDetail(scope.row)">查看进度</el-button>
                                <el-button size="mini" @click="viewTasks(scope.row)">待办任务</el-button>
                                <el-button size="mini" type="success" @click="viewChecklist(scope.row)">Checklist</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-tab-pane>
                <el-tab-pane label="部门待办任务" name="tasks">
                    <el-row :gutter="20">
                        <el-col :span="8">
                            <el-card shadow="hover">
                                <div slot="header">
                                    <i class="el-icon-s-tools"></i> IT部门待办
                                    <el-badge :value="itTasks.length" class="item" style="float:right"></el-badge>
                                </div>
                                <el-table :data="itTasks" size="small">
                                    <el-table-column prop="employeeName" label="员工"></el-table-column>
                                    <el-table-column prop="taskName" label="任务"></el-table-column>
                                    <el-table-column label="操作" width="100">
                                        <template slot-scope="scope">
                                            <el-button size="mini" type="success" :disabled="scope.row.completed" @click="completeTask(scope.row, 'it')">完成</el-button>
                                        </template>
                                    </el-table-column>
                                </el-table>
                            </el-card>
                        </el-col>
                        <el-col :span="8">
                            <el-card shadow="hover">
                                <div slot="header">
                                    <i class="el-icon-s-shop"></i> 行政部门待办
                                    <el-badge :value="adminTasks.length" class="item" style="float:right"></el-badge>
                                </div>
                                <el-table :data="adminTasks" size="small">
                                    <el-table-column prop="employeeName" label="员工"></el-table-column>
                                    <el-table-column prop="taskName" label="任务"></el-table-column>
                                    <el-table-column label="操作" width="100">
                                        <template slot-scope="scope">
                                            <el-button size="mini" type="success" :disabled="scope.row.completed" @click="completeTask(scope.row, 'admin')">完成</el-button>
                                        </template>
                                    </el-table-column>
                                </el-table>
                            </el-card>
                        </el-col>
                        <el-col :span="8">
                            <el-card shadow="hover">
                                <div slot="header">
                                    <i class="el-icon-user"></i> 部门导师待办
                                    <el-badge :value="mentorTasks.length" class="item" style="float:right"></el-badge>
                                </div>
                                <el-table :data="mentorTasks" size="small">
                                    <el-table-column prop="employeeName" label="员工"></el-table-column>
                                    <el-table-column prop="taskName" label="任务"></el-table-column>
                                    <el-table-column label="操作" width="100">
                                        <template slot-scope="scope">
                                            <el-button size="mini" type="success" :disabled="scope.row.completed" @click="completeTask(scope.row, 'mentor')">完成</el-button>
                                        </template>
                                    </el-table-column>
                                </el-table>
                            </el-card>
                        </el-col>
                    </el-row>
                </el-tab-pane>
            </el-tabs>
        </el-card>

        <el-dialog :title="dialogTitle" :visible.sync="dialogVisible" width="60%">
            <el-form :model="onboardingForm" :rules="rules" ref="onboardingForm" label-width="100px">
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="员工姓名" prop="employeeName">
                            <el-input v-model="onboardingForm.employeeName"></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="性别" prop="gender">
                            <el-radio-group v-model="onboardingForm.gender">
                                <el-radio label="男">男</el-radio>
                                <el-radio label="女">女</el-radio>
                            </el-radio-group>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="入职部门" prop="departmentId">
                            <el-select v-model="onboardingForm.departmentId" placeholder="请选择部门" style="width: 100%">
                                <el-option v-for="dept in departments" :key="dept.id" :label="dept.name" :value="dept.id"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="职位" prop="positionId">
                            <el-select v-model="onboardingForm.positionId" placeholder="请选择职位" style="width: 100%">
                                <el-option v-for="pos in positions" :key="pos.id" :label="pos.name" :value="pos.id"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="预计入职日期" prop="expectedDate">
                            <el-date-picker v-model="onboardingForm.expectedDate" type="date" style="width: 100%" value-format="yyyy-MM-dd"></el-date-picker>
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="邮箱" prop="email">
                            <el-input v-model="onboardingForm.email" placeholder="将自动生成企业邮箱"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="联系电话" prop="phone">
                            <el-input v-model="onboardingForm.phone"></el-input>
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="部门导师">
                            <el-select v-model="onboardingForm.mentor" placeholder="请选择导师" style="width: 100%">
                                <el-option label="张三" :value="1"></el-option>
                                <el-option label="李四" :value="2"></el-option>
                                <el-option label="王五" :value="3"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取消</el-button>
                <el-button type="primary" @click="submitOnboarding">发起流程</el-button>
            </div>
        </el-dialog>

        <el-dialog title="流程进度详情" :visible.sync="detailVisible" width="50%">
            <el-steps :active="currentStepIndex" direction="vertical" finish-status="success">
                <el-step title="预约登记完成" description="HR已完成入职预约登记，工号和邮箱已自动生成"></el-step>
                <el-step title="IT部门处理" description="电脑配置、系统账号开通"></el-step>
                <el-step title="行政部门处理" description="工位安排、门禁卡办理"></el-step>
                <el-step title="部门导师引导" description="欢迎邮件、入职引导培训"></el-step>
                <el-step title="入职完成" description="新人正式入职，Checklist全部完成"></el-step>
            </el-steps>
        </el-dialog>

        <el-dialog title="入职Checklist" :visible.sync="checklistVisible" width="40%">
            <el-table :data="checklist" border>
                <el-table-column prop="item" label="检查项"></el-table-column>
                <el-table-column prop="department" label="负责部门" width="100"></el-table-column>
                <el-table-column label="状态" width="100">
                    <template slot-scope="scope">
                        <el-tag :type="scope.row.completed ? 'success' : 'warning'">{{scope.row.completed ? '已完成' : '待完成'}}</el-tag>
                    </template>
                </el-table-column>
            </el-table>
        </el-dialog>
    </div>
</template>

<script>
export default {
    name: 'OnboardingFlow',
    data() {
        return {
            activeTab: 'list',
            loading: false,
            dialogVisible: false,
            dialogTitle: '',
            detailVisible: false,
            checklistVisible: false,
            currentStepIndex: 2,
            statistics: {
                pending: 12,
                inProgress: 8,
                completed: 25,
                total: 45
            },
            onboardingList: [
                { id: 1, processNo: 'OB202404001', employeeName: '张三', department: {name: '技术部'}, position: {name: 'Java开发'}, expectedDate: '2024-04-15', currentStep: 'IT处理', status: '进行中' },
                { id: 2, processNo: 'OB202404002', employeeName: '李四', department: {name: '产品部'}, position: {name: '产品经理'}, expectedDate: '2024-04-18', currentStep: '行政处理', status: '进行中' },
                { id: 3, processNo: 'OB202404003', employeeName: '王五', department: {name: '市场部'}, position: {name: '市场专员'}, expectedDate: '2024-04-20', currentStep: '预约登记', status: '待入职' },
                { id: 4, processNo: 'OB202404004', employeeName: '赵六', department: {name: '人事部'}, position: {name: 'HR专员'}, expectedDate: '2024-04-10', currentStep: '入职完成', status: '已完成' }
            ],
            itTasks: [
                { id: 1, employeeName: '张三', taskName: '配置办公电脑', completed: false },
                { id: 2, employeeName: '张三', taskName: '开通系统账号', completed: true },
                { id: 3, employeeName: '李四', taskName: '配置办公电脑', completed: false }
            ],
            adminTasks: [
                { id: 1, employeeName: '李四', taskName: '安排工位', completed: false },
                { id: 2, employeeName: '李四', taskName: '办理门禁卡', completed: false },
                { id: 3, employeeName: '王五', taskName: '安排工位', completed: false }
            ],
            mentorTasks: [
                { id: 1, employeeName: '张三', taskName: '发送欢迎邮件', completed: true },
                { id: 2, employeeName: '张三', taskName: '入职引导', completed: false },
                { id: 3, employeeName: '赵六', taskName: '入职引导', completed: true }
            ],
            checklist: [
                { item: '入职资料签收', department: 'HR', completed: true },
                { item: '劳动合同签订', department: 'HR', completed: true },
                { item: '办公电脑领取', department: 'IT', completed: false },
                { item: '门禁卡领取', department: '行政', completed: false },
                { item: '系统权限开通', department: 'IT', completed: true },
                { item: '部门见面会', department: '导师', completed: false }
            ],
            departments: [
                { id: 1, name: '技术部' },
                { id: 2, name: '产品部' },
                { id: 3, name: '市场部' },
                { id: 4, name: '人事部' },
                { id: 5, name: '财务部' }
            ],
            positions: [
                { id: 1, name: 'Java开发工程师' },
                { id: 2, name: '产品经理' },
                { id: 3, name: '市场专员' },
                { id: 4, name: 'HR专员' }
            ],
            onboardingForm: {
                employeeName: '',
                gender: '男',
                departmentId: null,
                positionId: null,
                expectedDate: '',
                email: '',
                phone: '',
                mentor: null
            },
            rules: {
                employeeName: [{ required: true, message: '请输入员工姓名', trigger: 'blur' }],
                departmentId: [{ required: true, message: '请选择入职部门', trigger: 'change' }],
                positionId: [{ required: true, message: '请选择职位', trigger: 'change' }],
                expectedDate: [{ required: true, message: '请选择预计入职日期', trigger: 'change' }]
            }
        }
    },
    methods: {
        getStepType(step) {
            const map = {
                '预约登记': 'warning',
                'IT处理': 'primary',
                '行政处理': 'info',
                '导师引导': '',
                '入职完成': 'success'
            }
            return map[step] || ''
        },
        showAddDialog() {
            this.dialogTitle = '发起入职流程'
            this.onboardingForm = {
                employeeName: '',
                gender: '男',
                departmentId: null,
                positionId: null,
                expectedDate: '',
                email: '',
                phone: '',
                mentor: null
            }
            this.dialogVisible = true
        },
        submitOnboarding() {
            this.$refs.onboardingForm.validate(valid => {
                if (valid) {
                    const processNo = 'OB' + new Date().getFullYear() + (new Date().getMonth() + 1).toString().padStart(2, '0') + (this.onboardingList.length + 1).toString().padStart(3, '0')
                    const workID = 'EMP' + (10000 + this.onboardingList.length + 1)
                    const email = this.onboardingForm.employeeName + '@company.com'
                    this.onboardingList.unshift({
                        id: this.onboardingList.length + 1,
                        processNo,
                        employeeName: this.onboardingForm.employeeName,
                        department: this.departments.find(d => d.id === this.onboardingForm.departmentId),
                        position: this.positions.find(p => p.id === this.onboardingForm.positionId),
                        expectedDate: this.onboardingForm.expectedDate,
                        currentStep: '预约登记',
                        status: '待入职',
                        workID,
                        email
                    })
                    this.$message.success(`入职流程发起成功！自动生成工号：${workID}，企业邮箱：${email}`)
                    this.dialogVisible = false
                    this.statistics.pending++
                    this.statistics.total++
                }
            })
        },
        viewDetail(row) {
            const stepMap = { '预约登记': 0, 'IT处理': 1, '行政处理': 2, '导师引导': 3, '入职完成': 4 }
            this.currentStepIndex = stepMap[row.currentStep] || 0
            this.detailVisible = true
        },
        viewChecklist() {
            this.checklistVisible = true
        },
        viewTasks() {
            this.activeTab = 'tasks'
        },
        completeTask(task, type) {
            const taskList = type === 'it' ? this.itTasks : type === 'admin' ? this.adminTasks : this.mentorTasks
            const index = taskList.findIndex(t => t.id === task.id)
            if (index > -1) {
                taskList[index].completed = true
                this.$message.success('任务已完成！')
            }
        }
    }
}
</script>

<style scoped>
.el-statistic {
    text-align: center;
}
</style>