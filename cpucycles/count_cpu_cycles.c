#include <stdio.h>
#include <stdlib.h>

inline volatile long long RDTSC() {
   register long long TSC asm("eax");
   asm volatile (".byte 15, 49" : : : "eax", "edx");
   return TSC;
}

int main(void) {
	unsigned long mask = 1;

	if (sched_setaffinity(0, sizeof(mask), &mask) < 0) {
		perror("sched_setaffinity");
		exit(1);
	}

	long long start = RDTSC();
	// here goes the code
	long long end = RDTSC();
	printf("%lld clock cycles\n",  end - start);

	return 0;
}
