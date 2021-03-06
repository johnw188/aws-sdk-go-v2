// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kmsiface provides an interface to enable mocking the AWS Key Management Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kmsiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

// KMSAPI provides an interface to enable mocking the
// kms.KMS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Key Management Service.
//    func myFunc(svc kmsiface.KMSAPI) bool {
//        // Make svc.CancelKeyDeletion request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := kms.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKMSClient struct {
//        kmsiface.KMSAPI
//    }
//    func (m *mockKMSClient) CancelKeyDeletion(input *kms.CancelKeyDeletionInput) (*kms.CancelKeyDeletionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKMSClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type KMSAPI interface {
	CancelKeyDeletionRequest(*kms.CancelKeyDeletionInput) kms.CancelKeyDeletionRequest

	CreateAliasRequest(*kms.CreateAliasInput) kms.CreateAliasRequest

	CreateGrantRequest(*kms.CreateGrantInput) kms.CreateGrantRequest

	CreateKeyRequest(*kms.CreateKeyInput) kms.CreateKeyRequest

	DecryptRequest(*kms.DecryptInput) kms.DecryptRequest

	DeleteAliasRequest(*kms.DeleteAliasInput) kms.DeleteAliasRequest

	DeleteImportedKeyMaterialRequest(*kms.DeleteImportedKeyMaterialInput) kms.DeleteImportedKeyMaterialRequest

	DescribeKeyRequest(*kms.DescribeKeyInput) kms.DescribeKeyRequest

	DisableKeyRequest(*kms.DisableKeyInput) kms.DisableKeyRequest

	DisableKeyRotationRequest(*kms.DisableKeyRotationInput) kms.DisableKeyRotationRequest

	EnableKeyRequest(*kms.EnableKeyInput) kms.EnableKeyRequest

	EnableKeyRotationRequest(*kms.EnableKeyRotationInput) kms.EnableKeyRotationRequest

	EncryptRequest(*kms.EncryptInput) kms.EncryptRequest

	GenerateDataKeyRequest(*kms.GenerateDataKeyInput) kms.GenerateDataKeyRequest

	GenerateDataKeyWithoutPlaintextRequest(*kms.GenerateDataKeyWithoutPlaintextInput) kms.GenerateDataKeyWithoutPlaintextRequest

	GenerateRandomRequest(*kms.GenerateRandomInput) kms.GenerateRandomRequest

	GetKeyPolicyRequest(*kms.GetKeyPolicyInput) kms.GetKeyPolicyRequest

	GetKeyRotationStatusRequest(*kms.GetKeyRotationStatusInput) kms.GetKeyRotationStatusRequest

	GetParametersForImportRequest(*kms.GetParametersForImportInput) kms.GetParametersForImportRequest

	ImportKeyMaterialRequest(*kms.ImportKeyMaterialInput) kms.ImportKeyMaterialRequest

	ListAliasesRequest(*kms.ListAliasesInput) kms.ListAliasesRequest

	ListAliasesPages(*kms.ListAliasesInput, func(*kms.ListAliasesOutput, bool) bool) error
	ListAliasesPagesWithContext(aws.Context, *kms.ListAliasesInput, func(*kms.ListAliasesOutput, bool) bool, ...aws.Option) error

	ListGrantsRequest(*kms.ListGrantsInput) kms.ListGrantsRequest

	ListGrantsPages(*kms.ListGrantsInput, func(*kms.ListRetirableGrantsOutput, bool) bool) error
	ListGrantsPagesWithContext(aws.Context, *kms.ListGrantsInput, func(*kms.ListRetirableGrantsOutput, bool) bool, ...aws.Option) error

	ListKeyPoliciesRequest(*kms.ListKeyPoliciesInput) kms.ListKeyPoliciesRequest

	ListKeyPoliciesPages(*kms.ListKeyPoliciesInput, func(*kms.ListKeyPoliciesOutput, bool) bool) error
	ListKeyPoliciesPagesWithContext(aws.Context, *kms.ListKeyPoliciesInput, func(*kms.ListKeyPoliciesOutput, bool) bool, ...aws.Option) error

	ListKeysRequest(*kms.ListKeysInput) kms.ListKeysRequest

	ListKeysPages(*kms.ListKeysInput, func(*kms.ListKeysOutput, bool) bool) error
	ListKeysPagesWithContext(aws.Context, *kms.ListKeysInput, func(*kms.ListKeysOutput, bool) bool, ...aws.Option) error

	ListResourceTagsRequest(*kms.ListResourceTagsInput) kms.ListResourceTagsRequest

	ListRetirableGrantsRequest(*kms.ListRetirableGrantsInput) kms.ListRetirableGrantsRequest

	PutKeyPolicyRequest(*kms.PutKeyPolicyInput) kms.PutKeyPolicyRequest

	ReEncryptRequest(*kms.ReEncryptInput) kms.ReEncryptRequest

	RetireGrantRequest(*kms.RetireGrantInput) kms.RetireGrantRequest

	RevokeGrantRequest(*kms.RevokeGrantInput) kms.RevokeGrantRequest

	ScheduleKeyDeletionRequest(*kms.ScheduleKeyDeletionInput) kms.ScheduleKeyDeletionRequest

	TagResourceRequest(*kms.TagResourceInput) kms.TagResourceRequest

	UntagResourceRequest(*kms.UntagResourceInput) kms.UntagResourceRequest

	UpdateAliasRequest(*kms.UpdateAliasInput) kms.UpdateAliasRequest

	UpdateKeyDescriptionRequest(*kms.UpdateKeyDescriptionInput) kms.UpdateKeyDescriptionRequest
}

var _ KMSAPI = (*kms.KMS)(nil)
