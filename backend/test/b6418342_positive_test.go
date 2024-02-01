package test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/kasama03/sut-final-lab/entity"
	"github.com/onsi/gomega"
)

func TestCustomer(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run(`save success`, func(t *testing.T) {
		customer := entity.Customers{
			Name:       "b",
			Age:        "12",
			CustomerID: "CM00000000",
		}
		ok, err := govalidator.ValidateStruct(customer)

		g.Expect(ok).To(gomega.BeTrue())
		g.Expect(err).To(gomega.BeNil())
	})
}
