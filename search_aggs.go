// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"fmt"
)

// Aggregations can be seen as a unit-of-work that build
// analytic information over a set of documents. It is
// (in many senses) the follow-up of facets in Elasticsearch.
// For more details about aggregations, visit:
// https://www.elastic.co/guide/en/elasticsearch/reference/6.2/search-aggregations.html
type Aggregation interface {
	// Source returns a JSON-serializable aggregation that is a fragment
	// of the request sent to Elasticsearch.
	Source() (interface{}, error)
}

// Aggregations is a list of aggregations that are part of a search result.
type Aggregations map[string]*AggregationResult

func (a *Aggregations) UnmarshalJSON(data []byte) error {
	var agg map[string]interface{}
	err := jsoniter.Unmarshal(data, &agg)
	if err != nil {
		return err
	}

	if (*a) == nil {
		(*a) = make(Aggregations)
	}

	for k, v := range agg {
		fmt.Printf("%+v, %T \n", k, v)
		(*a)[k] = createResult(v)
	}
	return nil
}

func createResult(value interface{}) *AggregationResult {
	res := new(AggregationResult)
	switch typeVal := value.(type) {
	case map[string]interface{}:
		for k, v := range typeVal {
			switch k {
			case "meta":
				res.Meta = v.(map[string]interface{})
			case "value":
				res.Value = v.(float64)
			case "doc_count":
				res.DocCount = int64(v.(float64))
			case "doc_count_error_upper_bound":
				res.DocCountErrorUpperBound = int64(v.(float64))
			case "sum_other_doc_count":
				res.SumOfOtherDocCount = int64(v.(float64))
			case "buckets":
				switch bucketVal := v.(type) {
				// NamedBuckets.
				case map[string]interface{}:
					res.NamedBuckets = createNamedBuckets(bucketVal)
				// Buckets.
				case []interface{}:
					res.Buckets = createBuckets(bucketVal)
				}
			default:
				if res.Aggregations == nil {
					res.Aggregations = make(Aggregations)
				}
				// is not part of result, is another leaf of Aggregation
				res.Aggregations[k] = createResult(v)
			}
		}
	default:
		fmt.Printf("NOT A MAP: %T \n", typeVal)
		return res
	}
	return res
}

func createBuckets(value []interface{}) []AggregationBucketKeyItem {
	var res []AggregationBucketKeyItem
	for _, bucket := range value {
		res = append(res, createBucket(bucket))
	}
	return res
}

func createNamedBuckets(value map[string]interface{}) map[string]AggregationBucketKeyItem {
	var res = make(map[string]AggregationBucketKeyItem)
	for k, v := range value {
		res[k] = createBucket(v)
	}
	return res
}

func createBucket(value interface{}) AggregationBucketKeyItem {
	var res AggregationBucketKeyItem
	switch typeVal := value.(type) {
	case map[string]interface{}:
		for k, v := range typeVal {
			switch k {
			case "key":
				res.Key = v
			case "key_as_string":
				res.KeyAsString = v.(string)
			case "doc_count":
				res.DocCount = int64(v.(float64))
			default:
				if res.Aggregations == nil {
					res.Aggregations = make(Aggregations)
				}
				// is not part of result, is another leaf of Aggregation
				res.Aggregations[k] = createResult(v)
			}
		}
	default:
		fmt.Printf("NOT A MAP: %T \n", typeVal)
		return res
	}
	return res
}

type AggregationResult struct {
	Aggregations
	Meta map[string]interface{} `json:"meta,omitempty"`

	// AggregationValueMetric
	Value float64 `json:"value"`

	// AggregationSingleBucket
	DocCount int64 `json:"doc_count"`

	// AggregationBucketFilters
	Buckets      []AggregationBucketKeyItem          `json:"buckets"`
	NamedBuckets map[string]AggregationBucketKeyItem `json:"buckets"`

	// AggregationBucketKeyItems
	DocCountErrorUpperBound int64 `json:"doc_count_error_upper_bound"`
	SumOfOtherDocCount      int64 `json:"sum_other_doc_count"`
}

// AggregationBucketKeyItem is a single bucket of an AggregationBucketKeyItems structure.
type AggregationBucketKeyItem struct {
	Aggregations

	Key         interface{} //`json:"key"`
	KeyAsString string      //`json:"key_as_string"`
	DocCount    int64       //`json:"doc_count"`
}

// Sum returns sum aggregation results.
// See: https://www.elastic.co/guide/en/elasticsearch/reference/6.2/search-aggregations-metrics-sum-aggregation.html
func (a Aggregations) Sum(name string) (*AggregationResult, bool) {
	_, ok := a[name]
	return nil, ok
}

// Filter returns filter results.
// See: https://www.elastic.co/guide/en/elasticsearch/reference/6.2/search-aggregations-bucket-filter-aggregation.html
func (a Aggregations) Filter(name string) (*AggregationResult, bool) {
	_, ok := a[name]
	return nil, ok
}

// Filters returns filters results.
// See: https://www.elastic.co/guide/en/elasticsearch/reference/6.2/search-aggregations-bucket-filters-aggregation.html
func (a Aggregations) Filters(name string) (*AggregationResult, bool) {
	_, ok := a[name]
	return nil, ok
}

// Terms returns terms aggregation results.
// See: https://www.elastic.co/guide/en/elasticsearch/reference/6.2/search-aggregations-bucket-terms-aggregation.html
func (a Aggregations) Terms(name string) (*AggregationResult, bool) {
	_, ok := a[name]
	return nil, ok
}
