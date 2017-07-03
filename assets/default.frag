#version 330

out vec4 frag_colour;

in vec4 gl_FragCoord;

uniform vec3 colour;

void main() {


    float r = sin(gl_FragCoord.x / 4) / 2 + 0.5;
    float g = sin(gl_FragCoord.y / 4) / 2 + 0.5;
    float b = sin(gl_FragCoord.z / 4) / 2 + 0.5;
    //float b = 0.5;

    //frag_colour = (vec4(r, g, b, 1) + vec4(colour, 1)) / 2;
    frag_colour = vec4(colour, 1);
}