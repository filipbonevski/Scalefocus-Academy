package main

import "testing"

func Benchmark100PrimesWith0msSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(100, 0)
	}
}

func Benchmark100PrimesWith5msSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(100, 5000000)
	}
}

func Benchmark100PrimesWith10msSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(100, 10000000)
	}
}

func Benchmark100GoPrimesWith0msSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrimesAndSleep(100, 0)
	}
}

func Benchmark100GoPrimesWith5msSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrimesAndSleep(100, 5000000)
	}
}

func Benchmark100GoPrimesWith10msSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrimesAndSleep(100, 10000000)
	}
}
