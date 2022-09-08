<template>
  <el-form ref="form" :model="form"  label-width="100px"  style="max-width: 460px">
    <el-form-item label="标题">
      <el-input    md="4" class="w-50 m-2"  v-model="formData.name" placeholder="标题"/>
    </el-form-item>
    <el-form-item label="排序">
      <el-input      class="w-50 m-2" v-model="formData.sort" placeholder="排序"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onEdit()">提交</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import {ElMessage, ElLoading} from "element-plus";
import {getVideoTutorialCateInfo, postVideoTutorialCateEdit} from "../../common/api";

export default {
  name: "VideoTutorialCateEdit",
  data() {
    return {
      formData: {
        id: this.$route.params.id,
        title: '',
        sort: ''
      },
    };
  },
  methods: {
    onEdit() {
      let isLoading = ElLoading.service({
        lock: true,
        text: 'Loading',
      })
      postVideoTutorialCateEdit(this.formData).then((res) => {
        ElMessage({
          message: '保存成功',
          type: 'success',
        })
        this.$router.push('/video-tutorial-cate-list')
      }).catch(() => {
        ElMessage({
          message: '保存失败',
          type: 'warning',
        })
      })
      isLoading.close()
    }
  },
  created() {
    if (this.formData.id > 0) {
      getVideoTutorialCateInfo({id: this.formData.id}).then((res) => {
        this.formData = res.data
      })
    }
  }
}
</script>

<style scoped>

</style>