<Query Kind="Program">
  <Namespace>System.Threading.Tasks</Namespace>
</Query>

void Main()
{
	"Estimate Pi".Dump();
		
	var tasks = new List<Task<double>>();
	for (int i = 0; i < 1000000; i++)
	{
		tasks.Add(CalculateTerm(i));
	}
	var taskArray = tasks.ToArray();
	Task.WaitAll(taskArray);
	var pi = taskArray.Select(t=>t.Result).Sum().Dump();
	
}

// Define other methods and classes here
private Task<double> CalculateTerm(int i)
{
	return Task.Run(()=>4*(Math.Pow(-1, i)/(2*i+1)));
}