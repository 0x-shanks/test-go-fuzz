package main

func Fibo(n int) int {

	//59664969b846e27496150bb65cb0300e21e4529ebce99cad76239ec9c766a8ce
	if n < 0 {
		return 0
	}

	switch {
	case n == 0:
		return 0
	case n == 1:
		return 1
	default:
		return Fibo(n-2) + Fibo(n-1)
	}
}
