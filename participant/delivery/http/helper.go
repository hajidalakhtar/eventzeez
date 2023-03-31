package http

import (
	"github.com/google/uuid"
)

func stringToUUIDS(bulkID []string) []uuid.UUID {
	var uuids []uuid.UUID
	for _, id := range bulkID {
		uuid, _ := uuid.Parse(id)
		uuids = append(uuids, uuid)
	}
	return uuids
}
