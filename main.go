package main

import (
	"campaign-coupon-system/configuration"
	"campaign-coupon-system/controller"
	"campaign-coupon-system/domain"
	"campaign-coupon-system/repository"
	"campaign-coupon-system/usecase"
	"log"
	"net"
	"net/http"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"

	proto "campaign-coupon-system/controller/gen/campaign/v1" // Import generated proto package
)

func main() {
	loadConfig, err := configuration.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load environment variables", err)
	}

	db := configuration.ConnectionMySQL(&loadConfig)
	cache := configuration.ConnectionRedis(&loadConfig)

	startServer(db, cache)
}

func startServer(db *gorm.DB, cache *redis.Client) {
	campaignRepo := repository.NewCampaignRepository(db, cache)
	couponRepo := repository.NewCouponRepository(db, cache)
	campaignService := domain.NewCampaignService(campaignRepo)
	couponService := domain.NewCouponService(couponRepo)
	campaignUseCase := usecase.NewCampaignUseCase(db, campaignService, couponService)

	grpcServer := grpc.NewServer()
	campaignRpcController := controller.NewCampaignRpcController(campaignUseCase)
	proto.RegisterCampaignServiceServer(grpcServer, campaignRpcController)
	reflection.Register(grpcServer)

	go func() {
		// gRPC 서버의 포트
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen on port 50051: %v", err)
		}

		log.Println("🚀 gRPC Server listening on :50051")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to start gRPC server: %v", err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Println("🚀 HTTP Server listening on :3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
