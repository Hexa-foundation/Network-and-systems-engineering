## This is a tutorial that I was following that teaches how to to write a simple kernel for 32-bit x86 and boot it. from:


[Bare bones](https://wiki.osdev.org/Bare_Bones)

The Netwide Assembler (NASM) is a popular assembler for the x86/x64 architecture.

**$ nasm -felf32 name.asm -o name.o** generates the executable. **-felf32** or **-felf64** depends on the target architecture.
