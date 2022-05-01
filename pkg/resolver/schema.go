package resolver

import (
	"go-elaborate/pkg/document"
)

const (
	statusSchemaUnknown = "unknown"
	statusSchemaPassed  = "passed"
	statusSchemaFailed  = "failed"
)

type SchemaConfig struct {
	includeSchema bool
}

func Schema(resolver Resolver, schema document.Schema, config SchemaConfig) Resolver {
	res := &schemaResolver{
		Typ:      resolverTypeSchema,
		Resolver: resolver,
		Status:   statusSchemaUnknown,
		schema:   schema,
	}
	if config.includeSchema {
		res.Schema = schema
	}
	return res
}

type schemaResolver struct {
	Typ      string      `json:"type" yaml:"type"`
	Status   string      `json:"status" yaml:"status"`
	Schema   interface{} `json:"schema,omitempty" yaml:"schema,omitempty"`
	Resolver Resolver    `json:"resolver" yaml:"resolver"`
	schema   document.Schema
}

func (s *schemaResolver) Get() (interface{}, error) {
	data, err := s.Resolver.Get()
	if err != nil || !s.schema.Verify(data) {
		s.Status = statusSchemaFailed
	} else {
		s.Status = statusSchemaPassed
	}
	return data, err
}
