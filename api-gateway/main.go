package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func initTracer(jaegerURL string) (*sdktrace.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerURL)))
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("api-gateway"),
		)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func main() {
	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Initialize tracer
	tp, err := initTracer("http://jaeger:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	// Create Gin router
	r := gin.New()

	// Middlewares
	r.Use(gin.Recovery())
	r.Use(loggingMiddleware(logger))
	r.Use(tracingMiddleware())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	logger.Infof("Starting API Gateway on port %s", port)
	if err := r.Run(":" + port); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}

func loggingMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		logger.WithFields(logrus.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       path,
			"query":      query,
			"ip":         c.ClientIP(),
			"user-agent": c.Request.UserAgent(),
			"latency":    latency,
		}).Info("request")
	}
}

func tracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := otel.GetTextMapPropagator().Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))
		tr := otel.Tracer("api-gateway")
		ctx, span := tr.Start(ctx, c.Request.URL.Path)
		defer span.End()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
