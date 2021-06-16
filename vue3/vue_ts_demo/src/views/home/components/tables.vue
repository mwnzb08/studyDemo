<template>
    <!--            stripe -- 显示斑马线 -->
    <!--            show-summary -- 顯示页尾统计 -->
    <el-card>
        <template #header>Table</template>
    <el-table
            :data="tableData"
            height="300"
            style="overflow-x: hidden"
            :fit="false"
            max-height="300"
            :span-method="arraySpanMethod"
            highlight-current-row
            @current-change="handleCurrentChange"
            :border="true"
            :row-class-name="theHeightLine"
            :default-sort="{prop: 'name', order: 'ascending'}"
            size="small"
            :row-style="{height: '15px', 'overflow-x': 'hidden'}"
            :cell-style="{padding: '10px'}"
    >
        <el-table-column
                fixed
                type="selection">
        </el-table-column>
        <el-table-column
                type="index"
                fixed
                label="id">
        </el-table-column>
        <el-table-column
                fixed
                :sortable="true"
                :filters="[{text: 'mowei', value: 'mowei'}, {text: 'aa', value: 'aa'}]"
                :filter-method="sortName"
                label="Name"
                prop="name">
        </el-table-column>
        <el-table-column
                label="FullName"
        >
            <el-table-column
                    label="LastName"
                    prop="name"
            >
            </el-table-column>
            <el-table-column
                    label="FirstName"
                    prop="name"
            >
            </el-table-column>

        </el-table-column>
        <el-table-column
                label="pwd"
                prop="pwd">
        </el-table-column>
        <el-table-column
                label="remark"
                prop="remark">
        </el-table-column>
        <el-table-column
                label="remark"
                prop="remark">
        </el-table-column>
        <el-table-column
                label="remark"
                prop="remark">
        </el-table-column>
        <el-table-column
                label="remark"
                prop="remark">
        </el-table-column>
        <el-table-column
                label="remark"
                prop="remark">
        </el-table-column>
    </el-table>
    <div>
        <el-pagination
                small
                layout="prev, pager, next"
                :total="500"
                :page-size="10"
                :pager-count="10"
        >
        </el-pagination>
    </div>
    </el-card>
</template>

<script>
    import {ref, defineComponent} from "vue";
    export default defineComponent({
        name: "tables",
        data () {
            return {
                tableData:[
                    {name:'aa', pwd:123456, remark: 'guangxi'},
                    {name:'mowei', pwd:123456, remark: 'guangxi'},
                    {name:'mowei', pwd:123456, remark: 'guangxi'},
                    {name:'mowei', pwd:123456, remark: 'guangxi'},
                    {name:'mowei', pwd:123456, remark: 'guangxi'},
                    {name:'mowei', pwd:123456, remark: 'guangxi'},
                    {name:'mowei', pwd:123456, remark: 'guangxi'},
                    {name:'mowei', pwd:123456, remark: 'guangxi'},
                    {name:'mowei', pwd:123456, remark: 'guangxi'},
                    {name:'mowei', pwd:123456, remark: 'guangxi'},
                    {name:'mowei', pwd:123456, remark: 'guangxi'}
                ]
            }
        },
        setup () {
            const hello = ref("my name is mowei")
            return {
                hello
            }
        },
        methods: {
            theHeightLine ({row, rowIndex}) {
                if (rowIndex %2 === 0) {
                    return 'warning-row'
                } else {
                    return 'success-row'
                }
            },
            handleCurrentChange (oldVal, newVal) {
                alert(JSON.stringify(oldVal) + JSON.stringify(newVal))
            },
            sortName (value, row, col) {
                const property = col['property'];
                return row[property] === value;
            }
            // arraySpanMethod () {
            //     return [2,1]
            // }
        }
    })
</script>

<style>
.el-card .el-table .rowh {
    height: 15px;
}
.el-card .el-table .warning-row {
    background-color: orange;
    height: 10px;
    color: white;
}
.el-card .el-table .success-row {
    background-color: blanchedalmond;
    height: 10px;
}
.el-table__body-wrapper::-webkit-scrollbar {
    width: 5px;
    height: 10px;
}

.el-table__body-wrapper::-webkit-scrollbar-thumb {
   background-color: #cccccc;
    border-radius: 10px;
}
.el-table__body-wrapper::-webkit-scrollbar-track {
    border-radius: 2px;
}
</style>
