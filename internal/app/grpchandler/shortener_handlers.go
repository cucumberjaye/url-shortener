package grpchandler

import (
	"context"

	"github.com/cucumberjaye/url-shortener/configs"
	"github.com/cucumberjaye/url-shortener/internal/app/handler"
	"github.com/cucumberjaye/url-shortener/internal/app/pb"
	"github.com/cucumberjaye/url-shortener/models"
	"github.com/cucumberjaye/url-shortener/pkg/token"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// структура для grpc сервиса
type ShortenerServer struct {
	pb.UnimplementedShotenerServiceServer

	Service handler.URLService
	Ch      chan models.DeleteData
}

// получения токена аутентификации
func (s *ShortenerServer) Authentication(ctx context.Context, in *pb.Empty) (*pb.AuthToken, error) {
	var authToken string
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		values := md.Get("authentication")
		if len(values) > 0 {
			authToken = values[0]
			_, err := token.CheckToken(authToken)
			if err == nil {
				return &pb.AuthToken{Token: authToken}, nil
			}
		}
	}

	id := uuid.New().String()
	authToken, err := token.GenerateNewToken(id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.AuthToken{Token: authToken}, nil
}

// получаем оригинальный url по сокращенному
func (s *ShortenerServer) GetFullURL(ctx context.Context, in *pb.Short) (*pb.Original, error) {
	short := baseURL() + in.GetShortUrl()
	fullURL, err := s.Service.GetFullURL(short)
	if err != nil && err.Error() == "URL was deleted" {
		return nil, status.Errorf(codes.DataLoss, "URL %s was deleted", in.GetShortUrl())

	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Original{OriginalUrl: fullURL}, nil
}

// принимает ориганальный url и возвращает сокращенный
func (s *ShortenerServer) Shortener(ctx context.Context, in *pb.Original) (*pb.Short, error) {
	var id string
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		values := md.Get("user_id")
		if len(values) > 0 {
			id = values[0]
		}
	}
	if id == "" {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	shortURL, err := s.Service.ShortingURL(in.GetOriginalUrl(), baseURL(), id)
	if err != nil && err.Error() != "url already exists" {
		return nil, status.Error(codes.Internal, err.Error())

	}
	return &pb.Short{ShortUrl: shortURL}, nil
}

// проверяет работоспособность хранилища (postgreSQL или файла).
func (s *ShortenerServer) Ping(ctx context.Context, in *pb.Empty) (*pb.CommonResponse, error) {
	err := s.Service.CheckDBConn()
	if err != nil {
		return &pb.CommonResponse{Status: pb.CommonResponse_FAIL}, status.Error(codes.Internal, err.Error())
	}

	return &pb.CommonResponse{Status: pb.CommonResponse_OK}, nil
}

// сокращение пачки url
func (s *ShortenerServer) BatchShortnener(ctx context.Context, in *pb.BatchRequest) (*pb.BatchResponse, error) {
	var id string
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		values := md.Get("user_id")
		if len(values) > 0 {
			id = values[0]
		}
	}
	if id == "" {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	input := make([]models.BatchInputJSON, len(in.OriginalUrls))
	for i, el := range in.OriginalUrls {
		input[i].CorrelationID = el.CorrelationId
		input[i].OriginalURL = el.OriginalUrl
	}

	tmp, err := s.Service.BatchSetURL(input, baseURL(), id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	length := len(tmp)
	out := make([]*pb.BatchShort, length)

	for i := 0; i < length; i++ {
		out[i] = &pb.BatchShort{
			CorrelationId: tmp[i].CorrelationID,
			ShortUrl:      tmp[i].OriginalURL,
		}
	}

	return &pb.BatchResponse{ShortUrls: out}, nil
}

// получение всех сокращенных ссылок пользователя
func (s *ShortenerServer) GetUserURL(ctx context.Context, in *pb.Empty) (*pb.URLsResponse, error) {
	var id string
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		values := md.Get("user_id")
		if len(values) > 0 {
			id = values[0]
		}
	}
	if id == "" {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	out, err := s.Service.GetAllUserURL(id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if len(out) == 0 {
		return nil, status.Error(codes.DataLoss, "empty")
	}

	res := make([]*pb.URLs, len(out))
	for i := 0; i < len(out); i++ {
		res[i] = &pb.URLs{
			OriginalUrl: out[i].OriginalURL,
			ShortUrl:    out[i].ShortURL,
		}
	}

	return &pb.URLsResponse{Urls: res}, nil
}

// удаление массива сокращенных ссылок
func (s *ShortenerServer) DeleteUserURL(ctx context.Context, in *pb.BatchDelete) (*pb.CommonResponse, error) {
	var id string
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		values := md.Get("user_id")
		if len(values) > 0 {
			id = values[0]
		}
	}
	if id == "" {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	for _, elem := range in.GetShortUrls() {
		short := baseURL() + elem.ShortUrl
		s.Ch <- models.DeleteData{
			ID:       id,
			ShortURL: short,
		}
	}

	return &pb.CommonResponse{Status: pb.CommonResponse_OK}, nil
}

// статистика по сервису
func (s *ShortenerServer) Stats(ctx context.Context, in *pb.Empty) (*pb.StatsInfo, error) {
	if len(configs.TrustedSubnet) == 0 {
		return nil, status.Error(codes.Internal, "forbiden in configuration")
	}
	var xRealIP string
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		values := md.Get("X-Real-IP")
		if len(values) > 0 {
			xRealIP = values[0]
		}
	}

	if configs.TrustedSubnet != xRealIP {
		return nil, status.Error(codes.Internal, "forbiden for ip")
	}

	stats, err := s.Service.GetStats()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.StatsInfo{
		UrlsCount:  int32(stats.URLs),
		UsersCount: int32(stats.Users),
	}, nil
}

// устанавливаем перед сокращенным url
func baseURL() string {
	if configs.BaseURL != "" {
		return configs.BaseURL + "/"
	}

	return ""
}
