
REM ----------------------------------------------
REM  Bitprim Go (Windows, MinGW_W64)
REM  ----------------------------------------------



REM git clone -b c-api https://github.com/bitprim/secp256k1.git
REM git clone -b c-api https://github.com/bitprim/bitprim-core.git
REM git clone -b c-api https://github.com/bitprim/bitprim-consensus.git
REM git clone -b c-api https://github.com/bitprim/bitprim-database.git
REM git clone -b c-api https://github.com/bitprim/bitprim-blockchain.git
REM git clone -b c-api https://github.com/bitprim/bitprim-network.git
REM git clone -b c-api https://github.com/bitprim/bitprim-node.git
REM git clone https://github.com/bitprim/bitprim-node-cint.git


REM Open MinGW_W64 Console - x86_64-7.1.0-posix-seh-rt_v5-rev0 -

REM  ----------------------------------------------
    cd secp256k1
    rd /s /q build
    mkdir build
    cd build
    cmake -G "MinGW Makefiles" -DENABLE_MODULE_RECOVERY=ON -DENABLE_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    REM cmake -G "MinGW Makefiles" -DENABLE_MODULE_RECOVERY=ON -DENABLE_TESTS=OFF  -DCMAKE_BUILD_TYPE=Debug ..

    mingw32-make.exe -j8
    cd ..\..
REM  ----------------------------------------------

REM  ----------------------------------------------
    cd bitprim-core
    rd /s /q build
    mkdir build
    cd build
    cmake -G "MinGW Makefiles" -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    REM  cmake -G "MinGW Makefiles" -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Debug ..

    mingw32-make.exe -j8
    cd ..\..
REM  ----------------------------------------------

REM  ----------------------------------------------
    cd bitprim-consensus
    rd /s /q build
    mkdir build
    cd build
    cmake -G "MinGW Makefiles" -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    REM  cmake -G "MinGW Makefiles" -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Debug ..

    mingw32-make.exe -j8
    cd ..\..
REM  ----------------------------------------------

REM  ----------------------------------------------
    cd bitprim-database
    rd /s /q build
    mkdir build
    cd build
    cmake -G "MinGW Makefiles" -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    REM  cmake -G "MinGW Makefiles" -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Debug ..

    mingw32-make.exe -j8
    cd ..\..
REM  ----------------------------------------------

REM  ----------------------------------------------
    cd bitprim-blockchain
    rd /s /q build
    mkdir build
    cd build
    cmake -G "MinGW Makefiles" -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    REM  cmake -G "MinGW Makefiles" -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Debug ..

    mingw32-make.exe -j8
    cd ..\..
REM  ----------------------------------------------

REM  ----------------------------------------------
    cd bitprim-network
    rd /s /q build
    mkdir build
    cd build
    cmake -G "MinGW Makefiles" -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    REM  cmake -G "MinGW Makefiles" -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Debug ..

    mingw32-make.exe -j8
    cd ..\..
REM  ----------------------------------------------

REM  ----------------------------------------------
    cd bitprim-node
    rd /s /q build
    mkdir build
    cd build
    cmake -G "MinGW Makefiles" -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    REM  cmake -G "MinGW Makefiles" -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Debug ..

    mingw32-make.exe -j8
    cd ..\..
REM  ----------------------------------------------

REM  ----------------------------------------------
    REM  master branch
    cd bitprim-node-cint
    rd /s /q build
    mkdir build
    cd build
    cmake -G "MinGW Makefiles" -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    REM  cmake -G "MinGW Makefiles" -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Debug ..

    mingw32-make.exe -j8
    cd ..\..
REM  ----------------------------------------------