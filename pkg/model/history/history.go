package history

import (
	"github.com/gofiber/fiber/v2"
	"gorm-rest/pkg/ftools/fdb"
	"gorm-rest/pkg/ftools/flog"
	"strings"
	"time"
)

type History struct {
	ID        int       `json:"id" gorm:"primary_key" schema:"method:get,post,put"`
	Model     string    `json:"model" gorm:"type:varchar(40)"`
	Action    string    `json:"action" gorm:"type:varchar(15)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
type ContentStatistic struct {
	Email     string      `json:"email"`
	Statistic []Statistic `json:"statistic,omitempty"`
}
type Statistic struct {
	Date        string   `json:"date"`
	Update      []string `json:"update,omitempty"`
	Create      []string `json:"create,omitempty"`
	UpdateCount int      `json:"update_count"`
	CreateCount int      `json:"create_count"`
}

func New(c *fiber.Ctx) error {
	defer func(c *fiber.Ctx) {
		if c.Response().StatusCode() == fiber.StatusOK {
			h := History{Action: string(c.Request().Header.Method())}
			_ = h
			p := string(c.Request().URI().Path())

			pSlice := strings.Split(p, "/")
			if len(pSlice) > 1 {
				h.Model = pSlice[len(pSlice)-1]
			} else {
				h.Model = p
			}
			if result := fdb.Client.Model(&h).Create(&h).Error; result != nil {
				flog.ErrorCtx(c).Msg("error create new row history")
			}
		}
	}(c)
	return c.Next()
}
