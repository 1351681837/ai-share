<template>
  <el-form ref="form" :model="form" label-width="100px" style="max-width: 460px">
    <el-form-item label="标题">
      <el-input md="4" class="w-50 m-2" v-model="formData.name" placeholder="标题"/>
    </el-form-item>

    <el-form-item label="分类:">
      <el-select v-model="formData.cate_id" placeholder="Please Select is type">
        <el-option v-for="item in cateName" :value="item.id" :label="item.name"/>
      </el-select>
    </el-form-item>

    <el-form-item label="视频：">
      <el-upload
          :auto-upload="false"
          :show-file-list="false"
          :on-change="onPptFileChange"
      >
        <template #trigger>
          <el-input type="primary" v-model="formData.videoUrl" disabled/>
        </template>
      </el-upload>
    </el-form-item>
    <el-form-item label="封面：">
      <el-upload
          class="avatar-uploader"
          action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
          :show-file-list="false"
          :on-success="onCoverFileChange"
      >
        <img v-if="file_cover_url" :src="file_cover_url" class="avatar" width="500"/>
        <el-icon v-else class="my-avatar-uploader-icon">
          <Plus/>
        </el-icon>
      </el-upload>
    </el-form-item>
    <el-form-item label="排序">
      <el-input class="w-50 m-2" v-model="formData.sort" placeholder="排序"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onEdit()">提交</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import {ElLoading, ElMessage} from "element-plus";
import {getVideoTutorialCateAllName, getVideoTutorialInfo, postVideoTutorialEdit} from "../../common/api";

export default {
  name: "VideoTutorialEdit",
  data() {
    return {
      formData: {
        id: this.$route.params.id,
        cate_id: "",
        title: '',
        coverUrl: '',
        videoUrl: '',
        sort: ''
      },
      cateName: [],
      file_cover_url: '',
      file_video_url: '',
    };
  },
  methods: {
    onEdit() {
      let isLoading = ElLoading.service({
        lock: true,
        text: 'Loading',
      })
      postVideoTutorialEdit(this.formData).then((res) => {
        ElMessage({
          message: '保存成功',
          type: 'success',
        })
        this.$router.push('/video-tutorial-list')
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
      getVideoTutorialInfo({id: this.formData.id}).then((res) => {
        this.formData = res.data
      })
    }
    getVideoTutorialCateAllName().then((res) => {
      console.log(res.data)
      this.cateName = res.data
    })
  }
}
</script>

<style scoped>

</style>