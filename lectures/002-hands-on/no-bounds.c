
int main()
{
  int ar[3] = { 3, 2, 5 };
  ar[2] = 20; // OK accessing last element (3rd element)
  ar[4] = 20; // BAD! accessing 5th element in a 3 element array -> out of bounds
}
