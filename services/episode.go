package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/Krystian19/cactus-core/models"
	"github.com/Krystian19/cactus-core/proto"
	"github.com/jinzhu/gorm"
)

// Episode : Get a single Episode based on the provided params
func (s *Server) Episode(ctx context.Context, request *proto.EpisodeRequest) (*proto.EpisodeResponse, error) {
	var result models.Episode
	query := s.db

	if request.Id != 0 {
		query = query.Where("id = ?", request.Id)
	}

	if request.ReleaseId != 0 {
		query = query.Where("release_id = ?", request.ReleaseId)
	}

	if request.OrderBy != nil {
		if len(strings.TrimSpace(request.OrderBy.Field)) != 0 {

			sorting := ""
			if request.OrderBy.Descending {
				sorting = models.DESC
			} else {
				sorting = models.ASC
			}

			// Example : SORT BY - "fieldname ASC"
			query = query.Order(fmt.Sprintf("%s %s", request.OrderBy.Field, sorting))
		}
	}

	if request.LessThan != nil {
		if len(strings.TrimSpace(request.LessThan.Field)) != 0 {
			// Example : WHERE fieldname < value
			query = query.Where(
				fmt.Sprintf("%s < %d", request.LessThan.Field, request.LessThan.Value),
			)
		}
	}

	if request.GreaterThan != nil {
		if len(strings.TrimSpace(request.GreaterThan.Field)) != 0 {
			// Example : WHERE fieldname > value
			query = query.Where(
				fmt.Sprintf("%s > %d", request.GreaterThan.Field, request.GreaterThan.Value),
			)
		}
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
	var result []models.Episode
	var resultCount uint
	query := s.db

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
				sorting = models.DESC
			} else {
				sorting = models.ASC
			}

			// Example : SORT BY - "fieldname ASC"
			query = query.Order(
				fmt.Sprintf("%s %s", request.OrderBy.Field, sorting),
			)
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
func (s *Server) HottestEpisodes(ctx context.Context, request *proto.PaginationRequest) (*proto.EpisodesResponse, error) {
	finalRes := []*proto.Episode{}
	var resultCount uint

	// SELECT E.*, count(ES.id) as seen_count FROM public."Episodes" E
	// 		LEFT JOIN public."EpisodesSeen" ES on ES.episode_id = E.id
	// 		group by E.id
	// 		order by seen_count DESC;

	queryString := fmt.Sprintf(`
		SELECT E.*, count(ES.id) as seen_count FROM public."%s" E
			LEFT JOIN public."%s" ES on ES.episode_id = E.id
			group by E.id
			order by seen_count DESC`,
		models.Episode.TableName(models.Episode{}),
		models.EpisodeSeen.TableName(models.EpisodeSeen{}),
	)

	query := s.db.Raw(queryString)

	if request != nil {
		if request.Limit != 0 {
			query = query.Limit(request.Limit)
		}

		if request.Offset != 0 {
			query = query.Offset(request.Offset)
		}
	}

	rows, err := query.Rows() // (*sql.Rows, error)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		var episode = models.Episode{}
		s.db.ScanRows(rows, &episode)
		finalRes = append(finalRes, episode.Episode)
	}

	// Remove pagination constraints before counting
	query = query.Limit(nil)
	query = query.Offset(nil)

	row := s.db.Raw(
		fmt.Sprintf("SELECT COUNT(ES.id) FROM (%s) as ES", queryString),
	).Row() // (*sql.Row)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	row.Scan(&resultCount)

	return &proto.EpisodesResponse{Episodes: finalRes, Count: uint64(resultCount)}, nil
}

// LatestEpisode : Get the Latest Episode of the specified Release
func (s *Server) LatestEpisode(ctx context.Context, request *proto.LatestEpisodeRequest) (*proto.EpisodeResponse, error) {
	return s.Episode(
		ctx,
		&proto.EpisodeRequest{
			ReleaseId: request.ReleaseId,
			OrderBy: &proto.OrderBy{
				Field:      "episode_order",
				Descending: true,
			},
		},
	)
}

// EpisodeCount : Get the Episode count of the specified Release
func (s *Server) EpisodeCount(ctx context.Context, request *proto.EpisodeCountRequest) (*proto.EpisodeCountResponse, error) {
	db := s.db

	var episodeCount uint

	if err := db.Where("release_id = ?", request.ReleaseId).Model(&models.Episode{}).Count(&episodeCount).Error; err != nil {
		episodeCount = 0
	}

	return &proto.EpisodeCountResponse{Count: int64(episodeCount)}, nil
}

// EpisodeSubtitles : Get a list of EpisodeSubtitles associated with the specified EpisodeId
func (s *Server) EpisodeSubtitles(ctx context.Context, request *proto.EpisodeSubtitlesRequest) (*proto.EpisodeSubtitlesListResponse, error) {
	var result []models.EpisodeSubtitle
	query := s.db

	query = query.Where("episode_id = ?", request.EpisodeId)

	if err := query.Find(&result).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	finalRes := []*proto.EpisodeSubtitle{}

	for i := range result {
		finalRes = append(finalRes, result[i].EpisodeSubtitle)
	}

	return &proto.EpisodeSubtitlesListResponse{EpisodeSubtitles: []*proto.EpisodeSubtitle{}}, nil
}

// EpisodeSeen : Marks x episode as "seen", with a timestamp
func (s *Server) EpisodeSeen(ctx context.Context, request *proto.EpisodeSeenRequest) (*proto.Empty, error) {
	// TODO : When users are implemented the EpisodeSeen record should include an user_id
	if err := s.db.Create(&models.EpisodeSeenFields{EpisodeID: request.EpisodeId}).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &proto.Empty{}, nil
}
