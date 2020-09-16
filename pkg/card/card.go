// Пакет управления банковскими картами
package card

import (
	"errors"
	"strings"
)

// Ошибки
var ErrCardNotFound = errors.New("card not found")

// Описание банковской карты
type Card struct {
	Id           cardId
	Owner
	Issuer  	 string
	Balance      int
	Currency     string
	Number       string
	Icon         string
	Transactions []Transaction
}

// Описание транзакции
type Transaction struct {
	Id     string
	Bill   int
	Time   int
	MCC    string
	Status string
}

// Идентификат банковской карты
type cardId string

// Инициалы владельца банковской карты
type Owner struct {
	FirstName string // Имя владельца карты
	LastName  string // Фамилия владельца карты
}

// Сервис банка
type Service struct {
	BankName string
	Cards    []*Card
}

const prefix = "5106 21" //Первые 6 цифр нашего банка


// Метод создания экземпляра банковской карты
func (s *Service) CardIssue(
	id cardId,
	fistName,
	lastName,
	issuer string,
	balance int,
	currency string,
	number string,
) *Card {
	var card = &Card{
		Id: id,
		Owner: Owner{
			FirstName: fistName,
			LastName:  lastName,
		},
		Issuer:   issuer,
		Balance:  balance,
		Currency: currency,
		Number:   number,
		Icon:     "https://.../logo.png",
	}
	s.Cards = append(s.Cards, card)
	return card
}

// Функция добавляет транзакцию совершенную по карте
func AddTransaction(card *Card, transaction Transaction) {
	card.Transactions = append(card.Transactions, transaction)
}

// Функция вычисления суммы транзакций по категории
func SumByMCC(transactions []Transaction, mcc []string) int {
	var mmcSum int

	for _, code := range mcc {
		for _, t := range transactions {
			if code == t.MCC {
				mmcSum += t.Bill
			}
		}
	}

	return mmcSum

}

// Функция определения категории по MCC
func TranslateMCC(code string) string {
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5812": "Рестораны",
	}

	const errCategoryUndef = "Категория не указана"

	if value, ok := mcc[code]; ok {
		return value
	}

	return errCategoryUndef

}

// Метод поиска банковской карты по номеру платежной системы
func (s *Service) Card(number string) (*Card, error) {

	for _, с := range s.Cards {
		if strings.HasPrefix(с.Number, prefix) == true {
			return с, nil
		}
	}
	return nil, ErrCardNotFound
}