package main

import "fmt"

const (
	espressoWater       = 250
	espressoMilk        = 0
	espressoCoffeeBeans = 16
	espressoCost        = 4
)

const (
	latteWater       = 350
	latteMilk        = 75
	latteCoffeeBeans = 20
	latteCost        = 7
)

const (
	cappuccinoWater       = 200
	cappuccinoMilk        = 100
	cappuccinoCoffeeBeans = 12
	cappuccinoCost        = 6
)

func printCounts(waterCount, milkCount, coffeeBeansCount, cupsCount, moneyCount int) {
	fmt.Printf("The coffee machine has:\n")
	fmt.Printf("%d of water\n", waterCount)
	fmt.Printf("%d of milk\n", milkCount)
	fmt.Printf("%d of coffee beans\n", coffeeBeansCount)
	fmt.Printf("%d of disposable cups\n", cupsCount)
	if moneyCount > 0 {
		fmt.Printf("$%d of money\n", moneyCount)
	} else {
		fmt.Printf("%d of money\n", moneyCount)
	}
}

func main() {
	availableWaterCount,
		availableMilkCount,
		availableCoffeeBeansCount,
		availableCupsCount,
		moneyCount := 400, 540, 120, 9, 550

	showMenu(&availableWaterCount, &availableMilkCount, &availableCoffeeBeansCount, &availableCupsCount, &moneyCount)
}

func showMenu(waterCount, milkCount, coffeeBeansCount, cupsCount, moneyCount *int) {
	var action string
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	fmt.Scan(&action)

	switch action {
	case "buy":
		fmt.Println("")
		accessBuyMode(waterCount, milkCount, coffeeBeansCount, cupsCount, moneyCount)
	case "fill":
		fmt.Println("")
		accessFillMode(waterCount, milkCount, coffeeBeansCount, cupsCount)
	case "take":
		fmt.Println("")
		accessTakeMode(moneyCount)
	case "remaining":
		fmt.Println("")
		printCounts(*waterCount, *milkCount, *coffeeBeansCount, *cupsCount, *moneyCount)
	case "exit":
		return
	default:
		fmt.Println("Not a valid action")
		return
	}

	fmt.Println("")
	showMenu(waterCount, milkCount, coffeeBeansCount, cupsCount, moneyCount)
}

func accessTakeMode(moneyCount *int) {
	fmt.Printf("I gave you $%d\n", *moneyCount)
	*moneyCount = 0
}

func accessFillMode(waterCount, milkCount, coffeeBeansCount, cupsCount *int) {
	var waterToAdd, milkToAdd, coffeeBeansToAdd, cupsToAdd int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&waterToAdd)
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&milkToAdd)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&coffeeBeansToAdd)
	fmt.Println("Write how many disposable coffee cups you want to add:")
	fmt.Scan(&cupsToAdd)

	*waterCount += waterToAdd
	*milkCount += milkToAdd
	*coffeeBeansCount += coffeeBeansToAdd
	*cupsCount += cupsToAdd
}

func accessBuyMode(waterCount, milkCount, coffeeBeansCount, cupsCount, moneyCount *int) {
	var option string
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	fmt.Scan(&option)
	if option != "1" && option != "2" && option != "3" {
		switch option {
		case "back":
			showMenu(waterCount, milkCount, coffeeBeansCount, cupsCount, moneyCount)
		default:
			fmt.Println("Invalid option. Please try again.")
			accessBuyMode(waterCount, milkCount, coffeeBeansCount, cupsCount, moneyCount)
		}
	}

	var quantity int
	fmt.Printf("How many of option '%s' would you like?\n", option)
	fmt.Scan(&quantity)
	if quantity < 1 {
		fmt.Println("Invalid quantity. Please try again.")
		accessBuyMode(waterCount, milkCount, coffeeBeansCount, cupsCount, moneyCount)
	}

	switch option {
	case "1":
		buyEspresso(waterCount, milkCount, coffeeBeansCount, cupsCount, moneyCount, quantity)
	case "2":
		buyLatte(waterCount, milkCount, coffeeBeansCount, cupsCount, moneyCount, quantity)
	case "3":
		buyCappuccino(waterCount, milkCount, coffeeBeansCount, cupsCount, moneyCount, quantity)
	}
}

func buyEspresso(waterCount, milkCount, coffeeBeansCount, cupsCount, moneyCount *int, quantity int) {
	if !hasEnoughWater(*waterCount, espressoWater, quantity) ||
		!hasEnoughMilk(*milkCount, espressoMilk, quantity) ||
		!hasEnoughCups(*cupsCount, quantity) ||
		!hasEnoughCoffeeBeans(*coffeeBeansCount, espressoCoffeeBeans, quantity) {
		return
	}
	fmt.Println("I have enough resources, making you a coffee!")
	*waterCount -= espressoWater
	*milkCount -= espressoMilk
	*coffeeBeansCount -= espressoCoffeeBeans
	*moneyCount += espressoCost
	*cupsCount--
}

func buyLatte(waterCount, milkCount, coffeeBeansCount, cupsCount, moneyCount *int, quantity int) {
	if !hasEnoughWater(*waterCount, latteWater, quantity) ||
		!hasEnoughMilk(*milkCount, latteMilk, quantity) ||
		!hasEnoughCups(*cupsCount, quantity) ||
		!hasEnoughCoffeeBeans(*coffeeBeansCount, latteCoffeeBeans, quantity) {
		return
	}
	fmt.Println("I have enough resources, making you a coffee!")
	*waterCount -= latteWater
	*milkCount -= latteMilk
	*coffeeBeansCount -= latteCoffeeBeans
	*moneyCount += latteCost
	*cupsCount--
}

func buyCappuccino(waterCount, milkCount, coffeeBeansCount, cupsCount, moneyCount *int, quantity int) {
	if !hasEnoughWater(*waterCount, cappuccinoWater, quantity) ||
		!hasEnoughMilk(*milkCount, cappuccinoMilk, quantity) ||
		!hasEnoughCups(*cupsCount, quantity) ||
		!hasEnoughCoffeeBeans(*coffeeBeansCount, cappuccinoCoffeeBeans, quantity) {
		return
	}
	fmt.Println("I have enough resources, making you a coffee!")
	*waterCount -= cappuccinoWater
	*milkCount -= cappuccinoMilk
	*coffeeBeansCount -= cappuccinoCoffeeBeans
	*moneyCount += cappuccinoCost
	*cupsCount--
}

func hasEnough(available, required int, resource string) bool {
	if required > available {
		fmt.Printf("Sorry, not enough %s!\n", resource)
		return false
	}

	return true
}

func hasEnoughWater(availableWater, requiredWater, quantity int) bool {
	return hasEnough(availableWater, quantity*requiredWater, "water")
}

func hasEnoughMilk(availableMilk, requiredMilk, quantity int) bool {
	return hasEnough(availableMilk, quantity*requiredMilk, "milk")
}

func hasEnoughCoffeeBeans(availableCoffeeBeans, requiredCoffeeBeans, quantity int) bool {
	return hasEnough(availableCoffeeBeans, quantity*requiredCoffeeBeans, "coffee beans")
}

func hasEnoughCups(availableCups, quantity int) bool {
	return hasEnough(availableCups, quantity, "disposable cups")
}
