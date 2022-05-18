

1.查询客户接口

请求地址：/services/transaction/select_customer

请求头：
Content-Type: application/x-www-form-urlencoded
token: xxxx
staff_email: xxx@xx.xx

请求参数：
1. customer_name  			字符串类型（选传）                       # 客户名称
2. customer_id				数字类型（选传）                           #  客户ID
3. customer_social_credit_code	字符串类型（选传）                       # 客户社会信用代码
4. page_number    				数字类型（选传） 默认值 1
5. page_size            				数字类型（选传） 默认值 5

返回值：

{
"count": 1,                                                              # 总数
"list": [
{
"rec_id": 1,                                                                 # 主键
"created_at": "2022-03-09T23:50:42+08:00",      # 记录创建时间
"updated_at": "2022-03-09T23:50:43+08:00",     # 记录上次修改时间
"deleted_at": null,                                                # 无用字段
"staff_name": "超级管理员",                                 # 职工名称
"staff_alias": "hello",                                             # 职工别名
"staff_email": "645171033@qq.com",                   # 职工邮箱
"staff_phone": null,                                               # 职工电话
"staff_password": "********",                                 # 职工密码（均被后端隐藏为8个*）
"staff_position": null,                                            # 职工职位
"staff_department_id": null,                                  # 职工所属部门ID
"staff_department_name": null,                            # 职工所属部门名称
"staff_role_id": 1,                                                   # 职工系统角色ID
"staff_role_name": "root"                                      # 职工系统角色名称
}
]
}



请求案例：



2.创建客户接口

请求地址：/transaction/staff/create_customer

请求头：
Content-Type: application/json
token: xxxx
staff_email: xxx@xx.xx

请求参数：
1.       customer_name				字符串类型（必传）                                    # 客户名称
2.       customer_legal_name		字符串类型（选传）                                    # 客户公司全称 
3.       customer_alias				字符串类型（选传）                                     # 客户昵称
4.       customer_logo_url                       字符串类型（选传）                                     # 客户logo 
5.       customer_address			字符串类型（选传）                                     # 客户地址
6.       customer_social_credit_code	字符串类型（选传）                                     # 客户社会信用代码
7.       customer_contact			字符串类型（选传）                                    # 客户方联系人
8.       customer_contact_phone		字符串类型（选传）                                   # 客户方联系人电话
9.       customer_contact_wechat	字符串类型（选传）                                    # 客户方联系人微信
10.       customer_owner_id			数字类型（选传）                                     	# 我方客户负责人ID
11.      customer_owner_name		字符串类型（选传）                                   # 我方客户负责人ID
12.      remark						字符串类型（选传）                                   # 备注

返回值：

{
"message": "创建客户成功!"
}



请求案例：





3.修改客户接口

请求地址：/transaction/staff/update_customer

请求头：
Content-Type: application/json
token: xxxx
staff_email: xxx@xx.xx

请求参数：

1.      rec_id                                        	数字类型  （必传）                                  # 客户ID
2.     customer_name				字符串类型（选传）                                    # 客户名称
3.       customer_legal_name		字符串类型（选传）                                    # 客户公司全称 
4.       customer_alias				字符串类型（选传）                                     # 客户昵称
5.       customer_logo_url                       字符串类型（选传）                                     # 客户logo 
6.       customer_address			字符串类型（选传）                                     # 客户地址
7.       customer_social_credit_code	字符串类型（选传）                                     # 客户社会信用代码
8.       customer_contact			字符串类型（选传）                                    # 客户方联系人
9.       customer_contact_phone		字符串类型（选传）                                   # 客户方联系人电话
10.       customer_contact_wechat	字符串类型（选传）                                    # 客户方联系人微信
11.       customer_owner_id			数字类型（选传）                                     	# 我方客户负责人ID
12.      customer_owner_name		字符串类型（选传）                                   # 我方客户负责人ID
13.      remark						字符串类型（选传）                                   # 备注

返回值：

{
"message": "修改客户成功!"
}



4.删除客户接口

请求地址：/transaction/staff/delete_customer

请求头：
Content-Type: application/x-www-form-urlencoded
token: xxxx
staff_email: xxx@xx.xx

请求参数：

1.      customer_id                                        数字类型  （必传）                                  # 客户ID


返回值：
{
"message": "删除客户成功!"
}
