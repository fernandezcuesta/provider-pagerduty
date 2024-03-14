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

type ActionsInitParameters struct {

	// Note added to the event.
	Annotate []AnnotateInitParameters `json:"annotate,omitempty" tf:"annotate,omitempty"`

	// An object with a single value field. The value sets whether the resulting alert status is trigger or resolve.
	EventAction []EventActionInitParameters `json:"eventAction,omitempty" tf:"event_action,omitempty"`

	// Allows you to copy important data from one event field to another. Extraction objects may use either of the following field structures:
	Extractions []ExtractionsInitParameters `json:"extractions,omitempty" tf:"extractions,omitempty"`

	// The ID of the priority applied to the event.
	Priority []PriorityInitParameters `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the service where the event will be routed.
	Route []RouteInitParameters `json:"route,omitempty" tf:"route,omitempty"`

	// The severity level of the event. Can be either info,warning,error, or critical.
	Severity []SeverityInitParameters `json:"severity,omitempty" tf:"severity,omitempty"`

	// Controls whether an alert is suppressed (does not create an incident). Note: If a threshold is set, the rule must also have a route action.
	Suppress []SuppressInitParameters `json:"suppress,omitempty" tf:"suppress,omitempty"`

	// An object with a single value field. The value sets the length of time to suspend the resulting alert before triggering. Note: A rule with a suspend action must also have a route action.
	Suspend []SuspendInitParameters `json:"suspend,omitempty" tf:"suspend,omitempty"`
}

type ActionsObservation struct {

	// Note added to the event.
	Annotate []AnnotateObservation `json:"annotate,omitempty" tf:"annotate,omitempty"`

	// An object with a single value field. The value sets whether the resulting alert status is trigger or resolve.
	EventAction []EventActionObservation `json:"eventAction,omitempty" tf:"event_action,omitempty"`

	// Allows you to copy important data from one event field to another. Extraction objects may use either of the following field structures:
	Extractions []ExtractionsObservation `json:"extractions,omitempty" tf:"extractions,omitempty"`

	// The ID of the priority applied to the event.
	Priority []PriorityObservation `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the service where the event will be routed.
	Route []RouteObservation `json:"route,omitempty" tf:"route,omitempty"`

	// The severity level of the event. Can be either info,warning,error, or critical.
	Severity []SeverityObservation `json:"severity,omitempty" tf:"severity,omitempty"`

	// Controls whether an alert is suppressed (does not create an incident). Note: If a threshold is set, the rule must also have a route action.
	Suppress []SuppressObservation `json:"suppress,omitempty" tf:"suppress,omitempty"`

	// An object with a single value field. The value sets the length of time to suspend the resulting alert before triggering. Note: A rule with a suspend action must also have a route action.
	Suspend []SuspendObservation `json:"suspend,omitempty" tf:"suspend,omitempty"`
}

type ActionsParameters struct {

	// Note added to the event.
	// +kubebuilder:validation:Optional
	Annotate []AnnotateParameters `json:"annotate,omitempty" tf:"annotate,omitempty"`

	// An object with a single value field. The value sets whether the resulting alert status is trigger or resolve.
	// +kubebuilder:validation:Optional
	EventAction []EventActionParameters `json:"eventAction,omitempty" tf:"event_action,omitempty"`

	// Allows you to copy important data from one event field to another. Extraction objects may use either of the following field structures:
	// +kubebuilder:validation:Optional
	Extractions []ExtractionsParameters `json:"extractions,omitempty" tf:"extractions,omitempty"`

	// The ID of the priority applied to the event.
	// +kubebuilder:validation:Optional
	Priority []PriorityParameters `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the service where the event will be routed.
	// +kubebuilder:validation:Optional
	Route []RouteParameters `json:"route,omitempty" tf:"route,omitempty"`

	// The severity level of the event. Can be either info,warning,error, or critical.
	// +kubebuilder:validation:Optional
	Severity []SeverityParameters `json:"severity,omitempty" tf:"severity,omitempty"`

	// Controls whether an alert is suppressed (does not create an incident). Note: If a threshold is set, the rule must also have a route action.
	// +kubebuilder:validation:Optional
	Suppress []SuppressParameters `json:"suppress,omitempty" tf:"suppress,omitempty"`

	// An object with a single value field. The value sets the length of time to suspend the resulting alert before triggering. Note: A rule with a suspend action must also have a route action.
	// +kubebuilder:validation:Optional
	Suspend []SuspendParameters `json:"suspend,omitempty" tf:"suspend,omitempty"`
}

type ActiveBetweenInitParameters struct {
	EndTime *float64 `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// A Unix timestamp in milliseconds which is combined with the timezone to determine the time this rule will start on each specified weekday. Note that the date of the timestamp you specify does not matter, except that it lets you determine whether daylight saving time is in effect so that you use the correct UTC offset for the timezone you specify. In practice, you may want to use the  to generate this value, as demonstrated in the resource.pagerduty_ruleset_rule.foo code example at the top of this page. To generate this timestamp manually, if you want your rule to apply starting at 9:30am in the America/New_York timezone, use your programing language of choice to determine a Unix timestamp that represents 9:30am in that timezone, like 1554989400000.
	StartTime *float64 `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type ActiveBetweenObservation struct {
	EndTime *float64 `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// A Unix timestamp in milliseconds which is combined with the timezone to determine the time this rule will start on each specified weekday. Note that the date of the timestamp you specify does not matter, except that it lets you determine whether daylight saving time is in effect so that you use the correct UTC offset for the timezone you specify. In practice, you may want to use the  to generate this value, as demonstrated in the resource.pagerduty_ruleset_rule.foo code example at the top of this page. To generate this timestamp manually, if you want your rule to apply starting at 9:30am in the America/New_York timezone, use your programing language of choice to determine a Unix timestamp that represents 9:30am in that timezone, like 1554989400000.
	StartTime *float64 `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type ActiveBetweenParameters struct {

	// +kubebuilder:validation:Optional
	EndTime *float64 `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// A Unix timestamp in milliseconds which is combined with the timezone to determine the time this rule will start on each specified weekday. Note that the date of the timestamp you specify does not matter, except that it lets you determine whether daylight saving time is in effect so that you use the correct UTC offset for the timezone you specify. In practice, you may want to use the  to generate this value, as demonstrated in the resource.pagerduty_ruleset_rule.foo code example at the top of this page. To generate this timestamp manually, if you want your rule to apply starting at 9:30am in the America/New_York timezone, use your programing language of choice to determine a Unix timestamp that represents 9:30am in that timezone, like 1554989400000.
	// +kubebuilder:validation:Optional
	StartTime *float64 `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type AnnotateInitParameters struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AnnotateObservation struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AnnotateParameters struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConditionsInitParameters struct {

	// Operator to combine sub-conditions. Can be and or or.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// List of sub-conditions that define the condition.
	Subconditions []SubconditionsInitParameters `json:"subconditions,omitempty" tf:"subconditions,omitempty"`
}

type ConditionsObservation struct {

	// Operator to combine sub-conditions. Can be and or or.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// List of sub-conditions that define the condition.
	Subconditions []SubconditionsObservation `json:"subconditions,omitempty" tf:"subconditions,omitempty"`
}

type ConditionsParameters struct {

	// Operator to combine sub-conditions. Can be and or or.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// List of sub-conditions that define the condition.
	// +kubebuilder:validation:Optional
	Subconditions []SubconditionsParameters `json:"subconditions,omitempty" tf:"subconditions,omitempty"`
}

type EventActionInitParameters struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EventActionObservation struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EventActionParameters struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ExtractionsInitParameters struct {

	// The conditions that need to be met for the extraction to happen. Must use valid RE2 regular expression syntax.
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// Field where the data is being copied from. Must be a PagerDuty Common Event Format (PD-CEF) field.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Field where the data is being copied to. Must be a PagerDuty Common Event Format (PD-CEF) field.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// A customized field message. This can also include variables extracted from the payload by using string interpolation.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type ExtractionsObservation struct {

	// The conditions that need to be met for the extraction to happen. Must use valid RE2 regular expression syntax.
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// Field where the data is being copied from. Must be a PagerDuty Common Event Format (PD-CEF) field.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Field where the data is being copied to. Must be a PagerDuty Common Event Format (PD-CEF) field.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// A customized field message. This can also include variables extracted from the payload by using string interpolation.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type ExtractionsParameters struct {

	// The conditions that need to be met for the extraction to happen. Must use valid RE2 regular expression syntax.
	// +kubebuilder:validation:Optional
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// Field where the data is being copied from. Must be a PagerDuty Common Event Format (PD-CEF) field.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Field where the data is being copied to. Must be a PagerDuty Common Event Format (PD-CEF) field.
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// A customized field message. This can also include variables extracted from the payload by using string interpolation.
	// +kubebuilder:validation:Optional
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type ParameterInitParameters struct {
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParameterObservation struct {
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParameterParameters struct {

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParametersInitParameters struct {
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParametersObservation struct {
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParametersParameters struct {

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PriorityInitParameters struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PriorityObservation struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PriorityParameters struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RouteInitParameters struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RouteObservation struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RouteParameters struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RuleInitParameters struct {

	// Actions to apply to an event if the conditions match.
	Actions []ActionsInitParameters `json:"actions,omitempty" tf:"actions,omitempty"`

	// Indicates whether the Event Rule is the last Event Rule of the Ruleset that serves as a catch-all. It has limited functionality compared to other rules and always matches.
	CatchAll *bool `json:"catchAll,omitempty" tf:"catch_all,omitempty"`

	// Conditions evaluated to check if an event matches this event rule. Is always empty for the catch-all rule, though.
	Conditions []ConditionsInitParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Indicates whether the rule is disabled and would therefore not be evaluated.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Position/index of the rule within the ruleset.
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// The ID of the ruleset that the rule belongs to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/ruleset/v1alpha1.Ruleset
	Ruleset *string `json:"ruleset,omitempty" tf:"ruleset,omitempty"`

	// Reference to a Ruleset in ruleset to populate ruleset.
	// +kubebuilder:validation:Optional
	RulesetRef *v1.Reference `json:"rulesetRef,omitempty" tf:"-"`

	// Selector for a Ruleset in ruleset to populate ruleset.
	// +kubebuilder:validation:Optional
	RulesetSelector *v1.Selector `json:"rulesetSelector,omitempty" tf:"-"`

	// Settings for scheduling the rule.
	TimeFrame []TimeFrameInitParameters `json:"timeFrame,omitempty" tf:"time_frame,omitempty"`

	// Populate variables from event payloads and use those variables in other event actions. NOTE: A rule can have multiple
	Variable []VariableInitParameters `json:"variable,omitempty" tf:"variable,omitempty"`
}

type RuleObservation struct {

	// Actions to apply to an event if the conditions match.
	Actions []ActionsObservation `json:"actions,omitempty" tf:"actions,omitempty"`

	// Indicates whether the Event Rule is the last Event Rule of the Ruleset that serves as a catch-all. It has limited functionality compared to other rules and always matches.
	CatchAll *bool `json:"catchAll,omitempty" tf:"catch_all,omitempty"`

	// Conditions evaluated to check if an event matches this event rule. Is always empty for the catch-all rule, though.
	Conditions []ConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Indicates whether the rule is disabled and would therefore not be evaluated.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The ID of the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Position/index of the rule within the ruleset.
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// The ID of the ruleset that the rule belongs to.
	Ruleset *string `json:"ruleset,omitempty" tf:"ruleset,omitempty"`

	// Settings for scheduling the rule.
	TimeFrame []TimeFrameObservation `json:"timeFrame,omitempty" tf:"time_frame,omitempty"`

	// Populate variables from event payloads and use those variables in other event actions. NOTE: A rule can have multiple
	Variable []VariableObservation `json:"variable,omitempty" tf:"variable,omitempty"`
}

type RuleParameters struct {

	// Actions to apply to an event if the conditions match.
	// +kubebuilder:validation:Optional
	Actions []ActionsParameters `json:"actions,omitempty" tf:"actions,omitempty"`

	// Indicates whether the Event Rule is the last Event Rule of the Ruleset that serves as a catch-all. It has limited functionality compared to other rules and always matches.
	// +kubebuilder:validation:Optional
	CatchAll *bool `json:"catchAll,omitempty" tf:"catch_all,omitempty"`

	// Conditions evaluated to check if an event matches this event rule. Is always empty for the catch-all rule, though.
	// +kubebuilder:validation:Optional
	Conditions []ConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Indicates whether the rule is disabled and would therefore not be evaluated.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Position/index of the rule within the ruleset.
	// +kubebuilder:validation:Optional
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// The ID of the ruleset that the rule belongs to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/ruleset/v1alpha1.Ruleset
	// +kubebuilder:validation:Optional
	Ruleset *string `json:"ruleset,omitempty" tf:"ruleset,omitempty"`

	// Reference to a Ruleset in ruleset to populate ruleset.
	// +kubebuilder:validation:Optional
	RulesetRef *v1.Reference `json:"rulesetRef,omitempty" tf:"-"`

	// Selector for a Ruleset in ruleset to populate ruleset.
	// +kubebuilder:validation:Optional
	RulesetSelector *v1.Selector `json:"rulesetSelector,omitempty" tf:"-"`

	// Settings for scheduling the rule.
	// +kubebuilder:validation:Optional
	TimeFrame []TimeFrameParameters `json:"timeFrame,omitempty" tf:"time_frame,omitempty"`

	// Populate variables from event payloads and use those variables in other event actions. NOTE: A rule can have multiple
	// +kubebuilder:validation:Optional
	Variable []VariableParameters `json:"variable,omitempty" tf:"variable,omitempty"`
}

type ScheduledWeeklyInitParameters struct {

	// Length of time the schedule will be active in milliseconds. For example duration = 2 * 60 * 60 * 1000 if you want your rule to apply for 2 hours, from the specified start_time.
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// A Unix timestamp in milliseconds which is combined with the timezone to determine the time this rule will start on each specified weekday. Note that the date of the timestamp you specify does not matter, except that it lets you determine whether daylight saving time is in effect so that you use the correct UTC offset for the timezone you specify. In practice, you may want to use the  to generate this value, as demonstrated in the resource.pagerduty_ruleset_rule.foo code example at the top of this page. To generate this timestamp manually, if you want your rule to apply starting at 9:30am in the America/New_York timezone, use your programing language of choice to determine a Unix timestamp that represents 9:30am in that timezone, like 1554989400000.
	StartTime *float64 `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The name of the timezone for the given schedule, which will be used to determine UTC offset including adjustment for daylight saving time. For example: timezone = "America/Toronto"
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`

	// An integer array representing which days during the week the rule executes. For example weekdays = [1,3,7] would execute on Monday, Wednesday and Sunday.
	Weekdays []*float64 `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type ScheduledWeeklyObservation struct {

	// Length of time the schedule will be active in milliseconds. For example duration = 2 * 60 * 60 * 1000 if you want your rule to apply for 2 hours, from the specified start_time.
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// A Unix timestamp in milliseconds which is combined with the timezone to determine the time this rule will start on each specified weekday. Note that the date of the timestamp you specify does not matter, except that it lets you determine whether daylight saving time is in effect so that you use the correct UTC offset for the timezone you specify. In practice, you may want to use the  to generate this value, as demonstrated in the resource.pagerduty_ruleset_rule.foo code example at the top of this page. To generate this timestamp manually, if you want your rule to apply starting at 9:30am in the America/New_York timezone, use your programing language of choice to determine a Unix timestamp that represents 9:30am in that timezone, like 1554989400000.
	StartTime *float64 `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The name of the timezone for the given schedule, which will be used to determine UTC offset including adjustment for daylight saving time. For example: timezone = "America/Toronto"
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`

	// An integer array representing which days during the week the rule executes. For example weekdays = [1,3,7] would execute on Monday, Wednesday and Sunday.
	Weekdays []*float64 `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type ScheduledWeeklyParameters struct {

	// Length of time the schedule will be active in milliseconds. For example duration = 2 * 60 * 60 * 1000 if you want your rule to apply for 2 hours, from the specified start_time.
	// +kubebuilder:validation:Optional
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// A Unix timestamp in milliseconds which is combined with the timezone to determine the time this rule will start on each specified weekday. Note that the date of the timestamp you specify does not matter, except that it lets you determine whether daylight saving time is in effect so that you use the correct UTC offset for the timezone you specify. In practice, you may want to use the  to generate this value, as demonstrated in the resource.pagerduty_ruleset_rule.foo code example at the top of this page. To generate this timestamp manually, if you want your rule to apply starting at 9:30am in the America/New_York timezone, use your programing language of choice to determine a Unix timestamp that represents 9:30am in that timezone, like 1554989400000.
	// +kubebuilder:validation:Optional
	StartTime *float64 `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The name of the timezone for the given schedule, which will be used to determine UTC offset including adjustment for daylight saving time. For example: timezone = "America/Toronto"
	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`

	// An integer array representing which days during the week the rule executes. For example weekdays = [1,3,7] would execute on Monday, Wednesday and Sunday.
	// +kubebuilder:validation:Optional
	Weekdays []*float64 `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type SeverityInitParameters struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type SeverityObservation struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type SeverityParameters struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type SubconditionsInitParameters struct {

	// Operator to combine sub-conditions. Can be and or or.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Parameter for the sub-condition. It requires both a path and value to be set.
	Parameter []ParameterInitParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`
}

type SubconditionsObservation struct {

	// Operator to combine sub-conditions. Can be and or or.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Parameter for the sub-condition. It requires both a path and value to be set.
	Parameter []ParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`
}

type SubconditionsParameters struct {

	// Operator to combine sub-conditions. Can be and or or.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Parameter for the sub-condition. It requires both a path and value to be set.
	// +kubebuilder:validation:Optional
	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`
}

type SuppressInitParameters struct {

	// The number value of the threshold_time_unit before an incident is created. Must be greater than 0.
	ThresholdTimeAmount *float64 `json:"thresholdTimeAmount,omitempty" tf:"threshold_time_amount,omitempty"`

	// The seconds,minutes, or hours the threshold_time_amount should be measured.
	ThresholdTimeUnit *string `json:"thresholdTimeUnit,omitempty" tf:"threshold_time_unit,omitempty"`

	// The number of alerts that should be suppressed. Must be greater than 0.
	ThresholdValue *float64 `json:"thresholdValue,omitempty" tf:"threshold_value,omitempty"`

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *bool `json:"value,omitempty" tf:"value,omitempty"`
}

type SuppressObservation struct {

	// The number value of the threshold_time_unit before an incident is created. Must be greater than 0.
	ThresholdTimeAmount *float64 `json:"thresholdTimeAmount,omitempty" tf:"threshold_time_amount,omitempty"`

	// The seconds,minutes, or hours the threshold_time_amount should be measured.
	ThresholdTimeUnit *string `json:"thresholdTimeUnit,omitempty" tf:"threshold_time_unit,omitempty"`

	// The number of alerts that should be suppressed. Must be greater than 0.
	ThresholdValue *float64 `json:"thresholdValue,omitempty" tf:"threshold_value,omitempty"`

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *bool `json:"value,omitempty" tf:"value,omitempty"`
}

type SuppressParameters struct {

	// The number value of the threshold_time_unit before an incident is created. Must be greater than 0.
	// +kubebuilder:validation:Optional
	ThresholdTimeAmount *float64 `json:"thresholdTimeAmount,omitempty" tf:"threshold_time_amount,omitempty"`

	// The seconds,minutes, or hours the threshold_time_amount should be measured.
	// +kubebuilder:validation:Optional
	ThresholdTimeUnit *string `json:"thresholdTimeUnit,omitempty" tf:"threshold_time_unit,omitempty"`

	// The number of alerts that should be suppressed. Must be greater than 0.
	// +kubebuilder:validation:Optional
	ThresholdValue *float64 `json:"thresholdValue,omitempty" tf:"threshold_value,omitempty"`

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	// +kubebuilder:validation:Optional
	Value *bool `json:"value,omitempty" tf:"value,omitempty"`
}

type SuspendInitParameters struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type SuspendObservation struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type SuspendParameters struct {

	// Boolean value that indicates if the alert should be suppressed before the indicated threshold values are met.
	// +kubebuilder:validation:Optional
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type TimeFrameInitParameters struct {

	// Values for executing the rule during a specific time period.
	ActiveBetween []ActiveBetweenInitParameters `json:"activeBetween,omitempty" tf:"active_between,omitempty"`

	// Values for executing the rule on a recurring schedule.
	ScheduledWeekly []ScheduledWeeklyInitParameters `json:"scheduledWeekly,omitempty" tf:"scheduled_weekly,omitempty"`
}

type TimeFrameObservation struct {

	// Values for executing the rule during a specific time period.
	ActiveBetween []ActiveBetweenObservation `json:"activeBetween,omitempty" tf:"active_between,omitempty"`

	// Values for executing the rule on a recurring schedule.
	ScheduledWeekly []ScheduledWeeklyObservation `json:"scheduledWeekly,omitempty" tf:"scheduled_weekly,omitempty"`
}

type TimeFrameParameters struct {

	// Values for executing the rule during a specific time period.
	// +kubebuilder:validation:Optional
	ActiveBetween []ActiveBetweenParameters `json:"activeBetween,omitempty" tf:"active_between,omitempty"`

	// Values for executing the rule on a recurring schedule.
	// +kubebuilder:validation:Optional
	ScheduledWeekly []ScheduledWeeklyParameters `json:"scheduledWeekly,omitempty" tf:"scheduled_weekly,omitempty"`
}

type VariableInitParameters struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Parameters []ParametersInitParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type VariableObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Parameters []ParametersObservation `json:"parameters,omitempty" tf:"parameters,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type VariableParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters []ParametersParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// RuleSpec defines the desired state of Rule
type RuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleParameters `json:"forProvider"`
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
	InitProvider RuleInitParameters `json:"initProvider,omitempty"`
}

// RuleStatus defines the observed state of Rule.
type RuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Rule is the Schema for the Rules API. Creates and manages a ruleset rule in PagerDuty.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuleSpec   `json:"spec"`
	Status            RuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleList contains a list of Rules
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rule `json:"items"`
}

// Repository type metadata.
var (
	Rule_Kind             = "Rule"
	Rule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rule_Kind}.String()
	Rule_KindAPIVersion   = Rule_Kind + "." + CRDGroupVersion.String()
	Rule_GroupVersionKind = CRDGroupVersion.WithKind(Rule_Kind)
)

func init() {
	SchemeBuilder.Register(&Rule{}, &RuleList{})
}
