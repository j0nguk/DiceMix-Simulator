package main

import (
	"fmt"
	"strconv"
	"rand"
)

const NRELAYS = 4

func main() {
//	pTable := []int{3,1,4,15,9,2,6,5,35,8,97,93,23,84,62,64,33,83,27,950,28,841,971,69,39,937,510,58,20,974,94,45,92,30,78,16}
/*406286208998628034825342117067982148086513282306647093844609550582231725359408128481117450284102701938521105559644622948954930381964428810975665933446128475648233786783165271201909145648566923460348610454326648213393607260249141273724587006606315588174881520920962829254091715364367892590360011330530548820466521384146951941511609 */
/*	msg := "Hello";
	msg2 := permutation(msg, &pTable);
	fmt.Println(msg2);
	msg3 := reversePermutation(msg, &pTable);
	fmt.Println(msg3);*/
// Sender
	sout := make(chan string)
	go sender(sout)
// Relays
	var rin [NRELAYS]chan string
	var rout [NRELAYS]chan string
	for i:=0; i<NRELAYS; i++ {	
		rin[i] = make(chan string)
		rout[i] = make(chan string)
		go relay(rin[i], rout[i])
	}
// Messages
	for i:=0; i<10; i++ {
		msg := <- sout
		dest := rand.Intn(NRELAYS)
		rin[dest] <- msg
		fmt.Println("Sender sent ", msg, " to relay #" , dest);
		for j:=0; j<NRELAYS; j++ {
			msg = <- rout[j]
			dest = rand.Intn(NRELAYS)
			rin[dest] <- msg
			fmt.Println("Relay #", j, " got a ", msg)
		}
	}
}

func sender(outgoing chan string) {
	for i:=0; ; i++ {
		outgoing <- "msg #" + strconv.Itoa(i)
	}
}

func relay(incoming chan string, outgoing chan string) {
	for {
		candidate := <-incoming
		outgoing <- candidate
	}
}


func permutation(msg string) string {
	return "hello";
}

func reversePermutation(msg string) string {
	return "hello";
}


