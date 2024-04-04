package app

type App struct {
	repo Repository
}

func NewApp(repo Repository) *App {
	return &App{
		repo: repo,
	}
}
