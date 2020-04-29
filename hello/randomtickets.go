
package main

import (
    "fmt" 
    "math/rand"
    "time"
    "strings"
)

type speed int32

func inputCost(defaultMin uint32, defaultMax uint32) (uint32, uint32){

    var minCost uint32
    var maxCost uint32

    fmt.Print("Введите минимальную и максимальную стоимость: ")
    fmt.Scan(&minCost, &maxCost)

    if minCost < maxCost{
        return minCost, maxCost
    }

    return defaultMin, defaultMax    
}

func inputNumOfTickets(defaultNum uint8) (uint8){
   
    var tickets uint8
    fmt.Print("Введите количество билетов: ")
    fmt.Scan(&tickets)

    if tickets > 0{
        return tickets
    }
    return defaultNum
}

func randRange(min speed, max speed) speed {

    return speed(rand.Intn(int(max - min)) + int(min))
}

func (trSpeed speed) calcDurationAndCost(minCost uint32, maxCost uint32) (uint32, float64) {
    
    return 345, 45.0
}


func main() {
    rand.Seed(time.Now().UnixNano())

    const EarthToMars = 62.1e6
    const secondsPerDay = 86400
    

    var minSpeed speed = 16
    var maxSpeed speed = 30    // in seconds

    minCost, maxCost    := inputCost(30, 50)      // in millions dollars USA
    tickets             := inputNumOfTickets(10)
    

    fmt.Printf("\t%-25v %4v %-12v $ %6v\n", "Spaceline", "Days", "Trip name", "Cost")
    fmt.Print("\t")

    // Вывод подчеркивающего знака равенства
    equalies := 25 + 4 + 12 + 6 + 5
    fmt.Println(strings.Repeat("=", equalies))

    for ticket := 0; uint8(ticket) < tickets; ticket++ {

        travelSpeed := randRange(minSpeed, maxSpeed)
        durationInDays := (EarthToMars / travelSpeed) / secondsPerDay
        
        percentRangeSpeed := float64(maxSpeed - minSpeed) / 100.0
        difSpeed := travelSpeed - minSpeed

        // На сколько процентов выше скорость от минимальной
        speedPercentToRange := float64(difSpeed) / percentRangeSpeed

        percentRangeCost := float64(maxCost - minCost) / 100.0

        // Полная стоимость 
        // минимальная стоимость плюс процент за скорость
        fullCost := float64(minCost) + (percentRangeCost * speedPercentToRange)

        //durationInDays, fullCost := calcDurationAndCost(minCost, maxCost)


        var companyName string
        switch company := rand.Intn(3); company {
            case 0:
                companyName = "Virginia Galactic"
                break;
            case 1:
                companyName = "SpaceX"
                break;
            default:
                companyName = "Space Adventures"
                break;
        }
        

        var tripTypeName string
        switch tripType := rand.Intn(2); tripType {
            case 0:
                tripTypeName = "Round-trip"
                durationInDays *= 2
                fullCost *= 2
                break;
            default:
                tripTypeName = "One-way"
                break;
        }

        fmt.Printf("\t%-25v %-4v %-12v $ %6.2f\n", companyName, durationInDays, tripTypeName, fullCost)
    }

    var temp string
    fmt.Scan(&temp)

}
