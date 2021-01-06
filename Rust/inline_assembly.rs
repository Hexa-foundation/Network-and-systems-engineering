//This will add 5 to the input in variable i 
//and write the result to variable o.

//The particular way this assembly does this is first copying the value from i to the output, and then adding 5 to it.

//an error will be generated
//error:error[E0658]: use of unstable library feature 'asm': inline assembly is not stable enough for use and is subject to change
fn main()
{
let i: u64 = 3;
let o: u64;
unsafe {
    asm!(
        "mov {0}, {1}",
        "add {0}, {number}",
        out(reg) o,
        in(reg) i,
        number = const 5,
    );
}
assert_eq!(o, 8);
}
