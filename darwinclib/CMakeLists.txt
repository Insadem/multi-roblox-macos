cmake_minimum_required(VERSION 3.5.0)
project(global LANGUAGES C)

set(CMAKE_OSX_ARCHITECTURES "x86_64;arm64")
set(CMAKE_OSX_DEPLOYMENT_TARGET "11.0")
add_subdirectory(./syncbreaker_darwin)
add_subdirectory(./urlhandler_darwin)