/*
  Copyright (C) 2019 - 2020 MWSOFT
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
package reader

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/superhero-match/consumer-delete-media/internal/consumer/model"
	dbm "github.com/superhero-match/consumer-delete-media/internal/db/model"
)

// Read consumes the Kafka topic and stores the newly registered superhero to DB and Elasticsearch.
func (r *Reader) Read() error {
	ctx := context.Background()

	for {
		fmt.Print("before FetchMessage")
		m, err := r.Consumer.Consumer.FetchMessage(ctx)
		fmt.Print("after FetchMessage")
		if err != nil {
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}

		fmt.Printf(
			"message at topic/partition/offset \n%v/\n%v/\n%v: \n%s = \n%s\n",
			m.Topic,
			m.Partition,
			m.Offset,
			string(m.Key),
			string(m.Value),
		)

		var pp model.ProfilePicture
		if err := json.Unmarshal(m.Value, &pp); err != nil {
			_ = r.Consumer.Consumer.Close()
			if err != nil {
				fmt.Println("Unmarshal")
				fmt.Println(err)
				err = r.Consumer.Consumer.Close()
				if err != nil {
					return err
				}

				return err
			}
		}

		err = r.DB.DeleteProfilePicture(dbm.ProfilePicture{
			SuperheroID: pp.SuperheroID,
			Position:    pp.Position,
			DeletedAt:   pp.DeletedAt,
		}, )
		if err != nil {
			fmt.Println("DB")
			fmt.Println(err)
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}

		err = r.ES.DeleteProfilePicture(pp.SuperheroID, pp.Position, pp.DeletedAt)
		if err != nil {
			fmt.Println("ES")
			fmt.Println(err)
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}

		keys := make([]string, 0)
		keys = append(keys, fmt.Sprintf(r.Cache.SuggestionKeyFormat, pp.SuperheroID))

		err = r.Cache.DeleteSuperhero(keys)
		if err != nil {
			fmt.Println("Cache")
			fmt.Println(err)
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}

		err = r.Consumer.Consumer.CommitMessages(ctx, m)
		if err != nil {
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}
	}
}
