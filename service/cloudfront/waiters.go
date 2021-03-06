// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilDistributionDeployed uses the CloudFront API operation
// GetDistribution to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *CloudFront) WaitUntilDistributionDeployed(input *GetDistributionInput) error {
	return c.WaitUntilDistributionDeployedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilDistributionDeployedWithContext is an extended version of WaitUntilDistributionDeployed.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFront) WaitUntilDistributionDeployedWithContext(ctx aws.Context, input *GetDistributionInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDistributionDeployed",
		MaxAttempts: 25,
		Delay:       aws.ConstantWaiterDelay(60 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "Distribution.Status",
				Expected: "Deployed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *GetDistributionInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.GetDistributionRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilInvalidationCompleted uses the CloudFront API operation
// GetInvalidation to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *CloudFront) WaitUntilInvalidationCompleted(input *GetInvalidationInput) error {
	return c.WaitUntilInvalidationCompletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInvalidationCompletedWithContext is an extended version of WaitUntilInvalidationCompleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFront) WaitUntilInvalidationCompletedWithContext(ctx aws.Context, input *GetInvalidationInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilInvalidationCompleted",
		MaxAttempts: 30,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "Invalidation.Status",
				Expected: "Completed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *GetInvalidationInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.GetInvalidationRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilStreamingDistributionDeployed uses the CloudFront API operation
// GetStreamingDistribution to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *CloudFront) WaitUntilStreamingDistributionDeployed(input *GetStreamingDistributionInput) error {
	return c.WaitUntilStreamingDistributionDeployedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStreamingDistributionDeployedWithContext is an extended version of WaitUntilStreamingDistributionDeployed.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudFront) WaitUntilStreamingDistributionDeployedWithContext(ctx aws.Context, input *GetStreamingDistributionInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilStreamingDistributionDeployed",
		MaxAttempts: 25,
		Delay:       aws.ConstantWaiterDelay(60 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "StreamingDistribution.Status",
				Expected: "Deployed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *GetStreamingDistributionInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.GetStreamingDistributionRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
