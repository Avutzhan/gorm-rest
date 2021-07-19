package booking

import (
	"encoding/json"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"gorm-rest/pkg/ftools/fdb"
	"gorm-rest/pkg/ftools/flog"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

type Booking struct {
	User    string `json:"user"`
	Members int    `json:"members"`
}

func (Booking) DisplayName() string {
	return "Бронирование"
}

func (b Booking) ValidateCreate() interface{} {
	return validation.ValidateStruct(&b,
		validation.Field(&b.User, validation.Required.Error("Значение обязательное")),
	)
}

func (b Booking) ValidateUpdate() interface{} {
	return validation.ValidateStruct(&b,
		validation.Field(&b.User, validation.Required.Error("Значение обязательное")),
	)
}

func App() *fiber.App {
	app := fiber.New()
	app.Get("", List)       //tested http://127.0.0.1:8787/v1/football/booking
	app.Get("/:id", Select) //tested http://127.0.0.1:8787/v1/football/booking/21
	app.Post("", Create)    //tested http://127.0.0.1:8787/v1/football/booking
	app.Put("/:id", Update) //tested http://127.0.0.1:8787/v1/football/booking/21
	return app
}

func Select(c *fiber.Ctx) error {
	var booking Booking

	//выбираем бронь по приходящему айдишнику
	if result := fdb.Client.Model(&booking).Where("id = ?", c.Params("id")).Take(&booking).Error; result != nil {
		if errors.Is(result, gorm.ErrRecordNotFound) {
			flog.WarnCtx(c).Err(result).Msg("booking not found")

			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Booking не найден"})
		}
		flog.ErrorCtx(c).Err(result).Msg("failed find booking")

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Внутренние проблемы"})
	}
	return c.JSON(booking)
}

func List(c *fiber.Ctx) error {
	var booking []Booking

	//выводим в массиве все элементы брони
	if result := fdb.Client.Model(&Booking{}).Find(&booking).Error; result != nil {
		flog.ErrorCtx(c).Err(result).Msg("error find all booking")

		return c.JSON(fiber.Map{"message": "Внутренняя ошибка поиска списка booking"})
	}
	return c.JSON(booking)
}

func Create(c *fiber.Ctx) error {
	booking := new(Booking)

	//проверка и валидация реквеста
	if err := c.BodyParser(booking); err != nil {
		flog.ErrorCtx(c).Err(err).Msg("failed parse create request")

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка обработки запроса"})
	}
	if result := booking.ValidateCreate(); result != nil {
		flog.ErrorCtx(c).Interface("error", result).Msg("request not validate")

		return c.Status(fiber.StatusBadRequest).JSON(result)
	}

	flog.ErrorCtx(c).Msg(booking.User)
	//ищем дубликаты по имени
	duplicate := new(Booking)
	if result := fdb.Client.Model(duplicate).Where("user = ?", booking.User).Take(&duplicate).Error; result != nil {
		if !errors.Is(result, gorm.ErrRecordNotFound) {
			flog.ErrorCtx(c).Err(result).Msg("failed find users duplicate")

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Внутрение проблемы"})
		}
	} else {
		flog.WarnCtx(c).Str("error", "duplicate row").Msg("find duplicate")

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"name": "Такой user уже есть"})
	}

	//если нет дубликатов то создаем бронь
	if result := fdb.Client.Model(&booking).Create(&booking).Error; result != nil {
		flog.ErrorCtx(c).Err(result).Msg("failed create new Booking")

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка создания Booking"})
	}

	return c.JSON(booking)
}

func Update(c *fiber.Ctx) error {
	booking := new(Booking)

	//проверка и валидация реквеста
	if err := c.BodyParser(booking); err != nil {
		flog.ErrorCtx(c).Err(err).Msg("failed parse create request")

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка обработки запроса"})
	}
	if result := booking.ValidateCreate(); result != nil {
		flog.ErrorCtx(c).Interface("error", result).Msg("request not validate")

		return c.Status(fiber.StatusBadRequest).JSON(result)
	}

	//Апдейтим бронь
	if result := fdb.Client.Model(&booking).Where("id = ?", c.Params("id")).Updates(&booking).Take(&booking).Error; result != nil {
		flog.ErrorCtx(c).Err(result).Msg("failed update assignment")

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка обновления assignment"})
	}

	return c.JSON(booking)
}

func GetFromDMS(url string) ([]Booking, error) {
	client := http.Client{Timeout: 60 * time.Second}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	switch response.StatusCode {
	case 200:
	default:
		return nil, errors.New(response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var res []Booking
	if err = json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return res, nil
}
