package model

import (
	"context"
	"testing"
)

func TestMovieDao_FindMovies(t *testing.T) {
	dao := NewMovieDao()
	ids := []string{"100000000000000000000004", "100000000000000000000010"}
	res, err := dao.FindMovies(context.Background(), ids)
	// try hit cache
	for i := 0; i < 50; i++ {
		_, _ = dao.FindMovies(context.Background(), ids)
	}
	if err != nil {
		t.Fatal(err)
	}
	for _, movie := range res {
		t.Logf("%+v", movie)
		for _, participant := range movie.Participants {
			t.Logf("%+v", participant)
		}
	}
}
