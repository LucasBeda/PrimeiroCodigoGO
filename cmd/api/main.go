package main

import (
	"net/http"

	"github.com/devfullcycle/go-intensivo-jul/Internal/entity"
	"github.com/labstack/echo/v4"
)

func main() {
	//usando chi
	//r := chi.NewRouter()
	//r.Use(middleware.Logger)
	////http.HandleFunc("/order", OrderHandler) - Opção sem informar se é get/post...
	//r.Get("/order", OrderHandler)
	//http.ListenAndServe(":8888", r)

	//usando echo
	e := echo.New()
	e.GET("/order", OrderHandler)
	e.Logger.Fatal(e.Start(":8888"))
}

//metodo para o chi
//func OrderHandler(w http.ResponseWriter, r *http.Request) {
//	//maneira de fazer validação de método de requisição
//	//if r.Method != http.MethodGet {
//	//	w.WriteHeader(http.StatusMethodNotAllowed)
//	//	return
//	//}
//	order, _ := entity.NewOrder("4", 10, 1)
//	err := order.CalculateFinalPrice()
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//	}
//	json.NewEncoder(w).Encode(order)
//}

// metodo para utilizacao do echo
func OrderHandler(c echo.Context) error {
	order, _ := entity.NewOrder("5", 11, 3)
	err := order.CalculateFinalPrice()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, order)
}
