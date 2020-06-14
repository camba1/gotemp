package globalUtils

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/metadata"
	"log"
	"strconv"
)

type authUtils struct{}

//getCurrentUserFromContext: User is added to the context during authentication. this function extracts it so
//that it can be used sending audit records to the broker
func (a *authUtils) GetCurrentUserFromContext(ctx context.Context) (int64, error) {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return 0, fmt.Errorf("unable to get user from metadata")
	}
	userId, err := strconv.ParseInt(meta["Userid"], 10, 64)
	if err != nil {
		return 0, err
	}
	log.Printf("userid from metadata: %d\n", userId)
	return userId, nil
}
