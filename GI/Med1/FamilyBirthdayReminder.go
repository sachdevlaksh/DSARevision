package Med1

import (
	"container/heap"
	"fmt"
	"time"
)


type Person struct {
	name     string
	birthday time.Time
	father   *Person
	mother   *Person
}

type PersonPriorityQueue []*Person

func (pq PersonPriorityQueue) Len() int { return len(pq) }

func (pq PersonPriorityQueue) Less(i, j int) bool {
	return pq[i].birthday.Before(pq[j].birthday)
}

func (pq PersonPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PersonPriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Person))
}

func (pq *PersonPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type BirthdayReminder struct {
	currentYear *PersonPriorityQueue
	nextYear    *PersonPriorityQueue
}

func NewBirthdayReminder(person *Person) *BirthdayReminder {
	currentYear := &PersonPriorityQueue{}
	nextYear := &PersonPriorityQueue{}
	heap.Init(currentYear)
	heap.Init(nextYear)

	br := &BirthdayReminder{
		currentYear: currentYear,
		nextYear:    nextYear,
	}
	br.addFamily(person)
	return br
}

func (br *BirthdayReminder) getNextBirthday() *Person {
	if br.currentYear.Len() == 0 {
		// Swap queues
		br.currentYear, br.nextYear = br.nextYear, br.currentYear
	}
	if br.currentYear.Len() == 0 {
		return nil
	}
	nextBirthdayPerson := heap.Pop(br.currentYear).(*Person)
	heap.Push(br.nextYear, nextBirthdayPerson)
	return nextBirthdayPerson
}

func (br *BirthdayReminder) addFamily(person *Person) {
	if person == nil {
		return
	}
	br.addPerson(person)
	br.addFamily(person.father)
	br.addFamily(person.mother)
}

func (br *BirthdayReminder) addPerson(person *Person) {
	if br.getCurrentDate().Before(person.birthday) {
		heap.Push(br.currentYear, person)
	} else {
		heap.Push(br.nextYear, person)
	}
}

func (br *BirthdayReminder) getCurrentDate() time.Time {
	return time.Now()
}

func main() {
	// Note: time.Date uses year, month (1-12), day
	father := &Person{"Bob", time.Date(2024, 3, 10, 0, 0, 0, 0, time.UTC), nil, nil}
	mother := &Person{"Carol", time.Date(2024, 7, 20, 0, 0, 0, 0, time.UTC), nil, nil}
	mainPerson := &Person{"Alice", time.Date(2024, 5, 15, 0, 0, 0, 0, time.UTC), father, mother}

	reminder := NewBirthdayReminder(mainPerson)

	nextBirthday := reminder.getNextBirthday()
	fmt.Println("Next birthday:", nextBirthday.name)

	nextBirthday = reminder.getNextBirthday()
	fmt.Println("Next birthday:", nextBirthday.name)

	nextBirthday = reminder.getNextBirthday()
	fmt.Println("Next birthday:", nextBirthday.name)

	nextBirthday = reminder.getNextBirthday()
	fmt.Println("Next birthday:", nextBirthday.name)
}