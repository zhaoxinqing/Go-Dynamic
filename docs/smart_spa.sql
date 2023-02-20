-- 同一个用户的不同设备id
SELECT * FROM device 
WHERE device_id in ('ff01b339bcf71d61de15bf5f1bc73f59','f315aad34c11a9912e6a60245b02ed26')


-- 声卡
SELECT DISTINCT device_id,gpu_a,gpu_b,audio_a,audio_b FROM fingerprint 
WHERE device_id in ('ff01b339bcf71d61de15bf5f1bc73f59','f315aad34c11a9912e6a60245b02ed26') 


SELECT DISTINCT device_id, device_brand, device_model,os_type, os_version,
IFNULL((SELECT os_type FROM config_os WHERE id=device.os_type),'') os_str,browser_type, browser_version, create_time 
FROM device 
WHERE device_id 
IN ('db98ef1faa25cfcb0253429b5c92d395','1e20a723160450d67dedd54d952a07b5','a280b9ebfb62ed7e7e89cd70cbe664a6',
'7cd0f5e9421784df214937ea3cf6b1b7','c8179db4181b531868e8a8cfe7de3428','0eb06256a2037de1ae2c2c9b62d85a58',
'85baa0b6579ac0044d9b1ab84438fd94','e47825611a96e717c2a91653d7d7c33c','99e88370545e6c10a0e13b03e06428d4',
'c8b8928be46c6037507b90974b4762e8','8b7480143841df14295a5c240f3009aa','213381d30e2ab984e1c104ffefc35f41',
'87668c3112c7db1668f3cfc751fa547d','4b78badef24c01a69b52982998bfd2d3','458865056d5451df7240171db41d348f',
'75c8356a9e14a6c87d32316598bbe1d9','2505e652fb402f6b7bd363ecc78a519b','427e66b13b6aedd23c1fb0438381ae94')

-- 关联算法
Select DISTINCT device_id,ip,ua_os_type,ua_os_version,ua_browser_type,ua_browser_version,gpu_a,gpu_b,audio_a,audio_b  
FROM fingerprint 
WHERE device_id IN ('edf6e4e719758f0e72e54febfef75e19','6553f6f282a6bd39b9ce51f945dc753d','7f0686d4775427ffe5f2278fc300f143')


-- 统计计数排序
SELECT device_id,count(device_id) as num FROM analysis_device WHERE company_id = '%s'

-- 查询已关联账户
SELECT a.id
FROM user_device a
INNER JOIN user_device_group b
ON a.device_id = b.device_id AND a.user_id = b.user_id
ORDER BY a.id DESC

-- 查询未关联id
SELECT id FROM user_device WHERE id not IN (
SELECT a.id
FROM user_device a
INNER JOIN user_device_group b
ON a.device_id = b.device_id AND a.user_id = b.user_id
ORDER BY a.id DESC)


-- 时间计算
-- 1、精确到秒
SELECT TIMESTAMPDIFF(SECOND,"2017-09-30 18:04:01","2017-09-28 18:04:01")
-- 2、精确到分钟
SELECT TIMESTAMPDIFF(MINUTE,"2017-09-30 18:04:01","2017-09-28 18:04:01")


-- 设备在线时长
SELECT device_id,count(device_id) as count 
FROM user_device_location 
WHERE company_id = '5e4f18255fa6481514408ea1' AND 
device_id IN ('250075910cbb32a6ae0bde807387531f','64e2f038a447a30ed5bc146e05c58e7b','a7b40ba77760670fd34fbabc723a27b9',
'0d2b8c00be3bbac91d5f8de2d583acce','a9e887d38044d42900ec07ae7470c80a','0cac6243073c6b167669addd1edf44f8',
'55245754b48c5e733356bad4fbd050ad','99f1d5e036c35665987eb8087f2c06ef','29405f6a90666c3327b406106816f2fb','584b4992beaade8ae1004cab8a47e72e') 
GROUP BY device_id

SELECT MAX(id)id,user_id FROM `analysis_device` WHERE company_id = '5e4f18255fa6481514408ea1' GROUP BY user_id ORDER BY id DESC

SELECT user_id FROM `analysis_device` WHERE company_id = '5e4f18255fa6481514408ea1' ORDER BY id DESC



-- 关联分析
SELECT * FROM `user_device_group_copy1` WHERE correlation_device_id != linked_id

-- 1442641	120006	46ca5125dfa1c91df4ac266011e137bd	253028ec9e51b748c0376108014db18d	2022-07-12 15:07:43	2022-07-12 15:07:43

SELECT * FROM user_device WHERE id = 120006

-- a0dc74682d44690438e549828065eeb7

SELECT MAX(id)id FROM fingerprint WHERE device_id in ('46ca5125dfa1c91df4ac266011e137bd','253028ec9e51b748c0376108014db18d') AND user_id = 'a0dc74682d44690438e549828065eeb7'
AND counts = 0 GROUP BY device_id ORDER BY id DESC

SELECT * FROM fingerprint WHERE id in ('1843910','1842139')

1. SELECT 
2. DISTINCT <select_list>
3. FROM <left_table>
4. <join_type> JOIN <right_table>
5. ON <join_condition>
6. WHERE <where_condition>
7. GROUP BY <group_by_list>
8. HAVING <having_condition>
9. ORDER BY <order_by_condition>
10.LIMIT <limit_number>

FROM
<表名> # 选取表，将多个表数据通过笛卡尔积变成一个表。
ON
<筛选条件> # 对笛卡尔积的虚表进行筛选
JOIN <join, left join, right join...> 
<join表> # 指定join，用于添加数据到on之后的虚表中，例如left join会将左表的剩余数据添加到虚表中
WHERE
<where条件> # 对上述虚表进行筛选
GROUP BY
<分组条件> # 分组
<SUM()等聚合函数> # 用于having子句进行判断，在书写上这类聚合函数是写在having判断里面的
HAVING
<分组筛选> # 对分组后的结果进行聚合筛选
SELECT
<返回数据列表> # 返回的单列必须在group by子句中，聚合函数除外
DISTINCT
# 数据除重
ORDER BY
<排序条件> # 排序
LIMIT
<行数限制>


SELECT * FROM t WHERE username LIKE '%陈%'

SELECT * FROM t WHERE username LIKE '陈%'

SELECT * FROM t WHERE id IN (2,3)

SELECT * FROM t WHERE id BETWEEN 2 AND 3

-- 不走索引
select * from A where A.id in (select id from B);
-- 走索引
select * from A where exists (select * from B where B.id = A.id);

SELECT * FROM t WHERE id = 1 OR id = 3

SELECT * FROM t WHERE score = 0

SELECT * FROM t WHERE id = 1
   UNION
SELECT * FROM t WHERE id = 3

SELECT * FROM t WHERE score IS NULL

SELECT * FROM t WHERE score = 0

-- 全表扫描
SELECT * FROM T WHERE score/10 = 9
-- 走索引
SELECT * FROM T WHERE score = 10*9

SELECT username, age, sex FROM T WHERE 1=1

select col1 from table where key_part2=1 and key_part3=2

select col1 from table where col_varchar=123; 

-- 不走age索引
SELECT * FROM t order by age;
 
-- 走age索引
SELECT * FROM t where age > 0 order by age;

SELECT * FROM students FORCE INDEX (idx_class_id) WHERE class_id = 1 ORDER BY id DESC;

-- 方法一：
insert into T values(1,2);
insert into T values(1,3);
insert into T values(1,4);
-- 方法二：
Insert into T values(1,2),(1,3),(1,4);


Update t1 set time=now() where col1=1; 
Select time from t1 where id =1; 

Update t1 set time=now () where col1=1 and @now: = now (); 
Select @now; 

SELECT col1, col2, COUNT(*) FROM table GROUP BY col1, col2 ORDER BY NULL ;

SELECT col1 FROM customer_info WHERE customer_id NOT in (SELECT customer_id FROM sales_info )


SELECT col1 FROM customer_info 
   LEFT JOIN sales_info ON customer_info.customer_id=sales_info.customer_id 
       WHERE sales_info.customer_id IS NULL 


SELECT COL1, COL2, COL3 FROM TABLE WHERE COL1 = 10 
UNION ALL 
SELECT COL1, COL2, COL3 FROM TABLE WHERE COL3= 'TEST'; 

SELECT COL1, COL2, COL3 FROM TABLE WHERE COL1 = 10 
UNION 
SELECT COL1, COL2, COL3 FROM TABLE WHERE COL3= 'TEST';

select * from t where thread_id = 10000 and deleted = 0 
   order by gmt_create asc limit 0, 15;

select t.* from (select id from t where thread_id = 10000 and deleted = 0
   order by gmt_create asc limit 0, 15) a, t 
      where a.id = t.id; 

SELECT * FROM (SELECT ROW_NUMBER() OVER(ORDER BY ID ASC) AS rowid,* 
   FROM infoTab)t WHERE t.rowid > 100000 AND t.rowid <= 100050

SHOW create table abnormal_index_user

CREATE TABLE `abnormal_index_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_id` varchar(256) DEFAULT NULL,
  `user_id` varchar(256) DEFAULT NULL,
  `device_id` varchar(256) DEFAULT NULL,
  `index` int(11) DEFAULT NULL,
  `special_mark` int(11) DEFAULT NULL,
  `score_total` int(11) DEFAULT NULL,
  `score_location` int(11) DEFAULT NULL,
  `score_device` int(11) DEFAULT NULL,
  `score_time_history` int(11) DEFAULT NULL,
  `score_device_duration` int(11) DEFAULT NULL,
  `score_login_frequency` int(11) DEFAULT NULL,
  `score_special_base` int(11) DEFAULT NULL,
  `desc` varchar(256) DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ind_app_user` (`app_id`,`user_id`),
  UNIQUE KEY `ind_app_device` (`app_id`,`device_id`)
) ENGINE=InnoDB AUTO_INCREMENT=113040 DEFAULT CHARSET=utf8mb4

-- 删除：
DELETE FROM "t_log" WHERE "id" = '1194421845659742210';

-- 更新：
UPDATE "t_log" SET "create_time" = '2022-04-14 15:10:48' WHERE "id" = '1194423873890902017';

-- 排序：
1、ORDER BY "province" ASC
2、ORDER BY "province" DESC