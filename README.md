# appManagementServer

移动后端接口API

###  登录注册部分
- 登录  
   ```
  router: /app/login
  
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
  
  request:{
    material_name: "电缆"
    material_size: "规格"
    material_num: "数量"
    applier: "申请人id"
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
  router: /apply/list
  
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
### 