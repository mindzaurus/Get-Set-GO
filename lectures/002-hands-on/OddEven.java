

public class OddEven { // JAVA PROGRAM
	public static void main(String args[]) {
		for (int j=0; j<5; j++)
		{
			int even_counter = 0, odd_counter = 0, i = 0;
			for (i=0; i<50000001; i++) {
				if (i % 2 == 0 )
					even_counter += 1;
				else
					odd_counter += 1;
			}

			System.out.println("even_counter "+even_counter);
			System.out.println("odd_counter "+odd_counter);
		}
	}

}
