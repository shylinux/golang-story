package product

import (
	"context"
	"html/template"
	"path"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/development/proto"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/system"
)

type ProductCmds struct{ config *config.Config }

func (s *ProductCmds) Create(ctx context.Context, arg ...string) {
	conf := s.config.Product
	dir := path.Join("usr", conf.Name)
	if !system.Exists(dir) {
		system.CommandBuild("", "git", "clone", conf.Repos, dir)
	}
	dir = path.Join(dir, "src")
	for _, portal := range conf.Portal {
		for _, views := range portal.Views {
			for _, view := range views.View {
				service := config.WithDef(view.Service, views.Service)
				if view.Source != "" && !system.Exists(path.Join(view.Source)) {
					system.CommandBuild("", "./bin/matrix", "project", "create", view.Source)
					system.CommandBuild(view.Source, "go", "build", "-v", "-o", "bin/matrix", "cmd/cmds.go")
					system.CommandBuild(view.Source, "./bin/matrix", "service", "create", service)
					system.CommandBuild(view.Source, "make")
				}
				ls := strings.Split(service, ".")
				service = ls[len(ls)-1]
				file := path.Join(view.Source, dir, view.Display)
				if !system.Exists(file) {
					system.NewTemplateFile(file, viewTemplate, template.FuncMap{
						"Views": func() config.Views { return views },
					}, map[string]interface{}{
						"portal":  portal.Name,
						"api":     template.JS(service),
						"service": template.JS(proto.Capital(service)),
					})
				}
				if view.Source != "" {
					copyFile(file, path.Join(dir, view.Display))
					copyFile(path.Join(view.Source, dir, "api/"+service+".js"), path.Join(dir, "api/"+service+".js"))
				}
			}
		}
	}
}
func (s *ProductCmds) List(ctx context.Context, arg ...string) {
}
func NewProductCmds(config *config.Config, cmds *cmds.Cmds) *ProductCmds {
	s := &ProductCmds{config: config}
	cmds = cmds.Add("product", "product command", s.List)
	cmds.Add("create", "create path", s.Create)
	return s
}
func copyFile(src, dst string) {
	if src != dst {
		system.Command("", "cp", src, dst)
	}
}

var viewTemplate = `
<template>
  <div>
    <div class="operate" style="display:flex">
      <div style="display:flex">
        <el-button @click="refresh">刷新</el-button>
        <el-input v-model="filter" @keyup.enter="search" placeholder="请输入关键字" clearable></el-input>
        <el-button @click="search">查找</el-button>
      </div><div style="flex-grow:1"></div>
      <div>
        <el-button type="primary" @click="toCreate">创建</el-button>
      </div>
    </div>
    <el-table class="table" :max-height="height" :data="list">
      <el-table-column prop="{{ .service }}ID" label="{{ .service }}ID"></el-table-column>
      <el-table-column prop="name" label="名称"></el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button type="warning" @click="toModify(scope.row)" plain>编 辑</el-button>
          <el-button type="danger" @click="remove(scope.row)" plain>删 除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination v-model:current-page="page" v-model:page-size="count" :total="total" @size-change="getList" @current-change="getList" layout="total, sizes, prev, pager, next, jumper" />
  </div>
  <el-drawer v-model="isCreate" title="创建">
    <el-form>
      <el-form-item prop="name" label="名称">
        <el-input v-model="form.name"></el-input>
      </el-form-item>
      <el-form-item v-show="create">
        <el-button type="primary" @click="create">创 建</el-button>
        <el-button type="primary" @click="cancel">取 消</el-button>
      </el-form-item>
    </el-form>
  </el-drawer>
  <el-drawer v-model="isModify" title="编辑">
    <el-form>
      <el-form-item prop="name" label="名称">
        <el-input v-model="form.name"></el-input>
      </el-form-item>
      <el-form-item v-show="modify">
        <el-button type="primary" @click="modify">编 辑</el-button>
        <el-button type="primary" @click="cancel">取 消</el-button>
      </el-form-item>
    </el-form>
  </el-drawer>
</template>

<script>
import { ElMessage } from 'element-plus'
import { {{ .service }}Service } from '@/api/{{ .api }}';
import router from '@/router'

export default {
  data: () => {
    return {
      list: [],
      page: 1,
      count: 10,
      total: 100,
      filter: "",
      height: 600,
      isCreate: false,
      isModify: false,
      form: {
		{{ .service }}ID: "",
        name: "",
      },
    }
  },
  mounted() { this.getList() },
  methods: {
    async getList() {
      let res = await {{ .service }}Service.List(this.page, this.count, "name", this.filter)
      this.list = res.data||[], this.total = res.total||0
      return res
    },
    async refresh() {
      var res = await this.getList()
      if (!res.error) {
        ElMessage.info("刷新成功")
      }
    },
    async search() {
      var res = await {{ .service }}Service.Search("name", "*"+this.filter+"*")
      if (!res.error) {
        this.list = res.data||[], this.total = res.total||0
        ElMessage.info("查找成功")
      }
    },
    async remove(value) {
      var res = await {{ .service }}Service.Remove(value.{{ .service }}ID)
      if (!res.error) {
        ElMessage.info(value.name+" 删除成功")
        this.getList()
      }
    },
    toInfo(value) {
      router.push("/{{ .portal }}/{{ .service }}/info/"+value.{{ .service }}ID)
    },
    toCreate() {
      this.isCreate = true
    },
    toModify(value) {
      this.form.name = value.name
      this.form.{{ .service }}ID = value.{{ .service }}ID
      this.isModify = true
    },
    async create() {
      var res = await {{ .service }}Service.Create(this.form.name)
      if (!res.error) {
        ElMessage.info(this.form.name+" 添加成功")
        this.getList(), this.isCreate = false
      }
    },
    async modify() {
      var res = await {{ .service }}Service.Rename(this.form.{{ .service }}ID, this.form.name)
      if (!res.error) {
        ElMessage.info(this.form.name+" 修改成功")
        this.getList(), this.isModify = false
      }
    },
    cancel() {
      this.isModify = false
      this.isCreate = false
    },
  },
}
</script>
<style scoped>
.table {
  width: 100%;
  overflow: auto;
}
.operate * {
  float:left;
}
</style>
`
