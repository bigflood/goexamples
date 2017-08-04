cmake_minimum_required(VERSION 3.1)
project(mmlib VERSION 1.1.6 LANGUAGES CXX)

# set(CMAKE_BUILD_TYPE, "Release")

# if(NOT CMAKE_BUILD_TYPE)
#   set(CMAKE_BUILD_TYPE "Release" CACHE STRING
#       "Choose the type of build, options are: Debug Release RelWithDebInfo MinSizeRel." FORCE)
# endif(NOT CMAKE_BUILD_TYPE)

if (NOT CMAKE_BUILD_TYPE MATCHES Debug)
    add_definitions(-DNDEBUG)
endif()

# BUILD_SHARED_LIBS is a standard CMake variable, but we declare it here to make
# it prominent in the GUI.
option(BUILD_SHARED_LIBS "Build shared libraries(DLLs)." OFF)

# ADD_COMPILE_OPTIONS( -g -Wall )

add_library(mmlib STATIC "")
target_sources(mmlib
  PRIVATE
    "${PROJECT_SOURCE_DIR}/mmlib.cpp"
    "${PROJECT_BINARY_DIR}/mmlib.h"
)

target_compile_definitions(mmlib PRIVATE -DHAVE_CONFIG_H)
if(BUILD_SHARED_LIBS)
  set_target_properties(mmlib PROPERTIES WINDOWS_EXPORT_ALL_SYMBOLS ON)
endif(BUILD_SHARED_LIBS)
