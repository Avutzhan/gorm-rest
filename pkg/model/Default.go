package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

type Default struct {
	ID        string    `json:"id" gorm:"->;type:int;" schema:"method:get,post,put"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (a Default) UpdatesMap() map[string]interface{} {
	if a.ID == "" {
		return map[string]interface{}{
			//Write this
		}
	} else {
		return map[string]interface{}{
			//Write this
		}
	}
}

func (Default) DisplayName() string {
	return "Страны"
}
func (c *Default) AfterFind(tx *gorm.DB) (err error) {
	//c.Hex, _ = converter.ConvertInt(fmt.Sprint(c.Number), 10, 30)
	return
}
func (c Default) ValidateCreate() interface{} {
	return validation.ValidateStruct(&c) //validation.Field(&c.ID, validation.Required.Error("Значение обязательное")),

}
func (c Default) ValidateUpdate() interface{} {
	return validation.ValidateStruct(&c) //validation.Field(&c.Name, validation.Required.Error("Значение обязательное")),

}

func App() *fiber.App {

	app := fiber.New()
	app.Get("", nil)
	app.Get("/:id", nil)
	app.Post("", nil)
	app.Put("/:id", nil)

	return app
}
func ParamValidate(c *fiber.Ctx) error {

	return nil
}
func Select(c *fiber.Ctx) error {
	return nil
}
func List(c *fiber.Ctx) error {

	return nil
}
func Create(c *fiber.Ctx) error {

	return nil
}
func Update(c *fiber.Ctx) error {
	return nil
}
