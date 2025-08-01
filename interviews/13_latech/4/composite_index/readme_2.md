А индекс будет применен только на операцию where? На какие еще операции применяется индекс?

Нет, **индексы используются не только в `WHERE`** — они также применяются в ряде других операций в SQL. Ниже — список и пояснение:

---

### 📌 Где используется индекс в SQL?

#### 1. **`WHERE`** — фильтрация строк

Самый очевидный и распространённый случай:

```sql
SELECT * FROM users WHERE email = 'user@example.com';
```

#### 2. **`JOIN`** — соединение таблиц

Если `JOIN` происходит по колонке с индексом, то индекс может быть использован:

```sql
SELECT * FROM orders
JOIN customers ON orders.customer_id = customers.id;
```

Если `customers.id` — проиндексирован, индекс может использоваться для быстрого поиска.

#### 3. **`ORDER BY`** — сортировка

Если порядок сортировки совпадает с порядком в индексе, его можно использовать для **избежания сортировки вручную**:

```sql
SELECT * FROM products ORDER BY created_at DESC;
```

Если `created_at` индексирован — сортировка может быть "бесплатной".

> ⚠️ Однако: если ты сортируешь **в обратном порядке** по нескольким колонкам или не с самого левого элемента индекса, индекс может быть проигнорирован.

#### 4. **`GROUP BY`**

При группировке по индексированным колонкам (особенно если `GROUP BY` = `ORDER BY`), индекс может использоваться:

```sql
SELECT category_id, COUNT(*) FROM products GROUP BY category_id;
```

#### 5. **`SELECT DISTINCT`**

Похоже на `GROUP BY`. Если ты делаешь `SELECT DISTINCT` по индексированной колонке, индекс может помочь.

#### 6. **`IN (...)`**

Работает аналогично `WHERE col = value`:

```sql
SELECT * FROM users WHERE id IN (1, 2, 3);
```

#### 7. **`EXISTS` / `NOT EXISTS`**

Если подзапрос использует фильтрацию по индексированной колонке, индекс будет задействован.

#### 8. **`LIMIT` + `ORDER BY`**

Комбинация может очень эффективно работать с индексом:

```sql
SELECT * FROM logs ORDER BY created_at DESC LIMIT 100;
```

→ индекс по `created_at DESC` позволит быстро выдать первые 100 записей.

---

### 💡 Пример: составной индекс `(a, b)` может использоваться при

| Запрос                  | Индекс используется?      |
| ----------------------- | ------------------------- |
| `WHERE a = ?`           | ✅ Да                      |
| `WHERE a = ? AND b = ?` | ✅ Да                      |
| `WHERE b = ?`           | ❌ Нет (если `a` пропущен) |
| `ORDER BY a, b`         | ✅ Да                      |
| `ORDER BY b`            | ❌ Нет                     |

---

### 🧪 Как проверить?

`EXPLAIN` и `EXPLAIN ANALYZE` в PostgreSQL / MySQL покажут, **используется ли индекс** и **в каком виде** (`Index Scan`, `Index Only Scan`, `Bitmap Index Scan`, `Using Index`, и т.д.).

---

Если хочешь, могу разобрать твой конкретный запрос и сказать, какие индексы он будет использовать.
