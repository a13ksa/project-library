package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`                               // ID dokumenta u bazi (automatski generisan)
	InventoryNumber   int                `json:"inv_number" bson:"inv_number"`                                     // Inventarski broj
	Signature         string             `json:"signature" bson:"signature"`                                       // Signatura (lokacija knjige u biblioteci)
	Authors           []string           `json:"authors" bson:"authors"`                                           // Lista autora
	Title             string             `json:"title" bson:"title"`                                               // Naslov knjige
	Genres            []string           `json:"genres" bson:"genres"`                                             // Lista žanrova
	PublicationPlace  *string            `json:"publication_place,omitempty" bson:"publication_place,omitempty"`   // Mesto objavljivanja
	PublicationYear   *int               `json:"publication_year,omitempty" bson:"publication_year,omitempty"`     // Godina objavljivanja
	Publisher         *string            `json:"publisher,omitempty" bson:"publisher,omitempty"`                   // Izdavač
	AcquisitionMethod *string            `json:"acquisition_method,omitempty" bson:"acquisition_method,omitempty"` // Način nabavke
	BindingType       *string            `json:"binding_type,omitempty" bson:"binding_type,omitempty"`             // Tip poveza (npr. tvrdi, meki)
	ISBN              *string            `json:"isbn,omitempty" bson:"isbn,omitempty"`                             // ISBN broj
	Dimensions        *string            `json:"dimensions,omitempty" bson:"dimensions,omitempty"`                 // Dimenzije knjige
	Price             *float64           `json:"price,omitempty" bson:"price,omitempty"`                           // Cena
	Note              *string            `json:"note,omitempty" bson:"note,omitempty"`                             // Napomena
	UDCNumber         *string            `json:"udc_number,omitempty" bson:"udc_number,omitempty"`                 // UDK broj (Univerzalna decimalna klasifikacija)
	Edition           *string            `json:"edition,omitempty" bson:"edition,omitempty"`                       // Izdanje
	Description       *string            `json:"description,omitempty" bson:"description,omitempty"`               // Opis knjige
}
