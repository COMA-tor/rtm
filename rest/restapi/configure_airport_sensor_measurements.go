// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/COMA-tor/rtm/rest/models"
	rst "github.com/RedisTimeSeries/redistimeseries-go"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/COMA-tor/rtm/rest/restapi/operations"
)

//go:generate swagger generate server --target ../../rest --name AirportSensorMeasurements --spec ../swagger.yml --principal interface{}

var units = map[string]string{
	"temperature": "°C",
	"pressure":    "hPa",
	"windspeed":   "m/s",
}

func configureFlags(api *operations.AirportSensorMeasurementsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func measure(point *rst.DataPoint, measureType string) *models.Measure {
	unit := units[measureType]
	return &models.Measure{
		Timestamp: &point.Timestamp,
		Value:     &point.Value,
		Unit:      &unit,
	}
}

func configureAPI(api *operations.AirportSensorMeasurementsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	client := rst.NewClient("172.17.3.174:6379", "rest", nil)

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.GetAirportIATATypeLastHandler = operations.GetAirportIATATypeLastHandlerFunc(func(params operations.GetAirportIATATypeLastParams) middleware.Responder {
		point, err := client.Get(params.Type + ":" + params.IATA)
		if err != nil {
			return middleware.Error(500, err)
		}
		return operations.NewGetAirportIATATypeLastOK().WithPayload(measure(point, params.Type))
	})

	api.GetAirportIATATypeHandler = operations.GetAirportIATATypeHandlerFunc(func(params operations.GetAirportIATATypeParams) middleware.Responder {
		count := *params.Count
		if count > 300 {
			count = 300
		}

		ranges, err := client.RangeWithOptions(params.Type+":"+params.IATA, rst.TimeRangeMinimum, rst.TimeRangeFull,
			rst.RangeOptions{
				AggType:    rst.AvgAggregation,
				TimeBucket: int(*params.Step),
				Count:      count,
			})
		if err != nil {
			return middleware.Error(500, err)
		}

		measures := make([]*models.Measure, len(ranges))
		for i, p := range ranges {
			measures[i] = measure(&p, params.Type)
		}
		return operations.NewGetAirportIATATypeOK().WithPayload(measures)
	})

	api.GetAirportIATAHandler = operations.GetAirportIATAHandlerFunc(func(params operations.GetAirportIATAParams) middleware.Responder {
		measures := make(map[string]*models.Measure)
		for k := range units {
			point, err := client.Get(k + ":" + params.IATA)
			if err != nil {
				return middleware.Error(500, err)
			}
			measures[k] = measure(point, k)
		}
		return operations.NewGetAirportIATAOK().WithPayload(&operations.GetAirportIATAOKBody{
			Pressure:    measures["pressure"],
			Temperature: measures["temperature"],
			Windspeed:   measures["wind_speed"],
		})
	})

	api.PreServerShutdown = func() {}
	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
