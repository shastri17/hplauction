import MySQLdb
import csv

user = 'root' # your username
passwd = 'Root@123' # your password
host = 'localhost' # your host
db = 'hpl_auction' # database where your table is stored
table = 'player' # table you want to save

con = MySQLdb.connect(user=user, passwd=passwd, host=host, db=db)
cursor = con.cursor()

for x in range(10):
    team_id=x+1
    query = "SELECT name,whatsapp_number FROM %s where team_id= %d;" % (table,team_id)
    cursor.execute(query)

    with open('outfile_%d.csv'%team_id,'w') as f:
        writer = csv.writer(f)
        for row in cursor.fetchall():
            writer.writerow(row)