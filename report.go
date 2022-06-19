package main

import (
	"log"
	"strconv"

	"golang.org/x/sync/errgroup"
)

func reportFuncWrapper[T interface{}](errgroup *errgroup.Group, reportFns ...func(<-chan T) error) []chan<- T {
	var channels []chan<- T
	for _, reportFn := range reportFns {
		localFn := reportFn
		channel := make(chan T)
		errgroup.Go(func() error {
			return localFn(channel)
		})
		channels = append(channels, channel)
	}

	return channels
}

func analyseIngredients(channel <-chan Drug) error {
	ingredientMap := make(map[string]int64)
	for drug := range channel {
		for _, ingredient := range drug.ActiveIngredients {
			ingredientMap[ingredient.Name]++
		}
	}
	log.Println(ingredientMap)
	return nil
}

func analyseStartYear(channel <-chan Drug) error {
	startYearMap := make(map[int]int64)
	for drug := range channel {
		year, err := strconv.Atoi(drug.MarketingStartDate[0:4])
		if err != nil {
			year = 0
		}
		startYearMap[year]++
	}
	log.Println(startYearMap)
	return nil
}

func analyseLatestExpirationDate(channel <-chan Drug) error {
	var latestExpirationDate int = 0
	var latestExpirationDateDrug Drug
	for drug := range channel {
		expirationDate, err := strconv.Atoi(drug.ListingExpirationDate)
		if err != nil {
			expirationDate = 0
		}

		if expirationDate > latestExpirationDate {
			latestExpirationDate = expirationDate
			latestExpirationDateDrug = drug
		}
	}

	log.Println(latestExpirationDate)
	log.Println(latestExpirationDateDrug)

	return nil
}
