package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleRequests(w http.ResponseWriter,r *http.Response) {
	fmt.Fprintf(w,"Welcome to load balancer\n")
}
func main(){
	servers:= []string {
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}

	balancer:= NewRoundRobinBalancer(servers)

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		server:= balancer.NextServer()
		proxy:= http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
			http.Redirect(w,r,server,http.StatusTemporaryRedirect)
		})
		proxy.ServeHTTP(w,r);
	})
	fmt.Println("Load balancer started...");
	http.ListenAndServe(":8080",nil)

}
type RoundRobinBalancer struct {
    servers []string
    index   int
}
func NewRoundRobinBalancer(servers []string) *RoundRobinBalancer{
	return &RoundRobinBalancer{servers:servers}
}
func (b *RoundRobinBalancer) NextServer() string {
	server:=b.servers[b.index]
	b.index=(b.index+1)%len(b.servers)
	return server
}
