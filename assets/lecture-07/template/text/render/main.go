package main

import (
	"html/template"
	"log"
	"os"
)

// DATA START OMIT

type User struct {
	Name     string   `json:"name,omitempty"`
	Active   bool     `json:"-"`
	Age      uint     `json:"age,omitempty"`
	Address  Address  `json:"address,omitempty"`
	Phones   []Phone  `json:"phone_numbers,omitempty"`
	Children []string `json:"children,omitempty"`
}

type Address struct {
	Street     string `json:"street,omitempty"`
	City       string `json:"city,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

type Phone struct {
	Type   string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

// DATA END OMIT

var user User = User{
	Name:   "Bob",
	Active: true,
	Age:    27,
	Address: Address{
		Street:     "Botanická 68A",
		City:       "Brno - Královo pole",
		PostalCode: "60200",
	},
	Phones: []Phone{
		{
			Type:   "mobile",
			Number: "269 555-1234",
		},
		{
			Type:   "office",
			Number: "269 555-1269",
		},
	},
	Children: []string{"Bob", "Rob"},
}

var userTemplate string = `--- User ---
Name: {{ .Name }}
Status: {{ if .Active }}ONLINE{{ else }}OFFLINE{{ end }}
Age: {{ .Age }}
Street: {{ .Address.Street }}
Phones [{{ len .Phones }}]:
{{- range .Phones }}
	{{ .Type }}: {{ .Number -}}
{{- end }}
Children:
{{- range .Children }} 
	Name: {{ . }} {{ if eq . "Bob" }}[My favourite child]{{ end -}}
{{- end }}
--- User ---
`

// RENDER START OMIT

func main() {
	tmpl, err := template.New("user").Parse(userTemplate)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(os.Stdout, user)
	if err != nil {
		log.Fatal(err)
	}
}

// RENDER END OMIT
