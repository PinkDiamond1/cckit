// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token/service/allowance/allowance.proto

package allowance

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AllowanceRequest) Validate() error {
	if this.OwnerAddress == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OwnerAddress", fmt.Errorf(`value '%v' must not be an empty string`, this.OwnerAddress))
	}
	if this.SpenderAddress == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SpenderAddress", fmt.Errorf(`value '%v' must not be an empty string`, this.SpenderAddress))
	}
	return nil
}
func (this *ApproveRequest) Validate() error {
	if this.OwnerAddress == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OwnerAddress", fmt.Errorf(`value '%v' must not be an empty string`, this.OwnerAddress))
	}
	if this.SpenderAddress == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SpenderAddress", fmt.Errorf(`value '%v' must not be an empty string`, this.SpenderAddress))
	}
	if !(this.Amount > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Amount", fmt.Errorf(`value '%v' must be greater than '0'`, this.Amount))
	}
	return nil
}
func (this *TransferFromRequest) Validate() error {
	if this.OwnerAddress == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OwnerAddress", fmt.Errorf(`value '%v' must not be an empty string`, this.OwnerAddress))
	}
	if this.RecipientAddress == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("RecipientAddress", fmt.Errorf(`value '%v' must not be an empty string`, this.RecipientAddress))
	}
	if !(this.Amount > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Amount", fmt.Errorf(`value '%v' must be greater than '0'`, this.Amount))
	}
	return nil
}
func (this *TransferFromResponse) Validate() error {
	return nil
}
func (this *AllowanceId) Validate() error {
	return nil
}
func (this *Allowance) Validate() error {
	return nil
}
func (this *Allowances) Validate() error {
	for _, item := range this.Items {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Items", err)
			}
		}
	}
	return nil
}
func (this *Approved) Validate() error {
	return nil
}
func (this *TransferredFrom) Validate() error {
	return nil
}
