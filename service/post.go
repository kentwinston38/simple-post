package service

import (
	"context"
	"myapp/config"
	"myapp/model"
	"time"
)

func PostCreate(ctx context.Context, input model.NewPost) (*model.Post, error) {
	var (
		db = config.GetDB()
	)

	post := model.Post{
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: time.Now().UTC(),
	}

	if err := db.Model(&post).Create(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func PostGetAll(ctx context.Context) ([]*model.Post, error) {
	var (
		db    = config.GetDB()
		posts []*model.Post
	)

	if err := db.Model(&posts).Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}
