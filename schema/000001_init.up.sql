CREATE TABLE IF NOT EXISTS link
(
    id INTEGER PRIMARY KEY,
    short TEXT, -- Короткие ссылки должны основываться на id записи(sequence) в БД, переведённой в систему счисления с алфавитом [A-Za-z0-9]
    long TEXT
)
