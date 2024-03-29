/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package es

import (
	"context"
	"encoding/json"
	"fmt"

	elastic "github.com/olivere/elastic/v7"

	"github.com/superhero-match/consumer-delete-media/internal/es/model"
)

// GetSuperhero returns Superhero by id.
func (e *es) GetSuperhero(superheroID string) (s *model.Superhero, err error) {
	q := elastic.NewTermQuery("superhero_id", superheroID)

	fmt.Println()
	fmt.Printf("%+v", q)
	fmt.Println()

	searchResult, err := e.Client.Search().
		Index(e.Index).
		Query(q).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	fmt.Printf("SearchResult: %+v", searchResult)

	fmt.Println()

	fmt.Println(searchResult.TotalHits())

	if searchResult.TotalHits() > 0 {
		for _, hit := range searchResult.Hits.Hits {
			fmt.Printf("Hit: %+v", hit)

			err := json.Unmarshal(hit.Source, &s)
			if err != nil {
				return nil, err
			}

			fmt.Println()
			fmt.Printf("Superhero Unmarshalled: %+v", &s)
			fmt.Println()
		}
	}

	return s, nil
}
