/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EscalationRuleInitParameters struct {
}

type EscalationRuleObservation struct {

	// The number of minutes before an unacknowledged incident escalates away from this rule.
	EscalationDelayInMinutes *float64 `json:"escalationDelayInMinutes,omitempty" tf:"escalation_delay_in_minutes,omitempty"`

	// ID of the user defined as the responder
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The targets an incident should be assigned to upon reaching this rule.
	Target []TargetObservation `json:"target,omitempty" tf:"target,omitempty"`
}

type EscalationRuleParameters struct {
}

type PlayInitParameters struct {

	// The telephone number that will be set as the conference number for any incident on which this response play is run.
	ConferenceNumber *string `json:"conferenceNumber,omitempty" tf:"conference_number,omitempty"`

	// The URL that will be set as the conference URL for any incident on which this response play is run.
	ConferenceURL *string `json:"conferenceUrl,omitempty" tf:"conference_url,omitempty"`

	// A human-friendly description of the response play.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The email of the user attributed to the request. Needs to be a valid email address of a user in the PagerDuty account.
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// The name of the response play.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A user and/or escalation policy to be requested as a responder to any incident on which this response play is run. There can be multiple responders defined on a single response play.
	Responder []ResponderInitParameters `json:"responder,omitempty" tf:"responder,omitempty"`

	// The message body of the notification that will be sent to this response play's set of responders. If empty, a default response request notification will be sent.
	RespondersMessage *string `json:"respondersMessage,omitempty" tf:"responders_message,omitempty"`

	// String representing how this response play is allowed to be run. Valid options are:
	Runnability *string `json:"runnability,omitempty" tf:"runnability,omitempty"`

	// A user and/or team to be added as a subscriber to any incident on which this response play is run. There can be multiple subscribers defined on a single response play.
	Subscriber []SubscriberInitParameters `json:"subscriber,omitempty" tf:"subscriber,omitempty"`

	// The content of the notification that will be sent to all incident subscribers upon the running of this response play. Note that this includes any users who may have already been subscribed to the incident prior to the running of this response play. If empty, no notifications will be sent.
	SubscribersMessage *string `json:"subscribersMessage,omitempty" tf:"subscribers_message,omitempty"`

	// The ID of the team associated with the response play.
	Team *string `json:"team,omitempty" tf:"team,omitempty"`

	// A string that determines the schema of the object. If not set, the default value is "response_play".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PlayObservation struct {

	// The telephone number that will be set as the conference number for any incident on which this response play is run.
	ConferenceNumber *string `json:"conferenceNumber,omitempty" tf:"conference_number,omitempty"`

	// The URL that will be set as the conference URL for any incident on which this response play is run.
	ConferenceURL *string `json:"conferenceUrl,omitempty" tf:"conference_url,omitempty"`

	// A human-friendly description of the response play.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The email of the user attributed to the request. Needs to be a valid email address of a user in the PagerDuty account.
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// ID of the user defined as the responder
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the response play.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A user and/or escalation policy to be requested as a responder to any incident on which this response play is run. There can be multiple responders defined on a single response play.
	Responder []ResponderObservation `json:"responder,omitempty" tf:"responder,omitempty"`

	// The message body of the notification that will be sent to this response play's set of responders. If empty, a default response request notification will be sent.
	RespondersMessage *string `json:"respondersMessage,omitempty" tf:"responders_message,omitempty"`

	// String representing how this response play is allowed to be run. Valid options are:
	Runnability *string `json:"runnability,omitempty" tf:"runnability,omitempty"`

	// A user and/or team to be added as a subscriber to any incident on which this response play is run. There can be multiple subscribers defined on a single response play.
	Subscriber []SubscriberObservation `json:"subscriber,omitempty" tf:"subscriber,omitempty"`

	// The content of the notification that will be sent to all incident subscribers upon the running of this response play. Note that this includes any users who may have already been subscribed to the incident prior to the running of this response play. If empty, no notifications will be sent.
	SubscribersMessage *string `json:"subscribersMessage,omitempty" tf:"subscribers_message,omitempty"`

	// The ID of the team associated with the response play.
	Team *string `json:"team,omitempty" tf:"team,omitempty"`

	// A string that determines the schema of the object. If not set, the default value is "response_play".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PlayParameters struct {

	// The telephone number that will be set as the conference number for any incident on which this response play is run.
	// +kubebuilder:validation:Optional
	ConferenceNumber *string `json:"conferenceNumber,omitempty" tf:"conference_number,omitempty"`

	// The URL that will be set as the conference URL for any incident on which this response play is run.
	// +kubebuilder:validation:Optional
	ConferenceURL *string `json:"conferenceUrl,omitempty" tf:"conference_url,omitempty"`

	// A human-friendly description of the response play.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The email of the user attributed to the request. Needs to be a valid email address of a user in the PagerDuty account.
	// +kubebuilder:validation:Optional
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// The name of the response play.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A user and/or escalation policy to be requested as a responder to any incident on which this response play is run. There can be multiple responders defined on a single response play.
	// +kubebuilder:validation:Optional
	Responder []ResponderParameters `json:"responder,omitempty" tf:"responder,omitempty"`

	// The message body of the notification that will be sent to this response play's set of responders. If empty, a default response request notification will be sent.
	// +kubebuilder:validation:Optional
	RespondersMessage *string `json:"respondersMessage,omitempty" tf:"responders_message,omitempty"`

	// String representing how this response play is allowed to be run. Valid options are:
	// +kubebuilder:validation:Optional
	Runnability *string `json:"runnability,omitempty" tf:"runnability,omitempty"`

	// A user and/or team to be added as a subscriber to any incident on which this response play is run. There can be multiple subscribers defined on a single response play.
	// +kubebuilder:validation:Optional
	Subscriber []SubscriberParameters `json:"subscriber,omitempty" tf:"subscriber,omitempty"`

	// The content of the notification that will be sent to all incident subscribers upon the running of this response play. Note that this includes any users who may have already been subscribed to the incident prior to the running of this response play. If empty, no notifications will be sent.
	// +kubebuilder:validation:Optional
	SubscribersMessage *string `json:"subscribersMessage,omitempty" tf:"subscribers_message,omitempty"`

	// The ID of the team associated with the response play.
	// +kubebuilder:validation:Optional
	Team *string `json:"team,omitempty" tf:"team,omitempty"`

	// A string that determines the schema of the object. If not set, the default value is "response_play".
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResponderInitParameters struct {

	// A human-friendly description of the response play.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the user defined as the responder
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the response play.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A string that determines the schema of the object. If not set, the default value is "response_play".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResponderObservation struct {

	// A human-friendly description of the response play.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The escalation rules
	EscalationRule []EscalationRuleObservation `json:"escalationRule,omitempty" tf:"escalation_rule,omitempty"`

	// ID of the user defined as the responder
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the response play.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The number of times the escalation policy will repeat after reaching the end of its escalation.
	NumLoops *float64 `json:"numLoops,omitempty" tf:"num_loops,omitempty"`

	// Determines how on call handoff notifications will be sent for users on the escalation policy. Defaults to "if_has_services". Could be "if_has_services", "always
	OnCallHandoffNotifications *string `json:"onCallHandoffNotifications,omitempty" tf:"on_call_handoff_notifications,omitempty"`

	// There can be multiple services associated with a policy.
	Service []ServiceObservation `json:"service,omitempty" tf:"service,omitempty"`

	// The ID of the team associated with the response play.
	Team []TeamObservation `json:"team,omitempty" tf:"team,omitempty"`

	// A string that determines the schema of the object. If not set, the default value is "response_play".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResponderParameters struct {

	// A human-friendly description of the response play.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the user defined as the responder
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the response play.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A string that determines the schema of the object. If not set, the default value is "response_play".
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ServiceInitParameters struct {
}

type ServiceObservation struct {

	// ID of the user defined as the responder
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A string that determines the schema of the object. If not set, the default value is "response_play".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ServiceParameters struct {
}

type SubscriberInitParameters struct {

	// ID of the user defined as the responder
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A string that determines the schema of the object. If not set, the default value is "response_play".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SubscriberObservation struct {

	// ID of the user defined as the responder
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A string that determines the schema of the object. If not set, the default value is "response_play".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SubscriberParameters struct {

	// ID of the user defined as the responder
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A string that determines the schema of the object. If not set, the default value is "response_play".
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TargetInitParameters struct {
}

type TargetObservation struct {

	// ID of the user defined as the responder
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A string that determines the schema of the object. If not set, the default value is "response_play".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TargetParameters struct {
}

type TeamInitParameters struct {
}

type TeamObservation struct {

	// ID of the user defined as the responder
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A string that determines the schema of the object. If not set, the default value is "response_play".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TeamParameters struct {
}

// PlaySpec defines the desired state of Play
type PlaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PlayParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider PlayInitParameters `json:"initProvider,omitempty"`
}

// PlayStatus defines the observed state of Play.
type PlayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PlayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Play is the Schema for the Plays API. Creates and manages a response play in PagerDuty.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type Play struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.from) || (has(self.initProvider) && has(self.initProvider.from))",message="spec.forProvider.from is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   PlaySpec   `json:"spec"`
	Status PlayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlayList contains a list of Plays
type PlayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Play `json:"items"`
}

// Repository type metadata.
var (
	Play_Kind             = "Play"
	Play_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Play_Kind}.String()
	Play_KindAPIVersion   = Play_Kind + "." + CRDGroupVersion.String()
	Play_GroupVersionKind = CRDGroupVersion.WithKind(Play_Kind)
)

func init() {
	SchemeBuilder.Register(&Play{}, &PlayList{})
}
