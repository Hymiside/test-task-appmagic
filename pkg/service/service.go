package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hymiside/test-task-appmagic/pkg/cache"
)

type GasInfoDict struct {
	Ethereum Ethereum `json:"ethereum"`
}

type Ethereum struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Time           string  `json:"time"`
	GasPrice       float64 `json:"gasPrice"`
	GasValue       float64 `json:"gasValue"`
	Average        float64 `json:"average"`
	MaxGasPrice    float64 `json:"maxGasPrice"`
	MedianGasPrice float64 `json:"medianGasPrice"`
}

type Service struct {
	cache *cache.Cache
}

func NewService(c cache.Cache) *Service {
	return &Service{cache: &c}
}

// SetInfoGas кладет данные в кэш
func (s *Service) SetInfoGas() error {
	infoGas, err := s.GetInfoGasGit()
	if err != nil {
		return err
	}

	go s.SetInfoSumAllPeriod(infoGas)
	return nil
}

// GetInfoGasGit возвращает json из Git'а
func (s *Service) GetInfoGasGit() (GasInfoDict, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/CryptoRStar/GasPriceTestTask/main/gas_price.json")
	if err != nil {
		return GasInfoDict{}, fmt.Errorf("request error in Get: %w", err)
	}
	defer resp.Body.Close()

	var data GasInfoDict

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return GasInfoDict{}, fmt.Errorf("error decode json: %w", err)
	}

	return data, nil
}

// SetInfoGasPerMonth считает сколько было потрачено gas помесячно и кладет в кэш
func (s *Service) SetInfoGasPerMonth(infoGas GasInfoDict) {}

// SetInfoPricePerDay считает среднюю цену за день и кладет в кэш
func (s *Service) SetInfoPricePerDay(infoGas GasInfoDict) {}

// SetInfoHourlyPrice считает частотное распределние цены по часам (за весь период) и кладет в кэш
func (s *Service) SetInfoHourlyPrice(infoGas GasInfoDict) {}

// SetInfoSumAllPeriod считает сколько заплатили за весь период и кладет в кэш
func (s *Service) SetInfoSumAllPeriod(infoGas GasInfoDict) {
	var sum float64

	for _, value := range infoGas.Ethereum.Transactions {
		sum += value.GasPrice * value.GasValue
	}
	s.cache.Set("SumAllPeriod", sum)
}

// GetInfoGasPerMonth возвращает сколько было потрачено gas помесячно и кладет в кэш
func (s *Service) GetInfoGasPerMonth(infoGas GasInfoDict) {}

// GetInfoPricePerDay возвращает среднюю цену за день и кладет в кэш
func (s *Service) GetInfoPricePerDay(infoGas GasInfoDict) {}

// GetInfoHourlyPrice возвращает частотное распределние цены по часам (за весь период) и кладет в кэш
func (s *Service) GetInfoHourlyPrice(infoGas GasInfoDict) {}

// GetInfoSumAllPeriod возвращает сколько заплатили за весь период и кладет в кэш
func (s *Service) GetInfoSumAllPeriod() interface{} {
	sum, err := s.cache.Get("SumAllPeriod")
	if err != nil {

		infoGas, err := s.GetInfoGasGit()
		if err != nil {
			return err
		}
		s.SetInfoSumAllPeriod(infoGas)
		return infoGas
	}
	return sum
}
