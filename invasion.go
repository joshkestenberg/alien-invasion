package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//I've made the assumption that all cities will have one word names without spaces
//Additionally, I've assumed that we will pretend that all aliens move simultaneously during each round, so there can never be 2 aliens in a city in the middle of a round, and we do our housekeeping at the end of every round

type Alien struct {
	Trapped bool
	Dead    bool
	Moved   bool
	Index   int
	Moves   int
}

type City struct {
	Name      string
	North     string
	East      string
	South     string
	West      string
	Aliens    []Alien
	Destroyed bool
}

func makeMap(filename string) ([]City, error) {
	var cities []City
	var err error
	//open map file
	file, err := os.Open(filename)
	if err != nil {
		return cities, err
	}

	scanner := bufio.NewScanner(file)
	//scan file
	for scanner.Scan() {
		txt := scanner.Text()
		//split line into array of useful info
		info := strings.Split(txt, " ")
		city := City{}

		city.Name = info[0]
		//remove city name so that all we have left is surrounding cities
		info = info[1:]
		//add surrounding cities to city struct
		for _, neighbor := range info {
			if strings.Contains(neighbor, "north=") {
				city.North = strings.Replace(neighbor, "north=", "", 1)
			} else if strings.Contains(neighbor, "east=") {
				city.East = strings.Replace(neighbor, "east=", "", 1)
			} else if strings.Contains(neighbor, "south=") {
				city.South = strings.Replace(neighbor, "south=", "", 1)
			} else if strings.Contains(neighbor, "west=") {
				city.West = strings.Replace(neighbor, "west=", "", 1)
			}
		}
		cities = append(cities, city)
	}

	return cities, err
}

func unleashAliens(cities []City, num int) ([]Alien, error) {
	var aliens []Alien
	var err error

	for i := 0; i < num; i++ {
		//create alien
		alien := Alien{Index: i}
		//set range for random number and pick a number
		rng := len(cities)

		random := rand.Intn(rng)
		//use number to select city in which to place alien
		cities[random].Aliens = append(cities[random].Aliens, alien)

		aliens = append(aliens, alien)
	}

	return aliens, err
}

func houseKeeping(cities []City, aliens []Alien) ([]City, []Alien) {
	var ctGame bool
	var mvGame bool
	var tpGame bool
	var ddGame bool

	var purgeCities []string

	for i := 0; i < len(cities); i++ {
		if cities[i].Destroyed == false {
			name := cities[i].Name
			if len(cities[i].Aliens) >= 2 {
				//destroy city
				cities[i].Destroyed = true
				//add name to array for later
				purgeCities = append(purgeCities, name)
				//kill aliens in city
				count := len(cities[i].Aliens)
				for n := 0; n < count; n++ {
					aliens[cities[i].Aliens[n].Index].Dead = true
				}
				cities[i].Aliens = cities[i].Aliens[:0]
				fmt.Println(count, "out of the", len(aliens), "initial aliens have perished in battle, and the city of", name, "has been destroyed along with them!")
			}
		}
	}

	//purge all reference to city from other cities
	for i := 0; i < len(cities); i++ {
		if cities[i].Destroyed == false {
			for n := 0; n < len(purgeCities); n++ {
				if cities[i].North == purgeCities[n] {
					cities[i].North = ""
				} else if cities[i].East == purgeCities[n] {
					cities[i].East = ""
				} else if cities[i].South == purgeCities[n] {
					cities[i].South = ""
				} else if cities[i].West == purgeCities[n] {
					cities[i].West = ""
				}
			}
		}
	}

	//check if game is over; if there is any proof that game is not over, set respective var to true.
	for _, city := range cities {
		if city.Destroyed == false {
			ctGame = true
		}
	}
	for _, alien := range aliens {
		if alien.Dead == false {
			ddGame = true
		}
		if alien.Trapped == false {
			tpGame = true
		}
		if alien.Moves < 10000 && alien.Dead == false && alien.Trapped == false {
			mvGame = true
		}
	}
	if ctGame == false {
		fmt.Println("The world has been destroyed!")
		os.Exit(0)
	} else if ddGame == false {
		fmt.Println("All the aliens are dead. Perhaps we can rebuild..")
		os.Exit(0)
	} else if mvGame == false {
		fmt.Println("All remaining aliens have moved 10,000 or more times! I, for one, welcome our new alien overlords!")
		os.Exit(0)
	} else if tpGame == false {
		fmt.Println("All aliens are trapped! Attack now while they're immobile! God speed!")
		os.Exit(0)
	}
	//set all aliens moved back to false
	for i := 0; i < len(aliens); i++ {
		aliens[i].Moved = false
	}
	return cities, aliens
}

func move(cities []City, aliens []Alien) ([]City, []Alien) {
	var curCity *City
	var alien *Alien
	var trapped bool

	//remove alien from current city
	for i := 0; i < len(cities); i++ {
		//if the city has at least one alien
		if len(cities[i].Aliens) >= 1 {
			//select the first alien from the city
			curCity = &cities[i]
			alien = &aliens[curCity.Aliens[0].Index]
			//make sure this city and alien are eligible
			if curCity.Destroyed == false && alien.Trapped == false && alien.Moved == false {
				if len(curCity.Aliens) > 1 {
					curCity.Aliens[0] = curCity.Aliens[(len(curCity.Aliens) - 1)]
					curCity.Aliens = curCity.Aliens[:(len(curCity.Aliens) - 1)]
				} else {
					curCity.Aliens = curCity.Aliens[:0]
				}

				mvCity := ""
				north := true
				east := true
				south := true
				west := true

				for mvCity == "" {
					//set range for random number and pick a number
					rng := 4
					random := rand.Intn(rng)
					//use number to select to which city the alien moves
					switch random {
					case 0:
						if curCity.North != "" {
							mvCity = curCity.North
						} else {
							north = false
						}
					case 1:
						if curCity.East != "" {
							mvCity = curCity.East
						} else {
							east = false
						}
					case 2:
						if curCity.South != "" {
							mvCity = curCity.South
						} else {
							south = false
						}
					case 3:
						if curCity.West != "" {
							mvCity = curCity.West
						} else {
							west = false
						}
					}
					if north == false && east == false && south == false && west == false {
						//alien is trapped!
						//find alien in array of aliens
						alien.Trapped = true
						//put alien back in city
						curCity.Aliens = append(curCity.Aliens, *alien)
						fmt.Println("An alien is trapped in", curCity.Name)
						//break loop
						trapped = true
					}
					if trapped == true {
						break
					}
				}
				//add our alien to his new city
				if mvCity != "" {
					for i := 0; i < len(cities); i++ {
						if cities[i].Name == mvCity && cities[i].Destroyed == false {
							alien.Moved = true
							alien.Moves += 1

							cities[i].Aliens = append(cities[i].Aliens, *alien)
						}
					}
				}
			}
		}
	}
	return cities, aliens
}

func main() {
	//get command line arg and log err if invalid
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	} else if num <= 0 {
		fmt.Println("The map requires at least one alien!")
		os.Exit(0)
	}
	//create our cities
	cities, err := makeMap("./invasion_map")
	if err != nil {
		log.Fatal(err)
	}
	//create and distribute our aliens
	aliens, err := unleashAliens(cities, num)
	if err != nil {
		log.Fatal(err)
	}
	//preliminary check to see if any cities/aliens were immediately destroyed
	cities, aliens = houseKeeping(cities, aliens)

	//until we hit the game's end (infinite loop escaped by 'os.Exit once conditions are met')
	for {
		//move all eligible aliens
		cities, aliens = move(cities, aliens)
		//destroy towns/kill aliens
		cities, aliens = houseKeeping(cities, aliens)
	}
}
