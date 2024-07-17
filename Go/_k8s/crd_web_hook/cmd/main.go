package main

import (
	"context"
	"flag"
	"net/http"

	"crd_webhook/pkg/webhook"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/webhook/server"
)

var (
	addr = flag.String("addr", ":443", "Address to bind the webhook server")
)

func main() {
	flag.Parse()

	cfg, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err.Error())
	}

	webhookServer := server.NewServer(*addr)
	webhook.SetupWebhook(webhookServer, cfg)

	mux := http.NewServeMux()
	mux.Handle("/", webhookServer)

	server := &http.Server{
		Addr:    *addr,
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServeTLS("", ""); err != nil {
			panic(err.Error())
		}
	}()

	// Wait for shutdown signal to gracefully shutdown the server
	ctx := context.Background()
	<-ctx.Done()
	server.Shutdown(ctx)
}
