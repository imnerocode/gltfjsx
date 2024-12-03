package templates

func TemplateJSX() string {
	return `
	import React from 'react';

	export default function Model() {
		const attributesPosition = {{ .Position }};
		const attributesNormal = {{ .Normal }};
		const attributesIndices = {{ .Indices }};

		const countPosition = attributesPosition.length / 3;
		const countNormal = attributesNormal.length / 3;

		return (
			<group>
				<mesh>
					<bufferGeometry>
						<bufferAttribute
							attach="attributes-position"
							array={new Float32Array(attributesPosition)}
							count={countPosition}
							itemSize={3}
						/>
						<bufferAttribute
							attach="attributes-normal"
							array={new Float32Array(attributesNormal)}
							count={countNormal}
							itemSize={3}
						/>
						<index
							attach="index"
							array={new Uint16Array(attributesIndices)}
						/>
					</bufferGeometry>
					<meshStandardMaterial color="orange" />
				</mesh>
			</group>
		);
	}
	`
}
