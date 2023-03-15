#### go_gin_api.site 
网站信息

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | int(10) unsigned | PRI | NO | auto_increment |  |
| 2 | category_id |  | int(11) |  | YES |  |  |
| 3 | title |  | varchar(50) |  | YES |  |  |
| 4 | thumb |  | varchar(100) |  | YES |  |  |
| 5 | description |  | varchar(300) |  | YES |  |  |
| 6 | url |  | varchar(100) |  | YES |  |  |
| 7 | created_at |  | timestamp |  | NO |  | CURRENT_TIMESTAMP |
| 8 | updated_at |  | timestamp |  | NO |  |  |
| 9 | is_used | 是否启用 1:是  -1:否 | tinyint(1) |  | NO |  | 1 |
