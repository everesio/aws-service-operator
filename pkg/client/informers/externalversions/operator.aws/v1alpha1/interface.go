/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/christopherhein/aws-operator/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CloudFormationTemplates returns a CloudFormationTemplateInformer.
	CloudFormationTemplates() CloudFormationTemplateInformer
	// DynamoDBs returns a DynamoDBInformer.
	DynamoDBs() DynamoDBInformer
	// ECRRepositories returns a ECRRepositoryInformer.
	ECRRepositories() ECRRepositoryInformer
	// S3Buckets returns a S3BucketInformer.
	S3Buckets() S3BucketInformer
	// SNSTopics returns a SNSTopicInformer.
	SNSTopics() SNSTopicInformer
	// SQSQueues returns a SQSQueueInformer.
	SQSQueues() SQSQueueInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// CloudFormationTemplates returns a CloudFormationTemplateInformer.
func (v *version) CloudFormationTemplates() CloudFormationTemplateInformer {
	return &cloudFormationTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DynamoDBs returns a DynamoDBInformer.
func (v *version) DynamoDBs() DynamoDBInformer {
	return &dynamoDBInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ECRRepositories returns a ECRRepositoryInformer.
func (v *version) ECRRepositories() ECRRepositoryInformer {
	return &eCRRepositoryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// S3Buckets returns a S3BucketInformer.
func (v *version) S3Buckets() S3BucketInformer {
	return &s3BucketInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SNSTopics returns a SNSTopicInformer.
func (v *version) SNSTopics() SNSTopicInformer {
	return &sNSTopicInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SQSQueues returns a SQSQueueInformer.
func (v *version) SQSQueues() SQSQueueInformer {
	return &sQSQueueInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
