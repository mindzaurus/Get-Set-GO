
#include <stdio.h> // C PROGRAM

int main() {
	for (int j=0; j<5; j++)
	{
		int even_counter = 0, odd_counter = 0, i = 0;
		for (i=0; i<50000001; i++) {
			if (i % 2 == 0 )
				even_counter += 1;
			else
				odd_counter += 1;
		}

		printf("even_counter %d\n", even_counter);
		printf("odd_counter %d\n", odd_counter);
	}
}
