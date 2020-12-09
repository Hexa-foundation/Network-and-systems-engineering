#prompt user for a number

puts -nonewline "Enter n.o:"
flush stdout

#assign argument to variable
gets stdin value

#return double or error message
if {[catch {set doubled [expr $value * 2]} errmsg]} {
	puts "script failed - $errmsg"
} else {
	puts "$value doubled is: $doubled"
}
puts "Regardless of error script continues..."

#Result
#Enter n.o:50
#50 doubled is: 100
#Regardless of error script continues...
