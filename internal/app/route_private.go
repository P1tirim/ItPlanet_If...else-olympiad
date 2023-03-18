package app

import (
	"api/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(router *fiber.App, handlers *http.Handlers) {

	router.Get("/accounts/search", handlers.CheckAuthForFun, handlers.GetAccountSearch)
	router.Get("/accounts/:accountId", handlers.CheckAuthForFun, handlers.GetAccountById)

	router.Get("/animals/search", handlers.CheckAuthForFun, handlers.GetAnimalSearch)
	router.Get("/animals/:animalId", handlers.CheckAuthForFun, handlers.GetAnimalById)
	router.Get("/animals/:animalId/locations", handlers.CheckAuthForFun, handlers.GetVisitedLocations)

	router.Get("/animals/types/:typeId", handlers.CheckAuthForFun, handlers.GetTypeOfAnimal)

	router.Get("/locations/:pointId", handlers.CheckAuthForFun, handlers.GetLocations)

	accounts := router.Group("/accounts", handlers.CheckAuth())
	animals := router.Group("/animals", handlers.CheckAuth())
	animalsTypes := router.Group("/animals/types", handlers.CheckAuth())
	locations := router.Group("/locations", handlers.CheckAuth())

	accounts.Delete("/:accountId", handlers.CheckOwnAccount, handlers.DeleteAccount)
	accounts.Put("/:accountId", handlers.CheckOwnAccount, handlers.UpdateAccount)

	animals.Put("/:animalId/types", handlers.UpdateTypeOfAnimal)
	animals.Delete("/:animalId/types/:typeId", handlers.DeleteTypeOfAnimal)
	animals.Post("/:animalId/types/:typeId", handlers.AddTypeToAnimal)
	animals.Delete("/:animalId", handlers.DeleteAnimal)
	animals.Put("/:animalId", handlers.UpdateAnimal)
	animals.Delete("/:animalId/locations/:visitedPointId", handlers.DeleteVisitedLocation)
	animals.Put(":animalId/locations", handlers.UpdateVisitedLocation)
	animals.Post("/:animalId/locations/:pointId", handlers.AddVisitedLocation)
	animals.Post("", handlers.AddAnimal)

	animalsTypes.Delete("/:typeId", handlers.DeleteType)
	animalsTypes.Put("/:typeId", handlers.UpdateType)
	animalsTypes.Post("", handlers.AddType)

	locations.Delete("/:pointId", handlers.DeleteLocation)
	locations.Put("/:pointId", handlers.UpdateLocations)
	locations.Post("", handlers.AddLocation)

}
