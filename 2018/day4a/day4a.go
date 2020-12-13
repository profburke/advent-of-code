package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Action string

const (
	beginsShift Action = "begins shift"
	fallsAsleep Action = "falls asleep"
	wakesUp Action = "wakes up"
)

type Entry struct {
	Year int
	Month int
	Day int
	Hour int
	Minute int
	GuardNumber int
	WhichAction Action
}

func (e Entry) SortableDate() string {
	dateString := fmt.Sprintf("%02d%02d%02d%02d%02d",
		e.Year,
		e.Month,
		e.Day,
		e.Hour,
		e.Minute)
	return dateString
}

func (e Entry) String() string {
	return fmt.Sprintf("[%02d/%02d/%02d %02d:%02d] - Guard %4d: %s",
		e.Year, e.Month, e.Day, e.Hour, e.Minute, e.GuardNumber, e.WhichAction)
}

func processRawAction(rawAction string) (guardNumber int, action Action, err error) {
	err = nil
	
	if strings.HasPrefix(rawAction, string(wakesUp)) {
		action = wakesUp
		guardNumber = -1
	} else if strings.HasPrefix(rawAction, string(fallsAsleep)) {
		action = fallsAsleep
		guardNumber = -2
	} else if strings.HasSuffix(rawAction, string(beginsShift)) {
		action = beginsShift
		matches, err := fmt.Sscanf(rawAction, "Guard #%d", &guardNumber)
		if err != nil || matches != 1 {
			guardNumber = -3
		}
	} else {
		err = errors.New("Unknown action: '" + rawAction + "'")
	}
	return guardNumber, action, err
}

func ReadEntries(r io.Reader) (entries []Entry, err error) {
	scanner := bufio.NewScanner(r)
	
	for scanner.Scan() {
		var (
			year int
			month int
			day int
			hour int
			minute int
		)
		
		line := scanner.Text()
		matches, err := fmt.Sscanf(line, "[%d-%d-%d %d:%d]", &year, &month, &day, &hour, &minute)
		if err != nil || matches != 5 {
			fmt.Fprintln(os.Stderr, "error reading entry: ", err)
			os.Exit(1)
		}
		
		guardNumber, action, err := processRawAction(line[19:])
		if err != nil {
			fmt.Fprintln(os.Stderr, "error converting input: ", err)
			break
		}
		
		entry := Entry{year, month, day, hour, minute, guardNumber, action}
		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	
	return entries, nil
}

func process(entries []Entry) (totals map[int]int, chart map[int]map[int]int) {
	totals = make(map[int]int)
	chart = make(map[int]map[int]int)

	currentGuard := -1
	totalSlept := 0
	sleepStart := -1

	// for now, assume guard always wakes up before shift is over...
	
	for _, entry := range entries {
		if entry.WhichAction == beginsShift {
			totals[currentGuard] += totalSlept
			currentGuard = entry.GuardNumber

			_, ok := chart[currentGuard]
			if !ok {
				chart[currentGuard] = make(map[int]int)
			}
			
			totalSlept = 0
			sleepStart = -1
		} else if entry.WhichAction == fallsAsleep {
			sleepStart = entry.Minute
		} else if entry.WhichAction == wakesUp {
			sleepEnd := entry.Minute
			totalSlept += (sleepEnd - sleepStart)
			for x := sleepStart; x < sleepEnd; x++ {
				chart[currentGuard][x] += 1
			}
		}
	}
	
	return totals, chart
}

func pickSleepiest(totals map[int]int) (result int) {
	maxSleep := -1
	result = -1

	for guard, total := range totals {
		if total > maxSleep {
			maxSleep = total
			result = guard
		}
	}
	
	return result
}

func mostSlept(data map[int]int) (result int) {
	max := -1
	result = -1

	for minute, multiples := range data {
		if multiples > max {
			max = multiples
			result = minute
		}
	}
	
	return result
}

func main() {
	entries, err := ReadEntries(os.Stdin)
	
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading entries: ", err)
		os.Exit(1)
	}

	sort.Slice(entries, func(i, j int) bool { return entries[i].SortableDate() < entries[j].SortableDate() })

	totals, chart := process(entries)

	targetGuard := pickSleepiest(totals)
	targetMinute := mostSlept(chart[targetGuard])

	fmt.Println("Guard: ", targetGuard)
	fmt.Println("Minute: ", targetMinute)
	fmt.Println(targetGuard * targetMinute)
}
