# app Server
	为企业android应用提供后端接口服务

# API
移动后端接口api
注：参数请求返回都以json格式

###  登录注册部分
- 登录
   ```
  router: /app/login
  method: POST

  request:{
  	userID: "123456"
  	password: "abcdef"
  }

  response (成功):{
 	 status: 1
 	 msg: "登录成功"
 	 param:""
  }
   response(失败):{
 	 status: 0
 	 msg: "登录失败, 原因.."
 	 param: ""
  }
  ```
- 注册 (暂时应该用不到)
  ```
  router: /app/register
  method: POST

  request:{
  	userID: "123456"
  	password: "abcdef"
  }

  response (成功):{
 	 status: 1
 	 msg: "注册成功"
 	 param:""
  }
  response(失败):{
 	 status: 0
 	 msg: "注册失败,原因."
 	 param: ""
  }
  ```
### 物料申请
- 填写申请表
```
  router: /apply/material
  method: POST

  request:{
  	material_id: 123
  	applier: 100
  }

  response (成功):{
 	 status: 1
 	 msg: "提交成功"
 	 param:""
  }
  response(失败):{
 	 status: 0
 	 msg: "提交失败,原因."
 	 param: ""
  }
  ```
- 申请单 列表
```
  router: /apply/
  method: GET

  request:{
    // 暂时不需要填写时间
  	start_time: “开始时间”
  	end_time: "结束时间"
  }

  response (成功):{
 	 status: 1
 	 msg: "获取成功"
 	 param: [
 	 {
 	 	table_id:  100   //表单id
 	 	applier: "申请人"
 	 	apply_time:"2018-11-10"
 	 },
 	 {
 		 table_id:  101   //表单id
 	 	applier: "申请人"
 	 	apply_time:"2018-11-11"
 	 }
 	 ]
  }
  response(失败):{
 	 status: 0
 	 msg: "获取失败,原因."
 	 param: ""
  }
  ```
- 查看某申请表详细内容
```
  router: /apply/detail
  method: GET

  request:{
    table_id : 100
  }

  response (成功):{
 	 status: 1
 	 msg: "获取成功"
 	 param: {
 	 	material_name: "电缆"
  		material_size: "规格"
  		material_num: "数量"
  		applier: "申请人id"
  		apply_time "申请时间"
 	 }
  response(失败):{
 	 status: 0
 	 msg: "获取失败,原因."
 	 param: ""
  }
  ```
### 材料入库
- 新增材料（入仓库）
```
	router: /material/in_warehouse
	method: POST

	request:{
	     name : "材料名称"
	     description: "规格"
	     unit: “单位”
	     provider: "提供商"
	     num: "数量"
  }

   response(成功/失败):{
 	 status: 1/0
 	 msg: "获取成功/失败,原因."
 	 param: ""
  }
```

### 材料领取
- 填写领取单
```
router: /material/receive/write_table
method: POST

request:{
		 receiver : 领取人id
		 material_id: 材料id
		 num: "领取数量"
  }

   response(成功/失败):{
 	 status: 1/0
 	 msg: "填写成功/失败,原因."
 	 param: ""
  }
```
- 查看某一领取单
```
	router: /material/receive/detail
	method: GET

	request:{
		 table_id: "领料单的id"
  }
   response(成功/失败):{
 	 status: 1/0
 	 msg: "成功/失败,原因."
 	 param: {
 	 	table_id: 领料单id
 	 	reciever_name: 领取人名字
 	 	write_time: 填写时间
 	 	material_id: 材料id
 	 	material_name: ”材料名字“
 	 	material_size: ”材料规格“
 	 	material_provider: "提供商"
 	 	material_num: 材料数量
	}
  }
```
- 查看领取单列表
```
router: /material/receive/
	method: GET
	request:{
	}
   response(成功/失败):{
 	 status: 1/0
 	 msg: "成功/失败,原因."
 	 param: [
 	 {
 	 	table_id: 领料单id
 	 	reciever_name: 领取人名字
 	 	write_time: 填写时间
 	 	material_name: 材料名
  },
  {}
  ...
  ]
```
### 材料归还
- 填写归还单
```
router: /material/back/write_table
method: POST

request:{
		 table_id : 领料单id
		 backer: 归还人id
		 num: "归还数量"
  }

   response(成功/失败):{
 	 status: 1/0
 	 msg: "填写成功/失败,原因."
 	 param: ""
  }
```

### 材料审核
- 填写审核
```
router: /material/check/write_table
method: POST

request:{
		 table_id : 领料单id
		 checker: 审核人id
		 num: "审核数量"
  }

   response(成功/失败):{
 	 status: 1/0
 	 msg: "填写成功/失败,原因."
 	 param: ""
  }
```