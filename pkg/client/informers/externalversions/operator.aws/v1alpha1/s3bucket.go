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
	time "time"

	operator_aws_v1alpha1 "github.com/christopherhein/aws-operator/pkg/apis/operator.aws/v1alpha1"
	versioned "github.com/christopherhein/aws-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/christopherhein/aws-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/christopherhein/aws-operator/pkg/client/listers/operator.aws/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// S3BucketInformer provides access to a shared informer and lister for
// S3Buckets.
type S3BucketInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.S3BucketLister
}

type s3BucketInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewS3BucketInformer constructs a new informer for S3Bucket type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewS3BucketInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredS3BucketInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredS3BucketInformer constructs a new informer for S3Bucket type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredS3BucketInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1alpha1().S3Buckets(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1alpha1().S3Buckets(namespace).Watch(options)
			},
		},
		&operator_aws_v1alpha1.S3Bucket{},
		resyncPeriod,
		indexers,
	)
}

func (f *s3BucketInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredS3BucketInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *s3BucketInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operator_aws_v1alpha1.S3Bucket{}, f.defaultInformer)
}

func (f *s3BucketInformer) Lister() v1alpha1.S3BucketLister {
	return v1alpha1.NewS3BucketLister(f.Informer().GetIndexer())
}
