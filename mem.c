#include <stdio.h>
#include <sys/sysinfo.h>


int main() {
	struct sysinfo info;

	sysinfo(&info);

	printf("%ld kb", info.freeram/1024);

	return 0;
}

