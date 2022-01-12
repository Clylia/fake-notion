// mongo 数据库创建索引
db.account.createIndex(
  {
    id: 1, // 这里的 1 表示从小到大的意思
  },
  {
    unique: true,
  }
);
