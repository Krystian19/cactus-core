package services

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Krystian19/cactus-core/models"
	"github.com/Krystian19/cactus-core/proto"
	"github.com/jinzhu/gorm"
)

// Episode : Get a single Episode based on the provided params
func (s *Server) Episode(ctx context.Context, request *proto.EpisodeRequest) (*proto.EpisodeResponse, error) {
	db, err := db()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()

	var result models.Episode
	query := db

	if request.Id != 0 {
		query = query.Where("id = ?", request.Id)
	}

	if err := query.First(&result).Error; err != nil {

		// If nothing was found
		if gorm.IsRecordNotFoundError(err) {
			return &proto.EpisodeResponse{Episode: nil}, nil
		}

		fmt.Println(err)
		return nil, err
	}

	return &proto.EpisodeResponse{Episode: result.Episode}, nil
}

// Episodes : Get a list of Episodes based on the provided params
func (s *Server) Episodes(ctx context.Context, request *proto.EpisodesRequest) (*proto.EpisodesResponse, error) {
	db, err := db()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()

	var result []models.Episode
	var resultCount uint
	query := db

	if request != nil && request.Query != nil {
		if request.Query.ReleaseId != 0 {
			query = query.Where("release_id = ?", request.Query.ReleaseId)
		}

		if request.Query.Limit != 0 {
			query = query.Limit(request.Query.Limit)
		}

		if request.Query.Offset != 0 {
			query = query.Offset(request.Query.Offset)
		}
	}

	if request != nil && request.OrderBy != nil {
		if len(strings.TrimSpace(request.OrderBy.Field)) != 0 {

			sorting := ""
			if request.OrderBy.Descending {
				sorting = DESC
			}

			// Example : SORT BY - "fieldname ASC"
			query = query.Order(fmt.Sprintf("%s %s", request.OrderBy.Field, sorting))
		}
	}

	if err := query.Find(&result).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Remove pagination constraints before counting
	query = query.Limit(nil)
	query = query.Offset(nil)

	if err := query.Model(&models.Episode{}).Count(&resultCount).Error; err != nil {
		resultCount = 0
	}

	finalRes := []*proto.Episode{}

	for i := range result {
		finalRes = append(finalRes, result[i].Episode)
	}

	return &proto.EpisodesResponse{Episodes: finalRes, Count: uint64(resultCount)}, nil
}

// HottestEpisodes : Get a list of HottestEpisodes based on the provided params
func (s *Server) HottestEpisodes(ctx context.Context, request *proto.EpisodesRequest) (*proto.EpisodesResponse, error) {
	// TODO : Order Hottest Episodes by the amount of times a single episode has been seen
	return s.Episodes(
		ctx,
		&proto.EpisodesRequest{
			OrderBy: &proto.OrderBy{
				Field:      "created_at",
				Descending: false,
			},
			Query: request.Query,
		},
	)
}

// NewestEpisodes : Get a list of NewestEpisodes based on the provided params
func (s *Server) NewestEpisodes(ctx context.Context, request *proto.EpisodesRequest) (*proto.EpisodesResponse, error) {
	return s.Episodes(
		ctx,
		&proto.EpisodesRequest{
			OrderBy: &proto.OrderBy{
				Field:      "created_at",
				Descending: true,
			},
			Query: request.Query,
		},
	)
}