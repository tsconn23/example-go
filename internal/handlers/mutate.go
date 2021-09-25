/*******************************************************************************
 * Copyright 2021 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package handlers

import (
	"context"
	"encoding/json"
	"github.com/project-alvarium/alvarium-sdk-go/pkg/config"
	"github.com/project-alvarium/alvarium-sdk-go/pkg/interfaces"
	"github.com/project-alvarium/example-go/internal/db"
	"github.com/project-alvarium/example-go/internal/models"
	logInterface "github.com/project-alvarium/provider-logging/pkg/interfaces"
	"github.com/project-alvarium/provider-logging/pkg/logging"
	"sync"
)

type Mutator struct {
	cfg         config.SdkInfo
	chSubscribe chan []byte
	chPublish   chan []byte
	db          db.MongoProvider
	logger      logInterface.Logger
	sdk         interfaces.Sdk
}

func NewMutator(sdk interfaces.Sdk, chSub chan []byte, chPub chan []byte, cfg config.SdkInfo, db db.MongoProvider, logger logInterface.Logger) Mutator {
	return Mutator{
		cfg:         cfg,
		chSubscribe: chSub,
		chPublish:   chPub,
		db:          db,
		logger:      logger,
		sdk:         sdk,
	}
}

func (m *Mutator) BootstrapHandler(ctx context.Context, wg *sync.WaitGroup) bool {
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			msg, ok := <-m.chSubscribe
			if ok {
				data, err := models.NewSampleData(m.cfg.Signature.PrivateKey)
				if err != nil {
					m.logger.Error(err.Error())
					continue
				}
				err = m.db.Save(ctx, data)
				if err != nil {
					m.logger.Error(err.Error())
					continue
				}

				b, _ := json.Marshal(data)
				m.sdk.Mutate(ctx, msg, b)
				m.chPublish <- b
			} else { //channel has been closed. End goroutine.
				m.logger.Write(logging.InfoLevel, "mutator::chSubscribe closed, exiting")
				close(m.chPublish)
				return
			}
		}
	}()

	wg.Add(1)
	go func() { // Graceful shutdown
		defer wg.Done()

		<-ctx.Done()
		m.logger.Write(logging.InfoLevel, "shutdown received")
	}()
	return true
}
