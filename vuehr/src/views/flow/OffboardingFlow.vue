<template>
    <div>
        <el-card>
            <div slot="header" class="clearfix">
                <span>离职流程管理</span>
                <el-button style="float: right; padding: 3px 0" type="text" @click="showAddDialog">发起离职申请</el-button>
            </div>
            
            <el-row :gutter="20" style="margin-bottom: 20px;">
                <el-col :span="6">
                    <el-statistic title="待审批" :value="statistics.pending" value-style="color: #E6A23C">
                        <template slot="prefix">
                            <i class="el-icon-edit"></i>
                        </template>
                    </el-statistic>
                </el-col>
                <el-col :span="6">
                    <el-statistic title="交接中" :value="statistics.inProgress" value-style="color: #409EFF">
                        <template slot="prefix">
                            <i class="el-icon-refresh"></i>
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
                    <el-statistic title="待结算薪资" :value="statistics.salaryPending" value-style="color: #F56C6C">
                        <template slot="prefix">
                            <i class="el-icon-money"></i>
                        </template>
                        <template slot="suffix"> 人</template>
                    </el-statistic>
                </el-col>
            </el-row>

            <el-tabs v-model="activeTab">
                <el-tab-pane label="离职流程列表" name="list">
                    <el-table :data="offboardingList" border stripe v-loading="loading">
                        <el-table-column prop="processNo" label="流程编号" width="120"></el-table-column>
                        <el-table-column prop="employeeName" label="员工姓名" width="100"></el-table-column>
                        <el-table-column prop="department.name" label="所属部门" width="100"></el-table-column>
                        <el-table-column prop="position.name" label="职位" width="100"></el-table-column>
                        <el-table-column prop="lastDate" label="最后工作日" width="120"></el-table-column>
                        <el-table-column prop="currentStep" label="当前节点" width="120">
                            <template slot-scope="scope">
                                <el-tag :type="getStepType(scope.row.currentStep)">{{scope.row.currentStep}}</el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column prop="status" label="状态" width="100">
                            <template slot-scope="scope">
                                <el-tag :type="scope.row.status === '已完成' ? 'success' : scope.row.status === '交接中' ? 'primary' : 'warning'">{{scope.row.status}}</el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column label="操作" width="350" fixed="right">
                            <template slot-scope="scope">
                                <el-button size="mini" type="primary" @click="viewDetail(scope.row)">查看进度</el-button>
                                <el-button size="mini" @click="viewHandover(scope.row)">交接清单</el-button>
                                <el-button size="mini" type="warning" @click="viewAssets(scope.row)">资产归还</el-button>
                                <el-button size="mini" type="success" @click="viewSalary(scope.row)">薪资核算</el-button>
                                <el-button size="mini" type="info" @click="generateCertificate(scope.row)">离职证明</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-tab-pane>
                <el-tab-pane label="离职审批" name="approval">
                    <el-table :data="approvalList" border stripe>
                        <el-table-column type="selection"></el-table-column>
                        <el-table-column prop="employeeName" label="申请人" width="100"></el-table-column>
                        <el-table-column prop="department.name" label="部门" width="100"></el-table-column>
                        <el-table-column prop="reason" label="离职原因" show-overflow-tooltip></el-table-column>
                        <el-table-column prop="applyDate" label="申请日期" width="120"></el-table-column>
                        <el-table-column prop="lastDate" label="期望离职日" width="120"></el-table-column>
                        <el-table-column prop="approver" label="审批人" width="100"></el-table-column>
                        <el-table-column label="操作" width="200">
                            <template slot-scope="scope">
                                <el-button size="mini" type="success" @click="approve(scope.row)">通过</el-button>
                                <el-button size="mini" type="danger" @click="reject(scope.row)">驳回</el-button>
                                <el-button size="mini" @click="viewApprovalDetail(scope.row)">详情</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-tab-pane>
            </el-tabs>
        </el-card>

        <el-dialog :title="dialogTitle" :visible.sync="dialogVisible" width="60%">
            <el-form :model="offboardingForm" :rules="rules" ref="offboardingForm" label-width="120px">
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="离职员工">
                            <el-select v-model="offboardingForm.employeeId" placeholder="请选择员工" style="width: 100%">
                                <el-option v-for="emp in employees" :key="emp.id" :label="emp.name" :value="emp.id"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="离职类型">
                            <el-select v-model="offboardingForm.type" style="width: 100%">
                                <el-option label="主动离职" value="主动离职"></el-option>
                                <el-option label="被动离职" value="被动离职"></el-option>
                                <el-option label="退休" value="退休"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="12">
                        <el-form-item label="申请日期">
                            <el-date-picker v-model="offboardingForm.applyDate" type="date" style="width: 100%" value-format="yyyy-MM-dd"></el-date-picker>
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="最后工作日">
                            <el-date-picker v-model="offboardingForm.lastDate" type="date" style="width: 100%" value-format="yyyy-MM-dd"></el-date-picker>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-form-item label="离职原因">
                    <el-input type="textarea" v-model="offboardingForm.reason" :rows="3"></el-input>
                </el-form-item>
                <el-form-item label="直属上级确认">
                    <el-switch v-model="offboardingForm.managerConfirmed" active-text="已确认"></el-switch>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取消</el-button>
                <el-button type="primary" @click="submitOffboarding">提交申请</el-button>
            </div>
        </el-dialog>

        <el-dialog title="离职流程进度" :visible.sync="detailVisible" width="50%">
            <el-steps :active="currentStepIndex" direction="vertical" finish-status="success">
                <el-step title="离职申请" description="员工提交离职申请"></el-step>
                <el-step title="审批通过" description="直属上级、部门经理、HR审批"></el-step>
                <el-step title="工作交接" description="项目、文档、权限交接"></el-step>
                <el-step title="资产归还" description="电脑、办公用品归还确认"></el-step>
                <el-step title="薪资结算" description="薪资、补偿金自动核算"></el-step>
                <el-step title="离职完成" description="开具离职证明"></el-step>
            </el-steps>
        </el-dialog>

        <el-dialog title="工作交接清单" :visible.sync="handoverVisible" width="60%">
            <el-row :gutter="20">
                <el-col :span="12">
                    <h4>项目交接</h4>
                    <el-table :data="handover.projects" border size="small">
                        <el-table-column prop="name" label="项目名称"></el-table-column>
                        <el-table-column prop="receiver" label="接收人" width="100"></el-table-column>
                        <el-table-column label="状态" width="80">
                            <template slot-scope="scope">
                                <el-tag :type="scope.row.completed ? 'success' : 'warning'" size="mini">{{scope.row.completed ? '已交接' : '待交接'}}</el-tag>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-col>
                <el-col :span="12">
                    <h4>文档交接</h4>
                    <el-table :data="handover.documents" border size="small">
                        <el-table-column prop="name" label="文档名称"></el-table-column>
                        <el-table-column prop="receiver" label="接收人" width="100"></el-table-column>
                        <el-table-column label="状态" width="80">
                            <template slot-scope="scope">
                                <el-tag :type="scope.row.completed ? 'success' : 'warning'" size="mini">{{scope.row.completed ? '已交接' : '待交接'}}</el-tag>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-col>
            </el-row>
            <h4 style="margin-top: 20px;">权限交接</h4>
            <el-table :data="handover.permissions" border size="small">
                <el-table-column prop="name" label="系统/权限名称"></el-table-column>
                <el-table-column prop="handler" label="处理人" width="100"></el-table-column>
                <el-table-column label="状态" width="80">
                    <template slot-scope="scope">
                        <el-tag :type="scope.row.completed ? 'success' : 'warning'" size="mini">{{scope.row.completed ? '已回收' : '待回收'}}</el-tag>
                    </template>
                </el-table-column>
            </el-table>
        </el-dialog>

        <el-dialog title="资产归还确认" :visible.sync="assetsVisible" width="50%">
            <el-table :data="assetsList" border>
                <el-table-column prop="name" label="资产名称"></el-table-column>
                <el-table-column prop="code" label="资产编号" width="120"></el-table-column>
                <el-table-column prop="quantity" label="数量" width="80" align="center"></el-table-column>
                <el-table-column label="归还状态" width="100">
                    <template slot-scope="scope">
                        <el-checkbox v-model="scope.row.returned" @change="checkAssetsComplete"></el-checkbox>
                    </template>
                </el-table-column>
            </el-table>
            <div style="margin-top: 15px; text-align: right;">
                <el-button type="primary" :disabled="!allAssetsReturned" @click="confirmAssetsReturn">确认资产归还</el-button>
            </div>
        </el-dialog>

        <el-dialog title="薪资结算核算" :visible.sync="salaryVisible" width="50%">
            <el-descriptions :column="2" border>
                <el-descriptions-item label="基本工资">{{salaryDetail.basic}}</el-descriptions-item>
                <el-descriptions-item label="绩效工资">{{salaryDetail.performance}}</el-descriptions-item>
                <el-descriptions-item label="本月出勤天数">{{salaryDetail.attendanceDays}}天</el-descriptions-item>
                <el-descriptions-item label="实际出勤工资">{{salaryDetail.actualSalary}}</el-descriptions-item>
                <el-descriptions-item label="年假结算">{{salaryDetail.annualLeave}}</el-descriptions-item>
                <el-descriptions-item label="补偿金">{{salaryDetail.compensation}}</el-descriptions-item>
                <el-descriptions-item label="社保公积金扣除">{{salaryDetail.socialInsurance}}</el-descriptions-item>
                <el-descriptions-item label="个税扣除">{{salaryDetail.tax}}</el-descriptions-item>
                <el-descriptions-item label="合计发放" :span="2" label-class-name="total-salary">
                    <span style="font-size: 18px; color: #F56C6C; font-weight: bold;">{{salaryDetail.total}}</span>
                </el-descriptions-item>
            </el-descriptions>
            <div style="margin-top: 15px; text-align: right;">
                <el-button type="primary" @click="confirmSalary">确认结算</el-button>
            </div>
        </el-dialog>

        <el-dialog title="离职证明预览" :visible.sync="certificateVisible" width="50%">
            <div style="padding: 30px; border: 1px solid #ccc; background: #fff;">
                <h2 style="text-align: center; margin-bottom: 30px;">离职证明</h2>
                <p style="text-indent: 2em; line-height: 2;">兹证明 <strong>{{currentEmployee}}</strong> 先生/女士（身份证号：_________________），自______年____月____日入职我公司，原任______部门______职务。</p>
                <p style="text-indent: 2em; line-height: 2;">现已办理完所有离职手续，薪资已结清，劳动关系已解除。</p>
                <p style="text-indent: 2em; line-height: 2;">特此证明！</p>
                <p style="text-align: right; margin-top: 50px;">公司盖章：</p>
                <p style="text-align: right;">日期：{{new Date().toLocaleDateString()}}</p>
            </div>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" @click="downloadCertificate">下载PDF</el-button>
                <el-button @click="certificateVisible = false">关闭</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
export default {
    name: 'OffboardingFlow',
    data() {
        return {
            activeTab: 'list',
            loading: false,
            dialogVisible: false,
            dialogTitle: '',
            detailVisible: false,
            handoverVisible: false,
            assetsVisible: false,
            salaryVisible: false,
            certificateVisible: false,
            currentStepIndex: 2,
            currentEmployee: '张三',
            allAssetsReturned: false,
            statistics: {
                pending: 5,
                inProgress: 3,
                completed: 12,
                salaryPending: 2
            },
            offboardingList: [
                { id: 1, processNo: 'OFF202404001', employeeName: '陈小明', department: {name: '技术部'}, position: {name: '前端开发'}, lastDate: '2024-04-30', currentStep: '审批中', status: '待审批' },
                { id: 2, processNo: 'OFF202404002', employeeName: '刘小红', department: {name: '产品部'}, position: {name: '产品经理'}, lastDate: '2024-04-25', currentStep: '工作交接', status: '交接中' },
                { id: 3, processNo: 'OFF202404003', employeeName: '王大伟', department: {name: '市场部'}, position: {name: '市场主管'}, lastDate: '2024-04-20', currentStep: '薪资结算', status: '交接中' },
                { id: 4, processNo: 'OFF202403001', employeeName: '赵丽丽', department: {name: '人事部'}, position: {name: 'HR专员'}, lastDate: '2024-03-31', currentStep: '离职完成', status: '已完成' }
            ],
            approvalList: [
                { id: 1, employeeName: '陈小明', department: {name: '技术部'}, reason: '个人职业发展原因，寻求更好的发展机会', applyDate: '2024-04-01', lastDate: '2024-04-30', approver: '张经理' },
                { id: 2, employeeName: '李华', department: {name: '财务部'}, reason: '家庭原因，需要回老家发展', applyDate: '2024-04-05', lastDate: '2024-05-05', approver: '王总监' }
            ],
            employees: [
                { id: 1, name: '张三-技术部' },
                { id: 2, name: '李四-产品部' },
                { id: 3, name: '王五-市场部' }
            ],
            handover: {
                projects: [
                    { name: 'HR系统重构项目', receiver: '李工', completed: true },
                    { name: '移动端APP开发', receiver: '王工', completed: false }
                ],
                documents: [
                    { name: '技术架构文档', receiver: '李工', completed: true },
                    { name: '接口文档', receiver: '王工', completed: true },
                    { name: '数据库设计文档', receiver: '赵工', completed: false }
                ],
                permissions: [
                    { name: 'Gitlab代码权限', handler: 'IT部', completed: true },
                    { name: 'OA系统权限', handler: 'IT部', completed: false },
                    { name: '财务系统权限', handler: '财务部', completed: false }
                ]
            },
            assetsList: [
                { name: '笔记本电脑（MacBook Pro）', code: 'IT-2022-0015', quantity: 1, returned: false },
                { name: '显示器', code: 'IT-2022-0089', quantity: 1, returned: true },
                { name: '键盘鼠标套装', code: 'IT-2023-0156', quantity: 1, returned: false },
                { name: '门禁卡', code: 'AD-2023-0045', quantity: 1, returned: false }
            ],
            salaryDetail: {
                basic: '15,000元',
                performance: '3,000元',
                attendanceDays: '20',
                actualSalary: '12,000元',
                annualLeave: '1,500元（3天）',
                compensation: '0元',
                socialInsurance: '2,250元',
                tax: '580元',
                total: '10,670元'
            },
            offboardingForm: {
                employeeId: null,
                type: '主动离职',
                applyDate: '',
                lastDate: '',
                reason: '',
                managerConfirmed: false
            },
            rules: {
                employeeId: [{ required: true, message: '请选择离职员工', trigger: 'change' }],
                lastDate: [{ required: true, message: '请选择最后工作日', trigger: 'change' }],
                reason: [{ required: true, message: '请填写离职原因', trigger: 'blur' }]
            }
        }
    },
    methods: {
        getStepType(step) {
            const map = {
                '审批中': 'warning',
                '工作交接': 'primary',
                '资产归还': 'info',
                '薪资结算': '',
                '离职完成': 'success'
            }
            return map[step] || ''
        },
        showAddDialog() {
            this.dialogTitle = '发起离职申请'
            this.offboardingForm = {
                employeeId: null,
                type: '主动离职',
                applyDate: '',
                lastDate: '',
                reason: '',
                managerConfirmed: false
            }
            this.dialogVisible = true
        },
        submitOffboarding() {
            this.$refs.offboardingForm.validate(valid => {
                if (valid) {
                    const processNo = 'OFF' + new Date().getFullYear() + (new Date().getMonth() + 1).toString().padStart(2, '0') + (this.offboardingList.length + 1).toString().padStart(3, '0')
                    const emp = this.employees.find(e => e.id === this.offboardingForm.employeeId)
                    this.offboardingList.unshift({
                        id: this.offboardingList.length + 1,
                        processNo,
                        employeeName: emp ? emp.name.split('-')[0] : '',
                        department: {name: emp ? emp.name.split('-')[1] : ''},
                        position: {name: ''},
                        lastDate: this.offboardingForm.lastDate,
                        currentStep: '审批中',
                        status: '待审批'
                    })
                    this.$message.success('离职申请已提交，进入审批流程！')
                    this.dialogVisible = false
                    this.statistics.pending++
                }
            })
        },
        viewDetail(row) {
            const stepMap = { '审批中': 0, '工作交接': 2, '资产归还': 3, '薪资结算': 4, '离职完成': 5 }
            this.currentStepIndex = stepMap[row.currentStep] || 0
            this.detailVisible = true
        },
        viewHandover() {
            this.handoverVisible = true
        },
        viewAssets() {
            this.assetsVisible = true
        },
        viewSalary() {
            this.salaryVisible = true
        },
        generateCertificate(row) {
            this.currentEmployee = row.employeeName
            this.certificateVisible = true
        },
        approve(row) {
            this.$confirm('确认通过该离职申请？', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                const index = this.approvalList.findIndex(a => a.id === row.id)
                if (index > -1) {
                    this.approvalList.splice(index, 1)
                    const flowIndex = this.offboardingList.findIndex(f => f.employeeName === row.employeeName)
                    if (flowIndex > -1) {
                        this.offboardingList[flowIndex].currentStep = '工作交接'
                        this.offboardingList[flowIndex].status = '交接中'
                    }
                    this.statistics.pending--
                    this.statistics.inProgress++
                    this.$message.success('审批已通过，进入交接流程！')
                }
            })
        },
        reject() {
            this.$message.info('已驳回')
        },
        viewApprovalDetail() {
            this.$message.info('查看审批详情')
        },
        checkAssetsComplete() {
            this.allAssetsReturned = this.assetsList.every(a => a.returned)
        },
        confirmAssetsReturn() {
            this.$message.success('资产归还已确认！')
            this.assetsVisible = false
        },
        confirmSalary() {
            this.$message.success('薪资结算已确认！')
            this.salaryVisible = false
        },
        downloadCertificate() {
            this.$message.success('离职证明PDF已生成并下载！')
            this.certificateVisible = false
        }
    }
}
</script>

<style scoped>
.el-statistic {
    text-align: center;
}
h4 {
    margin: 10px 0;
    color: #409EFF;
}
.total-salary {
    font-weight: bold;
}
</style>