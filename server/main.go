package main

import (
	"context"
	"net/http"

	grpcweb "github.com/seankhliao/go-grpcweb"
	"google.golang.org/grpc"

	pb "github.com/seankhliao/readss/readss"
)

func main() {
	svr := grpc.NewServer()
	pb.RegisterListerServer(svr, &Server{})

	// wrap grpc handler in grpc-web handler
	handler := grpcweb.New(svr)
	// http.ListenAndServe(":8090", handler)

	// OPTIONAL:
	// handle cors if necessary:
	//  Headers:
	//    Access-Control-Allow-Origin
	//    Access-Control-Allow-Headers
	//  Request:
	//    method: OPTIONS
	//    response: 200
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("access-control-allow-origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		handler.ServeHTTP(w, r)
	})
	http.ListenAndServe(":8090", h)

}

type Server struct {
}

func (s *Server) List(context.Context, *pb.ListRequest) (*pb.ListReply, error) {
	return &pb.ListReply{
		Articles: []*pb.Article{
			&pb.Article{
				Title:   "this is title 1",
				Url:     "https://google.com",
				Source:  "Google",
				Time:    "2000-01-01",
				Reltime: "19y ago",
			},
			&pb.Article{
				Title:   "this is title 2",
				Url:     "https://ibm.com",
				Source:  "IBM",
				Time:    "1999-01-01",
				Reltime: "20y ago",
			},
		},
	}, nil
}