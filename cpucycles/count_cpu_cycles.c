#define _GNU_SOURCE

#include <stdio.h>
#include <stdlib.h>
#include <sched.h>


inline volatile long long RDTSC() {
   register long long TSC asm("eax");
   asm volatile (".byte 15, 49" : : : "eax", "edx");
   return TSC;
}

int main(void) {
	cpu_set_t cpu;
	CPU_ZERO(&cpu);
	CPU_SET(1, &cpu);

	if (sched_setaffinity(0, sizeof(cpu), &cpu) < 0) {
		perror("sched_setaffinity");
		exit(1);
	}

	long long start = RDTSC();
	// here goes the code
	long long end = RDTSC();
	printf("%llu clock cycles\n",  end - start);

	return 0;
}
