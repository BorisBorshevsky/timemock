# timemock

Inspired by [Timecop](https://github.com/travisjeffery/timecop) Ruby Gem.
It provides methods for Freezing time, Traveling in time, and scaling time for testing.

## Motivation
Since in golang we cant override time package function for mock purposes 
 
## Install
```bash
$ go get -u github.com/BorisBorshevsky/timemock
```

## How to use
There are 2 ways (at least) to use this package.
- Use your own clock (and pass it as needed)
- Use the standard clock as a shared one between all the packages

both will have the same interface

```go
type Clock interface {
	Now() time.Time
	Since(time.Time) time.Duration
	Freeze(time.Time)
	Travel(time.Time)
	Scale(float64)
	Return()
}
```

**Im order to use the second method (shared clock), all the usages of `time.Now` and `time.Since` in your code (including non test code) should be replaced with `timemock.Now` and `timemock.Since`.**


## Examples

#### Global 
```go
timemock.Now() //time now

timemock.Freeze(timemock.Now()) //time is frozen

timemock.Return() //time is unfrozen

dummyTime := time.Unix(1522549800, 0) // Sunday, April 1, 2018 2:30:00 AM

timemock.Travel(dummyTime) //we traveled to 1st April

timemock.Scale(5) //time runs 5 times faster now

timemock.Return() // all is back to normal again

```


#### Travel
```go
clock := timemock.New()

timeForTest := time.Unix(1522549800, 0)
clock.Travel(timeForTest) //April 1, 2018 02:30:00

//what ever code that runs 10 seconds...
// ...
time.Sleep(time.Second * 10)
// ...

clock.Now() //April 1, 2018 02:30:10
clock.Return() //time is untraveled and back to regular time
```

#### Freeze
```go
clock := timemock.New()

timeForTest := time.Unix(1522549800, 0)
clock.Freeze(timeForTest) //April 1, 2018 02:30:00

//what ever code that runs 10 seconds...
// ...
time.Sleep(time.Second * 10)
// ...


clock.Now() //April 1, 2018 02:30:00, time is still frozen
clock.Return() //time is unfrozen and back to regular time
```

#### Scale
```go
clock := timemock.New()

clock.Scale(6) 

//what ever code that runs 10 seconds...
// ...
time.Sleep(time.Second * 10)
// ...

//clock will think that 10 * 6 = 60 seconds passed
clock.Now() //April 1, 2018 02:31:00, time is still frozen
```

> Travel can be used together with scale

# Contact

- [BorisBorshevsky@gmail.com](mailto:BorisBorshevsky@gmail.com)

