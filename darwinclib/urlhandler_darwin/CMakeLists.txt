cmake_minimum_required(VERSION 3.5.0)
project(urlhandler_darwin LANGUAGES C)

set(CMAKE_C_COMPILER clang)
set(CMAKE_SHARED_LIBRARY_PREFIX "")
set(BUILD_SHARED_LIBS true)
#set(CMAKE_C_FLAGS_DEBUG "${CMAKE_C_FLAGS_DEBUG} -g")

set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ../../../../../lib)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ../../../../../lib)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ../../../../../lib)

add_library(
    urlhandler_darwin 
    urlhandler_darwin.c
    )
target_link_libraries(urlhandler_darwin PRIVATE "-framework CoreServices")