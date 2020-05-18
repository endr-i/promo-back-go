package engine

import "sync"

var app *App

type App struct {
	errorChan    chan error
	criticalChan chan error
}

func GetApp() *App {
	var once *sync.Once
	once.Do(func() {
		app = newApp()
	})
	return app
}

func newApp() *App {
	result := new(App)
	result.errorChan = make(chan error)
	result.criticalChan = make(chan error)
	return result
}

func (app App) Run() {

}
