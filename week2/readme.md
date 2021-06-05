#### Q: 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

#### A: dao 通常不用来处理业务逻辑, 因此应当将错误 Wrap  抛给上层处理, 代码如下:

```golang
func QueryRow(id int) (*SqlRet, error) {
    var res *SqlRet
    if err := db.QueryRow("select row from table where id = ?", id).Scan(res); err != nil {
        return nil, errors.Wrap(err, fmt.Sprintf("QueryRow error: id: %d", id))
    }
    return res, nil
}
```
