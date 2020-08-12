/*
Copyright 2020 The Kubernetes Authors.

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

package opentelekomcloud

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"k8s.io/klog"
)

// CloudConfig stores every thing for cloud controller manager.
type CloudConfig struct {
	Auth         AuthOpts `json:"Auth"`
	LoadBalancer LBConfig `json:"LoadBalancer"`
}

// AuthOpts stores authentication related configuration.
type AuthOpts struct {
	// SecretName stores the AK/SK.
	// If non-empty the access key and secret key will be obtained dynamically from secret,
	// that means the access key and secret key from configuration will be ignored.
	SecretName string `json:"SecretName"`
	// AccessKey is the permanent access key.
	// AK/SK generation description: Log in to the management console, choose My Credential,
	// and click Access Keys to create an AK and SK.
	AccessKey string `json:"AccessKey"`
	// SecretKey is the permanent secret key.
	// AK/SK generation description: Log in to the management console, choose My Credential,
	// and click Access Keys to create an AK and SK.
	SecretKey string `json:"SecretKey"`
	// IAMEndpoint is the IAM(Identity and Access Management) service's endpoint.
	// E.g. 'https://iam.eu-de.otc.t-systems.com/v3/'.
	IAMEndpoint string `json:"IAMEndpoint"`
	// DomainID is the account ID.
	// Please refer to https://docs.otc.t-systems.com/en-us/api/apiug/apig-en-api-180328009.html.
	// E.g. '052ca6e3530010490f52c0135f7ff501'
	DomainID string `json:"DomainID"`
	// ProjectID is the project ID your workload working on.
	// Please refer to https://docs.otc.t-systems.com/en-us/api/apiug/apig-en-api-180328009.html.
	// E.g. '052d4df9f8800f2f2f99c0134ed5b282'
	ProjectID string `json:"ProjectID"`
	// Region is the region name.
	// E.g. 'ap-southeast-1'
	Region string `json:"Region"`
	// Cloud is the cloud platform domain name.
	// E.G. 'myhwclouds.com'
	Cloud string `json:"Cloud"`
}

func ReadConf(config io.Reader) (*CloudConfig, error) {
	configBytes, err := ioutil.ReadAll(config)
	if err != nil {
		klog.Errorf("Read config failed with error: %v", err)
		return nil, err
	}

	var cfg CloudConfig

	err = json.Unmarshal(configBytes, &cfg)
	if err != nil {
		klog.Errorf("Unmarshal config failed with error: %v", err)
		return nil, err
	}

	return &cfg, nil
}

func LogConf(cfg *CloudConfig) {
	klog.Infof("Log conf, Auth.IAMEndpoint: %s", cfg.Auth.IAMEndpoint)
	klog.Infof("Log conf, LoadBalancer.SecretName: %s", cfg.LoadBalancer.SecretName)
}
