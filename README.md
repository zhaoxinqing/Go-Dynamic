# README

## Git

> git branch -m master master  
> git fetch origin  
> git branch -u origin/master master  
> git remote set-head origin -a

## 日志

有错误，立刻打印，然后回抛

---

## MySQL & ElasticSearch

- MySQL：擅长事务类型操作，可以确保数据的安全和一致性；
- ElasticSearch：擅长海量数据的搜索、分析、计算；

```json
// 查询
GET collectibles_period_per_minutes_1440/_search
{
  "query": {
    "match_all": {}
  },
  "sort": [
    {
      "created_at": {
        "order": "desc"
      }
    }
  ],
  "size":30
}
```

## 工作

- 事事有汇报
