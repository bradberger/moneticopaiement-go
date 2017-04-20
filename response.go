package moneticopaiement

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// 3D secure status results
var (
	Status3DSNotMade      Status3DS = -1
	Status3DSInvalid      Status3DS = 0
	Status3DSLowRisk      Status3DS = 1
	Status3DSIncomplete   Status3DS = 2
	Stauts3DSHighRisk     Status3DS = 3
	Status3DSVeryHighRisk Status3DS = 4
)

// Credit Card types
var (
	CardBrandAmericanExpress CardBrand = "AM"
	CardBrandUnknown         CardBrand = "na"
	CardBrandVisa            CardBrand = "VI"
	CardBrandGIECB           CardBrand = "CB"
	CardBrandMasterCard      CardBrand = "MC"
)

// Payment results
var (
	PaymentResultPaid     PaymentResult = "paiement"
	PaymentResultTest     PaymentResult = "payetest"
	PaymentResulCancelled PaymentResult = "Annulation paiement refusé"
)

type Status3DS int

type CardBrand string

type PaymentResult string

type PaymentResultDate string

func (d PaymentResultDate) Time() time.Time {
	t, _ := time.Parse("02/01/2006_PM_03:04:05", string(d))
	return t
}

type PaymentResponse struct {
	MAC        string
	Date       PaymentResultDate
	TPE        string
	Amount     string
	Reference  string
	FreeText   string
	ReturnCode string
	CVX        bool
	VLD        string
	Brand      CardBrand
	Status3DS  Status3DS
}

func (p *PaymentResponse) ParseFromRequest(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	return p.ParseFromForm(r.Form)
}

func (p *PaymentResponse) ParseFromForm(form url.Values) error {

	if form == nil {
		return errors.New("form is empty")
	}

	p.Amount = form.Get("montant")
	p.Date = PaymentResultDate(form.Get("date"))
	p.MAC = form.Get("MAC")
	p.TPE = form.Get("TPE")
	p.Reference = form.Get("reference")
	p.FreeText = form.Get("texte-libre")
	p.ReturnCode = form.Get("code-retour")
	p.CVX = form.Get("cvx") == "oui"
	p.VLD = form.Get("vld")
	p.Brand = CardBrand(form.Get("brand"))
	s, err := strconv.Atoi(form.Get("status3ds"))
	if err != nil {
		return err
	}
	p.Status3DS = Status3DS(s)

	return nil
}

// Currency returns the three-letter currency of the payment
func (p *PaymentResponse) Currency() string {
	l := len(p.Amount)
	if l < 4 {
		return ""
	}
	return strings.ToLower(p.Amount[l-3:])
}

// Value returns a float64 representation of the payment amount
func (p *PaymentResponse) Value() float64 {
	l := len(p.Amount)
	if l < 4 {
		return 0
	}
	v, _ := strconv.ParseFloat(p.Amount[0:l-3], 64)
	return v
}
