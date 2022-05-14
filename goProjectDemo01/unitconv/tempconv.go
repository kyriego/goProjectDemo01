package unitconv

func FtoC(f Fahrenheit) Centigrade {
	return Centigrade((f - 32) / 1.8)
}

func CtoF(c Centigrade) Fahrenheit {
	return Fahrenheit(32 + c*1.8)
}

func KtoC(k Kelvins) Centigrade {
	return Centigrade(k - 273.15)
}

func CtoK(c Centigrade) Kelvins {
	return Kelvins(c + 273.15)
}
