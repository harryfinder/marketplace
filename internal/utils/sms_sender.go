package utils

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

// SmsSender Номер должно быт в формате "992921234567" состоящие из 12 цифр.
// textMessage текст Сообщении.
// senderAddress имя отправителя который показывает пользователю
func SmsSender(numberPhone string, textMessage string) (int, error) {

	message := map[string]interface{}{
		"phoneNumber":   numberPhone,
		"text":          textMessage,
		"senderAddress": "alif.stock",
		//"priority":      0,
		//"scheduledAt":   nil,
		//"expiresIn":     0,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Println("ERROR", err)
		return http.StatusBadRequest, AnotherError("oshibka v json marshale")
	}
	// Создаем объект реквеста
	req, err := http.NewRequest("POST", viper.GetString("integration.url_by_sms_center"), bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Println("ERROR", err)
		return http.StatusBadRequest, AnotherError("oshibka v integratsii sms senter")
	}

	// Добавляем content-type и авторизационное данные к заголовку запроса
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", viper.GetString("integration.sms_center"))

	// Отправляем запрос
	client := &http.Client{}
	response, err := client.Do(req)
	if response.StatusCode != http.StatusOK {
		var result map[string]interface{}
		json.NewDecoder(response.Body).Decode(&result)
		return http.StatusBadRequest, AnotherError("oshibka v decodinge")
	}
	if response.StatusCode == http.StatusOK {
		log.Println("INFO", "sms на номер:", numberPhone, "отправлено")
	}

	return http.StatusOK, nil
}
