
# PYTHON PROGRAM

for j in range(0, 5) :
	even_counter = 0
	odd_counter = 0

	for i in range(0, 50000001):
		if i % 2 == 0 :
			even_counter += 1
		else :
			odd_counter += 1

	print ("even_counter "+ str(even_counter))
	print ("odd_counter "+ str(odd_counter))
