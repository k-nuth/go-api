from conans import ConanFile, CMake
# import os

class BitprimGoDeps(ConanFile):
    name = "bitprim-go-deps"
    version = "0.1"
    license = "http://www.boost.org/users/license.html"
    url = "https://github.com/bitprim/bitprim-go"
    description = "Bitcoin Full Node Library with Golang interface"
    settings = "os", "compiler", "build_type", "arch"
    options = {"shared": [True, False]}
    
    default_options = "shared=False"
    # default_options = "shared=True"

    generators = "cmake"
    # exports_sources = "src/*", "CMakeLists.txt", "cmake/*", "bitprim-node-cintConfig.cmake.in", "include/*", "test/*", "console/*"
    # package_files = "build/lbitprim-node-cint.so"
    # build_policy = "missing"

    # TODO(fernando): queda pendiente seleccionar el default Shared=False
    requires = (("bitprim-node-cint/0.2@bitprim/testing"))
    # default_options = "bitprim-node-cint:shared=True" #, "OpenSSL:shared=True"


    # conan install bitprim-node-cint/0.2@bitprim/testing -o shared=True
    # conan install -o Pkg:shared=True -o OtherPkg:option=value

    def imports(self):
        # self.copy("*.h", "./bitprim_c/include/bitprim", "include/bitprim")
        # self.copy("*.hpp", dst="./bitprim_c/include/bitprim", src="include/bitprim")
        # self.copy("*.dylib", dst="./bitprim_c/lib", src="lib")
        # self.copy("*.so", dst="./bitprim_c/lib", src="lib")
        # self.copy("*.lib", dst="./bitprim_c/lib", src="lib")
        # self.copy("*.dll", dst="./bitprim_c/lib", src="lib")
        self.copy("*.h", "./bitprim_c/include/bitprim", "include/bitprim")
        self.copy("*.hpp", dst="./bitprim_c/include/bitprim", src="include/bitprim")
        
        self.copy("*.lib", dst="./bitprim_c/lib", src="lib")
        # self.copy("*.a", dst="./bitprim_c/lib", src="lib")
        self.copy("*.dylib", dst="./bitprim_c/lib", src="lib")
        self.copy("*.so", dst="./bitprim_c/lib", src="lib")
        self.copy("*.dll", dst="./bitprim_c/lib", src="lib")
        self.copy("*.dll", dst="./bitprim_c/lib", src="bin")

        # self.copy("*.h", dst="/Users/fernando/fertest", src="include/bitprim")
        # self.copy("*.hpp", dst="/Users/fernando/fertest", src="include/bitprim")
        # self.copy("*.dylib", dst="/Users/fernando/fertest", src="lib")

    # def build(self):
    # def package(self):
    # def package_info(self):
