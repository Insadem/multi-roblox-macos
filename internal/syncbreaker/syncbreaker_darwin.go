package syncbreaker

/*
#include <semaphore.h>
#include <stdio.h>

int destroy_semaphore() {
    const char *sem_name = "/RobloxPlayerUniq";
    if (sem_unlink(sem_name) == -1)  // Attempt to destroy the semaphore
    {
        return 1;  // failed to destroy semaphore
    }
    return 0;
}
*/
import "C"

func Break() bool {
	return C.destroy_semaphore() == 0
}
