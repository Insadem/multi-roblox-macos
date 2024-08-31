#include <semaphore.h>
#include <stdio.h>

int destroySemaphore() {
    const char *sem_name = "/RobloxPlayerUniq";
    if (sem_unlink(sem_name) == -1)  // Attempt to destroy the semaphore
    {
        return 1;  // failed to destroy semaphore
    }
    return 0;
}