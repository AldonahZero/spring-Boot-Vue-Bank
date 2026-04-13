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
                    <el-button type="primary" size="mini" style="margin-left: 5px" icon="el-icon-search" @click="searchOnboarding">搜索</el-button>
                    <el-button type="primary" size="mini" style="margin-left: 5px" icon="el-icon-plus" @click="showAddDialog">新建入职</el-button>
                </div>
                <div>
                    <el-tag type="info">总入职: {{statistics.total}}</el-tag>
                    <el-tag type="warning">待处理: {{statistics.pending}}</el-tag>
                    <el-tag type="primary">进行中: {{statistics.inProgress}}</el-tag>
                    <el-tag type="success">已完成: {{statistics.completed}}</el-tag>
                </div>
            </el-header>
            <el-main style="padding-left: 0px;padding-top: 0px">
                <div>
                    <el-table
                            :data="onboardings"
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
                                prop="email"
                                label="企业邮箱"
                                width="180">
                        </el-table-column>
                        <el-table-column
                                prop="position"
                                label="职位"
                                width="120">
                        </el-table-column>
                        <el-table-column
                                prop="entryDate"
                                label="入职日期"
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
                                width="180">
                            <template slot-scope="scope">
                                <el-button size="mini" @click="showDetail(scope.row)">详情</el-button>
                                <el-button size="mini" type="primary" @click="showTaskDialog(scope.row)">任务</el-button>
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

        <!-- 新建入职对话框 -->
        <el-dialog title="新建入职流程" :visible.sync="addDialogVisible" width="50%">
            <el-form :model="newOnboarding" label-width="100px">
                <el-form-item label="员工姓名">
                    <el-input v-model="newOnboarding.employeeName" placeholder="请输入员工姓名"></el-input>
                </el-form-item>
                <el-form-item label="联系电话">
                    <el-input v-model="newOnboarding.phone" placeholder="请输入联系电话"></el-input>
                </el-form-item>
                <el-form-item label="部门">
                    <el-select v-model="newOnboarding.departmentId" placeholder="请选择部门">
                        <el-option label="技术部" value="tech"></el-option>
                        <el-option label="人事部" value="hr"></el-option>
                        <el-option label="财务部" value="finance"></el-option>
                        <el-option label="行政部" value="admin"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="职位">
                    <el-input v-model="newOnboarding.position" placeholder="请输入职位"></el-input>
                </el-form-item>
                <el-form-item label="入职日期">
                    <el-date-picker v-model="newOnboarding.entryDate" type="date" placeholder="选择日期" value-format="yyyy-MM-dd"></el-date-picker>
                </el-form-item>
                <el-form-item label="导师">
                    <el-select v-model="newOnboarding.mentorId" placeholder="请选择导师">
                        <el-option v-for="user in users" :key="user.ID" :label="user.nickName" :value="user.ID"></el-option>
                    </el-select>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="addDialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="addOnboarding">确 定</el-button>
            </div>
        </el-dialog>

        <!-- 任务管理对话框 -->
        <el-dialog title="入职任务管理" :visible.sync="taskDialogVisible" width="60%">
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
        </el-dialog>

        <!-- 详情对话框 -->
        <el-dialog title="入职详情" :visible.sync="detailDialogVisible" width="50%">
            <el-descriptions :column="2" border v-if="currentOnboarding">
                <el-descriptions-item label="员工姓名">{{currentOnboarding.employeeName}}</el-descriptions-item>
                <el-descriptions-item label="工号">{{currentOnboarding.employeeNo}}</el-descriptions-item>
                <el-descriptions-item label="企业邮箱">{{currentOnboarding.email}}</el-descriptions-item>
                <el-descriptions-item label="联系电话">{{currentOnboarding.phone}}</el-descriptions-item>
                <el-descriptions-item label="部门">{{currentOnboarding.departmentId}}</el-descriptions-item>
                <el-descriptions-item label="职位">{{currentOnboarding.position}}</el-descriptions-item>
                <el-descriptions-item label="入职日期">{{currentOnboarding.entryDate}}</el-descriptions-item>
                <el-descriptions-item label="当前进度">{{currentOnboarding.progress}}%</el-descriptions-item>
            </el-descriptions>
        </el-dialog>
    </div>
</template>

<script>
    import {getRequest, postRequest, putRequest} from '../../utils/api'

    export default {
        name: "PerOnboarding",
        data() {
            return {
                keywords: '',
                onboardings: [],
                tableLoading: false,
                pageSize: 10,
                totalCount: 0,
                currentPage: 1,
                statistics: {
                    total: 0,
                    pending: 0,
                    inProgress: 0,
                    completed: 0
                },
                addDialogVisible: false,
                taskDialogVisible: false,
                detailDialogVisible: false,
                newOnboarding: {
                    employeeName: '',
                    phone: '',
                    departmentId: '',
                    position: '',
                    entryDate: '',
                    mentorId: null
                },
                currentOnboarding: null,
                currentTasks: [],
                users: []
            }
        },
        mounted() {
            this.loadOnboardings();
            this.loadStatistics();
        },
        methods: {
            loadOnboardings() {
                this.tableLoading = true;
                getRequest('/onboarding/list?page=' + this.currentPage + '&pageSize=' + this.pageSize).then(resp => {
                    this.tableLoading = false;
                    if (resp && resp.data) {
                        this.onboardings = resp.data.list;
                        this.totalCount = resp.data.total;
                    }
                });
            },
            loadStatistics() {
                getRequest('/onboarding/statistics').then(resp => {
                    if (resp && resp.data) {
                        this.statistics = resp.data;
                    }
                });
            },
            searchOnboarding() {
                this.currentPage = 1;
                this.loadOnboardings();
            },
            currentChange(val) {
                this.currentPage = val;
                this.loadOnboardings();
            },
            showAddDialog() {
                this.newOnboarding = {
                    employeeName: '',
                    phone: '',
                    departmentId: '',
                    position: '',
                    entryDate: '',
                    mentorId: null
                };
                this.addDialogVisible = true;
            },
            addOnboarding() {
                postRequest('/onboarding/create', this.newOnboarding).then(resp => {
                    if (resp && resp.status === 200) {
                        this.$message.success('创建成功');
                        this.addDialogVisible = false;
                        this.loadOnboardings();
                        this.loadStatistics();
                    }
                });
            },
            showDetail(row) {
                this.currentOnboarding = row;
                this.detailDialogVisible = true;
            },
            showTaskDialog(row) {
                this.currentOnboarding = row;
                this.currentTasks = row.onboardingTasks || [];
                this.taskDialogVisible = true;
            },
            completeTask(task) {
                putRequest('/onboarding/task/update?taskId=' + task.ID + '&status=2').then(resp => {
                    if (resp && resp.status === 200) {
                        this.$message.success('任务完成');
                        task.status = 2;
                        this.loadOnboardings();
                        this.loadStatistics();
                    }
                });
            },
            getStatusType(status) {
                const types = ['warning', 'primary', 'success', 'info'];
                return types[status] || 'info';
            },
            getStatusText(status) {
                const texts = ['待处理', '进行中', '已完成', '已取消'];
                return texts[status] || '未知';
            },
            getProgressStatus(status) {
                if (status === 2) return 'success';
                if (status === 1) return '';
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
                    'IT': 'IT部门',
                    'ADMIN': '行政部',
                    'MENTOR': '导师'
                };
                return texts[type] || type;
            },
            handleSelectionChange(val) {
                // 多选处理
            }
        }
    }
</script>

<style scoped>
</style>
