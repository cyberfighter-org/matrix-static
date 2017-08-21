// Copyright 2017 Michael Telatynski <7t3chguy@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import log "github.com/Sirupsen/logrus"

type RoomInitialSyncResp struct {
	err error
}

type RoomInitialSyncJob struct {
	roomID string
}

func (job RoomInitialSyncJob) Work(w *Worker) {
	resp := &RoomInitialSyncResp{}

	if _, exists := w.rooms[job.roomID]; !exists {
		loggerWithFields := log.WithField("worker", w.ID).WithField("roomID", job.roomID)
		loggerWithFields.Info("Started Initial Syncing Room")
		if newRoom, err := w.client.NewRoom(job.roomID); err == nil {
			loggerWithFields.Info("Finished Initial Syncing Room")
			w.rooms[job.roomID] = newRoom
		} else {
			loggerWithFields.WithError(err).Error("Failed Initial Syncing Room")
			resp.err = err
		}
	}

	w.Output <- resp
}
