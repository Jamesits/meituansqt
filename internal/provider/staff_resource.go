package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/jamesits/meituansqt/pkg/sqt"
	"strconv"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &StaffResource{}
var _ resource.ResourceWithImportState = &StaffResource{}

func NewStaffResource() resource.Resource {
	return &StaffResource{}
}

// StaffResource defines the resource implementation.
type StaffResource struct {
	sqtClient *sqt.SQT
}

// StaffResourceModel describes the resource data model.
type StaffResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Phone       types.String `tfsdk:"phone"`
	EntStaffNum types.String `tfsdk:"ent_staff_num"`
	Email       types.String `tfsdk:"email"`

	Id types.Int64 `tfsdk:"id"`
}

func (r *StaffResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_staff"
}

func (r *StaffResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Staff",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "姓名",
				Required:            true,
			},
			"phone": schema.StringAttribute{
				MarkdownDescription: "手机号",
				Optional:            true,
				// TODO: fix import
				//PlanModifiers: []planmodifier.String{
				//	stringplanmodifier.RequiresReplace(),
				//},
			},
			"ent_staff_num": schema.StringAttribute{
				MarkdownDescription: "工号",
				Optional:            true,
				//PlanModifiers: []planmodifier.String{
				//	stringplanmodifier.RequiresReplace(),
				//},
			},
			"email": schema.StringAttribute{
				MarkdownDescription: "邮箱",
				Optional:            true,
				//PlanModifiers: []planmodifier.String{
				//	stringplanmodifier.RequiresReplace(),
				//},
			},
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "用户 ID",
			},
		},
	}
}

func (r *StaffResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	sqtClient, ok := req.ProviderData.(*sqt.SQT)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sqt.SQT, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.sqtClient = sqtClient
}

func (r *StaffResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data StaffResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiResp, _, err := r.sqtClient.StaffBatchAdd(ctx, &sqt.StaffBatchAddRequest{
		StaffInfos: []sqt.StaffInfo{
			{
				Name:        data.Name.ValueString(),
				Phone:       data.Phone.ValueString(),
				EntStaffNum: data.EntStaffNum.ValueString(),
				Email:       data.Email.ValueString(),
			},
		},
	})
	if err != nil {
		resp.Diagnostics.AddError("HTTP Request Error", fmt.Sprintf("Unable to create staff, got error: %s", err))
		return
	}
	if apiResp == nil || apiResp.StaffAddResultItems == nil || len(apiResp.StaffAddResultItems) != 1 {
		resp.Diagnostics.AddError("API Error", fmt.Sprintf("Unable to create staff, API returns: %v", apiResp))
		return
	}

	data.Id = types.Int64Value(apiResp.StaffAddResultItems[0].StaffId)

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created staff")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StaffResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data StaffResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var queryType sqt.StaffIdType
	var queryValue string
	if !data.Id.IsNull() {
		queryType = sqt.StaffIdTypeStaffId
		queryValue = strconv.FormatInt(data.Id.ValueInt64(), 10)
	} else if !data.Email.IsNull() {
		queryType = sqt.StaffIdTypeEmail
		queryValue = data.Email.ValueString()
	} else if !data.Phone.IsNull() {
		queryType = sqt.StaffIdTypePhone
		queryValue = data.Phone.ValueString()
	} else {
		resp.Diagnostics.AddError("No identifier found for the object", "")
		return
	}

	apiResp, _, err := r.sqtClient.StaffBatchQuery(ctx, &sqt.StaffBatchQueryRequest{
		StaffIdType:      queryType,
		StaffIdentifiers: []string{queryValue},
	})
	if err != nil {
		resp.Diagnostics.AddError("HTTP Request Error", fmt.Sprintf("Unable to create staff, got error: %s", err))
		return
	}
	if apiResp == nil || apiResp.StaffQueryResultItems == nil || len(apiResp.StaffQueryResultItems) != 1 {
		resp.Diagnostics.AddError("API Error", fmt.Sprintf("Unable to query staff, API returns: %v", apiResp))
		return
	}

	data.Id = types.Int64Value(apiResp.StaffQueryResultItems[0].StaffId)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StaffResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data StaffResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiResp, _, err := r.sqtClient.StaffBatchUpdate(ctx, &sqt.StaffBatchUpdateRequest{
		StaffInfos: []sqt.StaffInfo{
			{
				StaffId:     data.Id.ValueInt64(),
				Name:        data.Name.ValueString(),
				Phone:       data.Phone.ValueString(),
				EntStaffNum: data.EntStaffNum.ValueString(),
				Email:       data.Email.ValueString(),
			},
		},
	})
	if err != nil {
		resp.Diagnostics.AddError("HTTP Request Error", fmt.Sprintf("Unable to update staff, got error: %s", err))
		return
	}
	if apiResp == nil || apiResp.StaffBatchUpdateResultItems == nil || len(apiResp.StaffBatchUpdateResultItems) != 1 {
		resp.Diagnostics.AddError("API Error", fmt.Sprintf("Unable to update staff, API returns: %v", apiResp))
		return
	}
	if apiResp.StaffBatchUpdateResultItems[0].Result != sqt.StaffUpdateResultSucceed {
		resp.Diagnostics.AddError("API Error", fmt.Sprintf("Unable to update staff, reason: %d", apiResp.StaffBatchUpdateResultItems[0].Result))
		return
	}

	data.Id = types.Int64Value(apiResp.StaffBatchUpdateResultItems[0].StaffId)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StaffResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data StaffResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiResp, _, err := r.sqtClient.StaffBatchDelete(ctx, &sqt.StaffBatchDeleteRequest{
		StaffIdType:      sqt.StaffIdTypeStaffId,
		StaffIdentifiers: []string{strconv.FormatInt(data.Id.ValueInt64(), 10)},
	})
	if err != nil {
		resp.Diagnostics.AddError("HTTP Request Error", fmt.Sprintf("Unable to create staff, got error: %s", err))
		return
	}
	if apiResp == nil || apiResp.StaffDeleteResultItems == nil || len(apiResp.StaffDeleteResultItems) != 1 {
		resp.Diagnostics.AddError("API Error", fmt.Sprintf("Unable to create staff, API returns: %v", apiResp))
		return
	}
}

func (r *StaffResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data StaffResourceModel

	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Unable to parse ID", fmt.Sprintf("error: %v", err))
		return
	}
	data.Id = types.Int64Value(id)

	// resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
