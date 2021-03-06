package cloudformation

import (
	"encoding/json"
)

// AWSBatchJobDefinition_Ulimit AWS CloudFormation Resource (AWS::Batch::JobDefinition.Ulimit)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ulimit.html
type AWSBatchJobDefinition_Ulimit struct {

	// HardLimit AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ulimit.html#cfn-batch-jobdefinition-ulimit-hardlimit
	HardLimit *Value `json:"HardLimit,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ulimit.html#cfn-batch-jobdefinition-ulimit-name
	Name *Value `json:"Name,omitempty"`

	// SoftLimit AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ulimit.html#cfn-batch-jobdefinition-ulimit-softlimit
	SoftLimit *Value `json:"SoftLimit,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_Ulimit) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.Ulimit"
}

func (r *AWSBatchJobDefinition_Ulimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
