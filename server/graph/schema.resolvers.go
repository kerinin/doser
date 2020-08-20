package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/google/uuid"
	"github.com/kerinin/doser/service/graph/generated"
	"github.com/kerinin/doser/service/graph/model"
)

const (
	firmataPrefix = "firmata"
)

func (r *mutationResolver) CreateFirmata(ctx context.Context, input model.NewFirmataInput) (*model.Firmata, error) {
	f := &model.Firmata{
		ID:         uuid.New().String(),
		SerialPort: input.SerialPort,
	}

	err := r.db.Update(func(txn *badger.Txn) error {
		var key = fmt.Sprintf("%s-%s", firmataPrefix, f.ID)

		value, err := json.Marshal(f)
		if err != nil {
			return fmt.Errorf("serializing firmata: %w", err)
		}

		return txn.Set([]byte(key), value)
	})

	return f, err
}

func (r *mutationResolver) CreatePump(ctx context.Context, input model.NewPumpInput) (*model.Pump, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CalibratePump(ctx context.Context, input model.CalibratePumpInput) (*model.Pump, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateWaterLevelSensor(ctx context.Context, input model.CreateWaterLevelSensor) (*model.WaterLevelSensor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAutoTopOff(ctx context.Context, input model.NewAutoTopOff) (*model.AutoTopOff, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAutoWaterChange(ctx context.Context, input model.NewAutoWaterChangeInput) (*model.AutoWaterChange, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateDosers(ctx context.Context, input model.NewDosersInput) (*model.Dosers, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Firmata(ctx context.Context) (*model.Firmata, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Pumps(ctx context.Context) ([]*model.Pump, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) WaterLevelSensors(ctx context.Context) ([]*model.WaterLevelSensor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AutoTopOff(ctx context.Context) ([]*model.AutoTopOff, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AutoWaterChanges(ctx context.Context) ([]*model.AutoWaterChange, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Dosers(ctx context.Context) ([]*model.Dosers, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
