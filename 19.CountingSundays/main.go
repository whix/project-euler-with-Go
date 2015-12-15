/*
You are given the following information, but you may prefer to do some research for yourself.

a)1 Jan 1900 was a Monday.
b)Thirty days has September,
  April, June and November.
  All the rest have thirty-one,
  Saving February alone,
  Which has twenty-eight, rain or shine.
c)And on leap years, twenty-nine.

A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
（第二十个世纪（1901年1月1日到2000年12月31日）一共有多少个星期日落在了当月的第一天？）
*/
package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	if year%100 == 0 {
		if year%400 == 0 {
			return true
		}
	} else {
		if year%4 == 0 {
			return true
		}
	}
	return false
}

func daysOfAYear(year int) (days int) {
	days = 30*4 + 31*7 + 28
	if isLeapYear(year) {
		days += 1
	}
	return
}

func daysFrom1900(year int) (days int) {
	for i := 1900; i < year; i++ {
		days += daysOfAYear(i)
	}
	return
}

// days begin from 1900/01/01
func isSunday(days int) bool {
	if days%7 == 0 {
		return true
	}
	return false
}

func countTheSundayFirst() (days int) {
	theDays := 0
	for i := 1901; i <= 2000; i++ {
		theDays = 0
		for month := 1; month < 13; month++ {
			switch month {
			case 1:
				{
					theDays = daysFrom1900(i) + 1
					if isSunday(theDays) {
						days++
					}
				}
			case 2:
				{
					theDays += 31
					if isSunday(theDays) {
						days++
					}
				}
			case 3:
				{
					if isLeapYear(i) {
						theDays += 29
					} else {
						theDays += 28
					}
					if isSunday(theDays) {
						days++
					}
				}
			case 4:
				{
					theDays += 31
					if isSunday(theDays) {
						days++
					}
				}
			case 5:
				{
					theDays += 30
					if isSunday(theDays) {
						days++
					}
				}
			case 6:
				{
					theDays += 31
					if isSunday(theDays) {
						days++
					}
				}
			case 7:
				{
					theDays += 30
					if isSunday(theDays) {
						days++
					}
				}
			case 8:
				{
					theDays += 31
					if isSunday(theDays) {
						days++
					}
				}
			case 9:
				{
					theDays += 31
					if isSunday(theDays) {
						days++
					}
				}
			case 10:
				{
					theDays += 30
					if isSunday(theDays) {
						days++
					}
				}
			case 11:
				{
					theDays += 31
					if isSunday(theDays) {
						days++
					}
				}
			case 12:
				{
					theDays += 30
					if isSunday(theDays) {
						days++
					}
				}
			}
		}
	}
	return
}

func main() {
	days := countTheSundayFirst()
	fmt.Println(days)
}
