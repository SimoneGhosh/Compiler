# Compiler
Bytecode Compiler &amp; Virtual Machine

Architecture of virtual machine
- Why stack machine?
    - easier to understand and build than register machine
    - register machine being faster does not play big role
    - right now just have to do stack arithmetic 

- How does stack machine work, example 1 + 2?
    - translate expressions to bytecode instructions that use a stack
    1. push operands 1 and 2 on stack
    2. add
    3. push result on stack
    - this means 2 instructions types
        1. pushing onto stack
        2. adding things on stack

- Implementation of stack
    - define opcodes and how they are encoded in bytecode
    - VM decodes and execute 
    - intructions to tell VM to push on to stack