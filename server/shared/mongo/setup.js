db.account.createIndex(
  {
    "account.email": 1, 
  },
  {
    unique: true,
  }
);

db.account.createIndex(
  {
    "account.username": 1, 
  },
  {
    unique: true,
  }
);