import csv
import MySQLdb

mydb = MySQLdb.connect (host = "localhost",user = "root", passwd = "Root@123",db = "hpl_auction")
cursor = mydb.cursor()

csv_data=csv.reader(file("team2020.csv"))
for row in csv_data:
	print row
	cursor.execute('INSERT INTO team(username,password,team_name,owners_name, icon1,icon2)VALUES(%s,%s,%s,%s,%s,%s)',row)
#close the connection to the database.
mydb.commit()
cursor.close()
print "one"
