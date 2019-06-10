import sqlite3
conn = sqlite3.connect("tv_storage.db")
cursor = conn.cursor()
select = "SELECT * FROM tvs"
cursor.execute(select)
temp = cursor.fetchall()
for i in temp:
    print(i)
conn.close()