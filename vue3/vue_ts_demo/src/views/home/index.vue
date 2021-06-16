<template>
<div>
    <el-row :gutter="24" style="height: 600px; overflow-y: hidden;padding: 15px">
        <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="6">
            <tables></tables>
        </el-col>
        <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="6">
           <echart></echart>
        </el-col>
        <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="6">
            sss
        </el-col>
        <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="6">
            sss
        </el-col>
    </el-row>
    <el-row :gutter="24">
        <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
           ssss
        </el-col>
        <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
          ssss
        </el-col>
    </el-row>
    <el-row :gutter="24">
        <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
          ssss
        </el-col>
        <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
          ssss
        </el-col>
    </el-row>
    <el-button type="primary">sss</el-button>
    <el-input></el-input>
</div>
</template>

<script>
    import {defineAsyncComponent, defineComponent} from "vue"
    import PageLoading from "@/views/PageLoading"
    // 防止刷新時候沒有載入完成出現未處理的代碼
    const Tables = defineAsyncComponent({
        loader: () => import("@/views/home/components/tables"),
        loadingComponent: PageLoading
    })
    const Echart = defineAsyncComponent({
        loader: () => import("@/views/home/components/echart"),
        loadingComponent: PageLoading
    })
    export default defineComponent({
        name: "index.vue",
        components: {Tables, Echart},
    })
</script>

<style>
</style>
