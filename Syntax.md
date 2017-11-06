
# Syntax

**Types:**

Numbers
Vectors
Strings
UDF Structs

## Builtin

## Constants

const c = 299792458 m/s
const G = 6.67408E-11 (m^3 * kg^-2 * s^-2)

## Functions

```

const G = 6.67E-1
const ℏ = 1.054E-35 (J * s)

let a = 5

func in2m(x float) -> x * 0.0254
func in2ft(x float) -> x / 12

func fNewton (M kg, m kg, r m) N = {
	G * m * M / r**2
}

func fNewtonMod (fN N, Vs m/s, Vo m/s) N = {
	if Vs/Vo > 0 {
		- fN * sqrt(Vs/Vo - 1)
	} else Vs/Vo < 0 {
		fN * sqrt(1 - Vs/Vo)
	} else {
		0
	}
}

func <A, B, C> (a A, b B) C
type Filter = func [A] (a A, i int) -> bool


```

## Structs

```

unit Kilogram (kg) {
	1000 = 1 g
}

unit Gram (g) {
	1000 = 1 kg
}

unit Second (s) {
	86400 = 1 Day
	1 = 1E9 ns
	1 = 1E6 μs
}

unit Meter (m)
unit Kelvin (K)
unit Mole (mol)
unit Current (A)
unit Intensity (cd)

unit Newton (N) {
	1 = 1 kg * m / s^2
}
unit Coulomb (C) {
	1 = 1 A * m
}
unit Volt (V) {
	1 = 1 J / C
}
unit Inches (in) {
	1 = 0.0254 m
	1 = 1/12 ft
}
unit Feet (ft) {
	1 = 12 in
	1 = 0.3048 m
	5280 = 1 mi
}

conversion 1 m = 3.28084 ft

1 m to ft

struct Planet = {
	Name string
	Mass (kg)
	Radius (m)
}

struct SolarSystem = {
	Name string
	Planets []Planet
}

struct Person = {
	Name string
}

const Earth = Planet{Name: "Earth", Mass: 5.972E24}

```

## Loops

```
var squares = for i in range(5) {
	i**2
}

var filtered = filter(array, (a, i) -> {
	a > 5
})

var filtered = array.filter((a, i) -> {
	a > 5
}) 

```

## Enums

```
enum Level = {
	LOW
	MEDIUM
	HIGH
}

```


