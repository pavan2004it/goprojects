package bookstore_test

import (
	"bookstore"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"sort"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2}
	want := 1
	result, err := bookstore.Buy(b)
	got := result.Copies
	if err != nil {
		t.Fatalf("Buy(%d): want no error for valid input, got %v", b.Copies, err)
	}
	if want != got {
		t.Errorf("want %d copies after buying "+
			"1 copy from a stock of 2, got %d", want, got)
	}

}

func TestInvalidBook(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("Negative values not supported")
	}

}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := bookstore.ExtendedCatalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}
	got := catalog.GetAllBooks()
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}

}
func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {Title: "For the Love of Go", ID: 1},
		2: {Title: "The Power of Go: Tools", ID: 2},
	}
	want := bookstore.Book{Title: "The Power of Go: Tools", ID: 2}
	got, err := catalog.GetBook(2)

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestGetBookBadIDReturnsError(t *testing.T) {
	catalog := bookstore.Catalog{}
	_, err := catalog.GetBook(999)
	if err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:           "For the Love of Go",
		PriceCents:      4000,
		DiscountPercent: 25,
	}

	want := 3000
	got := b.NetPriceCents()
	if want != got {
		t.Errorf("with price %d, after %d%% discount want net %d, got %d",
			b.PriceCents, b.DiscountPercent,
			want, got)
	}

}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "For the Love of Go",
		PriceCents: 4000,
	}
	want := 9000
	err := b.SetPriceCents(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.PriceCents
	if want != got {
		t.Errorf("with price %d, after setting the price to %d, got %d", b.PriceCents, want, got)
	}
}

func TestPriceCentsInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "For the Love of Go",
		PriceCents: 4000,
	}
	err := b.SetPriceCents(-1)
	if err == nil {
		t.Fatal("wanted error setting invalid price -1, got nil")
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{Title: "A midsummer nights dream"}
	cats := []bookstore.Category{
		bookstore.CategoryAutobiography,
		bookstore.CategoryParticlePhysics,
		bookstore.CategoryLargePrintRomance,
	}
	for _, cat := range cats {
		err := b.SetCategory(cat)
		if err != nil {
			t.Fatal(err)
		}
		want := cat
		got := b.Category()
		if want != got {
			t.Errorf("wanted %q, got %q", want, got)
		}
	}
}

func TestInvalidCategory(t *testing.T) {
	b := bookstore.Book{Title: "A tale of two cities"}
	err := b.SetCategory(99)
	if err == nil {
		t.Fatal("expected err got nil")
	}
}
