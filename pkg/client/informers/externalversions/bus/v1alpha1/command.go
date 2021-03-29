/*
Copyright 2021 The Volcano Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	busv1alpha1 "volcano.sh/apis/pkg/apis/bus/v1alpha1"
	versioned "volcano.sh/apis/pkg/client/clientset/versioned"
	internalinterfaces "volcano.sh/apis/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "volcano.sh/apis/pkg/client/listers/bus/v1alpha1"
)

// CommandInformer provides access to a shared informer and lister for
// Commands.
type CommandInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CommandLister
}

type commandInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCommandInformer constructs a new informer for Command type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCommandInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCommandInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCommandInformer constructs a new informer for Command type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCommandInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BusV1alpha1().Commands(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BusV1alpha1().Commands(namespace).Watch(options)
			},
		},
		&busv1alpha1.Command{},
		resyncPeriod,
		indexers,
	)
}

func (f *commandInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCommandInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *commandInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&busv1alpha1.Command{}, f.defaultInformer)
}

func (f *commandInformer) Lister() v1alpha1.CommandLister {
	return v1alpha1.NewCommandLister(f.Informer().GetIndexer())
}