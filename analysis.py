import re
import os
import csv
import pandas as pd

data = []

for fn in os.listdir("./data"):
	yr = int(re.sub("[^0-9]","", fn))
	print("scanning year %d\n" % (yr))
	with open("./data/"+fn, 'rt') as f:
		reader = csv.reader(f)
		for row in reader:
			data.append(row+[yr])

print("done scanning files")

data = pd.DataFrame(data, columns=['name','gender','n','year'])    
data["n"] = data["n"].apply(lambda n: int(n))

census = {}
for yr in range(1880, 2020):
	census[yr] = data[data["year"]==yr]["n"].sum()

data["census"] = data["year"].apply(lambda yr: census[yr])
data["prop"] = data["n"] / data["census"]




print(data.head())

# get year where peak "n"
data[(data["name"]=="Lula")&(data["n"]==data[(data["name"]=="Lula")]["n"].max())]

# get year where peak "prop"
data[(data["name"]=="Lula")&(data["prop"]==data[(data["name"]=="Lula")]["prop"].max())]
