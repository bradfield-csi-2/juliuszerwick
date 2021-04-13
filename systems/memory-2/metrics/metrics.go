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

type Address struct {
	fullAddress string // sequence of bytes, each char is 1 byte
	zip         int32  // zip is 5 digits long - fit in int32
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
	name     string // string
	age      int8   // age will never exceed 150-200 - fit in int 8
	address  Address
	payments []Payment
}

func AverageAge(users UserMap) float64 {
	average, count := float32(0.0), float32(0.0)
	//ages := make([]int8, 0)
	var ages [100000]int8
	for i, u := range users {
		ages[i] = u.age
	}

	//fmt.Println("============================")
	//fmt.Printf("len(ages): %v\n", len(ages))
	//fmt.Println("============================")

	//for _, u := range users {
	for _, a := range ages {
		count += 1
		//average += (float32(u.age) - average) / count
		average += (float32(a) - average) / count
	}
	return float64(average)
}

func AveragePaymentAmount(users UserMap) float64 {
	average, count := 0.0, 0.0
	for _, u := range users {
		for _, p := range u.payments {
			count += 1
			amount := float64(p.amount.dollars) + float64(p.amount.cents)/100
			average += (amount - average) / count
		}
	}
	return average
}

// Compute the standard deviation of payment amounts
func StdDevPaymentAmount(users UserMap) float64 {
	mean := AveragePaymentAmount(users)
	squaredDiffs, count := 0.0, 0.0
	for _, u := range users {
		for _, p := range u.payments {
			count += 1
			amount := float64(p.amount.dollars) + float64(p.amount.cents)/100
			diff := amount - mean
			squaredDiffs += diff * diff
		}
	}
	return math.Sqrt(squaredDiffs / count)
}

func LoadData() UserMap {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Unable to read users.csv", err)
	}
	reader := csv.NewReader(f)
	userLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse users.csv as csv", err)
	}

	users := make(UserMap, len(userLines))
	for _, line := range userLines {
		id, _ := strconv.Atoi(line[0])
		name := line[1]
		age, _ := strconv.Atoi(line[2])
		address := line[3]
		zip, _ := strconv.Atoi(line[3])
		users[UserId(id)] = &User{UserId(id), name, int8(age), Address{address, int32(zip)}, []Payment{}}
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
		userId, _ := strconv.Atoi(line[2])
		paymentCents, _ := strconv.Atoi(line[0])
		datetime, _ := time.Parse(time.RFC3339, line[1])
		users[UserId(userId)].payments = append(users[UserId(userId)].payments, Payment{
			DollarAmount{uint32(paymentCents / 100), uint32(paymentCents % 100)},
			datetime,
		})
	}

	return users
}
