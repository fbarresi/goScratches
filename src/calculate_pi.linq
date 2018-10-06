<Query Kind="Program" />

void Main()
{
	"Estimate Pi".Dump();
		
	var pi = 0d;
	for (int i = 0; i < 1000000; i++)
	{
		CalculateTerm(ref pi, i);
	}
	pi.Dump();
}

// Define other methods and classes here
private void CalculateTerm(ref double pi, int i)
{
	pi += 4*(Math.Pow(-1, i)/(2*i+1));
}