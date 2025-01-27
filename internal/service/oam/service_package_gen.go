// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package oam

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	oam_sdkv2 "github.com/aws/aws-sdk-go-v2/service/oam"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceLink,
			TypeName: "aws_oam_link",
		},
		{
			Factory:  DataSourceLinks,
			TypeName: "aws_oam_links",
		},
		{
			Factory:  DataSourceSink,
			TypeName: "aws_oam_sink",
		},
		{
			Factory:  DataSourceSinks,
			TypeName: "aws_oam_sinks",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceLink,
			TypeName: "aws_oam_link",
			Name:     "Link",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceSink,
			TypeName: "aws_oam_sink",
			Name:     "Sink",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceSinkPolicy,
			TypeName: "aws_oam_sink_policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ObservabilityAccessManager
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*oam_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return oam_sdkv2.NewFromConfig(cfg, func(o *oam_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.EndpointResolver = oam_sdkv2.EndpointResolverFromURL(endpoint)
		}
	}), nil
}

var ServicePackage = &servicePackage{}
