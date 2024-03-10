#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv){
	printf("我收到了 %d 個參數\n",argc);
	for (int i = 0;i<argc;i++){
		printf("%s\n",argv[i]);
	}
	return 0;
}
