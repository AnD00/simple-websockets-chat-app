package apigw

import (
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

var config aws.Config

func init() {
	var err error
	config, err = external.LoadDefaultAWSConfig()
	if err != nil {
		panic(err)
	}
}

func NewAPIGatewayManagementClient(domain, stage string) *apigatewaymanagementapi.Client {
	cp := config.Copy()
	cp.EndpointResolver = aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if service != "execute-api" {
			return config.EndpointResolver.ResolveEndpoint(service, region)
		}

		var endpoint url.URL
		endpoint.Path = stage
		endpoint.Host = domain
		endpoint.Scheme = "https"
		return aws.Endpoint{
			SigningRegion: region,
			URL:           endpoint.String(),
		}, nil
	})

	return apigatewaymanagementapi.New(cp)
}
