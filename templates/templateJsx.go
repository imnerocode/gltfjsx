package templates

func TemplateJSX() string {
	templateJsx := `
	import React from 'react';

	export default function Model(){
		return (
			<group>
				{{range .Meshes}}
				<mesh
					rotation={[{{.Node.Rotation 0}}, {{.Node.Rotation 1}}, {{.Node.Rotation 2}}]}
					scale={[{{.Node.Scale 0}}, {{.Node.Scale 1}}, {{.Node.Scale 2}}]}
					position={[{{.Node.Position 0}}, {{.Node.Position 1}}, {{.Node.Position 2}}]}
				>
					<bufferGeometry>
						<bufferAttribute
							attach="attributes-position"
							array={new Float32Array({{.Attributes.Position}})}
							count={{len .Attributes.Position}}
							itemSize={3}
						/>
						<bufferAttribute
							attach="attributes-normal"
							array={new Float32Array({{.Attributes.Normal}})}
							count={{len .Attributes.Normal}}
							itemSize={3}
						/>
						<bufferAttribute
							attach="attributes-uv"
							array={new Float32Array({{.Attributes.TexCoord}})}
							count={{len .Attributes.TexCoord}}
							itemSize={2}
						/>
					</bufferGeometry>
				</mesh>
				{{end}}
			</group>
		)
	}
	`
	return templateJsx
}
