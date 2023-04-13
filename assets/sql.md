#

## MySQL 常用命令

- 时区设置：MySQL服务器当前的全局时区设置、当前会话的时区设置、系统的时区设置

> SELECT @@global.time_zone;  
> SELECT @@session.time_zone;  
> SELECT @@system_time_zone;

- 事务清空表数据，保留表结构

> START TRANSACTION;  
> TRUNCATE TABLE [table_name];  
> -- 如果出现错误，执行ROLLBACK回滚事务  
> ROLLBACK;  
> -- 如果没有错误，则提交事务  
> COMMIT;

- 添加、删除字段

> ALTER TABLE [table_name] ADD COLUMN [column_name] [data_type];  
> ALTER TABLE [table_name] DROP COLUMN [column_name];  
> ALTER TABLE [table_name] ADD COLUMN [column_name] [data_type] AFTER [existing_column_name];

- 分组查询：SUM、AVG、MAX、MIN

> SELECT [column1], [column2], ..., [aggregate_function]([column]) FROM [table_name] GROUP BY [column1], [column2], ...;  
> SELECT customer_name, SUM(order_amount) as total_amount FROM orders GROUP BY customer_name;  
> customer_name是分组的列，SUM(order_amount)是对order_amount列进行求和计算

- 更新

> UPDATE [table_name] SET [column_name] = [new_value] WHERE id = [id_value];  
> UPDATE students SET score = 90 WHERE id = 1;

- 关联查询

> SELECT [column1], [column2], ... FROM [table1] JOIN [table2] ON [table1].[column] = [table2].[column];  
> SELECT students.name, scores.score FROM students JOIN scores ON students.id = scores.student_id
WHERE scores.subject = 'Math';  
> 在这个命令中，students和scores是要连接的表，students.id和scores.student_id是要连接的列，WHERE子句用于过滤只显示数学成绩。查询结果将显示每个学生的姓名和数学成绩。

- 增加、删除索引

> CREATE INDEX [index_name] ON [table_name] ([column_name]);  
> CREATE INDEX idx_age ON students (age);  
> DROP INDEX [index_name] ON [table_name];  
> DROP INDEX idx_age ON students;

- MVCC 多版本并发控制
在使用 `READ COMMITTD`、`REPEATABLE READ` 这两种隔离级别的事务在执行普通的 `SELECT` 操作时访问记录的 `版本链` 的过程，这样子可以使不同事务的 `读-写`、`写-读` 操作 `并发执行`，从而 `提高系统性能`。

- 按日计统计

> SELECT DATE_FORMAT(created_at,'%Y-%m-%d') time,SUM(id)FROM user WHERE id>10 GROUP BY time
