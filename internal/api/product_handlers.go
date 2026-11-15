package api

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/rocketseat-education/gobid/internal/jsonutils"
	"github.com/rocketseat-education/gobid/internal/usecase/product"
)

func (api *Api) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] Iniciando creación de producto...")

	data, problemns, err := jsonutils.DecodeValidJson[product.CreateProductReq](r)
	if err != nil {
		log.Printf("[ERROR] Falló la validación o decodificación del JSON: %v | Problemas: %+v", err, problemns)
		jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problemns)
		return
	}
	log.Printf("[DEBUG] Datos decodificados correctamente: %+v", data)

	userID, ok := api.Sessions.Get(r.Context(), "AuthenticatedUserId").(uuid.UUID)
	if !ok {
		log.Println("[ERROR] No se pudo obtener el 'AuthenticatedUserId' de la sesión o el tipo no coincide con uuid.UUID")
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "unexpected error, try again later",
		})
		return
	}
	log.Printf("[INFO] Usuario autenticado detectado con ID: %s", userID)

	id, err := api.ProductService.CreateProduct(
		r.Context(),
		userID,
		data.ProductName,
		data.Description,
		data.Baseprice,
		data.AuctionEnd,
	)
	if err != nil {
		log.Printf("[ERROR] Falló al insertar el producto en la base de datos: %v", err)
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "failed to create product auction try again later",
		})
		return
	}

	log.Printf("[SUCCESS] Producto creado exitosamente con ID: %s", id)
	jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"message":    "product created with success",
		"product_id": id,
	})
}
