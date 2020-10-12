/*
Package proto contains the definition of protobuffer messages and services related to the promotion service. Note that the
.pb.go and .pb.micro.go files are generated and should not be manually modified.
Note that the promotion message contains one atypical field: readOnlyLookup: This field does not actually exist in the
promotion DB, but it is pulled from the customer service when needed. To reduce the time we need to pull from that
service, we cache the value in the redis store
*/
package proto
