package booking

import "github.com/gofiber/fiber/v2"

const schemaName = "booking.json"

func SchemaApp() *fiber.App {
	app := fiber.New()
	sch := DeclareSchema().DeclareOperation()
	app.Get(schemaName, func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(sch)
	})
	return app
}

type Schema struct {
	Model            string     `json:"model"`
	ModelDisplayName string     `json:"model_display_name"`
	AdditionalField  bool       `json:"additional_field"`
	IsCloneable      bool       `json:"is_cloneable"`
	Operations       Operations `json:"operations"`
}
type Operations struct {
	List   *Data `json:"list,omitempty"`
	Read   *Data `json:"read,omitempty"`
	Create *Data `json:"create,omitempty"`
	Update *Data `json:"update,omitempty"`
}
type Data struct {
	Method string  `json:"method"`
	Action string  `json:"action"`
	Fields []Field `json:"fields"`
}
type Field struct {
	Name        string     `json:"name"`
	DisplayName string     `json:"display_name"`
	Type        string     `json:"type"`
	Multi       bool       `json:"multi"`
	Rich        bool       `json:"rich"`
	IsClearable bool       `json:"is_clearable"`
	Reference   *Reference `json:"reference,omitempty"`
}
type Reference struct {
	Model   string `json:"model"`
	Field   string `json:"field"`
	Display string `json:"display"`
}

func DeclareSchema() *Schema {
	return &Schema{
		Model:            "booking",
		ModelDisplayName: "Бронирование",
		AdditionalField:  false,
		IsCloneable:      false,
		Operations:       Operations{},
	}
}
func (s *Schema) DeclareOperation() *Schema {
	s.Operations = Operations{
		List:   ListData(),
		Read:   ReadData(),
		Create: CreateData(),
		Update: UpdateData(),
	}
	return s
}
func ListData() *Data {
	return &Data{
		Method: "GET",
		Action: ":id",
		Fields: ListField(),
	}
}
func ReadData() *Data {
	return &Data{
		Method: "GET",
		Action: "",
		Fields: ReadField(),
	}
}
func CreateData() *Data {
	return &Data{
		Method: "POST",
		Action: "",
		Fields: CreateField(),
	}
}
func UpdateData() *Data {
	return &Data{
		Method: "PUT",
		Action: ":id",
		Fields: UpdateField(),
	}
}
func ListField() []Field {
	return []Field{
		{Name: "user", DisplayName: "Наименование", Type: "string", Multi: false, Rich: false, IsClearable: false},
	}
}
func ReadField() []Field {
	return []Field{
		{Name: "user", DisplayName: "Наименование", Type: "string", Multi: false, Rich: false, IsClearable: false},
	}
}
func CreateField() []Field {
	return []Field{
		{Name: "user", DisplayName: "Наименование", Type: "string", Multi: false, Rich: false, IsClearable: false},
	}
}
func UpdateField() []Field {
	return []Field{
		{Name: "user", DisplayName: "Наименование", Type: "string", Multi: false, Rich: false, IsClearable: false},
	}
}
