// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elbv2_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleELBV2_AddTags() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.AddTagsInput{
		ResourceArns: []*string{ // Required
			aws.String("ResourceArn"), // Required
			// More values...
		},
		Tags: []*elbv2.Tag{ // Required
			{ // Required
				Key:   aws.String("TagKey"), // Required
				Value: aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.AddTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_CreateListener() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.CreateListenerInput{
		DefaultActions: []*elbv2.Action{ // Required
			{ // Required
				TargetGroupArn: aws.String("TargetGroupArn"), // Required
				Type:           aws.String("ActionTypeEnum"), // Required
			},
			// More values...
		},
		LoadBalancerArn: aws.String("LoadBalancerArn"), // Required
		Port:            aws.Int64(1),                  // Required
		Protocol:        aws.String("ProtocolEnum"),    // Required
		Certificates: []*elbv2.Certificate{
			{ // Required
				CertificateArn: aws.String("CertificateArn"),
			},
			// More values...
		},
		SslPolicy: aws.String("SslPolicyName"),
	}
	resp, err := svc.CreateListener(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_CreateLoadBalancer() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.CreateLoadBalancerInput{
		Name: aws.String("LoadBalancerName"), // Required
		Subnets: []*string{ // Required
			aws.String("SubnetId"), // Required
			// More values...
		},
		Scheme: aws.String("LoadBalancerSchemeEnum"),
		SecurityGroups: []*string{
			aws.String("SecurityGroupId"), // Required
			// More values...
		},
		Tags: []*elbv2.Tag{
			{ // Required
				Key:   aws.String("TagKey"), // Required
				Value: aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateLoadBalancer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_CreateRule() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.CreateRuleInput{
		Actions: []*elbv2.Action{ // Required
			{ // Required
				TargetGroupArn: aws.String("TargetGroupArn"), // Required
				Type:           aws.String("ActionTypeEnum"), // Required
			},
			// More values...
		},
		Conditions: []*elbv2.RuleCondition{ // Required
			{ // Required
				Field: aws.String("ConditionFieldName"),
				Values: []*string{
					aws.String("StringValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		ListenerArn: aws.String("ListenerArn"), // Required
		Priority:    aws.Int64(1),              // Required
	}
	resp, err := svc.CreateRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_CreateTargetGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.CreateTargetGroupInput{
		Name:     aws.String("TargetGroupName"), // Required
		Port:     aws.Int64(1),                  // Required
		Protocol: aws.String("ProtocolEnum"),    // Required
		VpcId:    aws.String("VpcId"),           // Required
		HealthCheckIntervalSeconds: aws.Int64(1),
		HealthCheckPath:            aws.String("Path"),
		HealthCheckPort:            aws.String("HealthCheckPort"),
		HealthCheckProtocol:        aws.String("ProtocolEnum"),
		HealthCheckTimeoutSeconds:  aws.Int64(1),
		HealthyThresholdCount:      aws.Int64(1),
		Matcher: &elbv2.Matcher{
			HttpCode: aws.String("HttpCode"), // Required
		},
		UnhealthyThresholdCount: aws.Int64(1),
	}
	resp, err := svc.CreateTargetGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_DeleteListener() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.DeleteListenerInput{
		ListenerArn: aws.String("ListenerArn"), // Required
	}
	resp, err := svc.DeleteListener(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_DeleteLoadBalancer() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.DeleteLoadBalancerInput{
		LoadBalancerArn: aws.String("LoadBalancerArn"), // Required
	}
	resp, err := svc.DeleteLoadBalancer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_DeleteRule() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.DeleteRuleInput{
		RuleArn: aws.String("RuleArn"), // Required
	}
	resp, err := svc.DeleteRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_DeleteTargetGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.DeleteTargetGroupInput{
		TargetGroupArn: aws.String("TargetGroupArn"), // Required
	}
	resp, err := svc.DeleteTargetGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_DeregisterTargets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.DeregisterTargetsInput{
		TargetGroupArn: aws.String("TargetGroupArn"), // Required
		Targets: []*elbv2.TargetDescription{ // Required
			{ // Required
				Id:   aws.String("TargetId"), // Required
				Port: aws.Int64(1),
			},
			// More values...
		},
	}
	resp, err := svc.DeregisterTargets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_DescribeListeners() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.DescribeListenersInput{
		ListenerArns: []*string{
			aws.String("ListenerArn"), // Required
			// More values...
		},
		LoadBalancerArn: aws.String("LoadBalancerArn"),
		Marker:          aws.String("Marker"),
		PageSize:        aws.Int64(1),
	}
	resp, err := svc.DescribeListeners(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_DescribeLoadBalancerAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.DescribeLoadBalancerAttributesInput{
		LoadBalancerArn: aws.String("LoadBalancerArn"), // Required
	}
	resp, err := svc.DescribeLoadBalancerAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_DescribeLoadBalancers() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.DescribeLoadBalancersInput{
		LoadBalancerArns: []*string{
			aws.String("LoadBalancerArn"), // Required
			// More values...
		},
		Marker: aws.String("Marker"),
		Names: []*string{
			aws.String("LoadBalancerName"), // Required
			// More values...
		},
		PageSize: aws.Int64(1),
	}
	resp, err := svc.DescribeLoadBalancers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_DescribeRules() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.DescribeRulesInput{
		ListenerArn: aws.String("ListenerArn"),
		RuleArns: []*string{
			aws.String("RuleArn"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeRules(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_DescribeSSLPolicies() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.DescribeSSLPoliciesInput{
		Marker: aws.String("Marker"),
		Names: []*string{
			aws.String("SslPolicyName"), // Required
			// More values...
		},
		PageSize: aws.Int64(1),
	}
	resp, err := svc.DescribeSSLPolicies(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_DescribeTags() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.DescribeTagsInput{
		ResourceArns: []*string{ // Required
			aws.String("ResourceArn"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_DescribeTargetGroupAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.DescribeTargetGroupAttributesInput{
		TargetGroupArn: aws.String("TargetGroupArn"), // Required
	}
	resp, err := svc.DescribeTargetGroupAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_DescribeTargetGroups() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.DescribeTargetGroupsInput{
		LoadBalancerArn: aws.String("LoadBalancerArn"),
		Marker:          aws.String("Marker"),
		Names: []*string{
			aws.String("TargetGroupName"), // Required
			// More values...
		},
		PageSize: aws.Int64(1),
		TargetGroupArns: []*string{
			aws.String("TargetGroupArn"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTargetGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_DescribeTargetHealth() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.DescribeTargetHealthInput{
		TargetGroupArn: aws.String("TargetGroupArn"), // Required
		Targets: []*elbv2.TargetDescription{
			{ // Required
				Id:   aws.String("TargetId"), // Required
				Port: aws.Int64(1),
			},
			// More values...
		},
	}
	resp, err := svc.DescribeTargetHealth(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_ModifyListener() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.ModifyListenerInput{
		ListenerArn: aws.String("ListenerArn"), // Required
		Certificates: []*elbv2.Certificate{
			{ // Required
				CertificateArn: aws.String("CertificateArn"),
			},
			// More values...
		},
		DefaultActions: []*elbv2.Action{
			{ // Required
				TargetGroupArn: aws.String("TargetGroupArn"), // Required
				Type:           aws.String("ActionTypeEnum"), // Required
			},
			// More values...
		},
		Port:      aws.Int64(1),
		Protocol:  aws.String("ProtocolEnum"),
		SslPolicy: aws.String("SslPolicyName"),
	}
	resp, err := svc.ModifyListener(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_ModifyLoadBalancerAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.ModifyLoadBalancerAttributesInput{
		Attributes: []*elbv2.LoadBalancerAttribute{ // Required
			{ // Required
				Key:   aws.String("LoadBalancerAttributeKey"),
				Value: aws.String("LoadBalancerAttributeValue"),
			},
			// More values...
		},
		LoadBalancerArn: aws.String("LoadBalancerArn"), // Required
	}
	resp, err := svc.ModifyLoadBalancerAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_ModifyRule() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.ModifyRuleInput{
		RuleArn: aws.String("RuleArn"), // Required
		Actions: []*elbv2.Action{
			{ // Required
				TargetGroupArn: aws.String("TargetGroupArn"), // Required
				Type:           aws.String("ActionTypeEnum"), // Required
			},
			// More values...
		},
		Conditions: []*elbv2.RuleCondition{
			{ // Required
				Field: aws.String("ConditionFieldName"),
				Values: []*string{
					aws.String("StringValue"), // Required
					// More values...
				},
			},
			// More values...
		},
	}
	resp, err := svc.ModifyRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_ModifyTargetGroup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.ModifyTargetGroupInput{
		TargetGroupArn:             aws.String("TargetGroupArn"), // Required
		HealthCheckIntervalSeconds: aws.Int64(1),
		HealthCheckPath:            aws.String("Path"),
		HealthCheckPort:            aws.String("HealthCheckPort"),
		HealthCheckProtocol:        aws.String("ProtocolEnum"),
		HealthCheckTimeoutSeconds:  aws.Int64(1),
		HealthyThresholdCount:      aws.Int64(1),
		Matcher: &elbv2.Matcher{
			HttpCode: aws.String("HttpCode"), // Required
		},
		UnhealthyThresholdCount: aws.Int64(1),
	}
	resp, err := svc.ModifyTargetGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_ModifyTargetGroupAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.ModifyTargetGroupAttributesInput{
		Attributes: []*elbv2.TargetGroupAttribute{ // Required
			{ // Required
				Key:   aws.String("TargetGroupAttributeKey"),
				Value: aws.String("TargetGroupAttributeValue"),
			},
			// More values...
		},
		TargetGroupArn: aws.String("TargetGroupArn"), // Required
	}
	resp, err := svc.ModifyTargetGroupAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_RegisterTargets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.RegisterTargetsInput{
		TargetGroupArn: aws.String("TargetGroupArn"), // Required
		Targets: []*elbv2.TargetDescription{ // Required
			{ // Required
				Id:   aws.String("TargetId"), // Required
				Port: aws.Int64(1),
			},
			// More values...
		},
	}
	resp, err := svc.RegisterTargets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_RemoveTags() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.RemoveTagsInput{
		ResourceArns: []*string{ // Required
			aws.String("ResourceArn"), // Required
			// More values...
		},
		TagKeys: []*string{ // Required
			aws.String("TagKey"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_SetRulePriorities() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.SetRulePrioritiesInput{
		RulePriorities: []*elbv2.RulePriorityPair{ // Required
			{ // Required
				Priority: aws.Int64(1),
				RuleArn:  aws.String("RuleArn"),
			},
			// More values...
		},
	}
	resp, err := svc.SetRulePriorities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_SetSecurityGroups() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.SetSecurityGroupsInput{
		LoadBalancerArn: aws.String("LoadBalancerArn"), // Required
		SecurityGroups: []*string{ // Required
			aws.String("SecurityGroupId"), // Required
			// More values...
		},
	}
	resp, err := svc.SetSecurityGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleELBV2_SetSubnets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := elbv2.New(sess)

	params := &elbv2.SetSubnetsInput{
		LoadBalancerArn: aws.String("LoadBalancerArn"), // Required
		Subnets: []*string{ // Required
			aws.String("SubnetId"), // Required
			// More values...
		},
	}
	resp, err := svc.SetSubnets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
