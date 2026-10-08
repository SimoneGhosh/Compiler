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

- Bytes
    - Constants are constant expressions
    - refers to expressions whose vals don't change
    - Val is determined at compile time
    - why?
        - while ints can easily be encoded and put into bytecode directly and push onto stack
        - but for string literals this causes bloat
    - what does this do?
        - don’t need to run the program to know what these expressions evaluate to
        - compiler can find them in the code and store the value they evaluate to
        - reference constants in instructions
        - when there is integer literal while compiling, evaluate, keep strack of *obj.Int by storing in memory and giving number
        - after compiling and pass instructs to VM, put constants in constant pool
    - Explanation of Opcode def
        - Each definition will have an Op prefix and the value it refers to will be determined by iota
        - iota generate increasing byte values
    - what does Make function do?
        - makes bytecode ( create a single bytecode instruction that’s made up of an Opcode and an optional number of operands)
        1. finds how long resulting instruction will be
        2. Allocate byte slice of correct length (instructionLen)
        3. Does not use lookup
            - more usable function signature
            - can use Make to easily build up bytecode instructions without having to check for errors after every call
            - risk to produce empty bte slices
        4. allocate the instruction []byte and add the Opcode as its first byte – by casting it into one
        5. iterate over defined OperandWidths
        6. take the matching element from operands and put it in the instruction (switch)
        7. after encoding operand, increase offset by width

