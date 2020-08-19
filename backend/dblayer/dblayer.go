package dblayer

type DBLayer interface {
	GetPlan(Plan models.Plan)
}
