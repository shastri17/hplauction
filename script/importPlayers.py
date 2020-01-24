import xlrd
import MySQLdb

# Open the workbook and define the worksheet
book = xlrd.open_workbook("players.xls")
sheet = book.sheet_by_name("HPL-2019")

# Establish a MySQL connection
database = MySQLdb.connect (host="localhost", user = "root", passwd = "Root@123", db = "hpl_auction")

# Get the cursor, which is used to traverse the database, line by line
cursor = database.cursor()

# Create the INSERT INTO sql query
query = """INSERT INTO player (name, nick_name, skill_area, batting_hand, bowling_hand, mobile_number, whatsapp_number, previously_played_teams, image) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s)"""

# Create a For loop to iterate through each row in the XLS file, starting at row 2 to skip the headers
for r in range(1, sheet.nrows):
		name		= sheet.cell(r,0).value
		nick_name	= sheet.cell(r,1).value
		skill_area  = sheet.cell(r,2).value
		batting_hand        = sheet.cell(r,3).value
		bowling_hand		= sheet.cell(r,4).value
		mobile_number		= sheet.cell(r,5).value
		whatsapp_number	= sheet.cell(r,6).value
		previously_played_teams		= sheet.cell(r,7).value
		image		= sheet.cell(r,8).value

		# Assign values from each row
		values = (name, nick_name, skill_area, batting_hand, bowling_hand, mobile_number, whatsapp_number, previously_played_teams, image)

		# Execute sql Query
		cursor.execute(query, values)

# Close the cursor
cursor.close()

# Commit the transaction
database.commit()

# Close the database connection
database.close()

# Print results
print ""
print "All Done! Bye, for now."
print ""
columns = str(sheet.ncols)
rows = str(sheet.nrows)