<template>
  <div style="display: flex;">
    <div style="order: 0;width: 300px;">
      <el-card style="height: calc(100vh);overflow-y: scroll;">
        <div slot="header" class="clearfix">
          <span>目录</span>
          <el-button style="float: right; padding: 3px 0" type="text">添加</el-button>
        </div>
        <div>
          <el-tree :data="tree" :props="{children: 'children', label: 'text'}"
                   node-key="id"
                   highlight-current
                   default-expand-all
                   @node-click="clickNode"
                   :expand-on-click-node="false">
      <span slot-scope="{ node, data }" class="custom-tree-node">
        <template>
           {{data.text}}
        </template>
      </span>
          </el-tree>
        </div>

      </el-card>
    </div>
    <div style="order: 1;width: calc(100vw - 300px);">
      <!--      <mavon-editor style="height: 100%;" v-model="content"/>-->
      <Editor style="height: 100%;" :value="content" :plugins="plugins" @change="handleChange"/>
    </div>
  </div>
</template>

<script>
  import { getContent, getTree } from '@/api/docApi.js'
  import { Editor } from '@bytemd/vue'
  import gfm from '@bytemd/plugin-gfm'

  const plugins = [
    gfm()
    // Add more plugins here
  ]
  export default {
    name: 'EditDoc',
    components: { Editor },
    data() {
      return {
        key: this.$route.params.key,
        id: this.$route.params.id,
        tree: undefined,
        content: '',
        plugins
      }
    },
    mounted() {
      this.reqTree()
      this.reqContent()
    },
    methods: {
      reqTree() {
        getTree({ key: this.key }, {}).then(resp => {
          console.log(resp.data)
          this.tree = this.listToTreeList(JSON.parse(resp.data.data))
          console.log(this.tree)
        })
      },
      reqContent() {
        if (this.id == undefined) {
          return
        }
        getContent({ key: this.key }, {}).then(resp => {
          console.log(resp)
        })
      },
      clickNode(data) {
        console.log(data)
        getContent({ key: this.key, id: data.id }, {}).then(resp => {
          console.log(resp.data.data)
          this.content = resp.data.data.content
        })
      },
      /**
       * 将普通列表转换为树结构的列表
       * @param list
       * @returns {{length}|*|Array}
       */
      listToTreeList(list) {
        if (!list || !list.length) {
          return []
        }
        let treeListMap = {}
        for (let item of list) {
          treeListMap[item.id] = item
        }
        for (let i = 0; i < list.length; i++) {
          if (list[i].parent && treeListMap[list[i].parent]) {
            if (!treeListMap[list[i].parent].children) {
              treeListMap[list[i].parent].children = []
            }
            treeListMap[list[i].parent].children.push(list[i])
            list.splice(i, 1)
            i--
          }
        }
        return list
      }
    }
  }
</script>

<style scoped lang="scss">
  .custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    /*justify-content: space-between;*/
    font-size: 20px;
    padding-right: 4px;
    width: 100%;
  }

  /deep/ .el-tree-node__content {
    height: 50px;

  }

  /deep/ .el-tree-node {
    white-space: normal;
  }

  .tree {
    overflow-x: hidden;
  }

  /deep/ .el-tree {
    width: 100%;

    .el-tree-node {
      .el-tree-node__content {
        height: 100%;
        /*align-items: start;*/
        word-break: break-all;
      }
    }
  }

  /deep/ .el-tree--highlight-current .el-tree-node.is-current > .el-tree-node__content {
    background-color: #1094fa;
    color: white;
  }

  .file-img {
    width: 24px;
    height: 24px;
  }
</style>
