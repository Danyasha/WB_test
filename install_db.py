import sqlite3
import os

if os.path.isfile("tv_storage.db"):
    print("File exist")
    exit(0)
conn = sqlite3.connect("tv_storage.db")
cursor = conn.cursor()
cursor.execute("""CREATE TABLE tvs
                  (id INTEGER PRIMARY KEY, 
                  brand CHAR DEFAULT NULL,
                  manufacturer CHAR NOT NULL,
                  model CHAR NOT NULL,
                  model_year TEXT NULL
                  )
               """)
print("DB created!")
tvs =  [(1, "Sony", "China", "SomeSonyTV", "2010-12-22"),
          (2, "Panasonic", "China", "PanasonicTV", "2022-6-14"),
          (3, "Sony", "China", "OneMoreSomeSonyTV", "2018-10-23"),
          (4, " ", "China", "JustTV", "2012-4-22"),]
cursor.executemany("INSERT INTO tvs VALUES (?,?,?,?,?)", tvs)
conn.commit()
select = "SELECT * FROM tvs"
cursor.execute(select)
print(cursor.fetchall())
conn.close()
exit()