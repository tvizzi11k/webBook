package recommendations_

import (
	"github.com/sajari/regression"
	"webBooks/internal/models "
)

type Recommender struct {
	Books []*models_.Books
}

func NewRecommender(books []*models_.Books) *Recommender {
	return &Recommender{Books: books}
}

func (r *Recommender) Recommend(prefs map[string]float64) []*models_.Books {
	var regressionModel regression.Regression
	var recommendedBooks []*models_.Books

	regressionModel.SetObserved("Rating")
	regressionModel.SetVar(0, "Genre")
	regressionModel.SetVar(1, "Length")

	for _, book := range r.Books {
		genreScore := prefs[book.Genre]
		regressionModel.Train(
			regression.DataPoint(
				book.Rating,
				[]float64{genreScore, float64(len(book.Description))},
			),
		)
	}
	regressionModel.Run()

	for _, book := range r.Books {
		genreScore := prefs[book.Genre]
		predicted, _ := regressionModel.Predict([]float64{genreScore, float64(len(book.Description))})

		if predicted > 3.5 {
			recommendedBooks = append(recommendedBooks, book)
		}
	}

	return recommendedBooks
}
