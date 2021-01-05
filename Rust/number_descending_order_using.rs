//Two part code
//Part one: Numbers in descending order with a funtion
//the sort funtion arranges in ascending order
//this code arranges it in descending order

use std::cmp::Ordering;

fn main()
{
let mut arr = [3,4,4,7,6,3,2,43,9];

fn desc(a:&i32, b:&i32) -> Ordering{ 
	if a < b { Ordering::Greater}
	
	else if a > b {Ordering::Less}
	else {Ordering::Equal}
}
arr.sort_by(desc);
print!("{:?}",arr);
}


//Part two:number descending order but with closures

use std::cmp::Ordering;

fn main()
{
	let mut arr = [4,8,1,10,0,45,12,7];
	
	
	//used a closure instead of a funtion
	let desc = |a:&i32, b:&i32| -> Ordering {
		if a < b {Ordering::Greater}
		
		else if a > b {Ordering::Less}
		
		else { Ordering::Equal }
		
};	
	arr.sort_by(desc);
	print!("{:?}",arr);
	
}
