package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{marshal .Schemes}},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {}
}`

var SwaggerInfo = &swag.Spec{
    Version:          "1.0.0",
    Host:             "localhost", // Change this to your actual host
    BasePath:         "/api/v1",         // Adjust the base path accordingly
    Schemes:          []string{"http", "https"},
    Title:            "BookStore Application",
    Description:      "A BookStore Application in GoLang",
    InfoInstanceName: "swagger",
    SwaggerTemplate:  docTemplate,
    LeftDelim:        "{{",
    RightDelim:       "}}",
}

func init() {
    swag.Register(SwaggerInfo.InfoInstanceName, SwaggerInfo)
}