package metrics

import (
	"encoding/csv"
	_ "fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

type UserId int32 // csv file of user data has ID go up to 100,000 - fit in int32

type UserMap map[UserId]*User

type UsersData struct {
	//	id       UserId // 4 bytes
	ages     []int8 // age will never exceed 150-200 - fit in int 8
	payments []uint32
}

type Address struct {
	zip         int32  // zip is 5 digits long - fit in int32
	fullAddress string // sequence of bytes, each char is 1 byte
}

type DollarAmount struct {
	dollars, cents uint32 // csv file shows payment amount can fit in uint 32
}

// Payment could be huge as `time.Time` is a struct containing a struct.
type Payment struct {
	amount DollarAmount
	time   time.Time
}

// This struct can be very big given the number of elements in `payments` slice.
type User struct {
	id       UserId // 4 bytes
	age      int8   // age will never exceed 150-200 - fit in int 8
	name     string // string
	address  Address
	payments []Payment
}

func AverageAge(users UsersData) float64 {
	//average, count := float32(0.0), float32(0.0)
	sum := uint64(0)
	//ages := make([]int8, 0)
	//var ages [100000]int8
	//for i, u := range users {
	//	ages[i] = u.age
	//}

	//fmt.Println("============================")
	//fmt.Printf("len(ages): %v\n", len(ages))
	//fmt.Println("============================")
	ages := users.ages

	//for _, u := range users {
	for _, a := range ages {
		//count += 1
		//average += (float32(u.age) - average) / count
		sum += uint64(a)
		//average += (float32(a) - average) / count
	}
	count := len(ages)
	//return float64(average)
	return float64(sum) / float64(count)
}

func AveragePaymentAmount(users UsersData) float64 {
	average, count := 0.0, 0.0
	payments := users.payments
	//for _, u := range users {
	//for _, p := range u.payments {
	for _, p := range payments {
		count += 1
		//	amount := float64(p.amount.dollars) + float64(p.amount.cents)/100
		amount := float64(p/100) + float64(p%100)/100
		average += (amount - average) / count
	}
	//}
	return average
}

// Compute the standard deviation of payment amounts
func StdDevPaymentAmount(users UsersData) float64 {
	mean := AveragePaymentAmount(users)
	squaredDiffs, count := 0.0, 0.0
	payments := users.payments
	//for _, u := range users {
	//for _, p := range u.payments {
	for _, p := range payments {
		count += 1
		//amount := float64(p.amount.dollars) + float64(p.amount.cents)/100
		amount := float64(p/100) + float64(p%100)/100
		diff := amount - mean
		squaredDiffs += diff * diff
	}
	//}
	return math.Sqrt(squaredDiffs / count)
}

//func LoadData() UserMap {
func LoadData() UsersData {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Unable to read users.csv", err)
	}
	reader := csv.NewReader(f)
	userLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse users.csv as csv", err)
	}

	//	users := make(UserMap, len(userLines))
	users := UsersData{}
	for _, line := range userLines {
		//id, _ := strconv.Atoi(line[0])
		//name := line[1]
		age, _ := strconv.Atoi(line[2])
		//address := line[3]
		//zip, _ := strconv.Atoi(line[3])
		//users[UserId(id)] = &User{UserId(id), int8(age), name, Address{int32(zip), address}, []Payment{}}
		users.ages = append(users.ages, int8(age))
	}

	f, err = os.Open("payments.csv")
	if err != nil {
		log.Fatalln("Unable to read payments.csv", err)
	}
	reader = csv.NewReader(f)
	paymentLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse payments.csv as csv", err)
	}

	for _, line := range paymentLines {
		//userId, _ := strconv.Atoi(line[2])
		paymentCents, _ := strconv.Atoi(line[0])
		//datetime, _ := time.Parse(time.RFC3339, line[1])
		//users[UserId(userId)].payments = append(users[UserId(userId)].payments, Payment{
		//	DollarAmount{uint32(paymentCents / 100), uint32(paymentCents % 100)},
		//	datetime,
		//})
		users.payments = append(users.payments, uint32(paymentCents))
	}

	return users
}
