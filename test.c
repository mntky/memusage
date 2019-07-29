#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>

int fd;
static char buf[2048];

int main() {
	char fname[] = "/proc/meminfo";
	fd = open(fname, O_RDONLY);
	if (fd == -1) {
		printf("[ERR]file open error\n");
	}
	printf("%p\n", &buf);
	int size = read(fd, &buf, sizeof buf - 1);
	if (size == -1) {
		printf("[ERR]read error\n");
	}
	printf("---size---\n%d\n---data---\n%s\n",size, buf);
	
	return 0;
}

