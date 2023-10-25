package order

import (
	"example.com/stdin_wrapper"
	"fmt"
	"strconv"
	"log"
)

type Order struct {
	name	string
	count	int
	fio		string
	phone	int
	address Address
}

type Address struct {
	index	int
	city	string
	street	string
	house	string
	flat 	string
}

func NewOrderStdin() (Order, error) {
	var name string
	var count int
	var fio string
	var phone int

	var index int
	var city string
	var street string
	var home string
	var flat string

	var readerWrapper stdin_wrapper.StdinReader

	readerWrapper = stdin_wrapper.NewStdinReader("Введите наименование товара", "", "")
	_name, _ := readerWrapper.ReadString()
	if readerWrapper.IsValidByRegex() { name = _name} 

	readerWrapper = stdin_wrapper.NewStdinReader("Введите количество товара", `^\d+$`, "(только числа)")
	_count, _ := readerWrapper.ReadString()
	if readerWrapper.IsValidByRegex() { count, _ = strconv.Atoi(_count)} else { log.Fatal() }

	readerWrapper = stdin_wrapper.NewStdinReader("Введите ФИО покупателя", `^([a-z]|[A-Z]|[а-я]|[А-Я])+$`, "(только буквы)")
	_fio, _ := readerWrapper.ReadString()
	if readerWrapper.IsValidByRegex() { fio = _fio} else { log.Fatal() }

	readerWrapper = stdin_wrapper.NewStdinReader("Введите контактный телефон", `^\d{10}$`, "(10 цифр)")
	_phone, _ := readerWrapper.ReadString()
	if readerWrapper.IsValidByRegex() { phone, _ = strconv.Atoi(_phone)} else { log.Fatal() }

	readerWrapper = stdin_wrapper.NewStdinReader("Введите почтовый индекс", `^\d{6}$`, "(ровно 6 цифр)")
	_index, _ := readerWrapper.ReadString()
	if readerWrapper.IsValidByRegex() { index, _ = strconv.Atoi(_index)} else { log.Fatal() }

	readerWrapper = stdin_wrapper.NewStdinReader("Введите город", "", "")
	_city, _ := readerWrapper.ReadString()
	if readerWrapper.IsValidByRegex() { city = _city}

	readerWrapper = stdin_wrapper.NewStdinReader("Введите улицу", "", "")
	_street, _ := readerWrapper.ReadString()
	if readerWrapper.IsValidByRegex() { street = _street}

	readerWrapper = stdin_wrapper.NewStdinReader("Введите дом", "", "")
	_home, _ := readerWrapper.ReadString()
	if readerWrapper.IsValidByRegex() { home = _home}

	readerWrapper = stdin_wrapper.NewStdinReader("Введите квартиру", "", "")
	_flat, _ := readerWrapper.ReadString()
	if readerWrapper.IsValidByRegex() { flat = _flat}

	var address = NewAddress(index, city, street, home, flat)

	return Order{
		name:   name,
		count: count,
		fio: fio,
		phone: phone,
		address: address,
	}, nil
}


func NewOrder(name string, count int, fio string, phone int, address Address) Order {
	return Order{
		name:   name,
		count: count,
		fio: fio,
		phone: phone,
		address: address,
	}
}

func NewAddress(index int, city string, street string, house string, flat string) Address {
	return Address{
		index:   index,
		city: city,
		street: street,
		house: house,
		flat: flat,
	}
}

func (ord Order) PrintFmt() {
	fmt.Printf("------\n")
	fmt.Printf("Наименование товара: %v\n", ord.name)
	fmt.Printf("Количество: %v\n", ord.count)
	fmt.Printf("ФИО покупателя: %v\n", ord.fio)
	fmt.Printf("Контактный телефон: %v\n", ord.phone)
	fmt.Printf("---\n")
	fmt.Printf("Адрес\n")
	fmt.Printf("Индекс: %v\n", ord.address.index)
	fmt.Printf("Город: %v\n", ord.address.city)
	fmt.Printf("Улица: %v\n", ord.address.street)
	fmt.Printf("Дом: %v\n", ord.address.house)
	fmt.Printf("Квартира: %v\n", ord.address.flat)
	fmt.Printf("------\n")
}

func (ord *Order) AddCount(other int) {
	ord.count += other
}
