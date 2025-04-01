package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(r *echo.Echo) {
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	g := r.Group("/api")

	// TODO
	auth := g.Group("/auth")
	_ = auth

	// TODO
	tasks := g.Group("/tasks")
	_ = tasks

	// TODO
	tests := g.Group("/test")
	_ = tests

	// TODO
	docs := g.Group("/docs")
	_ = docs

	// TODO
	lib := g.Group("/library")
	_ = lib

	// TODO
	grades := g.Group("/grades")
	_ = grades

	// TODO
	news := g.Group("/news")
	_ = news

	// TODO
	calendar := g.Group("/calendar")
	_ = calendar
}
