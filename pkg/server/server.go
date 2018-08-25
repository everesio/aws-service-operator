package server

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"time"

	awsclient "github.com/christopherhein/aws-operator/pkg/client/clientset/versioned/typed/operator.aws/v1alpha1"
	"github.com/christopherhein/aws-operator/pkg/config"
	"github.com/christopherhein/aws-operator/pkg/operator/cloudformationtemplate"
	"github.com/christopherhein/aws-operator/pkg/operator/dynamodb"
	"github.com/christopherhein/aws-operator/pkg/operator/ecrrepository"
	"github.com/christopherhein/aws-operator/pkg/operator/s3bucket"
	"github.com/christopherhein/aws-operator/pkg/operator/snstopic"
	"github.com/christopherhein/aws-operator/pkg/operator/sqsqueue"
	opkit "github.com/christopherhein/operator-kit"
	"k8s.io/api/core/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// New creates a new server from a config
func New(config *config.Config) *Server {
	return &Server{
		Config: config,
	}
}

// Run starts the server to listen to Kubernetes
func (c *Server) Run(stopChan chan struct{}) {
	config := c.Config
	logger := config.Logger
	logger.Info("Getting kubernetes context")
	context, restConfig, awsClientset, err := createContext(config.Kubeconfig)
	if err != nil {
		logger.Fatalf("failed to create context. %+v\n", err)
	}
	config.Context = context
	config.AWSClientset = awsClientset
	config.RESTConfig = restConfig

	// Create and wait for CRD resources
	logger.Info("Registering resources")
	resources := []opkit.CustomResource{
		cloudformationtemplate.Resource,
		s3bucket.Resource,
		dynamodb.Resource,
		sqsqueue.Resource,
		snstopic.Resource,
		ecrrepository.Resource,
	}
	err = opkit.CreateCustomResources(*context, resources)
	if err != nil {
		logger.Fatalf("failed to create custom resource. %+v\n", err)
	}

	// TODO: Figure out how to get the node tag so I can store the
	// `KubernetesCluster` attribute so that all resources can be cleaned up
	// using that tag. Also we can create an inventory of all aws resources that
	// are modifiable at any time and deletable using the kubectl cli

	region := ""
	ec2Session, err := session.NewSession()
	metadata := ec2metadata.New(ec2Session)
	if config.Region == "" {
		region, _ = metadata.Region()
	} else {
		region = config.Region
	}

	logger.Infof("Region: %+v", region)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		logger.Infof("Error creating AWS session '%s'\n", err)
	}
	config.AWSSession = sess

	// start watching the aws operator resources
	logger.Info("Watching the resources")
	cftcontroller := cloudformationtemplate.NewController(config, context, awsClientset)
	cftcontroller.StartWatch(v1.NamespaceAll, stopChan)

	s3controller := s3bucket.NewController(config, context, awsClientset)
	s3controller.StartWatch(v1.NamespaceAll, stopChan)

	ddbcontroller := dynamodb.NewController(config, context, awsClientset)
	ddbcontroller.StartWatch(v1.NamespaceAll, stopChan)

	sqscontroller := sqsqueue.NewController(config, context, awsClientset)
	sqscontroller.StartWatch(v1.NamespaceAll, stopChan)

	ecrcontroller := ecrrepository.NewController(config, context, awsClientset)
	ecrcontroller.StartWatch(v1.NamespaceAll, stopChan)

	snscontroller := snstopic.NewController(config, context, awsClientset)
	snscontroller.StartWatch(v1.NamespaceAll, stopChan)

}

func getClientConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}

func createContext(kubeconfig string) (*opkit.Context, *rest.Config, awsclient.OperatorV1alpha1Interface, error) {
	config, err := getClientConfig(kubeconfig)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get k8s config. %+v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get k8s client. %+v", err)
	}

	apiExtClientset, err := apiextensionsclient.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create k8s API extension clientset. %+v", err)
	}

	awsclientset, err := awsclient.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create object store clientset. %+v", err)
	}

	context := &opkit.Context{
		Clientset:             clientset,
		APIExtensionClientset: apiExtClientset,
		Interval:              500 * time.Millisecond,
		Timeout:               60 * time.Second,
	}
	return context, config, awsclientset, nil
}
