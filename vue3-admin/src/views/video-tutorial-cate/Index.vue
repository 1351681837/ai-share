<template>
  <el-container class="my-table-container">
    <div class="my-table-top">
      <el-form class="my-form">
        <el-form-item label="标题:">
          <el-input v-model="search.title" placeholder="请输入标题"/>
        </el-form-item>
        <el-form-item label="创建时间:">
          <div class="block">
            <el-date-picker
                v-model="search.start_time"
                type="datetime"
                placeholder="请输入开始时间"
            />
          </div>
          <div class="block">
            <el-date-picker
                v-model="search.end_time"
                type="datetime"
                placeholder="请输入结束时间"
            />
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSearch">搜索</el-button>
        </el-form-item>
      </el-form>
      <div class="actions">
        <el-button type="primary" @click="handleAdd">新增</el-button>
      </div>
    </div>

    <el-table :data="tableData" height="70vh" style="width: 100%">
      <el-table-column prop="id" label="ID"/>
      <el-table-column prop="name" label="分类"/>
      <el-table-column prop="sort" label="排序"/>
      <el-table-column label="状态">
        <template #default="scope">
          {{ scope.row.is_lock === 1 ? '正常' : '锁定' }}
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="创建时间"/>
      <el-table-column prop="updated_at" label="更新时间"/>
      <el-table-column prop="" label="操作">
        <template #default="scope">
          <el-row>
            <el-button size="small" type="primary" @click="handleEdit(scope.$index, scope.row)">修改</el-button>
          </el-row>
          <el-row>
            <Lock :row="scope.row"/>
          </el-row>
          <el-row>
            <Delete :row="scope.row" :index="scope.$index"/>
          </el-row>
          <el-row>
            <el-button size="small" type="primary" @click="handleContentEdit(scope.$index, scope.row)">编辑内容</el-button>
          </el-row>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination background layout="total, sizes, prev, pager, next, jumper" v-model:page-size="search.page_size"
                   v-model:currentPage="search.page" :page-sizes="[10, 30, 50, 100]" :total="total"/>

  </el-container>

  <!--  <PptEdit :rowInfo="rowInfo" :typeList="typeList" @changeRow="onChangeRow"/>-->
  <!--  <PptContentEdit />-->
</template>

<script>
import {getVideoTutorialCateList} from "../../common/api";
import Delete from "../../components/Delete.vue";
import Lock from "../../components/Lock.vue";
import {provide, ref} from "vue";

const mMark = ref("")

export default {
  name: "VideoTutorialCateList",
  components: {Delete, Lock},
  data() {
    return {
      tableData: [],
      total: 0,
      m_mark: '',
      search: {
        page: 1,
        page_size: 10,
        title: '',
        start_time: '',
        end_time: '',
      },
    }
  },
  methods: {
    onSearch() {
      this.search.page = 1
      this.getData()
    },
    handleAdd() {
      this.$router.push('/video-tutorial-cate-edit')
    },
    handleEdit(id) {
      // this.$router.push('/image-add/'+id)
    },
    getData() {
      // console.log(this.search);
      getVideoTutorialCateList(this.search).then((res) => {
        let data = this.$data;
        data.tableData = res.data.list
        data.total = res.data.total
        mMark.value = res.data.m_mark
      }).catch(() => {
        console.log('数据错误！')
      })
    }
  },
  created() {
    this.getData()

    provide('mMark', mMark)
  }
}
</script>

<style scoped>

</style>