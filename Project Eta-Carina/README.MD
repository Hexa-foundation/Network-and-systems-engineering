## ... development project

The first step in creating our own operating system kernel is to create a Rust executable that does not link the standard library. This makes it possible to run Rust code on the bare metal without an underlying operating system.


Steps / prerequisites


- Set up your version control system. Writing an operating system means you'll likely be creating hundreds (or thousands!) of lines of code. When you're working on revisions,

- Bare Bones is a tutorial that helps you write your first simple 32-bit kernel. After completing the tutorial, the next step is to use Meaty Skeleton to structure your own operating system.

- Linux from scratch.

- Operating_Systems_From_0_to_1 book.

- Decide on a bootloader.

- Choose a kernel design.

## (a). Kernel developement(monolithic/micro-kernel design)

![](https://github.com/Ruby-Raptor/Network-and-systems-engineering/blob/main/Project%20Eta-Carina/Extras/1200px-OS-structure.svg.png)

Advantages of Microkernel
Here, are the pros/benefits of using Microkernel

- Microkernel architecture is small and isolated therefore it can function better.
- Microkernels are secure because only those components are included that disrupt the functionality of the system otherwise.
- The expansion of the system is more accessible, so it can be added to the system application without disturbing the Kernel.
- Microkernels are modular, and the different modules can be replaced, reloaded, modified without even touching the Kernel.
- Fewer system crashes when compared with monolithic systems.
- Microkernel interface helps you to enforce a more modular system structure.
  Without recompiling, add new features.
- Server malfunction is also isolated as any other user program's malfunction.
- Microkernel system is flexible, so different strategies and APIs, implemented by different servers, which can coexist in the system.
- Increased security and stability will result in a decreased amount of code which runs on kernel mode.

Disadvantage of Microkernel
Here, are drawback/cons of using Microkernel:

- Providing services in a microkernel system are expensive compared to the normal monolithic system.
- Context switch or a function call needed when the drivers are implemented as procedures or processes, respectively.
- The performance of a microkernel system can be indifferent and may lead to some problems.

You need to learn how to write the following:

1. INTERRUPTS.
2. MEMORY MANAGEMENT.
3. MULTITASKING.
4. DEVICE MANAGEMENT.

**N/B:** for bootloaders / firmware and kernels, Rust should not link to **std**, so we use **#![no_std]**

(b). System programs
