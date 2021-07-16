package history

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm-rest/pkg/ftools/fdb"
	"gorm-rest/pkg/ftools/flog"
	"strings"
	"time"
)

type History struct {
	ID        string    `json:"id" gorm:"type:uuid;default:uuid_generate_v4()" schema:"method:get,post,put"`
	Email     string    `json:"email" gorm:"type:varchar(100)"`
	Model     string    `json:"model" gorm:"type:varchar(40)"`
	Action    string    `json:"action" gorm:"type:varchar(15)"`
	ModelID   string    `json:"model_id" gorm:"type:varchar(36)"`
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
			h := History{Email: string(c.Request().Header.Peek("X-User-Email")), Action: string(c.Request().Header.Method()), ModelID: c.Params("id")}
			_ = h
			p := string(c.Request().URI().Path())
			if h.ModelID != "" {
				p = strings.ReplaceAll(p, fmt.Sprintf("/%s", h.ModelID), "")
			}
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
