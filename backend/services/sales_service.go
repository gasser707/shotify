package services

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gasser707/go-gql-server/custom"
	dbModels "github.com/gasser707/go-gql-server/databases/models"
	"github.com/gasser707/go-gql-server/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)


type SalesServiceInterface interface {
	BuyImage(ctx context.Context, input model.BuyImageInput) (*custom.Sale, error)
	GetSales(ctx context.Context)([]*custom.Sale,error)
}

//SalesService implements the usersServiceInterface
var _ SalesServiceInterface = &salesService{}

type salesService struct {
	DB            *sql.DB
	AuthService   AuthServiceInterface
}

func NewSalesService(db *sql.DB, authSrv AuthServiceInterface) *salesService {
	return &salesService{DB: db, AuthService: authSrv}
}

func(s *salesService) BuyImage(ctx context.Context, input model.BuyImageInput)(*custom.Sale, error){
	userId, _, err := s.AuthService.validateCredentials(ctx)
	if err != nil {
		return nil, err
	}
	imgId,err:= strconv.Atoi(input.ImageID)
	if err != nil {
		return nil, err
	}

	img,err:= dbModels.FindImage(ctx, s.DB, imgId)
	if err != nil || !img.ForSale || img.UserID == int(userId) {
		return nil, fmt.Errorf("image doesn't exist or can't be sold to you")
	}
	sale := &dbModels.Sale{
		Price: input.Price,
		ImageID: imgId,
		BuyerID: int(userId),
		SellerID:img.UserID,		
	}
	err= sale.Insert(ctx,s.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &custom.Sale{
		Price: sale.Price,
		ImageID: input.ImageID,
		BuyerID: fmt.Sprintf("%v",userId),
		SellerID: fmt.Sprintf("%v",sale.SellerID),
		Time: &sale.CreatedAt,
		ID: fmt.Sprintf("%v",sale.ID),
	},nil
}

func(s *salesService) GetSales(ctx context.Context) ([]*custom.Sale, error){
	userId, _, err := s.AuthService.validateCredentials(ctx)
	if err != nil {
		return nil, err
	}
	dbSales,err:= dbModels.Sales(qm.Where("buyer_id = ? or seller_id = ?", userId, userId)).All(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	sales:= []*custom.Sale{}
	for _, s:= range dbSales{
		sale:= &custom.Sale{
			ID: fmt.Sprintf("%v",s.ID),
			Time: &s.CreatedAt,
			ImageID: fmt.Sprintf("%v",s.ImageID),
			BuyerID: fmt.Sprintf("%v",s.BuyerID),
			SellerID: fmt.Sprintf("%v",s.SellerID),
			Price: s.Price,
		}
		sales= append(sales, sale)
	}

	return sales, nil
	
}