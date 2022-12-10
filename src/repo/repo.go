package repo

type (
	YearMonth  string
	Amount     int64
	BudgetData struct {
		YearMonth YearMonth
		Amount    Amount
	}
)

// var BudgetDatas = map[YearMonth]Amount{
// 	"202201": 3100,
// }

type IBudgetRepo interface {
	GetAll() []BudgetData
}

type BudgetRepo struct {
	Data map[YearMonth]BudgetData
}

func (b *BudgetRepo) GetAll() (lst []BudgetData) {
	for _, v := range b.Data {
		lst = append(lst, v)
	}
	return
}

func NewBudgetRepo(budgetdata map[YearMonth]BudgetData) IBudgetRepo {
	return &BudgetRepo{
		Data: budgetdata,
	}
}
