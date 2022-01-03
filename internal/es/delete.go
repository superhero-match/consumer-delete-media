/*
  Copyright (C) 2019 - 2021 MWSOFT
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
	"github.com/superhero-match/consumer-delete-media/internal/es/model"
)

// DeleteProfilePicture deletes profile picture.
func (e *es) DeleteProfilePicture(superheroID string, position int64, deletedAt string) error {
	superhero, err := e.GetSuperhero(superheroID)
	if err != nil {
		return err
	}

	sourceID, err := e.GetDocumentID(superheroID)
	if err != nil {
		return err
	}

	// Delete profile picture at specified position.
	for i := 0; i < len(superhero.ProfilePictures); i++ {
		if position == superhero.ProfilePictures[i].Position {
			superhero.ProfilePictures = append(superhero.ProfilePictures[:i], superhero.ProfilePictures[i+1:]...)
		}
	}

	// After deleting picture at specified position, all existing pictures positions need to be updated.
	for i := 0; i < len(superhero.ProfilePictures); i++ {
		superhero.ProfilePictures[i].Position = int64(i) + int64(1)
	}

	return updateProfilePics(e, sourceID, superhero.ProfilePictures, deletedAt)
}

func updateProfilePics(e *es, sourceID string, pps []model.ProfilePicture, deletedAt string) error {
	_, err := e.Client.Update().
		Index(e.Index).
		Id(sourceID).
		Doc(map[string]interface{}{
			"profile_pics": pps,
			"updated_at":   deletedAt,
		}).
		Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
