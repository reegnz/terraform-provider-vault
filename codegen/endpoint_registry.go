package codegen

import "github.com/hashicorp/vault/sdk/framework"

// endpointRegistry is a registry of all the endpoints we'd
// like to have generated, along with the type of template
// we should use.
// IMPORTANT NOTE: To support high quality, only add one
// endpoint per PR.
var endpointRegistry = map[string]*additionalInfo{
	"/transform/role/{name}": {
		TemplateType: templateTypeResource,
	},
	"/transform/transformation/{name}": {
		TemplateType: templateTypeResource,
		AdditionalParameters: []*templatableParam{
			{
				OASParameter: &framework.OASParameter{
					Name:        "templates",
					Description: "Templates configured for transformation.",
					Schema: &framework.OASSchema{
						Type: "array",
						Items: &framework.OASSchema{
							Type: "string",
						},
						DisplayAttrs: &framework.DisplayAttributes{},
					},
				},
				Computed: true,
			},
		},
	},
}

type additionalInfo struct {
	TemplateType         templateType
	AdditionalParameters []*templatableParam
}
