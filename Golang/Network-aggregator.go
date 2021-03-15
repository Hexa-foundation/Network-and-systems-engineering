package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"sort"
	"strings"
	"time"
	"unsafe"
)

type subnet struct {
	network   uint32
	broadcast uint32
	mask      uint32
	cidr      uint8
	length    uint32
}

type SubnetList struct {
	Head *SubnetList
	Count int
	Next *SubnetList
	Prev *SubnetList
	Subnet *subnet
}

// Print the data from the subnet structure nicely
func (s *subnet) String() string {
	bytes := (*[4]byte)(unsafe.Pointer(&s.network))
	return fmt.Sprintf("%d.%d.%d.%d/%d", bytes[3], bytes[2], bytes[1], bytes[0], s.cidr)
}

// Some constants for masking with
const (
	MASK8  = 0x000000FF
	MASK16 = 0x0000FF00
	MASK24 = 0x00FF0000
	MASK32 = 0xFF000000
	Full32 = 0xFFFFFFFF
)

// swap32 is an internal function to swap 32bit Uints
func swap32(in uint32) uint32 {
	return ((in & MASK8) << 24) | ((in & MASK16) << 8) | ((in & MASK24) >> 8) | ((in & MASK32) >> 24)
}

// bitCount counts the number of unset lower order bits on a uint32
func bitCount(in uint32) uint8 {
	Count := uint8(32)
	for bc := 32; bc > 0; bc-- {
		if (in>>(32-bc))&1 == 1 {
			break
		}
		Count--
	}
	return Count
}

// IPNetworks parses a list of ip/cidr strings into a sorted double linked list of structures we can work with
func IPNetworks(in []string) *SubnetList {
	var networks = make([]*subnet, 0)
	for ln, x := range in {
		ipAddr, ipNet, err := net.ParseCIDR(x)
		if len(ipAddr) != 16 {
			fmt.Printf("Could not parse ip/cidr on line %d, skipping\n",ln)
			continue
		}
		if *(*uint16)(unsafe.Pointer(&ipAddr[10])) != 0xFFFF || *(*uint32)(unsafe.Pointer(&ipAddr[12])) == 0 {
			fmt.Printf("Invalid ip/cidr on line %d, skipping\n", ln+1)
			continue
		}
		if err != nil {
			fmt.Printf("Invalid ip/cidr on line %d, skipping\n", ln+1)
			continue
		}
		sn := new(subnet)
		sn.network = swap32(*(*uint32)(unsafe.Pointer(&ipAddr[12]))) & swap32(*(*uint32)(unsafe.Pointer(&ipNet.Mask[0])))
		sn.mask = swap32(*(*uint32)(unsafe.Pointer(&ipNet.Mask[0])))
		sn.length = (Full32 &^ sn.mask) + 1
		sn.broadcast = sn.network + sn.length
		sn.cidr = bitCount(sn.mask)
		networks = append(networks, sn)
	}
	sort.Slice(networks, func(i, j int) bool {
		if networks[i].network != networks[j].network {
			return networks[i].network < networks[j].network
		}
		return networks[i].cidr < networks[j].cidr
	})
	res := new(SubnetList)
	res.Head = res
	head := res.Head
	for _,x := range networks {
		res.Subnet = x
		res.Next = new(SubnetList)
		res.Next.Prev,res.Next.Head,res.Subnet = res, res.Head,x
		res = res.Next
		head.Count++
	}
	if res.Subnet == nil {
		res = res.Prev
		res.Next = nil
		head.Count--
	}
	return res.Head
}

// Remove smalls will remove any subnets covered by larger networks as well as any duplicates from a slice of
// subnet structures
func RemoveSmalls(sl *SubnetList) {
	for {
		if sl.Next == nil {
			return
		}
		sl1, sl2 := sl.Subnet, sl.Next.Subnet
		if sl2.network&sl1.mask == sl1.network {
			sl.Next = sl.Next.Next
			sl.Next.Prev = sl.Next
			sl.Head.Count--
		}
		if sl.Next != nil {
			RemoveSmalls(sl.Next)
		}
		return
	}
}

// AggregateNetworks recursively aggregates a linked list of networks
func AggregateNetworks(start *SubnetList) {
	sl := start
	for {
		if sl.Next == nil {
			return
		}
		s1, s2 := sl.Subnet, sl.Next.Subnet
		larger := s1.network & (Full32 << (32 - (s1.cidr - 1)))
		lb := larger+(s1.length*2)
		next := s2.network & (Full32 << (32 - (s1.cidr - 1)))
		if larger == s1.network && larger == next && lb <= s2.broadcast {
			s1.cidr--
			s1.mask = Full32 << (32 - (s1.cidr))
			s1.length *= 2
			s1.broadcast = s1.network + s1.length
			sl.Next = sl.Next.Next
			sl.Head.Count--
		}
		if sl.Next != nil {
			AggregateNetworks(sl.Next)
		}
		return
	}
}

func ReadSubnetFile(f string) ([]string, error) {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func main() {
	var outFile *os.File
	var err error
	var outputToFile bool
	input := flag.String("input", "", "inputfilename.txt")
	output := flag.String("output", "", "outputfilename.txt")
	overwrite := flag.Bool("overwrite", false, "if true, overwrite output file")
	baseStart := time.Now()

	flag.Parse()

	if *input == "" {
		fmt.Printf("Input file not specified\n")
		flag.PrintDefaults()
		return
	}
	if *output != "" {
		outputToFile = true
		if _, err = os.Stat(*output); err != nil {
			if os.IsNotExist(err) {
				if outFile, err = os.Create(*output); err != nil {
					fmt.Printf("Could not create output file %s: %v\n", *output, err)
					return
				}
			}
		} else {
			if !*overwrite {
				fmt.Printf("Output file exists and -overwrite not set to true\n")
				flag.PrintDefaults()
				return
			}
			outFile, err = os.OpenFile(*output, os.O_WRONLY|os.O_TRUNC, 0755)
			if err != nil {
				fmt.Printf("Failed to open and truncate output file %s: %v\n", *output, err)
				return
			}
		}
		fmt.Printf("Successfully opened output file...\n")
	}
	subnets, e := ReadSubnetFile(*input)
	inputLen := len(subnets)
	if e != nil {
		fmt.Printf("Could not read input file %s\n", *input)
		flag.PrintDefaults()
		return
	}
	fmt.Printf("Read %d input lines from input file\n", inputLen)
	if Networks := IPNetworks(subnets); Networks == nil {
		fmt.Printf("No valid networks found in input file\n")
		return
	} else {
		origLen := Networks.Head.Count
		fmt.Printf("Found %d valid IP networks, removing duplicates and covered subnets...\n", Networks.Head.Count)
		start := time.Now()
		RemoveSmalls(Networks)
		fmt.Printf("Removed %d covered and/or duplicate networks in %s\n",origLen-Networks.Head.Count, time.Since(start))
		start = time.Now()
		fmt.Printf("Starting aggregation...\n")
		for {
			origLen = Networks.Head.Count
			AggregateNetworks(Networks.Head)
			RemoveSmalls(Networks.Head)
			if Networks.Head.Count == origLen {
				fmt.Printf("Breaking, new length is %d\n",Networks.Head.Count)
				break
			}
		}
		fmt.Printf("Completed aggregation in %s\n",time.Since(start))
		if outputToFile {
			for sub := Networks.Head; sub.Next != nil; sub = sub.Next {
				_, err := fmt.Fprintf(outFile, "%s\n", sub.Subnet)
				if err != nil {
					fmt.Printf("Failed to write to output file: %v\n", err)
					return
				}
			}
			fmt.Printf("All done....aggregated to %d networks\n",Networks.Head.Count)
			fmt.Printf("Finished in %s\n", time.Since(baseStart))
			return
		}
		fmt.Printf("Aggregated networks: \n")
		for sub := Networks.Head; sub.Next != nil; sub = sub.Next {
				fmt.Printf("%s\n", sub.Subnet)
		}
	}
}