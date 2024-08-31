package deeplink

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Cocoa
#import <Foundation/Foundation.h>
#import "CustomProtocol.h"

#include <stdlib.h>
*/
import "C"
import "sync"

const bufferSize = 256

var (
	openURLChan chan string
	once        sync.Once
)

func Handler() <-chan string {
	once.Do(func() {
		openURLChan = make(chan string, bufferSize)
		C.StartCustomProtocolHandler()
	})
	return openURLChan
}

//export HandleCustomProtocol
func HandleCustomProtocol(url *C.char) {
	goUrl := C.GoString(url)
	openURLChan <- goUrl
}
