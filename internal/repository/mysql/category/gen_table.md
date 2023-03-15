#### go_gin_api.category 
站分类

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | int(10) unsigned | PRI | NO | auto_increment |  |
| 2 | parent_id |  | int(11) |  | NO |  | 0 |
| 3 | sort |  | int(11) |  | NO |  | 0 |
| 4 | title |  | varchar(50) |  | NO |  |  |
| 5 | icon |  | varchar(20) |  | NO |  |  |
| 6 | levels |  | int(11) |  | YES |  |  |
| 7 | created_at |  | timestamp |  | YES |  |  |
| 8 | updated_at |  | timestamp |  | YES |  |  |
