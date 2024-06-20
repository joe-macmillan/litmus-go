package types

import (
	clientTypes "k8s.io/apimachinery/pkg/types"
)

// ExperimentDetails is for collecting all the experiment-related details
type ExperimentDetails struct {
	AppKind                       string
	AppLabel                      string
	AppNS                         string
	AuxiliaryAppInfo              string
	ChaosDuration                 int
	ChaosNamespace                string
	ChaosPodName                  string
	ChaosUID                      clientTypes.UID
	Delay                         int
	EngineName                    string
	ExperimentName                string
	InstanceID                    string
	KillCommand                   string
	LIBImage                      string
	LIBImagePullPolicy            string
	NodeLabel                     string
	RampTime                      int
	RunID                         string
	SetHelperData                 string
	SSHUser                       string
	TargetContainer               string
	TargetNode                    string
	TargetNodeIP                  string
	TerminationGracePeriodSeconds int
	Timeout                       int
}
