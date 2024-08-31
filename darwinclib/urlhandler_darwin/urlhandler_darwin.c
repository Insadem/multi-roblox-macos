#include <CoreServices/CoreServices.h>
#include <stdio.h>

int set(char *bundleIdentifier, char *urlScheme) {
    OSStatus status = LSSetDefaultHandlerForURLScheme(
        CFStringCreateWithCString(NULL, urlScheme, kCFStringEncodingUTF8),
        CFStringCreateWithCString(NULL, bundleIdentifier,
                                  kCFStringEncodingUTF8));
    if (status == noErr) {
        return 0;
    } else {
        return 1;
    }
}

int check(char *bundleIdentifier, char *urlScheme) {
    CFStringRef currentHandler = LSCopyDefaultHandlerForURLScheme(
        CFStringCreateWithCString(NULL, urlScheme, kCFStringEncodingUTF8));

    if (!currentHandler) {
        return 1;
    }

    Boolean isDefault =
        CFStringCompare(currentHandler,
                        CFStringCreateWithCString(NULL, bundleIdentifier,
                                                  kCFStringEncodingUTF8),
                        0) == kCFCompareEqualTo;
    CFRelease(currentHandler);

    if (isDefault) {
        return 0;
    } else {
        return 1;
    }
}

// check("roblox-player", "com.roblox.RobloxPlayer");
// set("roblox", "com.roblox.RobloxPlayer"); roblox is valid deeplink too
