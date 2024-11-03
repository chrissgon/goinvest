package fund

type FundSearchRepo interface {
	Run(ID string) (FundEntity, error)
}
