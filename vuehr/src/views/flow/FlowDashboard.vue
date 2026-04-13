<template>
    <div>
        <el-card>
            <div slot="header">
                <span style="font-size: 18px; font-weight: bold;">智能入职/离职流程引擎 Dashboard</span>
            </div>

            <el-row :gutter="20" style="margin-bottom: 30px;">
                <el-col :span="6">
                    <el-card shadow="hover">
                        <div style="text-align: center;">
                            <i class="el-icon-user" style="font-size: 40px; color: #67C23A;"></i>
                            <el-statistic title="本月入职人数" :value="15" value-style="color: #67C23A; font-size: 28px;"></el-statistic>
                        </div>
                    </el-card>
                </el-col>
                <el-col :span="6">
                    <el-card shadow="hover">
                        <div style="text-align: center;">
                            <i class="el-icon-user-solid" style="font-size: 40px; color: #F56C6C;"></i>
                            <el-statistic title="本月离职人数" :value="5" value-style="color: #F56C6C; font-size: 28px;"></el-statistic>
                        </div>
                    </el-card>
                </el-col>
                <el-col :span="6">
                    <el-card shadow="hover">
                        <div style="text-align: center;">
                            <i class="el-icon-date" style="font-size: 40px; color: #409EFF;"></i>
                            <el-statistic title="待处理流程" :value="18" value-style="color: #409EFF; font-size: 28px;"></el-statistic>
                        </div>
                    </el-card>
                </el-col>
                <el-col :span="6">
                    <el-card shadow="hover">
                        <div style="text-align: center;">
                            <i class="el-icon-time" style="font-size: 40px; color: #E6A23C;"></i>
                            <el-statistic title="平均完成周期" value="3.2" value-style="color: #E6A23C; font-size: 28px;">
                                <template slot="suffix">天</template>
                            </el-statistic>
                        </div>
                    </el-card>
                </el-col>
            </el-row>

            <el-row :gutter="20">
                <el-col :span="12">
                    <el-card shadow="hover">
                        <div slot="header">
                            <i class="el-icon-data-line"></i> 近6个月入职/离职趋势
                        </div>
                        <div style="height: 300px; display: flex; align-items: flex-end; justify-content: space-around; padding: 0 20px;">
                            <div v-for="(item, index) in trendData" :key="index" style="text-align: center; width: 60px;">
                                <div style="display: flex; flex-direction: column; height: 200px; justify-content: flex-end;">
                                    <div :style="{height: item.onboarding * 8 + 'px', background: '#67C23A', marginBottom: '5px'}">
                                        <span style="color: #fff; font-size: 12px;">{{item.onboarding}}</span>
                                    </div>
                                    <div :style="{height: item.offboarding * 8 + 'px', background: '#F56C6C'}">
                                        <span style="color: #fff; font-size: 12px;">{{item.offboarding}}</span>
                                    </div>
                                </div>
                                <div style="margin-top: 10px; font-size: 12px;">{{item.month}}</div>
                            </div>
                        </div>
                        <div style="text-align: center; margin-top: 10px;">
                            <el-tag type="success">入职</el-tag>
                            <el-tag type="danger" style="margin-left: 20px;">离职</el-tag>
                        </div>
                    </el-card>
                </el-col>
                <el-col :span="12">
                    <el-card shadow="hover">
                        <div slot="header">
                            <i class="el-icon-pie-chart"></i> 各部门待办任务统计
                        </div>
                        <el-table :data="deptTaskData" border>
                            <el-table-column prop="dept" label="部门" width="120"></el-table-column>
                            <el-table-column label="入职待办" align="center">
                                <template slot-scope="scope">
                                    <el-progress :percentage="scope.row.onboarding" :stroke-width="12"></el-progress>
                                </template>
                            </el-table-column>
                            <el-table-column label="离职待办" align="center">
                                <template slot-scope="scope">
                                    <el-progress :percentage="scope.row.offboarding" :stroke-width="12" status="exception"></el-progress>
                                </template>
                            </el-table-column>
                            <el-table-column label="待处理" width="80" align="center">
                                <template slot-scope="scope">
                                    <el-badge :value="scope.row.pending" class="item"></el-badge>
                                </template>
                            </el-table-column>
                        </el-table>
                    </el-card>
                </el-col>
            </el-row>

            <el-row :gutter="20" style="margin-top: 20px;">
                <el-col :span="12">
                    <el-card shadow="hover">
                        <div slot="header">
                            <i class="el-icon-user-add"></i> 近期入职流程
                            <el-button style="float: right; padding: 3px 0" type="text" @click="$router.push('/onboarding')">查看全部</el-button>
                        </div>
                        <el-table :data="recentOnboarding" size="small">
                            <el-table-column prop="employeeName" label="姓名" width="80"></el-table-column>
                            <el-table-column prop="department" label="部门" width="100"></el-table-column>
                            <el-table-column prop="currentStep" label="当前进度">
                                <template slot-scope="scope">
                                    <el-tag size="mini" :type="getStepType(scope.row.currentStep)">{{scope.row.currentStep}}</el-tag>
                                </template>
                            </el-table-column>
                            <el-table-column prop="expectedDate" label="预计入职" width="100"></el-table-column>
                        </el-table>
                    </el-card>
                </el-col>
                <el-col :span="12">
                    <el-card shadow="hover">
                        <div slot="header">
                            <i class="el-icon-user-delete"></i> 近期离职流程
                            <el-button style="float: right; padding: 3px 0" type="text" @click="$router.push('/offboarding')">查看全部</el-button>
                        </div>
                        <el-table :data="recentOffboarding" size="small">
                            <el-table-column prop="employeeName" label="姓名" width="80"></el-table-column>
                            <el-table-column prop="department" label="部门" width="100"></el-table-column>
                            <el-table-column prop="currentStep" label="当前进度">
                                <template slot-scope="scope">
                                    <el-tag size="mini" :type="getStepType(scope.row.currentStep)">{{scope.row.currentStep}}</el-tag>
                                </template>
                            </el-table-column>
                            <el-table-column prop="lastDate" label="离职日期" width="100"></el-table-column>
                        </el-table>
                    </el-card>
                </el-col>
            </el-row>

            <el-row style="margin-top: 20px;">
                <el-col :span="24">
                    <el-card shadow="hover">
                        <div slot="header">
                            <i class="el-icon-bell"></i> 待办事项提醒 <el-badge :value="totalPending" type="danger" class="item"></el-badge>
                        </div>
                        <el-alert
                            title="自动化流程驱动，取代Excel和口头沟通"
                            description="系统已自动推送任务到各部门负责人，待办事项实时同步。标准化流程，新人入职体验大幅提升，离职交接风险降低90%"
                            type="success"
                            :closable="false"
                            show-icon>
                        </el-alert>
                        <el-timeline style="margin-top: 20px;">
                            <el-timeline-item
                                v-for="(activity, index) in activities"
                                :key="index"
                                :timestamp="activity.time"
                                placement="top"
                                :color="activity.color">
                                {{activity.content}}
                            </el-timeline-item>
                        </el-timeline>
                    </el-card>
                </el-col>
            </el-row>
        </el-card>
    </div>
</template>

<script>
export default {
    name: 'FlowDashboard',
    data() {
        return {
            totalPending: 18,
            trendData: [
                { month: '11月', onboarding: 12, offboarding: 3 },
                { month: '12月', onboarding: 18, offboarding: 5 },
                { month: '1月', onboarding: 8, offboarding: 2 },
                { month: '2月', onboarding: 5, offboarding: 1 },
                { month: '3月', onboarding: 20, offboarding: 4 },
                { month: '4月', onboarding: 15, offboarding: 5 }
            ],
            deptTaskData: [
                { dept: 'IT部门', onboarding: 75, offboarding: 60, pending: 5 },
                { dept: '行政部门', onboarding: 85, offboarding: 40, pending: 3 },
                { dept: '部门导师', onboarding: 60, offboarding: 30, pending: 4 },
                { dept: 'HR部门', onboarding: 95, offboarding: 80, pending: 6 }
            ],
            recentOnboarding: [
                { employeeName: '张三', department: '技术部', currentStep: 'IT处理', expectedDate: '2024-04-15' },
                { employeeName: '李四', department: '产品部', currentStep: '行政处理', expectedDate: '2024-04-18' },
                { employeeName: '王五', department: '市场部', currentStep: '预约登记', expectedDate: '2024-04-20' }
            ],
            recentOffboarding: [
                { employeeName: '陈小明', department: '技术部', currentStep: '工作交接', lastDate: '2024-04-30' },
                { employeeName: '刘小红', department: '产品部', currentStep: '资产归还', lastDate: '2024-04-25' },
                { employeeName: '王大伟', department: '市场部', currentStep: '薪资结算', lastDate: '2024-04-20' }
            ],
            activities: [
                {
                    content: '【IT待办】张三 - 配置办公电脑',
                    time: '5分钟前',
                    color: '#409EFF'
                },
                {
                    content: '【行政待办】李四 - 办理门禁卡',
                    time: '10分钟前',
                    color: '#E6A23C'
                },
                {
                    content: '【导师待办】张三 - 入职引导',
                    time: '30分钟前',
                    color: '#67C23A'
                },
                {
                    content: '赵六 - 入职流程全部完成，Checklist自动确认',
                    time: '1小时前',
                    color: '#909399'
                },
                {
                    content: '赵丽丽 - 离职证明已自动生成下载',
                    time: '2小时前',
                    color: '#909399'
                }
            ]
        }
    },
    methods: {
        getStepType(step) {
            const map = {
                '预约登记': 'warning',
                'IT处理': 'primary',
                '行政处理': 'info',
                '工作交接': 'primary',
                '资产归还': 'info',
                '薪资结算': 'warning'
            }
            return map[step] || ''
        }
    }
}
</script>

<style scoped>
.el-card {
    margin-bottom: 0;
}
</style>