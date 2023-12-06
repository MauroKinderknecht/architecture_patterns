package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/MauroKinderknecht/tech_talk/microservices/item_service/db"
	"github.com/google/uuid"
)

func (b *Backend) ListItems(ctx context.Context, request ListItemsRequestObject) (ListItemsResponseObject, error) {
	items, err := b.db.ListItems(ctx)
	fmt.Println(err)
	if err != nil {
		return ListItems500JSONResponse{Message: "Failed to retrieve items"}, nil
	}

	var res []Item
	for _, item := range items {
		res = append(res, Item{
			Id:    item.ID,
			Name:  item.Name,
			Price: item.Price,
		})
	}
	return ListItems200JSONResponse(res), nil
}

func (b *Backend) CreateItem(ctx context.Context, request CreateItemRequestObject) (CreateItemResponseObject, error) {
	item := &db.Item{
		ID:    generateID(),
		Name:  request.Body.Name,
		Price: request.Body.Price,
	}

	err := b.db.SaveItem(ctx, item)
	if err != nil {
		return CreateItem500JSONResponse{Message: "failed to save item"}, nil
	}

	res := Item{
		Id:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}
	return CreateItem201JSONResponse(res), nil
}

func (b *Backend) DeleteItem(ctx context.Context, request DeleteItemRequestObject) (DeleteItemResponseObject, error) {
	err := b.db.DeleteItem(ctx, request.ItemId)
	if err != nil {
		return DeleteItem500JSONResponse{Message: "failed to delete item"}, nil
	}

	return DeleteItem204Response{}, nil
}

func (b *Backend) GetItem(ctx context.Context, request GetItemRequestObject) (GetItemResponseObject, error) {
	item, err := b.db.GetItemById(ctx, request.ItemId)
	if err != nil {
		return GetItem500JSONResponse{Message: "failed to retrieve item"}, nil
	}

	res := Item{
		Id:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}
	return GetItem200JSONResponse(res), nil
}

func (b *Backend) UpdateItem(ctx context.Context, request UpdateItemRequestObject) (UpdateItemResponseObject, error) {
	item := &db.Item{
		ID:    request.ItemId,
		Name:  request.Body.Name,
		Price: request.Body.Price,
	}

	err := b.db.SaveItem(ctx, item)
	if err != nil {
		return UpdateItem500JSONResponse{Message: "failed to update item"}, nil
	}

	res := Item{
		Id:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}
	return UpdateItem200JSONResponse(res), nil
}

func generateID() string {
	return strings.ReplaceAll(uuid.Must(uuid.NewRandom()).String(), "-", "")
}
