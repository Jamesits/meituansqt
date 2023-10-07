package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jamesits/meituansqt/pkg/sqt"
)

// Ensure MeituanSqtProvider satisfies various provider interfaces.
var _ provider.Provider = &MeituanSqtProvider{}

// MeituanSqtProvider defines the provider implementation.
type MeituanSqtProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// MeituanSqtProviderModel describes the provider data model.
type MeituanSqtProviderModel struct {
	EntId     types.Int64  `tfsdk:"ent_id"`
	AccessKey types.String `tfsdk:"access_key"`
	SecretKey types.String `tfsdk:"secret_key"`
}

func (p *MeituanSqtProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "meituansqt"
	resp.Version = p.version
}

func (p *MeituanSqtProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ent_id": schema.Int64Attribute{
				MarkdownDescription: "Enterprise ID",
				Required:            true,
				Sensitive:           true,
			},
			"access_key": schema.StringAttribute{
				MarkdownDescription: "Access key",
				Required:            true,
				Sensitive:           true,
			},
			"secret_key": schema.StringAttribute{
				MarkdownDescription: "Secret key",
				Required:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *MeituanSqtProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data MeituanSqtProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.

	sqtClient := sqt.NewProduction(data.EntId.ValueInt64(), data.AccessKey.ValueString(), data.SecretKey.ValueString())
	resp.DataSourceData = sqtClient
	resp.ResourceData = sqtClient
}

func (p *MeituanSqtProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewStaffResource,
	}
}

func (p *MeituanSqtProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &MeituanSqtProvider{
			version: version,
		}
	}
}
