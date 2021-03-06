/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

type MacLearningCounters struct {

	// The number of packets with unknown source MAC address that are dispatched without learning the source MAC address. Applicable only when the MAC limit is reached and MAC Limit policy is MAC_LEARNING_LIMIT_POLICY_ALLOW.
	MacNotLearnedPacketsAllowed int64 `json:"mac_not_learned_packets_allowed,omitempty"`

	// The number of packets with unknown source MAC address that are dropped without learning the source MAC address. Applicable only when the MAC limit is reached and MAC Limit policy is MAC_LEARNING_LIMIT_POLICY_DROP.
	MacNotLearnedPacketsDropped int64 `json:"mac_not_learned_packets_dropped,omitempty"`

	// Number of MACs learned
	MacsLearned int64 `json:"macs_learned,omitempty"`
}
