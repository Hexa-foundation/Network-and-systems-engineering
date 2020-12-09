#if-statement in TCl
#argv 0 takes the very 1st argument
#argc takes the argument count
set x [lindex $argv 0 ]

if { $x == 1 } {
	puts "condition 1 you entered is $x"
	
} elseif { $x == 2 } {
	puts "condition 2 you entered: $x"
	
} else {
	puts "$x is not a valid argument"
}   
