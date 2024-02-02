<template>
  <div>编辑器</div>
</template>

<script>
import request from '@/utils/request'
import api from '@/api/API_LIST.js'
export default {
  name: 'EditDoc',
  mounted() {
    this.getContent()
  },
  methods: {
    getContent() {
      console.log(api.Content)
      request({
        url: api.Content,
        method: 'get',
        params: {}
      }).then( resp =>{
        console.log(resp)
      })
    }
  }
}
</script>

<style scoped>

</style>
