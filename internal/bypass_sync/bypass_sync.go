//go:build darwin

package bypass_sync

// #include <stdio.h>
// #include <semaphore.h>
// int destroy_semaphore()
// {
//     const char *sem_name = "/RobloxPlayerUniq";
//     if (sem_unlink(sem_name) == -1) // Attempt to destroy the semaphore
//     {
//         return 1;
//     }
//     return 0;
// }
import "C"
import "errors"

func BypassSync() error {
	num := int(C.destroy_semaphore())
	if num != 0 {
		return errors.New("couldn't destroy semaphore")
	}
	return nil
}
