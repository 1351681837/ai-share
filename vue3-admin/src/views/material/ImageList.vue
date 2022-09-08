<template>
  <div style="text-align: left">
    <el-button type="primary"   style="float: right" @click="addImage">新增</el-button>
    <el-form :inline="true" class="demo-form-inline">
      <el-form-item label="标题">
        <el-input v-model="search.title" placeholder="标题"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">查询</el-button>
      </el-form-item>
    </el-form>

  </div>

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
    </el-table-column>
    <el-table-column
        prop="title"
        label="标题">
    </el-table-column>
    <el-table-column
        label="图片">
      <template #default="scope">
        <div style="display: flex; align-items: center">
          <el-avatar shape="square" :size="100" fit="fill" :src="scope.row.img_url"/>
        </div>
      </template>
    </el-table-column>
    <el-table-column fixed="right" label="操作" width="120">
      <template  #default="scope">
        <el-button link type="primary" @click="edit(scope.row.id)" size="small">Edit</el-button>
      </template>
    </el-table-column>
  </el-table>
  <el-pagination background layout="total, sizes, prev, pager, next, jumper" v-model:page-size="search.page_size"
                 v-model:currentPage="search.page" :page-sizes="[10, 30, 50, 100]" :total="total"
                 @size-change="getData" @current-change="getData"
  />
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
      total: 0,
      search: {
        page: 1,
        page_size: 10,
        title: ''
      },
    }
  },
  methods: {
    onSubmit() {
      this.search.page = 1
      this.getData()
    },
    addImage() {
      this.$router.push('/image-add')
    },
    edit(id){
      this.$router.push('/image-add/'+id)
    },
    // handleSizeChange(){
    //   this.getData()
    // },
    // handleCurrentChange(){
    //   this.()
    // },
    getData() {
      // console.log(this.search);
      getImageList(this.search).then((res) => {
        let data = this.$data;
        data.tableData = res.data.list
        data.total = res.data.total
      }).catch(() => {
        console.log('数据错误！')
      })
    }
  },
  created() {
    this.getData()
  }
}
</script>

<style scoped>

</style>