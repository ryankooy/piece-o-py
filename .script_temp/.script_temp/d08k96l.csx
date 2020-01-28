using System;
using System.Text;
  
class Number
{
  static void Main(string[] args)
  {
    Program generator = new Program();
    int rand = generator.RandNum(5, 100);
    Console.WriteLine($"Random number between 5 and 100 is {rand}.");
    
    string str = generator.Randomness(10, false);
    Console.WriteLine($"Random string of 10 characters is {str}.");
    
    string pass = generator.RandomPW();
    Console.WriteLine($"Random string of 6 characters is {pass}.");
    
    Console.ReadKey();
  }
}

  public class Program
{
  public int RandNum(int min, int max)
  {
    Random random = new Random();
    return random.Next(min, max);
  }
  
  public string Randomness(int size, bool lowerCase)
  {
    StringBuilder builder = new StringBuilder();
    Random random = new Random();
    char ch;
    
    for (int i = 0; i < size; i++)
    {
      ch = Convert.ToChar(Convert.ToInt32(Math.Floor(26 * random.NextDouble() + 65)));
      
      builder.Append(ch);
    }
    
    if (lowerCase)
      return builder.ToString().ToLower();
    
    return builder.ToString();
  }
  
  public string RandomPW()
  {
    StringBuilder builder = new StringBuilder();
    builder.Append(Randomness(4, true));
    builder.Append(RandNum(1000, 9999));
    builder.Append(Randomness(2, false));
    return builder.ToString();
  }
}
