<template>
  <el-form ref="form" :model="form">
    <el-form-item label="标题">
      <el-input v-model="formData.title" style="width: 500px" placeholder="标题"></el-input>
    </el-form-item>
    <el-form-item label="封面：">
      <el-upload
          class="avatar-uploader"
          action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
          :show-file-list="false"
          :on-success="onCoverFileChange"
      >
        <img v-if="file_url" :src="file_url" class="avatar" width="500" />
        <el-icon v-else class="my-avatar-uploader-icon">
          <Plus/>
        </el-icon>
      </el-upload>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onEdit()">提交</el-button>
    </el-form-item>
  </el-form>
</template>

<script>

import {ElLoading} from "element-plus";
import {getImageInfo, getQiniuToken, postImageEdit} from "../../common/api";
import axios from "axios";
import localStorageCache from "../../common/localStorage";

export default {
  data() {
    return {
      formData: {
        id: this.$route.params.id,
        title: '',
        img_url: ''
      },
      file_url: "",
      coverFile: []
    };
  },
  methods: {
    onCoverFileChange(res, file) {
      this.file_url = URL.createObjectURL(res.raw);
      this.coverFile = res;
    },
    async onEdit() {
      let scene = localStorageCache.get("x-scene")
      let isLoading = ElLoading.service({
        lock: true,
        text: 'Loading',
      })

      let qiniuToken = null
      if (this.file_url.length > 0) {
        await getQiniuToken(scene.upload.image, 'png').then(res => {
          qiniuToken = res.data[0]
        })

        let qiniuFormData = new FormData()
        qiniuFormData.append("key", qiniuToken.key)
        qiniuFormData.append("token", qiniuToken.token)
        qiniuFormData.append("file", this.coverFile.raw, this.coverFile.raw.name)
        let axiosObj = axios.create({
          baseURL: "", // 所有的请求地址前缀部分
          timeout: 60000, // 请求超时时间毫秒
          withCredentials: false, // 异步请求携带cookie
          headers: {},
          data: {},
        })
        console.log(qiniuFormData)
        await axiosObj.post(qiniuToken.up_serve, qiniuFormData).then(res => {
          this.coverFile.value = []
          this.formData.img_url = res.data.key
        })
      }
      postImageEdit(this.formData).then((res) => {
        console.log('提交成功！')
        this.$router.push('/image-list')
      }).catch(() => {
        console.log('数据错误！')
      })
      isLoading.close()
    }
  },
  created() {
    console.log()
    console.log(this.formData.id)
    if (this.formData.id > 0) {
      getImageInfo({id: this.formData.id}).then((res) => {
        this.formData.title = res.data.title
        this.file_url = this.formData.img_url = res.data.img_url
      })
    }
  }
}
</script>

<style scoped>

</style>