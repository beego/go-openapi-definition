// Copyright 2022 beego
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package oai

type API struct {
	Openapi           string               `json:"openapi" yaml:"openapi"`
	Info              Info                 `json:"info" yaml:"info"`
	JsonSchemaDialect string               `json:"jsonSchemaDialect" yaml:"jsonSchemaDialect"`
	Servers           []Server             `json:"servers" yaml:"servers"`
	Paths             map[string]Path      `json:"paths" yaml:"paths"`
	Webhooks          map[string]PathValue `json:"webhooks" yaml:"webhooks"`
	Components        Components
}

type Components struct {
	Schemas         map[string]Schema           `json:"schemas" yaml:"schemas"`
	Responses       map[string]ResponseValue    `json:"responses" yaml:"responses"`
	Parameters      map[string]ParameterValue   `json:"parameters" yaml:"parameters"`
	Examples        map[string]ExampleValue     `json:"examples" yaml:"examples"`
	RequestBodies   map[string]RequestBodyValue `json:"requestBodies" yaml:"requestBodies"`
	Headers         map[string]HeaderValue      `json:"headers" yaml:"headers"`
	SecuritySchemes map[string]SecurityScheme   `json:"securitySchemes" yaml:"securitySchemes"`
	Links           map[string]LinkValue        `json:"links" yaml:"links"`
	Callbacks       map[string]CallbackValue    `json:"callbacks" yaml:"callbacks"`
	PathItems       map[string]PathValue        `json:"pathItems" yaml:"pathItems"`
	Security        SecurityRequirement         `json:"security" yaml:"security"`
	Tags            []Tag                       `json:"tags" yaml:"tags"`
	ExternalDocs    ExternalDoc                 `json:"externalDocs" yaml:"externalDocs"`
}

type Tag struct {
	Name         string      `json:"name" yaml:"name"`
	Description  string      `json:"description" yaml:"description"`
	ExternalDocs ExternalDoc `json:"externalDocs" yaml:"externalDocs"`
}

type SecurityScheme struct {
	Type             string     `json:"type" yaml:"type"`
	Description      string     `json:"description" yaml:"description"`
	Name             string     `json:"name" yaml:"name"`
	In               string     `json:"in" yaml:"in"`
	Scheme           string     `json:"scheme" yaml:"scheme"`
	BearerFormat     string     `json:"bearerFormat" yaml:"bearerFormat"`
	Flows            OAuthFlows `json:"flows" yaml:"flows"`
	OpenIdConnectUrl string     `json:"openIdConnectUrl" yaml:"openIdConnectUrl"`
}

type OAuthFlows struct {
	Implicit          OAuthFlow `json:"implicit" yaml:"implicit"`
	Password          OAuthFlow `json:"password" yaml:"password"`
	ClientCredentials OAuthFlow `json:"clientCredential" yaml:"clientCredential"`
	AuthorizationCode OAuthFlow `json:"authorizationCode" yaml:"authorizationCode"`
}

type OAuthFlow struct {
	AuthorizationUrl string            `json:"authorizationUrl" yaml:"authorizationUrl"`
	TokenUrl         string            `json:"tokenUrl" yaml:"tokenUrl"`
	RefreshUrl       string            `json:"RefreshUrl" yaml:"refreshUrl"`
	Scopes           map[string]string `json:"scopes" yaml:"scopes"`
}

type ResponseValue struct {
	Response
	Reference
}

type PathValue struct {
	Path
	Reference
}

type Path struct {
	Ref         string           `json:"$ref" yaml:"$ref"`
	Summary     string           `json:"summary" yaml:"summary"`
	Description string           `json:"description" yaml:"description"`
	Get         Operation        `json:"get" yaml:"get"`
	Put         Operation        `json:"put" yaml:"put"`
	Post        Operation        `json:"post" yaml:"post"`
	Delete      Operation        `json:"delete" yaml:"delete"`
	Options     Operation        `json:"options" yaml:"options"`
	Head        Operation        `json:"head" yaml:"head"`
	Patch       Operation        `json:"patch" yaml:"patch"`
	Trace       Operation        `json:"trace" yaml:"trace"`
	Servers     []Server         `json:"servers" yaml:"servers"`
	Parameters  []ParameterValue `json:"parameters" yaml:"parameters"`
}

type Operation struct {
	Tags         []string                 `json:"tags" yaml:"tags"`
	Summary      string                   `json:"summary" yaml:"summary"`
	Description  string                   `json:"description" yaml:"description"`
	ExternalDocs ExternalDoc              `json:"externalDocs" yaml:"externalDocs"`
	OperationId  string                   `json:"operationId" yaml:"operationId"`
	Parameters   []ParameterValue         `json:"parameters" yaml:"parameters"`
	RequestBody  RequestBodyValue         `json:"requestBody" yaml:"requestBody"`
	Responses    map[string]Response      `json:"responses" yaml:"responses"`
	Callbacks    map[string]CallbackValue `json:"callbacks" yaml:"callbacks"`
	Deprecated   bool                     `json:"deprecated" yaml:"deprecated"`
	Security     []SecurityRequirement    `json:"security" yaml:"security"`
	Servers      []Server                 `json:"servers" yaml:"servers"`
}

type SecurityRequirement map[string][]string

type CallbackValue struct {
	CallbackMap
	Reference
}
type CallbackMap map[string]Callback

type Callback struct {
	Path
	Reference
}

type Response struct {
	Description string                 `json:"description" yaml:"description"`
	Headers     map[string]HeaderValue `json:"headers" yaml:"headers"`
	Content     map[string]MediaType   `json:"content" yaml:"content"`
	Links       map[string]LinkValue   `json:"links" yaml:"links"`
}

type LinkValue struct {
	Link
	Reference
}

type Link struct {
	OperationRef string            `json:"operationRef" yaml:"operationRef"`
	OperationId  string            `json:"operationId" yaml:"operationId"`
	Parameters   map[string]string `json:"parameters" yaml:"parameters"`
	RequestBody  string            `json:"requestBody" yaml:"requestBody"`
	Description  string            `json:"description" yaml:"description"`
	Server       Server            `json:"server" yaml:"server"`
}

type RequestBodyValue struct {
	RequestBody
	Reference
}

type RequestBody struct {
	Description string               `json:"description" yaml:"description"`
	Content     map[string]MediaType `json:"content" yaml:"content"`
	Required    bool                 `json:"required" yaml:"required"`
}

type ParameterValue struct {
	Parameter
	Reference
}

type Parameter struct {
	Name            string                  `json:"name" yaml:"name"`
	In              string                  `json:"in" yaml:"in"`
	Description     string                  `json:"description" yaml:"description"`
	Required        bool                    `json:"required" yaml:"required"`
	Deprecated      bool                    `json:"deprecated" yaml:"deprecated"`
	AllowEmptyValue bool                    `json:"allowEmptyValue" yaml:"allowEmptyValue"`
	Style           string                  `json:"style" yaml:"style"`
	Explode         bool                    `json:"explode" yaml:"explode"`
	AllowReserved   bool                    `json:"allowReserved" yaml:"allowReserved"`
	Schema          Schema                  `json:"schema" yaml:"schema"`
	Example         ExampleValue            `json:"example" yaml:"example"`
	Examples        map[string]ExampleValue `json:"examples" yaml:"examples"`
	Content         map[string]MediaType    `json:"content" yaml:"content"`
}

type MediaType struct {
	Schema   Schema                  `json:"schema" yaml:"schema"`
	Example  ExampleValue            `json:"example" yaml:"example"`
	Examples map[string]ExampleValue `json:"examples" yaml:"examples"`
	Encoding map[string]Encoding     `json:"encoding" yaml:"encoding"`
}

type Encoding struct {
	ContentType   string                 `json:"contentType" yaml:"contentType"`
	Headers       map[string]HeaderValue `json:"headers" yaml:"headers"`
	Style         string                 `json:"style" yaml:"style"`
	Explode       bool                   `json:"explode" yaml:"explode"`
	AllowReserved bool                   `json:"allowReserved" yaml:"allowReserved"`
}

type HeaderValue struct {
	Header
	Reference
}

type Header Parameter

type ExampleValue struct {
	Example
	Reference
}

type Reference struct {
	Ref         string `json:"$ref" yaml:"$ref"`
	Summary     string `json:"summary" yaml:"summary"`
	Description string `json:"description" yaml:"description"`
}

type Example struct {
	Summary       string      `json:"summary" yaml:"summary"`
	Description   string      `json:"description" yaml:"description"`
	Value         interface{} `json:"value" yaml:"value"`
	ExternalValue string      `json:"externalValue" yaml:"externalValue"`
}

type Schema struct {
	Discriminator Discriminator `json:"discriminator" yaml:"discriminator"`
	XML           XML           `json:"xml" yaml:"xml"`
	ExternalDocs  ExternalDoc   `json:"externalDocs" yaml:"externalDocs"`
}

type XML struct {
	Name      string `json:"name" yaml:"name"`
	Namespace string `json:"namespace" yaml:"namespace"`
	Prefix    string `json:"prefix" yaml:"prefix"`
	Attribute bool   `json:"attribute" yaml:"attribute"`
	Wrapped   bool   `json:"wrapped" yaml:"wrapped"`
}

type Discriminator struct {
	PropertyName string            `json:"propertyName" yaml:"propertyName"`
	Mapping      map[string]string `json:"mapping" yaml:"mapping"`
}

type ExternalDoc struct {
	Description string `json:"description" yaml:"description"`
	Url         string `json:"url" yaml:"url"`
}

type Server struct {
	Url         string                    `json:"url" yaml:"url"`
	Description string                    `json:"description" yaml:"description"`
	Variables   map[string]ServerVariable `json:"variables" yaml:"variables"`
}

type ServerVariable struct {
	Enum        []string `json:"enum" yaml:"enum"`
	Default     string   `json:"default" yaml:"default"`
	Description string   `json:"description" yaml:"description"`
}

type Info struct {
	Title          string  `json:"title" yaml:"title"`
	Summary        string  `json:"summary" yaml:"summary"`
	Description    string  `json:"description" yaml:"description"`
	TermsOfService string  `json:"termsOfService" yaml:"termsOfService"`
	Contact        Contact `json:"contact" yaml:"contact"`
	License        License `json:"license" yaml:"license"`
	Version        string  `json:"version" yaml:"version"`
}

type License struct {
	Name       string `json:"name" yaml:"name"`
	Identifier string `json:"identifier" yaml:"identifier"`
	Url        string `json:"url" yaml:"url"`
}

type Contact struct {
	Name  string `json:"name" yaml:"name"`
	Url   string `json:"url" yaml:"url"`
	Email string `json:"email" yaml:"email"`
}
