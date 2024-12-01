package templates

func TemplateJSX() string {
	templateJsx := `
	import React from 'react';

	export default function Model(){
		return (
			<group>
				{{range .Nodes}}
				<mesh
					rotation=[{{index .Rotation 0}}, {{index .Rotation 1}}, {{index .Rotation 2}}]
					scale=[{{index .Scale 0}}, {{index .Scale 1}}, {{index .Scale 2}}]
					position=[{{index .Position 0}}, {{index .Position 1}}, {{index .Position 2}}]
				>
				</mesh>
				{{end}}
			</group>

		)
	}
	`

	return templateJsx
}
