package app

import (
	"errors"
	"fmt"
)

var (
	// *********************************************ERROR*********************************************
	ErrNotFound                = errors.New("не найден")
	ErrNotEmail                = errors.New(" Пользователь с такой почтой не существует!")
	ErrRecordNotFound          = errors.New(" Пользователь с такими данными не существует!")
	ErrNotPhone                = errors.New(" Пользователь с таким номером существует! Пожалуйста, используйте другой номер.")
	ErrInternal                = errors.New("что-то пошло не так. Пожалуйста, повторите попытку позже")
	ErrConflict                = errors.New("уже существует")
	ErrNotFoundUser            = errors.New("такой пользователь не существует в нашей баззе. Используйте корпоративную почту! ")
	ErrConflictEmail           = errors.New("адрес электронной почты уже используется другим пользователем")
	ErrConflictTransaction     = errors.New("ошибка транзакции")
	ErrConflictNumber          = errors.New("Номер неправильный, должно быт 9 цифр  ")
	ErrNotRightKey             = errors.New("Код подтверждения не правильный ")
	ErrNotRightRepeatPassword  = errors.New("подтверждённый пароль не совпадает с введённым")
	ErrUserActive              = errors.New("Пользователь уже активный ")
	ErrBadEmail                = errors.New("адрес электронной почты неправильно")
	ErrNotFoundEmailOrPassword = errors.New("не найден адрес электронной почты или пароль")
	ErrNotFoundEmail           = errors.New("пользователь с такой почтой не найден")
	ErrDontHaveAccess          = errors.New("у вас нет доступа")
	// *********************************************SUCCESS*********************************************
	ExcelFileSuccess = fmt.Sprintln("excel-sendNotificationLogics успешно добавлен")
)

// AnotherError ...
func AnotherError(err string) error {
	return errors.New(err)
}

// ErrWrong ...
func ErrWrong(field string) error {
	var res string

	switch field {
	case "email":
		res = "неправильный адрес электронной почты"
	case "password":
		res = "неправильный пароль"
	case "role":
		res = "неправильная роль"
	case "quantity":
		res = "количество не может быть больше количество продаваемый акции"
	case "token":
		res = "недействительный токен"
	case "old_password":
		res = "старый пароль неверный"
	case "match_passwords":
		res = "новый пароль и подтвержденный новый пароль не совпадают"
	case "price_0":
		res = "цена не можеть быть нулевым или меньше нуля"
	case "price+20%":
		res = "стоимость акции не можеть быть выше 20% от стоимости указанной акции"
	case "price-20%":
		res = "стоимость акции не можеть быть ниже 20% от стоимости указанной акции"
	case "stock_amount":
		res = "у вас нет столько акции"
	case "stock_amounts":
		res = "количество акции не может быть ровно или меньше 0"
	default:
		res = "неверные данные"
	}

	return errors.New(res)
}
