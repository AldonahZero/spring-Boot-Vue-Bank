<template>
    <div>
        <el-container>
            <el-header style="padding: 0px;display:flex;justify-content:space-between;align-items: center">
                <div style="display: inline">
                    <el-input
                            placeholder="通过员工姓名搜索..."
                            clearable
                            style="width: 300px;margin: 0px;padding: 0px;"
                            size="mini"
                            prefix-icon="el-icon-search"
                            v-model="keywords">
                    </el-input>
                    <el-button type="primary" size="mini" style="margin-left: 5px" icon="el-icon-search" @click="searchOffboarding">搜索</el-button>
                    <el-button type="primary" size="mini" style="margin-left: 5px" icon="el-icon-plus" @click="showAddDialog">新建离职</el-button>
                </div>
                <div>
                    <el-tag type="info">总离职: {{statistics.total}}</el-tag>
                    <el-tag type="warning">待审批: {{statistics.pending}}</el-tag>
                    <el-tag type="primary">进行中: {{statistics.inProgress}}</el-tag>
                    <el-tag type="success">已完成: {{statistics.completed}}</el-tag>
                </div>
            </el-header>
            <el-main style="padding-left: 0px;padding-top: 0px">
                <div>
                    <el-table
                            :data="offboardings"
                            v-loading="tableLoading"
                            border
                            stripe
                            @selection-change="handleSelectionChange"
                            size="mini"
                            style="width: 100%">
                        <el-table-column
                                type="selection"
                                width="30">
                        </el-table-column>
                        <el-table-column
                                prop="employeeName"
                                label="姓名"
                                width="100">
                        </el-table-column>
                        <el-table-column
                                prop="employeeNo"
                                label="工号"
                                width="120">
                        </el-table-column>
                        <el-table-column
                                prop="departmentId"
                                label="部门"
                                width="100">
                        </el-table-column>
                        <el-table-column
                                prop="position"
                                label="职位"
                                width="120">
                        </el-table-column>
                        <el-table-column
                                prop="lastWorkingDate"
                                label="最后工作日"
                                width="120">
                        </el-table-column>
                        <el-table-column
                                label="进度"
                                width="200">
                            <template slot-scope="scope">
                                <el-progress :percentage="scope.row.progress" :status="getProgressStatus(scope.row.status)"></el-progress>
                            </template>
                        </el-table-column>
                        <el-table-column
                                label="状态"
                                width="100">
                            <template slot-scope="scope">
                                <el-tag :type="getStatusType(scope.row.status)">{{getStatusText(scope.row.status)}}</el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column
                                label="操作"
                                width="250">
                            <template slot-scope="scope">
                                <el-button size="mini" @click="showDetail(scope.row)">详情</el-button>
                                <el-button size="mini" type="primary" @click="showTaskDialog(scope.row)">任务</el-button>
                                <el-button v-if="scope.row.status === 0" size="mini" type="success" @click="showApproveDialog(scope.row)">审批</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                    <div style="display: flex;justify-content: space-between;margin: 2px">
                        <el-pagination
                                background
                                :page-size="pageSize"
                                :current-page="currentPage"
                                @current-change="currentChange"
                                layout="prev, pager, next"
                                :total="totalCount">
                        </el-pagination>
                    </div>
                </div>
            </el-main>
        </el-container>

        <!-- 新建离职对话框 -->
        <el-dialog title="新建离职流程" :visible.sync="addDialogVisible" width="50%">
            <el-form :model="newOffboarding" label-width="100px">
                <el-form-item label="员工姓名">
                    <el-input v-model="newOffboarding.employeeName" placeholder="请输入员工姓名"></el-input>
                </el-form-item>
                <el-form-item label="工号">
                    <el-input v-model="newOffboarding.employeeNo" placeholder="请输入工号"></el-input>
                </el-form-item>
                <el-form-item label="部门">
                    <el-select v-model="newOffboarding.departmentId" placeholder="请选择部门">
                        <el-option label="技术部" value="tech"></el-option>
                        <el-option label="人事部" value="hr"></el-option>
                        <el-option label="财务部" value="finance"></el-option>
                        <el-option label="行政部" value="admin"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="职位">
                    <el-input v-model="newOffboarding.position" placeholder="请输入职位"></el-input>
                </el-form-item>
                <el-form-item label="申请离职日期">
                    <el-date-picker v-model="newOffboarding.resignationDate" type="date" placeholder="选择日期" value-format="yyyy-MM-dd"></el-date-picker>
                </el-form-item>
                <el-form-item label="最后工作日">
                    <el-date-picker v-model="newOffboarding.lastWorkingDate" type="date" placeholder="选择日期" value-format="yyyy-MM-dd"></el-date-picker>
                </el-form-item>
                <el-form-item label="离职原因">
                    <el-input type="textarea" v-model="newOffboarding.resignationReason" placeholder="请输入离职原因"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="addDialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="addOffboarding">确 定</el-button>
            </div>
        </el-dialog>

        <!-- 审批对话框 -->
        <el-dialog title="离职审批" :visible.sync="approveDialogVisible" width="40%">
            <el-form :model="approveForm" label-width="100px">
                <el-form-item label="审批结果">
                    <el-radio-group v-model="approveForm.approved">
                        <el-radio :label="true">批准</el-radio>
                        <el-radio :label="false">拒绝</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="审批备注">
                    <el-input type="textarea" v-model="approveForm.remarks" placeholder="请输入审批备注"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="approveDialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="submitApprove">确 定</el-button>
            </div>
        </el-dialog>

        <!-- 任务管理对话框 -->
        <el-dialog title="离职任务管理" :visible.sync="taskDialogVisible" width="70%">
            <el-tabs v-model="activeTab">
                <el-tab-pane label="任务进度" name="tasks">
                    <el-timeline>
                        <el-timeline-item
                                v-for="(task, index) in currentTasks"
                                :key="index"
                                :type="getTaskType(task.status)"
                                :icon="getTaskIcon(task.status)"
                                :timestamp="task.createdAt">
                            <el-card>
                                <h4>{{task.taskName}}</h4>
                                <p>{{task.description}}</p>
                                <p>负责部门: {{getTaskTypeText(task.taskType)}}</p>
                                <el-button v-if="task.status !== 2" size="mini" type="primary" @click="completeTask(task)">完成任务</el-button>
                                <el-tag v-else type="success">已完成</el-tag>
                            </el-card>
                        </el-timeline-item>
                    </el-timeline>
                </el-tab-pane>
                <el-tab-pane label="工作交接" name="handover">
                    <el-button size="mini" type="primary" @click="showAddHandoverDialog">添加交接事项</el-button>
                    <el-table :data="handoverItems" size="mini" style="margin-top: 10px">
                        <el-table-column prop="itemName" label="交接事项"></el-table-column>
                        <el-table-column prop="itemType" label="类型">
                            <template slot-scope="scope">
                                {{getHandoverTypeText(scope.row.itemType)}}
                            </template>
                        </el-table-column>
                        <el-table-column prop="receiverName" label="接收人"></el-table-column>
                        <el-table-column prop="status" label="状态">
                            <template slot-scope="scope">
                                <el-tag :type="scope.row.status === 1 ? 'success' : 'warning'">
                                    {{scope.row.status === 1 ? '已交接' : '待交接'}}
                                </el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column label="操作">
                            <template slot-scope="scope">
                                <el-button v-if="scope.row.status === 0" size="mini" type="primary" @click="completeHandover(scope.row)">确认交接</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-tab-pane>
                <el-tab-pane label="资产归还" name="assets">
                    <el-button size="mini" type="primary" @click="showAddAssetDialog">添加资产</el-button>
                    <el-table :data="assetReturns" size="mini" style="margin-top: 10px">
                        <el-table-column prop="assetName" label="资产名称"></el-table-column>
                        <el-table-column prop="assetNo" label="资产编号"></el-table-column>
                        <el-table-column prop="status" label="状态">
                            <template slot-scope="scope">
                                <el-tag :type="getAssetStatusType(scope.row.status)">
                                    {{getAssetStatusText(scope.row.status)}}
                                </el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column label="操作">
                            <template slot-scope="scope">
                                <el-button v-if="scope.row.status === 0" size="mini" type="primary" @click="completeAssetReturn(scope.row)">确认归还</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-tab-pane>
            </el-tabs>
        </el-dialog>

        <!-- 添加交接事项对话框 -->
        <el-dialog title="添加交接事项" :visible.sync="addHandoverDialogVisible" width="40%">
            <el-form :model="newHandover" label-width="100px">
                <el-form-item label="交接类型">
                    <el-select v-model="newHandover.itemType" placeholder="请选择类型">
                        <el-option label="项目" value="PROJECT"></el-option>
                        <el-option label="文档" value="DOCUMENT"></el-option>
                        <el-option label="账号" value="ACCOUNT"></el-option>
                        <el-option label="其他" value="OTHER"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="交接事项">
                    <el-input v-model="newHandover.itemName" placeholder="请输入交接事项名称"></el-input>
                </el-form-item>
                <el-form-item label="描述">
                    <el-input type="textarea" v-model="newHandover.description" placeholder="请输入描述"></el-input>
                </el-form-item>
                <el-form-item label="接收人">
                    <el-input v-model="newHandover.receiverName" placeholder="请输入接收人姓名"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="addHandoverDialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="addHandover">确 定</el-button>
            </div>
        </el-dialog>

        <!-- 添加资产对话框 -->
        <el-dialog title="添加资产归还" :visible.sync="addAssetDialogVisible" width="40%">
            <el-form :model="newAsset" label-width="100px">
                <el-form-item label="资产类型">
                    <el-select v-model="newAsset.assetType" placeholder="请选择类型">
                        <el-option label="电脑" value="COMPUTER"></el-option>
                        <el-option label="门禁卡" value="ACCESS_CARD"></el-option>
                        <el-option label="其他" value="OTHER"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="资产名称">
                    <el-input v-model="newAsset.assetName" placeholder="请输入资产名称"></el-input>
                </el-form-item>
                <el-form-item label="资产编号">
                    <el-input v-model="newAsset.assetNo" placeholder="请输入资产编号"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="addAssetDialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="addAsset">确 定</el-button>
            </div>
        </el-dialog>

        <!-- 详情对话框 -->
        <el-dialog title="离职详情" :visible.sync="detailDialogVisible" width="50%">
            <el-descriptions :column="2" border v-if="currentOffboarding">
                <el-descriptions-item label="员工姓名">{{currentOffboarding.employeeName}}</el-descriptions-item>
                <el-descriptions-item label="工号">{{currentOffboarding.employeeNo}}</el-descriptions-item>
                <el-descriptions-item label="部门">{{currentOffboarding.departmentId}}</el-descriptions-item>
                <el-descriptions-item label="职位">{{currentOffboarding.position}}</el-descriptions-item>
                <el-descriptions-item label="申请日期">{{currentOffboarding.resignationDate}}</el-descriptions-item>
                <el-descriptions-item label="最后工作日">{{currentOffboarding.lastWorkingDate}}</el-descriptions-item>
                <el-descriptions-item label="离职原因" :span="2">{{currentOffboarding.resignationReason}}</el-descriptions-item>
                <el-descriptions-item label="当前进度">{{currentOffboarding.progress}}%</el-descriptions-item>
                <el-descriptions-item label="审批备注">{{currentOffboarding.approvalRemarks}}</el-descriptions-item>
            </el-descriptions>
            <div style="margin-top: 20px">
                <el-button type="primary" @click="generateCertificate">生成离职证明</el-button>
            </div>
            <el-input v-if="certificateContent" type="textarea" :rows="10" v-model="certificateContent" style="margin-top: 10px" readonly></el-input>
        </el-dialog>
    </div>
</template>

<script>
    import {getRequest, postRequest, putRequest} from '../../utils/api'

    export default {
        name: "PerOffboarding",
        data() {
            return {
                keywords: '',
                offboardings: [],
                tableLoading: false,
                pageSize: 10,
                totalCount: 0,
                currentPage: 1,
                statistics: {
                    total: 0,
                    pending: 0,
                    inProgress: 0,
                    completed: 0,
                    rejected: 0
                },
                addDialogVisible: false,
                approveDialogVisible: false,
                taskDialogVisible: false,
                detailDialogVisible: false,
                addHandoverDialogVisible: false,
                addAssetDialogVisible: false,
                activeTab: 'tasks',
                newOffboarding: {
                    employeeName: '',
                    employeeNo: '',
                    departmentId: '',
                    position: '',
                    resignationDate: '',
                    lastWorkingDate: '',
                    resignationReason: ''
                },
                approveForm: {
                    approved: true,
                    remarks: ''
                },
                currentOffboarding: null,
                currentTasks: [],
                handoverItems: [],
                assetReturns: [],
                newHandover: {
                    itemType: '',
                    itemName: '',
                    description: '',
                    receiverName: ''
                },
                newAsset: {
                    assetType: '',
                    assetName: '',
                    assetNo: ''
                },
                certificateContent: ''
            }
        },
        mounted() {
            this.loadOffboardings();
            this.loadStatistics();
        },
        methods: {
            loadOffboardings() {
                this.tableLoading = true;
                getRequest('/offboarding/list?page=' + this.currentPage + '&pageSize=' + this.pageSize).then(resp => {
                    this.tableLoading = false;
                    if (resp && resp.data) {
                        this.offboardings = resp.data.list;
                        this.totalCount = resp.data.total;
                    }
                });
            },
            loadStatistics() {
                getRequest('/offboarding/statistics').then(resp => {
                    if (resp && resp.data) {
                        this.statistics = resp.data;
                    }
                });
            },
            searchOffboarding() {
                this.currentPage = 1;
                this.loadOffboardings();
            },
            currentChange(val) {
                this.currentPage = val;
                this.loadOffboardings();
            },
            showAddDialog() {
                this.newOffboarding = {
                    employeeName: '',
                    employeeNo: '',
                    departmentId: '',
                    position: '',
                    resignationDate: '',
                    lastWorkingDate: '',
                    resignationReason: ''
                };
                this.addDialogVisible = true;
            },
            addOffboarding() {
                postRequest('/offboarding/create', this.newOffboarding).then(resp => {
                    if (resp && resp.status === 200) {
                        this.$message.success('创建成功');
                        this.addDialogVisible = false;
                        this.loadOffboardings();
                        this.loadStatistics();
                    }
                });
            },
            showApproveDialog(row) {
                this.currentOffboarding = row;
                this.approveForm = {
                    approved: true,
                    remarks: ''
                };
                this.approveDialogVisible = true;
            },
            submitApprove() {
                putRequest('/offboarding/approve?id=' + this.currentOffboarding.ID + '&approved=' + this.approveForm.approved + '&remarks=' + this.approveForm.remarks).then(resp => {
                    if (resp && resp.status === 200) {
                        this.$message.success('审批完成');
                        this.approveDialogVisible = false;
                        this.loadOffboardings();
                        this.loadStatistics();
                    }
                });
            },
            showDetail(row) {
                this.currentOffboarding = row;
                this.certificateContent = '';
                this.detailDialogVisible = true;
            },
            showTaskDialog(row) {
                this.currentOffboarding = row;
                this.currentTasks = row.offboardingTasks || [];
                this.handoverItems = row.handoverItems || [];
                this.assetReturns = row.assetReturns || [];
                this.activeTab = 'tasks';
                this.taskDialogVisible = true;
            },
            completeTask(task) {
                putRequest('/offboarding/task/update?taskId=' + task.ID + '&status=2').then(resp => {
                    if (resp && resp.status === 200) {
                        this.$message.success('任务完成');
                        task.status = 2;
                        this.loadOffboardings();
                        this.loadStatistics();
                    }
                });
            },
            showAddHandoverDialog() {
                this.newHandover = {
                    itemType: '',
                    itemName: '',
                    description: '',
                    receiverName: ''
                };
                this.addHandoverDialogVisible = true;
            },
            addHandover() {
                this.newHandover.offboardingId = this.currentOffboarding.ID;
                postRequest('/offboarding/handover/add', this.newHandover).then(resp => {
                    if (resp && resp.status === 200) {
                        this.$message.success('添加成功');
                        this.addHandoverDialogVisible = false;
                        this.handoverItems.push({...this.newHandover, status: 0});
                    }
                });
            },
            completeHandover(item) {
                putRequest('/offboarding/handover/update?id=' + item.ID + '&status=1').then(resp => {
                    if (resp && resp.status === 200) {
                        this.$message.success('交接完成');
                        item.status = 1;
                    }
                });
            },
            showAddAssetDialog() {
                this.newAsset = {
                    assetType: '',
                    assetName: '',
                    assetNo: ''
                };
                this.addAssetDialogVisible = true;
            },
            addAsset() {
                this.newAsset.offboardingId = this.currentOffboarding.ID;
                postRequest('/offboarding/asset/add', this.newAsset).then(resp => {
                    if (resp && resp.status === 200) {
                        this.$message.success('添加成功');
                        this.addAssetDialogVisible = false;
                        this.assetReturns.push({...this.newAsset, status: 0});
                    }
                });
            },
            completeAssetReturn(asset) {
                putRequest('/offboarding/asset/update?id=' + asset.ID + '&status=1').then(resp => {
                    if (resp && resp.status === 200) {
                        this.$message.success('归还确认');
                        asset.status = 1;
                    }
                });
            },
            generateCertificate() {
                getRequest('/offboarding/certificate?id=' + this.currentOffboarding.ID).then(resp => {
                    if (resp && resp.data) {
                        this.certificateContent = resp.data.certificate;
                    }
                });
            },
            getStatusType(status) {
                const types = ['warning', 'primary', 'primary', 'success', 'danger'];
                return types[status] || 'info';
            },
            getStatusText(status) {
                const texts = ['待审批', '审批中', '交接中', '已完成', '已拒绝'];
                return texts[status] || '未知';
            },
            getProgressStatus(status) {
                if (status === 3) return 'success';
                if (status === 4) return 'exception';
                return '';
            },
            getTaskType(status) {
                return status === 2 ? 'success' : status === 1 ? 'primary' : 'warning';
            },
            getTaskIcon(status) {
                return status === 2 ? 'el-icon-check' : status === 1 ? 'el-icon-loading' : 'el-icon-time';
            },
            getTaskTypeText(type) {
                const texts = {
                    'HR': '人事部',
                    'HANDOVER': '工作交接',
                    'ASSET': '资产归还',
                    'PERMISSION': '权限回收',
                    'FINANCE': '财务结算'
                };
                return texts[type] || type;
            },
            getHandoverTypeText(type) {
                const texts = {
                    'PROJECT': '项目',
                    'DOCUMENT': '文档',
                    'ACCOUNT': '账号',
                    'OTHER': '其他'
                };
                return texts[type] || type;
            },
            getAssetStatusType(status) {
                return status === 1 ? 'success' : status === 2 ? 'danger' : 'warning';
            },
            getAssetStatusText(status) {
                return status === 1 ? '已归还' : status === 2 ? '损坏/丢失' : '待归还';
            },
            handleSelectionChange(val) {
                // 多选处理
            }
        }
    }
</script>

<style scoped>
</style>
