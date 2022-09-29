package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Hymiside/test-task-appmagic/pkg/cache"
	"github.com/Hymiside/test-task-appmagic/pkg/models"
)

type Service struct {
	cache *cache.Cache
}

var amountDays float64

func NewService(c cache.Cache) *Service {
	return &Service{cache: &c}
}

// SetInfoGas кладет данные в кэш
func (s *Service) SetInfoGas() error {
	infoGas, err := s.GetInfoGasGit()
	if err != nil {
		return err
	}
	amountDays = float64(len(infoGas.Ethereum.Transactions) / 24)

	s.SetInfoGasPerMonth(infoGas)
	s.SetInfoHourlyPrice(infoGas)
	s.SetInfoSumAllPeriod(infoGas)
	s.SetInfoPricePerDay(infoGas)
	return nil
}

// GetInfoGasGit возвращает json из Git'а
func (s *Service) GetInfoGasGit() (models.GasInfoDict, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/CryptoRStar/GasPriceTestTask/main/gas_price.json")
	if err != nil {
		return models.GasInfoDict{}, fmt.Errorf("request error in Get: %w", err)
	}
	defer resp.Body.Close()

	var data models.GasInfoDict
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return models.GasInfoDict{}, fmt.Errorf("error decode json: %w", err)
	}
	return data, nil
}

// SetInfoGasPerMonth считает сколько было потрачено gas помесячно и кладет в кэш
func (s *Service) SetInfoGasPerMonth(infoGas models.GasInfoDict) {
	gasPerMonthDict := map[string]float64{
		"January":   0.0,
		"February":  0.0,
		"March":     0.0,
		"April":     0.0,
		"May":       0.0,
		"June":      0.0,
		"July":      0.0,
		"August":    0.0,
		"September": 0.0,
		"October":   0.0,
		"November":  0.0,
		"December":  0.0,
	}

	for _, value := range infoGas.Ethereum.Transactions {
		t, _ := time.Parse("06-02-01 15:04", value.Time)
		gasPerMonthDict[t.Month().String()] += value.GasValue
	}
	s.cache.Set("GasPerMonth", gasPerMonthDict)
}

// SetInfoPricePerDay считает среднюю цену за день и кладет в кэш
func (s *Service) SetInfoPricePerDay(infoGas models.GasInfoDict) {
	var (
		count int
		sum   float64
	)
	averagePriceDict := make(map[string]float64)

	for _, value := range infoGas.Ethereum.Transactions {
		sum += value.GasPrice
		count++

		if count == 23 {
			averagePrice := sum / 24
			averagePriceDict[value.Time[:8]] = averagePrice

			count = 0
			sum = 0
		}
	}
	s.cache.Set("PricePerDay", averagePriceDict)
}

// SetInfoHourlyPrice считает частотное распределние цены по часам (за весь период) и кладет в кэш
func (s *Service) SetInfoHourlyPrice(infoGas models.GasInfoDict) {
	var count int
	hourlyPriceDict := make(map[string]float64)

	for _, value := range infoGas.Ethereum.Transactions {
		if count != 24 {
			hourlyPriceDict[value.Time[9:]] = value.GasPrice
			count++
		}
		hourlyPriceDict[value.Time[9:]] += value.GasPrice
	}

	for key, value := range hourlyPriceDict {
		hourlyPriceDict[key] = value / amountDays
	}
	s.cache.Set("HourlyPrice", hourlyPriceDict)
}

// SetInfoSumAllPeriod считает сколько заплатили за весь период и кладет в кэш
func (s *Service) SetInfoSumAllPeriod(infoGas models.GasInfoDict) {
	var sum float64

	for _, value := range infoGas.Ethereum.Transactions {
		sum += value.GasPrice * value.GasValue
	}
	s.cache.Set("SumAllPeriod", sum)
}

// GetInfoGas возвращает данные из кэша по ключу
func (s *Service) GetInfoGas(key string) (interface{}, error) {
	res, err := s.cache.Get(key)
	if err != nil {
		return nil, err
	}
	return res, nil
}
