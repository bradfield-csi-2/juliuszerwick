package metrics

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

// UserId is 8 bytes
type UserId int

type UserMap map[UserId]*User

// Address struct is 8 + num_chars bytes
type Address struct {
	fullAddress string // sequence of bytes, each char is 1 byte
	zip         int    // int is 4 bytes
}

// Dollar Amount is 16 bytes
type DollarAmount struct {
	dollars, cents uint64 // both are 8 bytes
}

// Payment could be huge as `time.Time` is a struct containing a struct.
type Payment struct {
	amount DollarAmount
	time   time.Time
}

// This struct can be very big given the number of elements in `payments` slice.
type User struct {
	id       UserId    // 8 bytes
	name     string    // num_chars bytes
	age      int       // 8 bytes
	address  Address   // 8 bytes + num_bytes
	payments []Payment // HUGE
}

// Passing in a map of Users (structs) when you only
// use the age field of each User.
// Better to simply pass in an array of the age values?
func AverageAge(users UserMap) float64 {
	average, count := 0.0, 0.0
	for _, u := range users {
		count += 1
		average += (float64(u.age) - average) / count
	}
	return average
}

// Only need the payments info, so why pass in map of Users?
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
// Same as AveragePaymentAmount, why pass in a map of Users?
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
		users[UserId(id)] = &User{UserId(id), name, age, Address{address, zip}, []Payment{}}
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
			DollarAmount{uint64(paymentCents / 100), uint64(paymentCents % 100)},
			datetime,
		})
	}

	return users
}
