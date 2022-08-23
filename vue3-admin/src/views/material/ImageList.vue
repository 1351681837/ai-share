<template>
  <el-table
      ref="multipleTable"
      :data="tableData"
      tooltip-effect="dark"
      style="width: 100%">
    <el-table-column
        type="selection">
    </el-table-column>
    <el-table-column
        prop="id"
        label="ID">
      <!--      <template slot-scope="scope">{{ scope.row.date }}</template>-->
    </el-table-column>
    <el-table-column
        prop="title"
        label="标题">
    </el-table-column>
    <el-table-column
        label="图片">
      <el-image
          style="width: 100px; height: 100px"
          src=""
          :fit="fit"></el-image>
    </el-table-column>
    <el-pagination background layout="total, sizes, prev, pager, next, jumper" v-model:page-size="pageSize"
                   v-model:currentPage="page" :page-sizes="[10, 30, 50, 100]" :total="total"/>
  </el-table>
  <!--  <div style="margin-top: 20px">-->
  <!--    <el-button @click="toggleSelection([tableData[1], tableData[2]])">切换第二、第三行的选中状态</el-button>-->
  <!--    <el-button @click="toggleSelection()">取消选择</el-button>-->
  <!--  </div>-->
</template>

<script>
import {getImageList} from "../../common/api"

export default {
  name: "ImageList",
  components: {},
  data() {
    return {
      tableData: [],
      page: 1,
      pageSize: 10,
      total: 0,
    }
  },
  created() {
    getImageList().then((res) => {
      let data = this.data();
      data.page++
      data.tableData = res.data.list
      data.total = res.data.total
    }).catch(() => {
      console.log('数据错误！')
    })
  }
}
</script>

<style scoped>

</style>