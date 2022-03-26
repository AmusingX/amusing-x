package model

import (
	charon2 "amusingx.fit/amusingx/mysqlstruct/charon"
	"amusingx.fit/amusingx/services/charon/mysql/charon"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestProductInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product := &charon2.Product{
		Name: "name",
		Desc: "desc",
	}
	product, err := ProductInsert(context.Background(), product)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("product: %s", logger.ToJson(product))
}

func TestProductQueryById(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product, err := ProductQueryById(context.Background(), 3)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("product: %s", logger.ToJson(product))
}

func TestProductDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	err := ProductDelete(context.Background(), []int64{3})
	if err != nil {
		t.Fatalf("some err: %s", err)
	}
}

func TestProductUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	charon.Mock()

	product := &charon2.Product{
		ID:   3,
		Name: "name999",
		Desc: "desc999",
	}
	err := ProductUpdate(context.Background(), product)
	if err != nil {
		t.Fatalf("some err: %s", err)
	}

	t.Logf("product: %s", logger.ToJson(product))
}
