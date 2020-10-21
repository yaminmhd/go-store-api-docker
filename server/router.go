package server

import (
	"github.com/yaminmhd/go-hardware-store/appcontext"
	"github.com/yaminmhd/go-hardware-store/handler"
	"github.com/yaminmhd/go-hardware-store/repository"
	"github.com/yaminmhd/go-hardware-store/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Dependency struct {
	Handlers     handlers
	Services     services
	Repositories repositories
}

type repositories struct {
	ProductRepository repository.Product
}
type services struct {
	ProductService service.ProductService
}

type handlers struct {
	ProductHandler handler.Product
}

func initRepositories() repositories {
	productRepository := repository.NewProductRepository(appcontext.GetDB())
	return repositories{
		ProductRepository: productRepository,
	}
}

func initServices(repositories repositories) services {
	productService := service.NewProductService(repositories.ProductRepository)
	return services{
		ProductService: productService,
	}
}

func initHandlers(services services) handlers {
	productHandler := handler.NewProductHandler(services.ProductService)
	return handlers{
		ProductHandler: productHandler,
	}
}

func Init() Dependency {
	repositories := initRepositories()
	services := initServices(repositories)
	handlers := initHandlers(services)
	return Dependency{
		Handlers: handlers,
		Services: services,
		Repositories: repositories,
	}
}

func Router(dependencies Dependency) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.Ping).Methods("GET")
	router.HandleFunc("/v1/products", dependencies.Handlers.ProductHandler.GetProducts).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)
	return router
}
