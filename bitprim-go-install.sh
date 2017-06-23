#!/bin/bash

# ----------------------------------------------
# Bitprim Go (Ubuntu, GCC)
# ----------------------------------------------

git clone -b c-api https://github.com/bitprim/secp256k1.git
git clone -b c-api https://github.com/bitprim/bitprim-core.git
git clone -b c-api https://github.com/bitprim/bitprim-consensus.git
git clone -b c-api https://github.com/bitprim/bitprim-database.git
git clone -b c-api https://github.com/bitprim/bitprim-blockchain.git
git clone -b c-api https://github.com/bitprim/bitprim-network.git
git clone -b c-api https://github.com/bitprim/bitprim-node.git
git clone https://github.com/bitprim/bitprim-node-cint.git
    

# ----------------------------------------------
    cd secp256k1
    mkdir build
    cd build
    cmake -DENABLE_MODULE_RECOVERY=ON -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    # cmake -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Debug ..
    make -j4
    cd ../..
# ----------------------------------------------

# ----------------------------------------------
    cd bitprim-core
    mkdir build
    cd build
    cmake -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    # cmake -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Debug ..

    make -j4
    cd ../..
# ----------------------------------------------

# ----------------------------------------------
    cd bitprim-consensus
    mkdir build
    cd build
    cmake -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    # cmake -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Debug ..

    make -j4
    cd ../..
# ----------------------------------------------

# ----------------------------------------------
    cd bitprim-database
    mkdir build
    cd build
    cmake -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    # cmake -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Debug ..

    make -j4
    cd ../..
# ----------------------------------------------

# ----------------------------------------------
    cd bitprim-blockchain
    mkdir build
    cd build
    cmake -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    # cmake -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Debug ..

    make -j4
    cd ../..
# ----------------------------------------------

# ----------------------------------------------
    cd bitprim-network
    mkdir build
    cd build
    cmake -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    # cmake -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Debug ..

    make -j4
    cd ../..
# ----------------------------------------------

# ----------------------------------------------
    cd bitprim-node
    mkdir build
    cd build
    cmake -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    # cmake -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Debug ..

    make -j4
    cd ../..
# ----------------------------------------------

# ----------------------------------------------
    # master branch
    git clone https://github.com/bitprim/bitprim-node-cint.git
    cd bitprim-node-cint
    mkdir build
    cd build
    cmake -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Release ..
    # cmake -DWITH_TESTS=OFF -DCMAKE_BUILD_TYPE=Debug ..

    make -j4
    cd ../..
# ----------------------------------------------

