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
	"github.com/superhero-match/consumer-delete-media/internal/cache"
	"github.com/superhero-match/consumer-delete-media/internal/config"
	"github.com/superhero-match/consumer-delete-media/internal/consumer"
	"github.com/superhero-match/consumer-delete-media/internal/db"
	"github.com/superhero-match/consumer-delete-media/internal/es"
)

// Reader holds all the data relevant.
type Reader struct {
	DB       *db.DB
	Cache    *cache.Cache
	ES       *es.ES
	Consumer *consumer.Consumer
}

// NewReader configures Reader.
func NewReader(cfg *config.Config) (r *Reader, err error) {
	dbs, err := db.NewDB(cfg)
	if err != nil {
		return nil, err
	}

	c, err := cache.NewCache(cfg)
	if err != nil {
		return nil, err
	}

	e, err := es.NewES(cfg)
	if err != nil {
		return nil, err
	}

	cs := consumer.NewConsumer(cfg)

	return &Reader{
		DB:       dbs,
		Cache:    c,
		ES:       e,
		Consumer: cs,
	}, nil
}
