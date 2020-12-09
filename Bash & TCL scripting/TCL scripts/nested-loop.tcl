#argv 0 takes the very 1st argument
#argc takes the argument count
#lindex $argv 0 takes the 1st user argument
if {$argc == 2} {
	set x [lindex $argv 0]
	set y [lindex $argv 1]
	
	puts "Beginning the while loop"
	for {set i $x} {$i <= $y} {incr i} { puts $i }
} else {
	puts "Invalid n.o of arguments"
}

#Terminal 
#tclsh8.6 practice.tcl 5 10
#result:5 6 7 8 9 10
