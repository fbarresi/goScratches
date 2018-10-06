import random
import multiprocessing
from multiprocessing import Pool


#term of Leibniz Series (pi/4 = 1 -1/3 +1/5 -1/7 ...)
def term(dict):
	approx = 0.0
	start = dict['start']
	end = start+dict['iterations']
	for x in range(int(start), int(end)):
		approx+=(4. * pow(-1, x) / (2.*x + 1))
	return approx

if __name__=='__main__':
    
    np = multiprocessing.cpu_count()
    #np=1
    print('You have {0:1d} CPUs'.format(np))
    # Nummber of points to use for the Pi estimation
    n = 1000000
    
    # iterable with a list of terms to calculate in each worker
    part_count=[{'start': float(i*(n/np)),'iterations': float(n/np)} for i in range(np)]

    #Create the worker pool
    pool = Pool(processes=np)   

    # parallel map
    count=pool.map(term, part_count)

    print("Esitmated value of Pi: "+ str(sum(count)))

