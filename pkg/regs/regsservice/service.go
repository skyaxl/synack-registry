package regsservice

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/kataras/golog"
	"github.com/skyaxl/synack-registry/db/pkg/models"
	"github.com/skyaxl/synack-registry/pkg/regs/regscontract"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

//Service instance of regservice
type Service struct {
	db boil.ContextExecutor
}

//New Service
func New(db boil.ContextExecutor) *Service {
	return &Service{db}
}

//Reg register a request
func (svc *Service) Reg(ctx context.Context, reg *regscontract.Request) (err error) {
	id, _ := uuid.NewGen().NewV4()
	regNew := models.Registry{
		ID:       null.StringFrom(id.String()),
		UserName: null.StringFrom(reg.Username),
		Request:  null.BytesFrom(reg.DumpReq),
		Response: null.BytesFrom(reg.DumpRes),
	}

	golog.Infof("Inserting registry for user: %s", reg.Username)
	if err = regNew.Insert(ctx, svc.db, boil.Infer()); err != nil {
		golog.Errorf("Error on inserting registry for user: %s, err: %v, dumps %s %s", reg.Username, err, string(reg.DumpReq), string(reg.DumpRes))
		return err
	}
	return
}
