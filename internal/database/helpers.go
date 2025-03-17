package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

func GetActiveUser(ctx context.Context, dbQueries *Queries) (GetUserByIDRow, error) {
	activeUserIDStr, err := dbQueries.GetActiveUserID(ctx)
	if err != nil {
		return GetUserByIDRow{}, fmt.Errorf("failed to get active user ID: %w", err)

	}

	activeUserID, err := strconv.ParseInt(activeUserIDStr, 10, 64)
	if err != nil {
		return GetUserByIDRow{}, fmt.Errorf("failed to convert active user ID to integer: %w", err)
	}

	user, err := dbQueries.GetUserByID(ctx, activeUserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return GetUserByIDRow{}, fmt.Errorf("no user found for active user ID: %w", err)
		}
		return GetUserByIDRow{}, fmt.Errorf("failed to retrieve user from DB: %w", err)
	}

	return user, nil
}
